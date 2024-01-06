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

type CouncilService struct {
	ctx     context.Context
	storage *mongo.Collection
}

func NewCouncilService() *CouncilService {
	return &CouncilService{
		ctx:     *db.GetContext(),
		storage: db.GetStorage("council"),
	}
}

func (this *CouncilService) Create(member models.CouncilMemberCreate) error {
	_, err := this.storage.InsertOne(this.ctx, member)
	return err
}

func (this *CouncilService) Update(memberIdStr string, memberUpdate models.CouncilMemberUpdate) error {
	id, err := primitive.ObjectIDFromHex(memberIdStr)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": memberUpdate}
	res, err := this.storage.UpdateOne(this.ctx, filter, update)
	if res.MatchedCount == 0 {
		return errors.New("Document not found.")
	}
	return err
}

func (this *CouncilService) Delete(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := this.storage.DeleteOne(this.ctx, filter)
	return err
}

func (this *CouncilService) GetImagePathId(id primitive.ObjectID) (imagePathId primitive.ObjectID, err error) {
	var member models.CouncilMemberRead
	filter := bson.M{"_id": id}
	err = this.storage.FindOne(this.ctx, filter).Decode(&member)
	imagePathId = member.ImagePathId
	return
}

func (this *CouncilService) GetAll() (members []models.CouncilMemberRead, err error) {
	filter := bson.M{}
	cur, err := this.storage.Find(this.ctx, filter)
	if err != nil {
		return
	}
	err = cur.All(this.ctx, &members)
	return
}

func (this *CouncilService) Get(memberIdStr string) (member models.CouncilMemberRead, err error) {
	memberId, err := primitive.ObjectIDFromHex(memberIdStr)
	if err != nil {
		return
	}
	filter := bson.M{"_id": memberId}
	err = this.storage.FindOne(this.ctx, filter).Decode(&member)
	return
}
