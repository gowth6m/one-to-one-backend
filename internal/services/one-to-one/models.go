package one_to_one

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WellbeingScores struct {
	WorkOverall           int `json:"workOverall" bson:"workOverall" validate:"required"`
	Wellbeing             int `json:"wellbeing" bson:"wellbeing" validate:"required"`
	Growth                int `json:"growth" bson:"growth" validate:"required"`
	WorkRelationships     int `json:"workRelationships" bson:"workRelationships" validate:"required"`
	ImpactAndProductivity int `json:"impactAndProductivity" bson:"impactAndProductivity" validate:"required"`
}

type GoneWell struct {
	Label string `json:"label" bson:"label" validate:"required"`
	Theme string `json:"theme" bson:"theme" validate:"required"`
}

type Challenges struct {
	Label string `json:"label" bson:"label" validate:"required"`
	Theme string `json:"theme" bson:"theme" validate:"required"`
}

type Agenda struct {
	Label string `json:"label" bson:"label" validate:"required"`
}

type UserSummary struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email" validate:"required"`
	FirstName string             `json:"firstName" bson:"firstName" validate:"required"`
	LastName  string             `json:"lastName" bson:"lastName" validate:"required"`
}

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ CREATE OBJECTS -----------------------------------------
// ---------------------------------------------------------------------------------------------------

type CreateWeeklyReportRequest struct {
	Week            int             `json:"week" binding:"required" bson:"week"`
	Year            int             `json:"year" binding:"required" bson:"year"`
	WellbeingScores WellbeingScores `json:"wellbeingScores" binding:"required" bson:"wellbeingScores"`
	Agendas         []Agenda        `json:"agendas" binding:"required" bson:"agendas"`
	GoneWell        []GoneWell      `json:"goneWell" binding:"required" bson:"goneWell"`
	Challenges      []Challenges    `json:"challenges" binding:"required" bson:"challenges"`
}

type UpdateWeeklyReportRequest struct {
	ID              primitive.ObjectID `json:"id" binding:"required" bson:"_id"`
	Week            int                `json:"week" binding:"required" bson:"week"`
	Year            int                `json:"year" binding:"required" bson:"year"`
	WellbeingScores WellbeingScores    `json:"wellbeingScores" binding:"required" bson:"wellbeingScores"`
	Agendas         []Agenda           `json:"agendas" binding:"required" bson:"agendas"`
	GoneWell        []GoneWell         `json:"goneWell" binding:"required" bson:"goneWell"`
	Challenges      []Challenges       `json:"challenges" binding:"required" bson:"challenges"`
}

// ---------------------------------------------------------------------------------------------------
// ----------------------------------------- RESPONSE OBJECTS ----------------------------------------
// ---------------------------------------------------------------------------------------------------

type WeeklyReportResponse struct {
	ID              primitive.ObjectID `json:"id,omitempty"`
	Reportee        primitive.ObjectID `json:"reportee"`
	ReportingTo     primitive.ObjectID `json:"reportingTo"`
	Week            int                `json:"week"`
	Year            int                `json:"year"`
	WellbeingScores WellbeingScores    `json:"wellbeingScores"`
	Agendas         []Agenda           `json:"agendas"`
	GoneWell        []GoneWell         `json:"goneWell"`
	Challenges      []Challenges       `json:"challenges"`
	CreatedAt       time.Time          `json:"createdAt,omitempty"`
	UpdatedAt       time.Time          `json:"updatedAt,omitempty"`
}

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ MONGO OBJECTS ------------------------------------------
// ---------------------------------------------------------------------------------------------------

type WeeklyReport struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	Reportee        primitive.ObjectID `json:"reportee" bson:"reportee" validate:"required"`
	ReportingTo     primitive.ObjectID `json:"reportingTo" bson:"reportingTo" validate:"required"`
	Week            int                `json:"week" bson:"week" validate:"required"`
	Year            int                `json:"year" bson:"year" validate:"required"`
	WellbeingScores WellbeingScores    `json:"wellbeingScores" bson:"wellbeingScores" validate:"required"`
	Agendas         []Agenda           `json:"agendas" bson:"agendas" validate:"required"`
	GoneWell        []GoneWell         `json:"goneWell" bson:"goneWell" validate:"required"`
	Challenges      []Challenges       `json:"challenges" bson:"challenges" validate:"required"`
	CreatedAt       primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt       primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
