package council

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

func (cs *CouncilService) Create(member models.CouncilMemberCreate) error {
	_, err := cs.storage.InsertOne(cs.ctx, member)
	return err
}

func (cs *CouncilService) Update(memberIdStr string, memberUpdate models.CouncilMemberUpdate) error {
	id, err := primitive.ObjectIDFromHex(memberIdStr)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": memberUpdate}
	res, err := cs.storage.UpdateOne(cs.ctx, filter, update)
	if res.MatchedCount == 0 {
		return errors.New("document not found")
	}
	return err
}

func (cs *CouncilService) Delete(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := cs.storage.DeleteOne(cs.ctx, filter)
	return err
}

func (cs *CouncilService) GetImagePathId(id primitive.ObjectID) (imagePathId primitive.ObjectID, err error) {
	var member models.CouncilMemberRead
	filter := bson.M{"_id": id}
	err = cs.storage.FindOne(cs.ctx, filter).Decode(&member)
	imagePathId = member.ImagePathId
	return
}

func (cs *CouncilService) GetAll() (members []models.CouncilMemberRead, err error) {
	filter := bson.M{}
	cur, err := cs.storage.Find(cs.ctx, filter)
	if err != nil {
		return
	}
	err = cur.All(cs.ctx, &members)
	return
}

func (cs *CouncilService) Get(memberIdStr string) (member models.CouncilMemberRead, err error) {
	memberId, err := primitive.ObjectIDFromHex(memberIdStr)
	if err != nil {
		return
	}
	filter := bson.M{"_id": memberId}
	err = cs.storage.FindOne(cs.ctx, filter).Decode(&member)
	return
}
