package models

import (
	"time"
)

type EditionCreate struct {
	Year       int       `json:"year" bson:"year" binding:"required"`
	Number     int       `json:"number" bson:"number" binding:"required"`
	Volume     int       `json:"volume" bson:"volume" binding:"required"`
	DocumentID string    `json:"documentID" bson:"documentID" binding:"required"`
	ImageID    string    `json:"imageID" bson:"imageID" binding:"required"`
	Date       time.Time `json:"date" bson:"date" binding:"required"`
}

type EditionRead struct {
	Id         int       `json:"id" bson:"_id" binding:"required"`
	Year       int       `json:"year" bson:"year" binding:"required"`
	Number     int       `json:"number" bson:"number" binding:"required"`
	Volume     int       `json:"volume" bson:"volume" binding:"required"`
	DocumentID string    `json:"documentID" bson:"documentID" binding:"required"`
	ImageID    string    `json:"imageID" bson:"imageID" binding:"required"`
	Date       time.Time `json:"date" bson:"date" binding:"required"`
}

type EditionUpdate struct {
	Id         int    `json:"id" binding:"required"`
	Year       int    `json:"year" bson:"year,omitempty" binding:"-"`
	Number     int    `json:"number" bson:"number,omitempty" binding:"-"`
	Volume     int    `json:"volume" bson:"volume,omitempty" binding:"-"`
	DocumentID string `json:"documentID" bson:"documentID,omitempty" binding:"-"`
	ImageID    string `json:"imageID" bson:"imageID,omitempty" binding:"-"`
}
