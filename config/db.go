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

	conn := fmt.Sprintf("mongodb://%s:%s@%s:%s/?replicaSet=replica-set", user, pass, host, port)

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(conn).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// connect to mongodb
	client, err := mongo.Connect(ctx, clientOptions)
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

	transactionIndexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "transaction_id", Value: -1}},
		Options: options.Index().SetUnique(true),
	}
	transactionIndexCreated, err := transactionCollection.Indexes().CreateOne(context.Background(), transactionIndexModel)
	if err != nil {
		log.Fatal(err)
	}

	// campaigns_campaigns_-1 index
	campaignCollection := OpenCollection(client, "campaigns")

	campaignIndexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "campaign_id", Value: -1}},
		Options: options.Index().SetUnique(true),
	}

	campaignIndexCreated, err := campaignCollection.Indexes().CreateOne(context.Background(), campaignIndexModel)
	if err != nil {
		log.Fatal(err)
	}

	// cards_cards_-1 index
	cardCollection := OpenCollection(client, "cards")

	cardIndexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "card_id", Value: -1}},
		Options: options.Index().SetUnique(true),
	}

	cardIndexCreated, err := cardCollection.Indexes().CreateOne(context.Background(), cardIndexModel)
	if err != nil {
		log.Fatal(err)
	}

	// user_users_-1 index
	userCollection := OpenCollection(client, "users")

	userIndexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "user_id", Value: -1}},
		Options: options.Index().SetUnique(true),
	}

	userIndexCreated, err := userCollection.Indexes().CreateOne(context.Background(), userIndexModel)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created Transaction Index %s\n", transactionIndexCreated)
	fmt.Printf("Created Campaign Index %s\n", campaignIndexCreated)
	fmt.Printf("Created Card Index %s\n", cardIndexCreated)
	fmt.Printf("Created User Index %s\n", userIndexCreated)
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("loyalty").Collection(collectionName)

	return collection
}
