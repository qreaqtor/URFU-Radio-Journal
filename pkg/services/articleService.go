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

type ArticleService struct {
	ctx     context.Context
	storage *mongo.Collection
}

func NewArticleService() *ArticleService {
	return &ArticleService{
		ctx:     *db.GetContext(),
		storage: db.GetStorage("articles"),
	}
}

func (this *ArticleService) Create(article models.ArticleCreate) (id primitive.ObjectID, err error) {
	res, err := this.storage.InsertOne(this.ctx, article)
	id = res.InsertedID.(primitive.ObjectID)
	return
}

func (this *ArticleService) GetAll() (articles []models.ArticleRead, err error) {
	filter := bson.M{}
	cur, err := this.storage.Find(this.ctx, filter)
	if err != nil {
		return
	}
	err = cur.All(this.ctx, &articles)
	return
}

func (this *ArticleService) GetArticlesFilePathsByEditionId(editionId primitive.ObjectID) (filePathsId []primitive.ObjectID, err error) {
	filter := bson.M{"editionid": editionId}
	cur, err := this.storage.Find(this.ctx, filter)
	if err != nil {
		return
	}
	var res []models.ArticleRead
	if err = cur.All(this.ctx, res); err != nil {
		return
	}
	for _, v := range res {
		filePathsId = append(filePathsId, v.FilePathId)
	}
	return
}

func (this *ArticleService) Update(newArticle models.ArticleUpdate) error {
	filter := bson.M{"_id": newArticle.Id}
	update := bson.M{"$set": newArticle}
	res, err := this.storage.UpdateOne(this.ctx, filter, update)
	if res.MatchedCount == 0 {
		return errors.New("Document not found.")
	}
	return err
}

func (this *ArticleService) Delete(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	res, err := this.storage.DeleteOne(this.ctx, filter)
	if res.DeletedCount == 0 {
		return errors.New("Document not found.")
	}
	return err
}

func (this *ArticleService) DeleteMany(editionId primitive.ObjectID) error {
	filter := bson.M{"editionId": editionId}
	res, err := this.storage.DeleteMany(this.ctx, filter)
	if res.DeletedCount == 0 {
		return errors.New("Documents not found.")
	}
	return err
}
