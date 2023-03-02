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

var transactionCollection *mongo.Collection = config.OpenCollection(config.Client, "transactions")

func RetrieveAllTransactions(userId string, skip int64, slice int64) (transaction []models.Transaction, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "user_id", Value: userId}}
	opts := options.Find().SetSort(bson.D{{"transaction_date", 1}}).SetLimit(slice).SetSkip(skip)

	cursor, err := transactionCollection.Find(ctx, filter, opts)

	if err != nil {
		panic(err)
	}

	err = cursor.All(ctx, &transaction)

	return transaction, err
}

func CreateTransactions(userId string, transactions models.TransactionList) (result *mongo.InsertManyResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// convert from slice of struct to slice of interface
	t := make([]interface{}, len(transactions.Transactions))
	for i, v := range transactions.Transactions {
		v.UserId = userId
		t[i] = v
	}

	result, err = transactionCollection.InsertMany(ctx, t)
	return result, err

}
