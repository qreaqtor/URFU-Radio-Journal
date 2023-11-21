package services

import (
	"context"
	"urfu-radio-journal/pkg/db"

	"go.mongodb.org/mongo-driver/mongo"
)

type CommentsService struct {
	ctx     context.Context
	storage *mongo.Collection
}

func NewCommentsService() *CommentsService {
	return &CommentsService{
		ctx:     *db.GetContext(),
		storage: db.GetStorage("comments"),
	}
}

func (this *CommentsService) Create() {

}

func (this *CommentsService) GetAll() {

}

func (this *CommentsService) Update() {

}

func (this *CommentsService) Delete() {

}
