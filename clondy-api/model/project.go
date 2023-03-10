package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Project .....
type Project struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id"`
	UserID      string             `json:"user_id" bson:"user_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdateAt    time.Time          `json:"update_at" bson:"update_at"`
	Properties  map[string]interface{}
}

// Properties model
type Properties interface {
}
