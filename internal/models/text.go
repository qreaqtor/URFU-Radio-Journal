package models

type text struct {
	Ru  string `json:"Ru" bson:"Ru" binding:"required"`
	Eng string `json:"Eng" bson:"Eng" binding:"required"`
}

// type textCommentCreate struct {
// 	Ru  string `json:"Ru" bson:"Ru" binding:"-"`
// 	Eng string `json:"Eng" bson:"Eng" binding:"-"`
// }
