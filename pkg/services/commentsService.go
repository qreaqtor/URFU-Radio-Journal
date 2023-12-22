package services

import (
	"context"
	"errors"
	"fmt"
	"unicode"
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
	unicodeRange, err := this.determineLanguage(comment.ContentPart)
	if err != nil {
		return err
	}
	if unicodeRange == unicode.Latin {
		comment.Content.Eng = comment.ContentPart
	} else {
		comment.Content.Ru = comment.ContentPart
	}
	_, err = this.storage.InsertOne(this.ctx, comment)
	return err
}

func (this *CommentsService) determineLanguage(str string) (unicodeRange *unicode.RangeTable, err error) {
	ruCount, engCount := 0, 0
	for _, r := range str {
		if unicode.Is(unicode.Cyrillic, r) {
			ruCount++
		} else if unicode.Is(unicode.Latin, r) {
			engCount++
		} else if !unicode.In(r, unicode.Cyrillic, unicode.Latin, unicode.Number, unicode.Space, unicode.Punct) {
			return nil, fmt.Errorf("The string contains an unsupported character: \"%s\".", string(r))
		}
	}
	if ruCount > engCount {
		unicodeRange = unicode.Cyrillic
	} else {
		unicodeRange = unicode.Latin
	}
	return unicodeRange, err
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
	_, err := this.storage.DeleteOne(this.ctx, filter)
	return err
}

func (this *CommentsService) DeleteManyHandler(filter primitive.M) error {
	_, err := this.storage.DeleteMany(this.ctx, filter)
	return err
}

func (this *CommentsService) Approve(commentApprove models.CommentApprove) error {
	unicodeRange, err := this.determineLanguage(commentApprove.ContentPart)
	if err != nil {
		return err
	}
	contentField := "content.Ru"
	if unicodeRange == unicode.Latin {
		contentField = "content.Eng"
	}
	filter := bson.M{
		"_id":        commentApprove.Id,
		contentField: "",
		"isApproved": false,
	}
	update := bson.M{"$set": bson.M{
		"isApproved": true,
		contentField: commentApprove.ContentPart,
	}}
	res, err := this.storage.UpdateOne(this.ctx, filter, update)
	if res.MatchedCount == 0 {
		return errors.New("Document not found. Check field content.")
	}
	return err
}
