package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Edition struct {
	Id        primitive.ObjectID `json:"id" bson:"_id" binding:"omitempty"`
	Title     string             `json:"title" bson:"title" binding:"required"`
	FileName  string             `json:"fileName" bson:"fileName" binding:"required"`
	CoverName string             `json:"coverName" bson:"coverName" binding:"required"`
	Articles  []Article          `json:"articles" bson:"articles" binding:"required"`
}
