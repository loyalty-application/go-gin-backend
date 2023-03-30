package collections

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/loyalty-application/go-gin-backend/config"
	"github.com/loyalty-application/go-gin-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/imdario/mergo"
)

var transactionCollection *mongo.Collection = config.OpenCollection(config.Client, "transactions")
var unprocessedCollection *mongo.Collection = config.OpenCollection(config.Client, "unprocessed")

func RetrieveAllTransactions(skip int64, slice int64) (transaction []models.Transaction, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetSort(bson.D{{Key: "transaction_date", Value: 1}}).SetLimit(slice).SetSkip(skip)

	cursor, err := transactionCollection.Find(ctx, bson.D{}, opts)
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &transaction); err != nil {
		panic(err)
	}

	return transaction, err
}

func RetrieveAllTransactionsForUser(cardIdList []string, skip int64, slice int64) (transaction []models.Transaction, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "card_id", Value: bson.M{"$in" : cardIdList}}}
	opts := options.Find().SetSort(bson.D{{Key: "transaction_date", Value: 1}}).SetLimit(slice).SetSkip(skip)

	cursor, err := transactionCollection.Find(ctx, filter, opts)

	if err != nil {
		panic(err)
	}

	err = cursor.All(ctx, &transaction)

	return transaction, err
}

func CreateTransactions(userId string, transactions models.TransactionList) (result interface{}, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// convert from slice of struct to slice of interface
	t := make([]interface{}, len(transactions.Transactions))
	for i, v := range transactions.Transactions {

		// Placeholders for testing
		v.Points = rand.Float64() * 100
		v.Cashback = rand.Float64() * 100
		v.Miles = rand.Float64() * 100

		t[i] = v
		
	}

	// convert from slice of interface to mongo's bulkWrite model
	models := make([]mongo.WriteModel, 0)
	for _, doc := range t {
		models = append(models, mongo.NewInsertOneModel().SetDocument(doc))
	}
	
	// If an error occurs during the processing of one of the write operations, MongoDB
	// will continue to process remaining write operations in the list.
	bulkWriteOptions := options.BulkWrite().SetOrdered(false)
	// log.Println("Bulk Writing", models)
	result, err = unprocessedCollection.BulkWrite(ctx, models, bulkWriteOptions)
    if err != nil && !mongo.IsDuplicateKeyError(err) {
        panic(err)
    }

	return result, err
}

func UpdateTransaction(transactionId string, transaction models.Transaction) (result *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get original data
	filter := bson.D{{Key: "transaction_id", Value: transactionId}}
	singleResult := transactionCollection.FindOne(ctx, filter)
	if singleResult.Err() != nil {
		log.Println(singleResult.Err().Error())
		return nil, singleResult.Err()
	}

	// Update original data with changed fields in transaction
	initialTransaction := models.Transaction{}
	err = singleResult.Decode(&initialTransaction)
	if err != nil {
		panic(err)
	}
	
	if err = mergo.Merge(&initialTransaction, transaction, mergo.WithOverride); err != nil {
		log.Println(err.Error())
		panic(err)
	}

	// Insert into db
	update := bson.D{{Key: "$set", Value: initialTransaction}}

	result, err = transactionCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		panic(err)
	}

	return result, err
}

func DeleteTransaction(transactionId string) (result *mongo.UpdateResult, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "transaction_id", Value: transactionId}}
	update := bson.D{{Key: "$set", Value: bson.M{"is_deleted": true}}}

	log.Println("Deleting", transactionId)
	result, err = transactionCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		panic(err)
	}

	return result, err
}

func UpdateCardPointsFromTransactions(card models.Card) (result models.Card, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := []bson.M{
		{"$match": bson.M{"card_id": card.CardId}},
		{"$group": bson.M{
			"_id": nil,
			"totalPoints": bson.M{"$sum": "$points"},
			"totalMiles": bson.M{"$sum": "$miles"},
			"totalCashback": bson.M{"$sum": "$cashback"},
		}},
	}

	cursor, err := transactionCollection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Println(err.Error())
		return result, err
	}

	temp := struct {
		TotalPoints   float64 `bson:"totalPoints"`
		TotalMiles    float64 `bson:"totalMiles"`
		TotalCashback float64 `bson:"totalCashback"`
	}{}

	if err = cursor.Decode(&temp); err != nil {
		log.Println(err.Error())
		return result, err
	}

	log.Println("Struct =", temp)

	card.Value += temp.TotalCashback + temp.TotalMiles + temp.TotalPoints

	log.Println("Card =", card)

	return card, err
}