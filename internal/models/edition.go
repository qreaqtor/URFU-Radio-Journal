package models

type Edition struct {
	Title    string    `json:"title" bson:"title" binding:"required"`
	Filename string    `json:"filename" bson:"filename" binding:"required"`
	Articles []Article `json:"articles" bson:"articles" binding:"required"`
	Cover    string    `json:"cover" bson:"cover" binding:"required"`
}
