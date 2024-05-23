package models

type Text struct {
	Ru  string `json:"Ru" bson:"Ru" binding:"required"`
	Eng string `json:"Eng" bson:"Eng" binding:"required"`
}
