package models

import "time"

type Article struct {
	Title       string    `json:"title" bson:"title" binding:"required"`
	Authors     []string  `json:"authors" bson:"authors" binding:"required"`
	Annotation  string    `json:"annotation" bson:"annotation" binding:"required"`
	Keywords    []string  `json:"keywords" bson:"keywords" binding:"required"`
	Filename    string    `json:"filename" bson:"filename" binding:"required"`
	Literature  []string  `json:"literature" bson:"literature" binding:"required"`
	Publication time.Time `json:"publication" bson:"publication" binding:"required"`
}
