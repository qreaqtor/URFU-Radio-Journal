package comments

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

func (cs *CommentsService) Create(comment models.CommentCreate) error {
	unicodeRange, err := cs.determineLanguage(comment.ContentPart)
	if err != nil {
		return err
	}
	if unicodeRange == unicode.Latin {
		comment.Content.Eng = comment.ContentPart
	} else {
		comment.Content.Ru = comment.ContentPart
	}
	_, err = cs.storage.InsertOne(cs.ctx, comment)
	return err
}

func (cs *CommentsService) determineLanguage(str string) (unicodeRange *unicode.RangeTable, err error) {
	ruCount, engCount := 0, 0
	for _, r := range str {
		if unicode.Is(unicode.Cyrillic, r) {
			ruCount++
		} else if unicode.Is(unicode.Latin, r) {
			engCount++
		} else if !unicode.In(r, unicode.Cyrillic, unicode.Latin, unicode.Number, unicode.Space, unicode.Punct) {
			return nil, fmt.Errorf("the string contains an unsupported character: \"%s\"", string(r))
		}
	}
	if ruCount > engCount {
		unicodeRange = unicode.Cyrillic
	} else {
		unicodeRange = unicode.Latin
	}
	return unicodeRange, err
}

func (cs *CommentsService) GetAll(onlyApproved bool, articleIdStr string) (comments []models.CommentRead, err error) {
	filter := bson.M{}
	if onlyApproved {
		filter = bson.M{"isApproved": true}
	}
	if articleIdStr != "" {
		var articleId primitive.ObjectID
		articleId, err = primitive.ObjectIDFromHex(articleIdStr)
		if err != nil {
			return
		}
		filter["articleId"] = articleId
	}
	cur, err := cs.storage.Find(cs.ctx, filter)
	if err != nil {
		return
	}
	err = cur.All(cs.ctx, &comments)
	return
}

func (cs *CommentsService) Update(comment models.CommentUpdate) error {
	filter := bson.M{"_id": comment.Id}
	update := bson.M{"$set": comment}
	res, err := cs.storage.UpdateOne(cs.ctx, filter, update)
	if res.MatchedCount == 0 {
		return errors.New("document not found")
	}
	return err
}

func (cs *CommentsService) Delete(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := cs.storage.DeleteOne(cs.ctx, filter)
	return err
}

func (cs *CommentsService) DeleteManyHandler(filter primitive.M) error {
	_, err := cs.storage.DeleteMany(cs.ctx, filter)
	return err
}

func (cs *CommentsService) Approve(commentApprove models.CommentApprove) error {
	unicodeRange, err := cs.determineLanguage(commentApprove.ContentPart)
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
	res, err := cs.storage.UpdateOne(cs.ctx, filter, update)
	if res.MatchedCount == 0 {
		return errors.New("document not found. Check field content")
	}
	return err
}
