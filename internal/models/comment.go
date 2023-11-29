package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentCreate struct {
	ArticleId  primitive.ObjectID `json:"editionId" bson:"editionId" binding:"required"`
	Content    text               `json:"content" bson:"content" binding:"required,dive"`
	Date       time.Time          `json:"date" bson:"date" binding:"required"`
	IsApproved bool               `bson:"isApproved"`
}

type CommentRead struct {
	Id         primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	ArticleId  primitive.ObjectID `json:"editionId" bson:"editionId" binding:"required"`
	Content    text               `json:"content" bson:"content" binding:"required,dive"`
	Date       time.Time          `json:"date" bson:"date" binding:"required"`
	IsApproved bool               `json:"isApproved" bson:"isApproved" binding:"required"`
}

type CommentUpdate struct {
	Id      primitive.ObjectID `json:"id" bson:"-" binding:"required"`
	Content text               `json:"content" bson:"content" binding:"required,dive"`
}
