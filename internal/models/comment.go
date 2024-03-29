package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentCreate struct {
	ArticleId   primitive.ObjectID `json:"articleId" bson:"articleId" binding:"required"`
	ContentPart string             `json:"content" bson:"-" binding:"required"`
	Content     text               `json:"-" bson:"content" binding:"-"`
	Date        time.Time          `json:"date" bson:"date" binding:"required"`
	IsApproved  bool               `json:"-" bson:"isApproved"`
}

type CommentRead struct {
	Id         primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	ArticleId  primitive.ObjectID `json:"articleId" bson:"articleId" binding:"required"`
	Content    text               `json:"content" bson:"content" binding:"required"`
	Date       time.Time          `json:"date" bson:"date" binding:"required"`
	IsApproved bool               `json:"isApproved" bson:"isApproved" binding:"required"`
}

type CommentUpdate struct {
	Id      primitive.ObjectID `json:"id" bson:"-" binding:"required"`
	Content text               `json:"content" bson:"content" binding:"required"`
}

type CommentApprove struct {
	Id          primitive.ObjectID `json:"id" binding:"required"`
	ContentPart string             `json:"content" binding:"required"`
}
