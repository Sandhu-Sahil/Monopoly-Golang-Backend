package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Property struct {
	PropertyID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	Price      int                `json:"price,string" bson:"price"`
	BuildPrice int                `json:"buildPrice,string" bson:"buildPrice"`
	Rent       int                `json:"rent,string" bson:"rent"`
	Rent1H     int                `json:"rent1H,string" bson:"rent1H"`
	Rent2H     int                `json:"rent2H,string" bson:"rent2H"`
	Rent3H     int                `json:"rent3H,string" bson:"rent3H"`
	Rent4H     int                `json:"rent4H,string" bson:"rent4H"`
	RentHotel  int                `json:"rentHotel,string" bson:"rentHotel"`
	Mortgage   int                `json:"mortgage,string" bson:"mortgage"`
	Color      string             `json:"color" bson:"color"`
}
