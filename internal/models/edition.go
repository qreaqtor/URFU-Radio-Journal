package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Edition struct {
	Id        primitive.ObjectID `json:"id" bson:"_id" binding:"omitempty"`
	Title     string             `json:"title" bson:"title" binding:"required"`
	FileName  string             `json:"fileName" bson:"fileName" binding:"required"`
	CoverName string             `json:"coverName" bson:"coverName" binding:"required"`
	Date      time.Time          `json:"date" bson:"date" binding:"required" time_format:"yyyy-mm-dd"`
	Articles  []Article          `json:"articles" bson:"articles" binding:"required"`
	Comments  []Comment          `json:"comments" bson:"comments" binding:"omitempty" default:"[]"`
}
