package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentCreate struct {
	Content    string    `json:"content" bson:"content" binding:"required"`
	Date       time.Time `json:"date" bson:"date" binding:"required"`
	IsApproved bool      `bson:"isApproved"`
}

type CommentRead struct {
	Id         primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Content    string             `json:"content" bson:"content" binding:"required"`
	Date       time.Time          `json:"date" bson:"date" binding:"required"`
	IsApproved bool               `json:"isApproved" bson:"isApproved" binding:"required"`
}

type CommentUpdate struct {
	Content    string `json:"content" bson:"content" binding:"required"`
	IsApproved bool   `bson:"isApproved"`
}
