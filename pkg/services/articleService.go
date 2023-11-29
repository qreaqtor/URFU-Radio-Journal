package services

import (
	"context"
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

func (this *ArticleService) CreateArticle(article models.ArticleCreate) error {
	_, err := this.storage.InsertOne(this.ctx, article)
	return err
}

func (this *ArticleService) GetAllArticles() (articles []models.ArticleRead, err error) {
	filter := bson.M{}
	cur, err := this.storage.Find(this.ctx, filter)
	if err != nil {
		return
	}
	err = cur.All(this.ctx, &articles)
	return
}

func (this *ArticleService) UpdateArticle(newArticle models.ArticleUpdate) error {
	filter := bson.M{"_id": newArticle.Id}
	update := bson.M{"$set": newArticle}
	_, err := this.storage.UpdateOne(this.ctx, filter, update)
	return err
}

func (this *ArticleService) DeleteArticle(id string) error {
	articleId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": articleId}
	_, err = this.storage.DeleteOne(this.ctx, filter)
	return err
}
