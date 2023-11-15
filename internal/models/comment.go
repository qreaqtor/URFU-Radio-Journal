package models

import "time"

type Comment struct {
	Content string    `json:"content" bson:"content" binding:"required"`
	Date    time.Time `json:"date" bson:"date" binding:"required"`
}
