package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ctx context.Context
	db  *mongo.Database
)

func init() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	user, password, dbName := os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"), os.Getenv("DB_NAME")
	ctx = context.TODO()
	link := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.5k1ygzv.mongodb.net/", user, password)

	mongoconn := options.Client().ApplyURI(link)
	client, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo: ", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}
	db = client.Database(dbName)
}

func GetStorage(collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}

func GetContext() *context.Context {
	return &ctx
}
