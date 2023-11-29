package models

type Article struct {
	Title      text     `json:"title" bson:"title" binding:"required,dive"`
	Authors    []string `json:"authors" bson:"authors" binding:"required"`
	Content    text     `json:"content" bson:"content" binding:"required,dive"`
	Keywords   []text   `json:"keywords" bson:"keywords" binding:"required,dive"`
	FileName   string   `json:"fileName" bson:"fileName" binding:"required"`
	Literature []string `json:"literature" bson:"literature" binding:"required"`
}
