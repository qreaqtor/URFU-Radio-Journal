package commentst

import (
	"context"
	"errors"
	"fmt"
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

func (cs *CommentsStorage) InsertOne(comment *models.CommentCreate) (string, error) {
	res, err := cs.collection.InsertOne(cs.ctx, comment)
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("bad inserted id")
	}
	return id.Hex(), err
}

func (cs *CommentsStorage) GetAll(onlyApproved bool, commentIdStr string) ([]*models.CommentRead, error) {
	articleId, err := primitive.ObjectIDFromHex(commentIdStr)
	if err != nil {
		return nil, err
	}
	filter := bson.M{
		"articleId": articleId,
	}
	if onlyApproved {
		filter = bson.M{"isApproved": true}
	}
	cur, err := cs.collection.Find(cs.ctx, filter)
	if err != nil {
		return nil, err
	}
	comments := make([]*models.CommentRead, 0)
	err = cur.All(cs.ctx, &comments)
	return comments, err
}

func (cs *CommentsStorage) UpdateOne(comment *models.CommentUpdate) error {
	filter := bson.M{"_id": comment.Id}
	update := bson.M{"$set": comment}
	res, err := cs.collection.UpdateOne(cs.ctx, filter, update)
	if res.MatchedCount == 0 {
		return errors.New("document not found")
	}
	return err
}

func (cs *CommentsStorage) Delete(idStr string) error {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	res, err := cs.collection.DeleteOne(cs.ctx, filter)
	if res.DeletedCount != 1 {
		return fmt.Errorf("deleted count was %d, baut want 1", res.DeletedCount)
	}
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
