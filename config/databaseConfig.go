package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func InitDatabase() {

	mongoUsername := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	mongoPassword := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	mongoHost := os.Getenv("MONGO_HOST")
	mongoPort := os.Getenv("MONGO_PORT")

	MongoDb := "mongodb://" + mongoUsername + ":" + mongoPassword + "@" + mongoHost + ":" + mongoPort

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	Client = client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)

	return collection
}
