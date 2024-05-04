package setupst

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetConnect(user, password, dbName string) *mongo.Database {
	ctx := context.TODO()
	link := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.eiptj34.mongodb.net/", user, password)

	mongoconn := options.Client().ApplyURI(link)
	client, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo: ", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}
	return client.Database(dbName)
}
