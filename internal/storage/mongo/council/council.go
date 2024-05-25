package councilst

import (
	"context"
	"errors"
	"fmt"
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

func (cs *CouncilStorage) InsertOne(member *models.CouncilMemberCreate) (string, error) {
	res, err := cs.collection.InsertOne(cs.ctx, member)
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("bad value for id")
	}
	return id.Hex(), err
}

func (cs *CouncilStorage) UpdateOne(memberIdStr string, memberUpdate *models.CouncilMemberUpdate) error {
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

func (cs *CouncilStorage) Delete(idStr string) error {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	res, err := cs.collection.DeleteOne(cs.ctx, filter)
	if res.DeletedCount != 1 {
		return fmt.Errorf("deleted count was %d, but want 1", res.DeletedCount)
	}
	return err
}

func (cs *CouncilStorage) GetImagePathId(idStr string) (string, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return "", err
	}
	var member models.CouncilMemberRead
	filter := bson.M{"_id": id}
	err = cs.collection.FindOne(cs.ctx, filter).Decode(&member)
	return member.ImagePathId, err
}

func (cs *CouncilStorage) GetAll() ([]*models.CouncilMemberRead, error) {
	filter := bson.M{}
	cur, err := cs.collection.Find(cs.ctx, filter)
	if err != nil {
		return nil, err
	}
	members := make([]*models.CouncilMemberRead, 0)
	err = cur.All(cs.ctx, &members)
	return members, err
}

func (cs *CouncilStorage) FindOne(memberIdStr string) (*models.CouncilMemberRead, error) {
	memberId, err := primitive.ObjectIDFromHex(memberIdStr)
	if err != nil {
		return nil, err
	}
	member := &models.CouncilMemberRead{}
	filter := bson.M{"_id": memberId}
	err = cs.collection.FindOne(cs.ctx, filter).Decode(member)
	return member, err
}
