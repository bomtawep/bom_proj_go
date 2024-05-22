package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Email     string             `json:"email,omitempty" validate:"required,email"`
	Password  string             `json:"password,omitempty" validate:"required"`
	Firstname string             `json:"firstname,omitempty" validate:"required"`
	Lastname  string             `json:"lastname,omitempty" validate:"required"`
	Gender    string             `json:"gender"`
	Phone     string             `json:"phone"`
	Birthdate string             `json:"birthdate"`
}

type Email struct {
	Email string `json:"email,omitempty" validate:"required,email"`
}
