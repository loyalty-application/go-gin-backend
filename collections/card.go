package collections

import (
	"context"
	"time"

	"github.com/loyalty-application/go-gin-backend/config"
	"github.com/loyalty-application/go-gin-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cardCollection *mongo.Collection = config.OpenCollection(config.Client, "cards")

func RetrieveAllCards(skip int64, slice int64) (cards []models.Card, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetSort(bson.D{{Key: "card_id", Value: 1}}).SetLimit(slice).SetSkip(skip)

	cursor, err := cardCollection.Find(ctx, bson.D{}, opts)
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &cards); err != nil {
		panic(err)
	}

	return cards, err
}

func RetrieveSpecificCard(cardId string) (result models.Card, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "card_id", Value: cardId}}
	err = cardCollection.FindOne(ctx, filter).Decode(&result)
	
	return result, err
}

func CreateCard(card models.Card) (result *mongo.InsertOneResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	result, err = cardCollection.InsertOne(ctx, card)

	return result, err
}