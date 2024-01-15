package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EditionCreate struct {
	Title       text               `json:"title" bson:"title" binding:"required"`
	FilePathId  primitive.ObjectID `json:"filePathId" bson:"filePathId" binding:"required"`
	CoverPathId primitive.ObjectID `json:"coverPathId" bson:"coverPathId" binding:"required"`
	VideoPathId primitive.ObjectID `json:"videoPathId" bson:"videoPathId" binding:"-"`
	Date        time.Time          `json:"date" bson:"date" binding:"required"`
}

type EditionRead struct {
	Id          primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Title       text               `json:"title" bson:"title" binding:"required"`
	FilePathId  primitive.ObjectID `json:"filePathId" bson:"filePathId" binding:"required"`
	CoverPathId primitive.ObjectID `json:"coverPathId" bson:"coverPathId" binding:"required"`
	VideoPathId primitive.ObjectID `json:"videoPathId" bson:"videoPathId" binding:"required"`
	Date        time.Time          `json:"date" bson:"date" binding:"required"`
}

type EditionUpdate struct {
	Id          primitive.ObjectID `json:"id" binding:"required"`
	Title       text               `json:"title" bson:"title,omitempty" binding:"-"`
	FilePathId  primitive.ObjectID `json:"filePathId" bson:"filePathId,omitempty" binding:"-"`
	CoverPathId primitive.ObjectID `json:"coverPathId" bson:"coverPathId,omitempty" binding:"-"`
	VideoPathId primitive.ObjectID `json:"videoPathId" bson:"videoPathId,omitempty" binding:"-"`
}
