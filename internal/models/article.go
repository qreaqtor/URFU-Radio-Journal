package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ArticleCreate struct {
	Title      text     `json:"title" bson:"title" binding:"required,dive"`
	Authors    []string `json:"authors" bson:"authors" binding:"required"`
	Content    text     `json:"content" bson:"content" binding:"required,dive"`
	Keywords   []text   `json:"keywords" bson:"keywords" binding:"required,dive"`
	FilePathId string   `json:"filePathId" bson:"filePathId" binding:"required"`
	Literature []string `json:"literature" bson:"literature" binding:"required"`
}

type ArticleRead struct {
	Id         primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Title      text               `json:"title" bson:"title" binding:"required,dive"`
	Authors    []string           `json:"authors" bson:"authors" binding:"required"`
	Content    text               `json:"content" bson:"content" binding:"required,dive"`
	Keywords   []text             `json:"keywords" bson:"keywords" binding:"required,dive"`
	FilePathId string             `json:"filePathId" bson:"filePathId" binding:"required"`
	Literature []string           `json:"literature" bson:"literature" binding:"required"`
}

type ArticleUpdate struct {
	Id         primitive.ObjectID `json:"id" binding:"required"`
	Title      text               `json:"title" bson:"title,omitempty" binding:"dive"`
	Authors    []string           `json:"authors" bson:"authors,omitempty" binding:"-"`
	Content    text               `json:"content" bson:"content,omitempty" binding:"dive"`
	Keywords   []text             `json:"keywords" bson:"keywords,omitempty" binding:"dive"`
	FilePathId string             `json:"filePathId" bson:"filePathId,omitempty" binding:"-"`
	Literature []string           `json:"literature" bson:"literature,omitempty" binding:"-"`
}
