package edition

import (
	"context"
	"errors"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/pkg/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EditionService struct {
	ctx     context.Context
	storage *mongo.Collection
}

func NewEditionService() *EditionService {
	return &EditionService{
		ctx:     *db.GetContext(),
		storage: db.GetStorage("editions"),
	}
}

func (es *EditionService) Create(edition models.EditionCreate) (id primitive.ObjectID, err error) {
	res, err := es.storage.InsertOne(es.ctx, edition)
	id = res.InsertedID.(primitive.ObjectID)
	return
}

func (es *EditionService) GetAll() (editions []models.EditionRead, err error) {
	filter := bson.M{}
	cur, err := es.storage.Find(es.ctx, filter)
	if err != nil {
		return
	}
	err = cur.All(es.ctx, &editions)
	return
}

func (es *EditionService) Get(id string) (edition models.EditionRead, err error) {
	editionId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": editionId}
	err = es.storage.FindOne(es.ctx, filter).Decode(&edition)
	return
}

func (es *EditionService) Update(newEdition models.EditionUpdate) error {
	filter := bson.M{"_id": newEdition.Id}
	update := bson.M{"$set": newEdition}
	res, err := es.storage.UpdateOne(es.ctx, filter, update)
	if res.MatchedCount == 0 {
		return errors.New("document not found")
	}
	return err
}

func (es *EditionService) Delete(editionId primitive.ObjectID) error {
	filter := bson.M{"_id": editionId}
	_, err := es.storage.DeleteOne(es.ctx, filter)
	return err
}
