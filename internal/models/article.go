package models

import (
	"time"
)

type Article struct {
	Title      string    `json:"title" bson:"title" binding:"required"`
	Authors    []string  `json:"authors" bson:"authors" binding:"required"`
	Annotation string    `json:"annotation" bson:"annotation" binding:"required"`
	Keywords   []string  `json:"keywords" bson:"keywords" binding:"required"`
	FileName   string    `json:"fileName" bson:"fileName" binding:"required"`
	Literature []string  `json:"literature" bson:"literature" binding:"required"`
	Date       time.Time `json:"publication" bson:"publication" binding:"required"`
}
