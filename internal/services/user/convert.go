package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func ConvertCreateUserRequestToUser(req CreateUserRequest) (User, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	return User{
		ID:        primitive.NewObjectID(),
		Password:  string(hashed),
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Reportees: []primitive.ObjectID{},
		ReportsTo: nil,
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
