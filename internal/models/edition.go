package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EditionCreate struct {
	Title text `json:"title" bson:"title" binding:"required,dive"`
	// FileName  string               `json:"fileName" bson:"fileName" binding:"required"`
	// CoverName string               `json:"coverName" bson:"coverName" binding:"required"`
	// VideoName string               `json:"videoName" bson:"videoName" binding:"-"`
	Date     time.Time            `json:"date" bson:"date" binding:"required"`
	Articles []primitive.ObjectID `json:"articles" bson:"articles" binding:"required"`
}

type EditionRead struct {
	Id    primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Title text               `json:"title" bson:"title" binding:"required,dive"`
	// FileName  string               `json:"fileName" bson:"fileName" binding:"required"`
	// CoverName string               `json:"coverName" bson:"coverName" binding:"required"`
	// VideoName string               `json:"videoName" bson:"videoName" binding:"required"`
	Date     time.Time            `json:"date" bson:"date" binding:"required"`
	Articles []primitive.ObjectID `json:"articles" bson:"articles" binding:"required"`
}

type EditionUpdate struct {
	Id    primitive.ObjectID `json:"id" binding:"required"`
	Title text               `json:"title" bson:"title,omitempty" binding:"dive"`
	// FileName  string               `json:"fileName" bson:"fileName,omitempty" binding:"-"`
	// CoverName string               `json:"coverName" bson:"coverName,omitempty" binding:"-"`
	// VideoName string               `json:"videoName" bson:"videoName,omitempty" binding:"-"`
	Articles []primitive.ObjectID `json:"articles" bson:"articles,omitempty" binding:"-"`
}
