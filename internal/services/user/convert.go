package user

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func ConvertCreateUserRequestToUser(req CreateUserRequest) (User, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	defaultReportsTo, err := primitive.ObjectIDFromHex("6695a2379a9e246dc998afc7")
	if err != nil {
		return User{}, fmt.Errorf("invalid default ReportsTo ObjectID: %v", err)
	}

	return User{
		ID:        primitive.NewObjectID(),
		Password:  string(hashed),
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Reportees: []primitive.ObjectID{},
		ReportsTo: &defaultReportsTo,
	}, nil
}

func ConvertUserToUserResponse(user User) UserResponse {
	reportees := make([]string, len(user.Reportees))

	return UserResponse{
		ID: user.ID.Hex(),

		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		ReportsTo: nil,
		Reportees: reportees,
	}
}
