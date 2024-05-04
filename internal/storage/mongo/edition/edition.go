package editionst

import (
	"context"
	"errors"
	"fmt"
	"urfu-radio-journal/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EditionStorage struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewEditionService(db *mongo.Database, collection string) *EditionStorage {
	return &EditionStorage{
		ctx:        context.Background(),
		collection: db.Collection(collection),
	}
}

func (es *EditionStorage) InsertOne(edition *models.EditionCreate) (string, error) {
	res, err := es.collection.InsertOne(es.ctx, edition)
	if err != nil {
		return "", nil
	}
	id := res.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func (es *EditionStorage) GetAll() ([]*models.EditionRead, error) {
	filter := bson.M{}
	cur, err := es.collection.Find(es.ctx, filter)
	if err != nil {
		return nil, err
	}
	editions := make([]*models.EditionRead, 0)
	err = cur.All(es.ctx, &editions)
	return editions, err
}

func (es *EditionStorage) FindOne(id string) (*models.EditionRead, error) {
	editionId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": editionId}
	edition := &models.EditionRead{}
	err = es.collection.FindOne(es.ctx, filter).Decode(&edition)
	return edition, err
}

func (es *EditionStorage) UpdateOne(newEdition *models.EditionUpdate) error {
	filter := bson.M{"_id": newEdition.Id}
	update := bson.M{"$set": newEdition}
	res, err := es.collection.UpdateOne(es.ctx, filter, update)
	if res.MatchedCount == 0 {
		return errors.New("document not found")
	}
	return err
}

func (es *EditionStorage) Delete(editionIdStr string) error {
	editionId, err := primitive.ObjectIDFromHex(editionIdStr)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": editionId}
	res, err := es.collection.DeleteOne(es.ctx, filter)
	if res.DeletedCount != 1 {
		return fmt.Errorf("deleted count was %d, but want 1", res.DeletedCount)
	}
	return err
}
