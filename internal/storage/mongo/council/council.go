package councilst

import (
	"context"
	"errors"
	"urfu-radio-journal/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CouncilStorage struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewCouncilStorage(db *mongo.Database, collection string) *CouncilStorage {
	return &CouncilStorage{
		ctx:        context.Background(),
		collection: db.Collection(collection),
	}
}

func (cs *CouncilStorage) Create(member *models.CouncilMemberCreate) error {
	_, err := cs.collection.InsertOne(cs.ctx, member)
	return err
}

func (cs *CouncilStorage) Update(memberIdStr string, memberUpdate models.CouncilMemberUpdate) error {
	id, err := primitive.ObjectIDFromHex(memberIdStr)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": memberUpdate}
	res, err := cs.collection.UpdateOne(cs.ctx, filter, update)
	if res.MatchedCount == 0 {
		return errors.New("document not found")
	}
	return err
}

func (cs *CouncilStorage) Delete(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := cs.collection.DeleteOne(cs.ctx, filter)
	return err
}

func (cs *CouncilStorage) GetImagePathId(id primitive.ObjectID) (imagePathId primitive.ObjectID, err error) {
	var member models.CouncilMemberRead
	filter := bson.M{"_id": id}
	err = cs.collection.FindOne(cs.ctx, filter).Decode(&member)
	imagePathId = member.ImagePathId
	return
}

func (cs *CouncilStorage) GetAll() (members []models.CouncilMemberRead, err error) {
	filter := bson.M{}
	cur, err := cs.collection.Find(cs.ctx, filter)
	if err != nil {
		return
	}
	err = cur.All(cs.ctx, &members)
	return
}

func (cs *CouncilStorage) Get(memberIdStr string) (member models.CouncilMemberRead, err error) {
	memberId, err := primitive.ObjectIDFromHex(memberIdStr)
	if err != nil {
		return
	}
	filter := bson.M{"_id": memberId}
	err = cs.collection.FindOne(cs.ctx, filter).Decode(&member)
	return
}
