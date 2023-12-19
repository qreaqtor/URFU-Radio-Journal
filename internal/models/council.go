package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CouncilMemberCreate struct {
	Name        text               `json:"name" bson:"name" binding:"required"`
	Email       string             `json:"email" bson:"email" binding:"required,email"`
	ImagePathId primitive.ObjectID `json:"imagePathId" bson:"imagePathId" binding:"required"`
	ScopusURL   string             `json:"scopus" bson:"scopus" binding:"required,url"`
	Description text               `json:"description" bson:"description" binding:"required"`
	Content     text               `json:"content" bson:"content" binding:"required"`
	Rank        string             `json:"rank" bson:"rank" binding:"required"`
}

type CouncilMemberRead struct {
	Id          primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Name        text               `json:"name" bson:"name" binding:"required"`
	Email       string             `json:"email" bson:"email" binding:"required,email"`
	ImagePathId primitive.ObjectID `json:"imagePathId" bson:"imagePathId" binding:"required"`
	ScopusURL   string             `json:"scopus" bson:"scopus" binding:"required,url"`
	Description text               `json:"description" bson:"description" binding:"required"`
	Content     text               `json:"content" bson:"content" binding:"required"`
	Rank        string             `json:"rank" bson:"rank" binding:"required"`
}

type CouncilMemberUpdate struct {
	Name        text               `json:"name" bson:"name,omitempty" binding:"-"`
	Email       string             `json:"email" bson:"email,omitempty" binding:"omitempty,email"`
	ImagePathId primitive.ObjectID `json:"imagePathId" bson:"imagePathId,omitempty" binding:"-"`
	ScopusURL   string             `json:"scopus" bson:"scopus,omitempty" binding:"omitempty,url"`
	Description text               `json:"description" bson:"description,omitempty" binding:"-"`
	Content     text               `json:"content" bson:"content,omitempty" binding:"-"`
	Rank        string             `json:"rank" bson:"rank,omitempty" binding:"-"`
}
