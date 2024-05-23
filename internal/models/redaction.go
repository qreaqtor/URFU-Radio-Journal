package models

import (
	"time"
)

type RedactionMemberCreate struct {
	Name        Text      `json:"name" bson:"name" binding:"required"`
	Email       string    `json:"email" bson:"email" binding:"required,email"`
	ImagePathId string    `json:"imagePathId" bson:"imagePathId" binding:"required"`
	Description Text      `json:"description" bson:"description" binding:"required"`
	Content     Text      `json:"content" bson:"content" binding:"required"`
	Rank        string    `json:"rank" bson:"rank" binding:"required"`
	Location    Text      `json:"location" bson:"location" binding:"required"`
	DateJoin    time.Time `json:"dateJoin" bson:"dateJoin" binding:"required"`
}

type RedactionMemberRead struct {
	Id          int       `json:"id" bson:"_id" binding:"required"`
	Name        Text      `json:"name" bson:"name" binding:"required"`
	Email       string    `json:"email" bson:"email" binding:"required,email"`
	ImagePathId string    `json:"imagePathId" bson:"imagePathId" binding:"required"`
	Description Text      `json:"description" bson:"description" binding:"required"`
	Content     Text      `json:"content" bson:"content" binding:"required"`
	Rank        string    `json:"rank" bson:"rank" binding:"required"`
	Location    Text      `json:"location" bson:"location" binding:"required"`
	DateJoin    time.Time `json:"dateJoin" bson:"dateJoin" binding:"required"`
}

type RedactionMemberUpdate struct {
	Name        Text      `json:"name" bson:"name,omitempty" binding:"-"`
	Email       string    `json:"email" bson:"email,omitempty" binding:"omitempty,email"`
	ImagePathId string    `json:"imagePathId" bson:"imagePathId,omitempty" binding:"-"`
	Description Text      `json:"description" bson:"description,omitempty" binding:"-"`
	Content     Text      `json:"content" bson:"content,omitempty" binding:"-"`
	Rank        string    `json:"rank" bson:"rank,omitempty" binding:"-"`
	Location    Text      `json:"location" bson:"location,omitempty" binding:"-"`
	DateJoin    time.Time `json:"dateJoin" bson:"dateJoin,omitempty" binding:"-"`
}
