package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email" bson:"email" binding:"required,email"`
	UserName string             `json:"user_name" bson:"user_name" binding:"required,alphanum"`
	Password string             `json:"password" bson:"password" binding:"required"`
}
