package users

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Username  string             `json:"username,omitempty" validate:"required"`
	Password  string             `json:"password,omitempty" validate:"required"`
	Firstname string             `json:"firstname,omitempty" validate:"required"`
	Lastname  string             `json:"lastname,omitempty" validate:"required"`
}
