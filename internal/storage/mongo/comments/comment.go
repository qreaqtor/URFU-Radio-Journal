package commentst

import (
	"context"
	"errors"
	"urfu-radio-journal/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentsStorage struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewCommentStorage(db *mongo.Database, collection string) *CommentsStorage {
	return &CommentsStorage{
		collection: db.Collection(collection),
	}
}

func (cs *CommentsStorage) Create(comment models.CommentCreate) error {
	_, err := cs.collection.InsertOne(cs.ctx, comment)
	return err
}

func (cs *CommentsStorage) GetAll(onlyApproved bool, articleIdStr string) ([]*models.CommentRead, error) {
	filter := bson.M{}
	if onlyApproved {
		filter = bson.M{"isApproved": true}
	}
	if articleIdStr != "" {
		var articleId primitive.ObjectID
		articleId, err := primitive.ObjectIDFromHex(articleIdStr)
		if err != nil {
			return nil, err
		}
		filter["articleId"] = articleId
	}
	cur, err := cs.collection.Find(cs.ctx, filter)
	if err != nil {
		return nil, err
	}
	comments := make([]*models.CommentRead, 0)
	err = cur.All(cs.ctx, &comments)
	return comments, err
}

func (cs *CommentsStorage) Update(comment models.CommentUpdate) error {
	filter := bson.M{"_id": comment.Id}
	update := bson.M{"$set": comment}
	res, err := cs.collection.UpdateOne(cs.ctx, filter, update)
	if res.MatchedCount == 0 {
		return errors.New("document not found")
	}
	return err
}

func (cs *CommentsStorage) Delete(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := cs.collection.DeleteOne(cs.ctx, filter)
	return err
}

func (cs *CommentsStorage) Approve(commentApprove *models.CommentApprove, contentField string) error {
	filter := bson.M{
		"_id":        commentApprove.Id,
		contentField: "",
		"isApproved": false,
	}
	update := bson.M{"$set": bson.M{
		"isApproved": true,
		contentField: commentApprove.ContentPart,
	}}
	res, err := cs.collection.UpdateOne(cs.ctx, filter, update)
	if res.MatchedCount == 0 {
		return errors.New("document not found. Check field content")
	}
	return err
}
