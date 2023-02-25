package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Property struct {
	PropertyID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name" binding:"required,alphanum"`
	Price      int                `json:"price,string" bson:"price" binding:"required,numeric"`
	BuildPrice int                `json:"buildPrice,string" bson:"buildPrice" binding:"required,numeric"`
	Rent       int                `json:"rent,string" bson:"rent" binding:"required,numeric"`
	Rent1H     int                `json:"rent1H,string" bson:"rent1H" binding:"required,numeric"`
	Rent2H     int                `json:"rent2H,string" bson:"rent2H" binding:"required,numeric"`
	Rent3H     int                `json:"rent3H,string" bson:"rent3H" binding:"required,numeric"`
	Rent4H     int                `json:"rent4H,string" bson:"rent4H" binding:"required,numeric"`
	RentHotel  int                `json:"rentHotel,string" bson:"rentHotel" binding:"required,numeric"`
	Mortgage   int                `json:"mortgage,string" bson:"mortgage" binding:"required,numeric"`
	Color      string             `json:"color" bson:"color" binding:"required,alphanum"`
}
