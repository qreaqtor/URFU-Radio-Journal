package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ArticleCreate struct {
	EditionId  primitive.ObjectID `json:"editionId" bson:"editionId" binding:"required"`
	Title      text               `json:"title" bson:"title" binding:"required"`
	Authors    []string           `json:"authors" bson:"authors" binding:"required"`
	Content    text               `json:"content" bson:"content" binding:"required"`
	Keywords   []text             `json:"keywords" bson:"keywords" binding:"required"`
	FilePathId primitive.ObjectID `json:"filePathId" bson:"filePathId" binding:"required"`
	Literature []string           `json:"literature" bson:"literature" binding:"required"`
}

type ArticleRead struct {
	Id         primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	EditionId  primitive.ObjectID `json:"editionId" bson:"editionId" binding:"required"`
	Title      text               `json:"title" bson:"title" binding:"required"`
	Authors    []string           `json:"authors" bson:"authors" binding:"required"`
	Content    text               `json:"content" bson:"content" binding:"required"`
	Keywords   []text             `json:"keywords" bson:"keywords" binding:"required"`
	FilePathId primitive.ObjectID `json:"filePathId" bson:"filePathId" binding:"required"`
	Literature []string           `json:"literature" bson:"literature" binding:"required"`
}

type ArticleUpdate struct {
	Id         primitive.ObjectID `json:"id" binding:"required"`
	Title      text               `json:"title" bson:"title,omitempty" binding:"-"`
	Authors    []string           `json:"authors" bson:"authors,omitempty" binding:"-"`
	Content    text               `json:"content" bson:"content,omitempty" binding:"-"`
	Keywords   []text             `json:"keywords" bson:"keywords,omitempty" binding:"-"`
	FilePathId primitive.ObjectID `json:"filePathId" bson:"filePathId,omitempty" binding:"-"`
	Literature []string           `json:"literature" bson:"literature,omitempty" binding:"-"`
}
