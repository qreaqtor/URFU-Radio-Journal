package services

import (
	"context"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/pkg/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const layout = "0000-00-00"

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

func (this *EditionService) CreateEdition(edition models.EditionCreate) error {
	edition.Comments = make([]models.Comment, 0)
	_, err := this.storage.InsertOne(this.ctx, edition)
	return err
}

func (this *EditionService) GetAllEditions() (editions []models.EditionRead, err error) {
	filter := bson.M{}
	cur, err := this.storage.Find(this.ctx, filter)
	if err != nil {
		return
	}
	err = cur.All(this.ctx, &editions)
	return
}

func (this *EditionService) UpdateEdition(newEdition models.EditionUpdate) error {
	filter := bson.M{"_id": newEdition.Id}
	update := bson.M{"$set": newEdition}
	_, err := this.storage.UpdateOne(this.ctx, filter, update)
	return err
}

func (this *EditionService) DeleteEdition(id string) error {
	editionId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": editionId}
	_, err = this.storage.DeleteOne(this.ctx, filter)
	return err
}

func (this *EditionService) AddComment(id string, newComment models.Comment) error {
	editionId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": editionId}
	update := bson.M{"$push": bson.M{"comments": newComment}}
	_, err = this.storage.UpdateOne(this.ctx, filter, update)
	return err
}
