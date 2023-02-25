package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Game struct {
	GameID  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserIDs User               `json:"UserIDs" bson:"UserIDs"`
}

type Player struct {
	PlayerID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	GameID        Game               `json:"gameID" bson:"gameID" binding:""`
	PlayerName    string             `json:"playerName" bson:"playerName" binding:"required"`
	Piece         string             `json:"piece" bson:"piece" binding:"required"`
	PlayerBalance string             `json:"playerBalance" bson:"playerBalance" binding:"required"`
	PropertyOwned []Property         `json:"playerProperties" bson:"playerProperties"`
}
