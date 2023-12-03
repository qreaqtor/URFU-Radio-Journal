package services

import (
	"context"
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

func (this *EditionService) Create(edition models.EditionCreate) (id primitive.ObjectID, err error) {
	res, err := this.storage.InsertOne(this.ctx, edition)
	id = res.InsertedID.(primitive.ObjectID)
	return
}

func (this *EditionService) GetAll() (editions []models.EditionRead, err error) {
	filter := bson.M{}
	cur, err := this.storage.Find(this.ctx, filter)
	if err != nil {
		return
	}
	err = cur.All(this.ctx, &editions)
	return
}

func (this *EditionService) Get(id string) (edition models.EditionRead, err error) {
	editionId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": editionId}
	err = this.storage.FindOne(this.ctx, filter).Decode(&edition)
	return
}

func (this *EditionService) Update(newEdition models.EditionUpdate) error {
	filter := bson.M{"_id": newEdition.Id}
	update := bson.M{"$set": newEdition}
	_, err := this.storage.UpdateOne(this.ctx, filter, update)
	return err
}

func (this *EditionService) Delete(editionId primitive.ObjectID) error {
	filter := bson.M{"_id": editionId}
	_, err := this.storage.DeleteOne(this.ctx, filter)
	return err
}
