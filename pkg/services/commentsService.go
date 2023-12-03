package services

import (
	"context"
	"errors"
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

func (this *CommentsService) Update(comment models.CommentUpdate) error {
	filter := bson.M{"_id": comment.Id}
	update := bson.M{"$set": comment}
	res, err := this.storage.UpdateOne(this.ctx, filter, update)
	if res.MatchedCount == 0 {
		return errors.New("Document not found.")
	}
	return err
}

func (this *CommentsService) Delete(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	res, err := this.storage.DeleteMany(this.ctx, filter)
	if res.DeletedCount == 0 {
		return errors.New("Documents not found.")
	}
	return err
}

func (this *CommentsService) DeleteHandler(data []primitive.ObjectID) error {
	filter := bson.M{"articleId": bson.M{"$in": data}}
	res, err := this.storage.DeleteMany(this.ctx, filter)
	if res.DeletedCount == 0 {
		return errors.New("Documents not found.")
	}
	return err
}

func (this *CommentsService) Approve(commentApprove models.CommentApprove) error {
	filter := bson.M{"_id": commentApprove.Id}
	update := bson.M{"$set": bson.M{
		"isApproved": true,
		"content": bson.M{
			"$set": bson.M{
				"Eng": commentApprove.ContentEng,
			},
		},
	}}
	res, err := this.storage.UpdateOne(this.ctx, filter, update)
	if res.MatchedCount == 0 {
		return errors.New("Document not found.")
	}
	return err
}
