package services

import (
	"context"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/pkg/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (this *CommentsService) Create(comment models.CommentCreate) error {
	_, err := this.storage.InsertOne(this.ctx, comment)
	return err
}

func (this *CommentsService) GetAll(onlyApproved bool) (comments []models.CommentRead, err error) {
	filter := bson.M{}
	if onlyApproved {
		filter = bson.M{"isApproved": true}
	}
	cur, err := this.storage.Find(this.ctx, filter)
	if err != nil {
		return
	}
	err = cur.All(this.ctx, &comments)
	return
}

func (this *CommentsService) Update(data []models.CommentUpdate) error {
	var filter, update bson.M
	for _, comment := range data {
		filter = bson.M{"_id": comment.Id}
		update = bson.M{"$set": comment}
		if _, err := this.storage.UpdateOne(this.ctx, filter, update); err != nil {
			return err
		}
	}
	return nil
}

func (this *CommentsService) Delete(data []primitive.ObjectID) error {
	filter := bson.M{"_id": bson.M{"$in": data}}
	_, err := this.storage.DeleteMany(this.ctx, filter)
	return err
}

func (this *CommentsService) Approve(data []primitive.ObjectID) error {
	filter := bson.M{"_id": bson.M{"$in": data}}
	update := bson.M{"$set": bson.M{"isApproved": true}}
	_, err := this.storage.UpdateMany(this.ctx, filter, update)
	return err
}
