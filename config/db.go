package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBinstance()

func DBinstance() (client *mongo.Client) {

	user := os.Getenv("MONGO_USERNAME")
	pass := os.Getenv("MONGO_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")

	uri := "mongodb://" + user + ":" + pass + "@" + host + ":" + port

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect to mongodb
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// initialise indexes
	InitIndexes(client)

	return client
}

func InitIndexes(client *mongo.Client) {

	// transactions_transactions_-1 index
	transactionCollection := OpenCollection(client, "transactions")

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"transaction_id", -1}},
		Options: options.Index().SetUnique(true),
	}
	indexCreated, err := transactionCollection.Indexes().CreateOne(context.Background(), indexModel)

	// campaigns_campaigns_-1 index
	campaignCollection := OpenCollection(client, "campaigns")

	campaignIndexModel := mongo.IndexModel{
		Keys:    bson.D{{"campaign_id", -1}},
		Options: options.Index().SetUnique(true),
	}
	campaignIndexCreated, err := campaignCollection.Indexes().CreateOne(context.Background(), campaignIndexModel)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created Index %s\n", indexCreated)
	fmt.Printf("Created Campaign Index %s\n", campaignIndexCreated)
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("loyalty").Collection(collectionName)

	return collection
}
