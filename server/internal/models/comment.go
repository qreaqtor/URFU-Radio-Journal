package models

import (
	"time"
)

type CommentCreate struct {
	ArticleId   int       `json:"articleId" bson:"articleId" binding:"required"`
	ContentPart string    `json:"content" bson:"-" binding:"required"`
	Content     Text      `json:"-" bson:"content" binding:"-"`
	Date        time.Time `json:"date" bson:"date" binding:"required"`
	IsApproved  bool      `json:"-" bson:"isApproved"`
	Author      string    `json:"author" bson:"author" binding:"required"`
}

type CommentRead struct {
	Id         int       `json:"id" bson:"_id" binding:"required"`
	ArticleId  int       `json:"articleId" bson:"articleId" binding:"required"`
	Content    Text      `json:"content" bson:"content" binding:"required"`
	Date       time.Time `json:"date" bson:"date" binding:"required"`
	IsApproved bool      `json:"isApproved" bson:"isApproved" binding:"required"`
	Author     string    `json:"author" bson:"author" binding:"required"`
}

type CommentUpdate struct {
	Id      int  `json:"id" bson:"-" binding:"required"`
	Content Text `json:"content" bson:"content" binding:"required"`
}

type CommentApprove struct {
	Id          int    `json:"id" binding:"required"`
	ContentPart string `json:"content" binding:"required"`
}

type CommentQuery struct {
	OnlyApproved bool `form:"onlyApproved"`
	ArticleID    int  `form:"articleId" binding:"required"`
	BatchArgs
}
