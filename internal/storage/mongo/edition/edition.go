package editionst

import (
	"context"
	"errors"
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

func (es *EditionStorage) Create(edition models.EditionCreate) (id primitive.ObjectID, err error) {
	res, err := es.collection.InsertOne(es.ctx, edition)
	id = res.InsertedID.(primitive.ObjectID)
	return
}

func (es *EditionStorage) GetAll() (editions []models.EditionRead, err error) {
	filter := bson.M{}
	cur, err := es.collection.Find(es.ctx, filter)
	if err != nil {
		return
	}
	err = cur.All(es.ctx, &editions)
	return
}

func (es *EditionStorage) Get(id string) (edition models.EditionRead, err error) {
	editionId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": editionId}
	err = es.collection.FindOne(es.ctx, filter).Decode(&edition)
	return
}

func (es *EditionStorage) Update(newEdition models.EditionUpdate) error {
	filter := bson.M{"_id": newEdition.Id}
	update := bson.M{"$set": newEdition}
	res, err := es.collection.UpdateOne(es.ctx, filter, update)
	if res.MatchedCount == 0 {
		return errors.New("document not found")
	}
	return err
}

func (es *EditionStorage) Delete(editionId primitive.ObjectID) error {
	filter := bson.M{"_id": editionId}
	_, err := es.collection.DeleteOne(es.ctx, filter)
	return err
}
