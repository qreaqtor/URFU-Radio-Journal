package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EditionCreate struct {
	Title     string    `json:"title" bson:"title" binding:"required"`
	FileName  string    `json:"fileName" bson:"fileName" binding:"required"`
	CoverName string    `json:"coverName" bson:"coverName" binding:"required"`
	Date      time.Time `json:"date" bson:"date" binding:"required"`
	Articles  []Article `json:"articles" bson:"articles" binding:"required,dive"`
}

type EditionRead struct {
	Id        primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Title     string             `json:"title" bson:"title" binding:"required"`
	FileName  string             `json:"fileName" bson:"fileName" binding:"required"`
	CoverName string             `json:"coverName" bson:"coverName" binding:"required"`
	Date      time.Time          `json:"date" bson:"date" binding:"required"`
	Articles  []Article          `json:"articles" bson:"articles" binding:"required"`
}

type EditionUpdate struct {
	Id        primitive.ObjectID `json:"id" binding:"required"`
	Title     string             `json:"title" bson:"title,omitempty" binding:"-"`
	FileName  string             `json:"fileName" bson:"fileName,omitempty" binding:"-"`
	CoverName string             `json:"coverName" bson:"coverName,omitempty" binding:"-"`
	Articles  []Article          `json:"articles" bson:"articles,omitempty" binding:"dive"`
}
