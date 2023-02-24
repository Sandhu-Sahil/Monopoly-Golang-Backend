package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email" bson:"email" binding:"required"`
	UserName string             `json:"user_name" bson:"user_name" binding:"required"`
	Password string             `json:"password" bson:"password" binding:"required"`
	GameIDs  []Game             `json:"gameIDs" bson:"gameIDs"`
}

type Game struct {
	GameID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PlayerIDs []Player           `json:"playerIDs" bson:"playerIDs"`
}

type Player struct {
	PlayerID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PlayerName    string             `json:"playerName" bson:"playerName" binding:"required"`
	Piece         string             `json:"piece" bson:"piece" binding:"required"`
	PlayerBalance string             `json:"playerBalance" bson:"playerBalance" binding:"required"`
}
