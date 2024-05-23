package models

type Author struct {
	FullName   Text   `json:"fullname" bson:"fullname" binding:"required"`
	Affilation string `json:"affilation" bson:"affilation" binding:"required"`
	Email      string `json:"email" bson:"email" binding:"required,email"`
}
