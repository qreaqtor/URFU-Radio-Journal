package models

type BatchArgs struct {
	Offset int `form:"offset" binding:"gte=0"`
	Limit  int `form:"limit" binding:"gte=0"`
}
