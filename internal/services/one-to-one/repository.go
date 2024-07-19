package one_to_one

import (
	"context"
	"fmt"
	"one-to-one/internal/db"
	user "one-to-one/internal/services/user"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OneToOneRepository interface {
	CreateWeeklyReport(c context.Context, report CreateWeeklyReportRequest, currentUserId primitive.ObjectID) (WeeklyReport, error)
	GetAllWeeklyReports(c context.Context, currentUserId primitive.ObjectID, isReportee bool) ([]WeeklyReport, error)
	UpdateWeeklyReport(c context.Context, report UpdateWeeklyReportRequest, currentUserId primitive.ObjectID, isReportee bool) (WeeklyReport, error)
	GetWeeklyReportByWeekAndYear(c context.Context, week int, year int, currentUserId primitive.ObjectID, isReportee bool) (WeeklyReport, error)
}

type repositoryImpl struct {
	collection     *mongo.Collection
	userCollection *mongo.Collection
}

func NewOneToOneRepository() OneToOneRepository {
	collection := db.Client.Database(db.DATABASE_NAME).Collection(db.COLLECTION_WEEKLY_REPORT)
	userCollection := db.Client.Database(db.DATABASE_NAME).Collection(db.COLLECTION_USER)
	return &repositoryImpl{collection: collection, userCollection: userCollection}
}

func (r *repositoryImpl) CreateWeeklyReport(c context.Context, report CreateWeeklyReportRequest, currentUserId primitive.ObjectID) (WeeklyReport, error) {

	var reportee user.User
	err := r.userCollection.FindOne(c, bson.M{"_id": currentUserId}).Decode(&reportee)
	if err != nil {
		return WeeklyReport{}, err
	}

	// Find reportingTo user by ID
	var reportingTo user.User
	err = r.userCollection.FindOne(c, bson.M{"_id": reportee.ReportsTo}).Decode(&reportingTo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return WeeklyReport{}, fmt.Errorf("no user to report to")
		}
		return WeeklyReport{}, err
	}

	// Create WeeklyReport
	mongoReport := WeeklyReport{
		ID:              primitive.NewObjectID(),
		Reportee:        reportee.ID,
		ReportingTo:     reportingTo.ID,
		Week:            report.Week,
		Year:            report.Year,
		WellbeingScores: report.WellbeingScores,
		Agendas:         report.Agendas,
		GoneWell:        report.GoneWell,
		Challenges:      report.Challenges,
		CreatedAt:       primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:       primitive.NewDateTimeFromTime(time.Now()),
	}

	// Insert the new WeeklyReport into the collection
	_, err = r.collection.InsertOne(c, mongoReport)
	if err != nil {
		return WeeklyReport{}, err
	}

	return mongoReport, nil
}

func (r *repositoryImpl) GetAllWeeklyReports(c context.Context, currentUserId primitive.ObjectID, isReportee bool) ([]WeeklyReport, error) {
	var filter bson.M
	if isReportee {
		filter = bson.M{"reportee": currentUserId}
	} else {
		filter = bson.M{"reportingTo": currentUserId}
	}

	findOptions := options.Find().SetSort(bson.D{
		{Key: "year", Value: -1},
		{Key: "week", Value: -1},
	})

	cursor, err := r.collection.Find(c, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	var reports []WeeklyReport
	if err = cursor.All(c, &reports); err != nil {
		return nil, err
	}

	return reports, nil
}

func (r *repositoryImpl) UpdateWeeklyReport(c context.Context, report UpdateWeeklyReportRequest, currentUserId primitive.ObjectID, isReportee bool) (WeeklyReport, error) {
	// Find current user by ID
	var currentUser user.User
	err := r.userCollection.FindOne(c, bson.M{"_id": currentUserId}).Decode(&currentUser)
	if err != nil {
		return WeeklyReport{}, err
	}

	// Find the object for reportingTo user
	var reportingTo user.User
	err = r.userCollection.FindOne(c, bson.M{"_id": currentUser.ReportsTo}).Decode(&reportingTo)
	if err != nil {
		return WeeklyReport{}, err
	}

	var reportObj WeeklyReport
	err = r.collection.FindOne(c, bson.M{"_id": report.ID}).Decode(&reportObj)
	if err != nil {
		return WeeklyReport{}, err
	}

	var filter bson.M
	if isReportee {
		filter = bson.M{
			"_id":      report.ID,
			"reportee": currentUserId,
		}
	} else {
		filter = bson.M{
			"_id":         report.ID,
			"reportingTo": currentUserId,
		}
	}

	updatedReport := WeeklyReport{
		ID:              report.ID,
		Reportee:        currentUserId,
		ReportingTo:     reportingTo.ID,
		Week:            report.Week,
		Year:            report.Year,
		WellbeingScores: report.WellbeingScores,
		Agendas:         report.Agendas,
		GoneWell:        report.GoneWell,
		Challenges:      report.Challenges,
		UpdatedAt:       primitive.NewDateTimeFromTime(time.Now()),
		CreatedAt:       reportObj.CreatedAt,
	}

	update := bson.M{
		"$set": updatedReport,
	}

	result, err := r.collection.UpdateOne(c, filter, update)
	if err != nil {
		return WeeklyReport{}, err
	}

	if result.MatchedCount == 0 {
		return WeeklyReport{}, mongo.ErrNoDocuments
	}

	return updatedReport, nil
}

func (r *repositoryImpl) GetWeeklyReportByWeekAndYear(c context.Context, week int, year int, currentUserId primitive.ObjectID, isReportee bool) (WeeklyReport, error) {
	var filter bson.M
	if isReportee {
		filter = bson.M{
			"week":     week,
			"year":     year,
			"reportee": currentUserId,
		}
	} else {
		filter = bson.M{
			"week":        week,
			"year":        year,
			"reportingTo": currentUserId,
		}
	}

	var report WeeklyReport
	err := r.collection.FindOne(c, filter).Decode(&report)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return WeeklyReport{}, err
		}
		return WeeklyReport{}, err
	}

	return report, nil
}
