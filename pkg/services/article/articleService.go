package article

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

func (as *ArticleService) Create(article models.ArticleCreate) (id primitive.ObjectID, err error) {
	res, err := as.storage.InsertOne(as.ctx, article)
	id = res.InsertedID.(primitive.ObjectID)
	return
}

func (as *ArticleService) GetAll(editionIdStr string) (articles []models.ArticleRead, err error) {
	filter := bson.M{}
	if editionIdStr != "" {
		var editionId primitive.ObjectID
		editionId, err = primitive.ObjectIDFromHex(editionIdStr)
		if err != nil {
			return
		}
		filter["editionId"] = editionId
	}
	cur, err := as.storage.Find(as.ctx, filter)
	if err != nil {
		return
	}
	err = cur.All(as.ctx, &articles)
	return
}

func (as *ArticleService) Get(articleIdStr string) (article models.ArticleRead, err error) {
	articleId, err := primitive.ObjectIDFromHex(articleIdStr)
	if err != nil {
		return
	}
	filter := bson.M{"_id": articleId}
	err = as.storage.FindOne(as.ctx, filter).Decode(&article)
	return
}

func (as *ArticleService) GetIdsByEditionId(editionId primitive.ObjectID) (articlesId, filePathsId []primitive.ObjectID, err error) {
	articlesId = make([]primitive.ObjectID, 0)
	filePathsId = make([]primitive.ObjectID, 0)
	filter := bson.M{"editionId": editionId}
	cur, err := as.storage.Find(as.ctx, filter)
	if err != nil {
		return
	}
	var res []models.ArticleRead
	if err = cur.All(as.ctx, &res); err != nil {
		return
	}
	for _, v := range res {
		articlesId = append(articlesId, v.Id)
		filePathsId = append(filePathsId, v.FilePathId)
	}
	return
}

func (as *ArticleService) GetFilePathId(id primitive.ObjectID) (filePathId primitive.ObjectID, err error) {
	filter := bson.M{"_id": id}
	var article models.ArticleRead
	err = as.storage.FindOne(as.ctx, filter).Decode(&article)
	if err != nil {
		return
	}
	filePathId = article.FilePathId
	return
}

func (as *ArticleService) Update(newArticle models.ArticleUpdate) error {
	filter := bson.M{"_id": newArticle.Id}
	update := bson.M{"$set": newArticle}
	res, err := as.storage.UpdateOne(as.ctx, filter, update)
	if res.MatchedCount == 0 {
		return errors.New("document not found")
	}
	return err
}

func (as *ArticleService) Delete(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := as.storage.DeleteOne(as.ctx, filter)
	return err
}

func (as *ArticleService) DeleteManyHandler(editionId primitive.ObjectID) error {
	filter := bson.M{"editionId": editionId}
	_, err := as.storage.DeleteMany(as.ctx, filter)
	return err
}
