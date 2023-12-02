package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EditionCreate struct {
	Title       text                 `json:"title" bson:"title" binding:"required,dive"`
	FilePathId  string               `json:"filePathId" bson:"filePathId" binding:"required"`
	CoverPathId string               `json:"coverPathId" bson:"coverPathId" binding:"required"`
	VideoPathId string               `json:"videoPathId" bson:"videoPathId" binding:"-"`
	Date        time.Time            `json:"date" bson:"date" binding:"required"`
	Articles    []primitive.ObjectID `json:"articles" bson:"articles" binding:"required"`
}

type EditionRead struct {
	Id          primitive.ObjectID   `json:"id" bson:"_id" binding:"required"`
	Title       text                 `json:"title" bson:"title" binding:"required,dive"`
	FilePathId  string               `json:"filePathId" bson:"filePathId" binding:"required"`
	CoverPathId string               `json:"coverPathId" bson:"coverPathId" binding:"required"`
	VideoPathId string               `json:"videoPathId" bson:"videoPathId" binding:"required"`
	Date        time.Time            `json:"date" bson:"date" binding:"required"`
	Articles    []primitive.ObjectID `json:"articles" bson:"articles" binding:"required"`
}

type EditionUpdate struct {
	Id          primitive.ObjectID   `json:"id" binding:"required"`
	Title       text                 `json:"title" bson:"title,omitempty" binding:"dive"`
	FilePathId  string               `json:"filePathId" bson:"filePathId,omitempty" binding:"-"`
	CoverPathId string               `json:"coverPathId" bson:"coverPathId,omitempty" binding:"-"`
	VideoPathId string               `json:"videoPathId" bson:"videoPathId,omitempty" binding:"-"`
	Articles    []primitive.ObjectID `json:"articles" bson:"articles,omitempty" binding:"-"`
}
