package user

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type Session struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ CREATE OBJECTS -----------------------------------------
// ---------------------------------------------------------------------------------------------------
type CreateUserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"firstName" binding:"required,alpha"`
	LastName  string `json:"lastName" binding:"required,alpha"`
}

// ---------------------------------------------------------------------------------------------------
// ----------------------------------------- RESPONSE OBJECTS ----------------------------------------
// ---------------------------------------------------------------------------------------------------
type UserResponse struct {
	ID        string   `json:"id,omitempty"`
	Email     string   `json:"email"`
	FirstName string   `json:"firstName,omitempty"`
	LastName  string   `json:"lastName,omitempty"`
	ReportsTo *string  `json:"reportsTo,omitempty"`
	Reportees []string `json:"reportees,omitempty"`
	CreatedAt string   `json:"createdAt,omitempty"`
	UpdatedAt string   `json:"updatedAt,omitempty"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ MONGO OBJECTS ------------------------------------------
// ---------------------------------------------------------------------------------------------------
type User struct {
	ID        primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	Password  string               `json:"-" bson:"password,omitempty" validate:"required"`
	Email     string               `json:"email" bson:"email" validate:"required,email"`
	FirstName string               `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName  string               `json:"lastName,omitempty" bson:"lastName,omitempty"`
	ReportsTo *primitive.ObjectID  `json:"reportsTo,omitempty" bson:"reportsTo,omitempty"`
	Reportees []primitive.ObjectID `json:"reportees,omitempty" bson:"reportees,omitempty"`
	CreatedAt primitive.DateTime   `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime   `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
