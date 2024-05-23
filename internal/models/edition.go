package models

import (
	"time"
)

type EditionCreate struct {
	Year        int       `json:"year" bson:"year" binding:"required"`
	Number      int       `json:"number" bson:"number" binding:"required"`
	Volume      int       `json:"volume" bson:"volume" binding:"required"`
	FilePathId  string    `json:"filePathId" bson:"filePathId" binding:"required"`
	CoverPathId string    `json:"coverPathId" bson:"coverPathId" binding:"required"`
	Date        time.Time `json:"date" bson:"date" binding:"required"`
}

type EditionRead struct {
	Id          int       `json:"id" bson:"_id" binding:"required"`
	Year        int       `json:"year" bson:"year" binding:"required"`
	Number      int       `json:"number" bson:"number" binding:"required"`
	Volume      int       `json:"volume" bson:"volume" binding:"required"`
	FilePathId  string    `json:"filePathId" bson:"filePathId" binding:"required"`
	CoverPathId string    `json:"coverPathId" bson:"coverPathId" binding:"required"`
	Date        time.Time `json:"date" bson:"date" binding:"required"`
}

type EditionUpdate struct {
	Id          int    `json:"id" binding:"required"`
	Year        int    `json:"year" bson:"year,omitempty" binding:"-"`
	Number      int    `json:"number" bson:"number,omitempty" binding:"-"`
	Volume      int    `json:"volume" bson:"volume,omitempty" binding:"-"`
	FilePathId  string `json:"filePathId" bson:"filePathId,omitempty" binding:"-"`
	CoverPathId string `json:"coverPathId" bson:"coverPathId,omitempty" binding:"-"`
}
