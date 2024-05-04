package articlest

import (
	"context"
	"fmt"
	"urfu-radio-journal/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ArticleStorage struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewArticleStorage(db *mongo.Database, collection string) *ArticleStorage {
	return &ArticleStorage{
		ctx:        context.Background(),
		collection: db.Collection(collection),
	}
}

func (a *ArticleStorage) InsertOne(article *models.ArticleCreate) (string, error) {
	res, err := a.collection.InsertOne(a.ctx, article)
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("wrong id of inserted doc")
	}
	return id.Hex(), err
}

func (as *ArticleStorage) Find(editionIdStr string) ([]*models.ArticleRead, error) {
	editionId, err := primitive.ObjectIDFromHex(editionIdStr)
	if err != nil {
		return nil, err
	}
	filter := bson.M{
		"editionId": editionId,
	}
	cur, err := as.collection.Find(as.ctx, filter)
	if err != nil {
		return nil, err
	}
	articles := make([]*models.ArticleRead, 0)
	err = cur.All(as.ctx, &articles)
	return articles, err
}

func (as *ArticleStorage) FindOne(articleIdStr string) (*models.ArticleRead, error) {
	articleId, err := primitive.ObjectIDFromHex(articleIdStr)
	if err != nil {
		return nil, err
	}
	article := &models.ArticleRead{}
	filter := bson.M{"_id": articleId}
	err = as.collection.FindOne(as.ctx, filter).Decode(article)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (as *ArticleStorage) Delete(IdStr string) error {
	id, err := primitive.ObjectIDFromHex(IdStr)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	res, err := as.collection.DeleteOne(as.ctx, filter)
	if res.DeletedCount != 1 {
		return fmt.Errorf("wrong deleted count, want 1, but was %d", res.DeletedCount)
	}
	return err
}

func (as *ArticleStorage) GetFilePathId(id primitive.ObjectID) (string, error) {
	filter := bson.M{"_id": id}
	var article models.ArticleRead
	err := as.collection.FindOne(as.ctx, filter).Decode(&article)
	if err != nil {
		return "", nil
	}
	return article.FilePathId.Hex(), nil
}

func (as *ArticleStorage) UpdateOne(newArticle *models.ArticleUpdate) error {
	filter := bson.M{"_id": newArticle.Id}
	update := bson.M{"$set": newArticle}
	res, err := as.collection.UpdateOne(as.ctx, filter, update)
	if res.MatchedCount == 0 {
		return fmt.Errorf("document not found")
	}
	return err
}
