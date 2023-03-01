package collections

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/loyalty-application/go-gin-backend/config"
	"github.com/loyalty-application/go-gin-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var transactionCollection *mongo.Collection = config.OpenCollection(config.Client, "transaction")

func CreateTransactions(transactions models.TransactionList) (result *mongo.InsertManyResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err = transactionCollection.InsertMany(ctx, transactions.Transactions)

	return result, err
}

func RetrieveAllTransactions(userId string) (transactions models.TransactionList, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": userId}
	cursor, err := transactionCollection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}

	//err = cursor.All(ctx, &transactions.Transactions)
	//if err != nil {
	//panic(err)
	//}

	if err = cursor.All(ctx, &transactions.Transactions); err != nil {
		panic(err)
	}

	for _, result := range transactions.Transactions {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}

	return transactions, err
}
