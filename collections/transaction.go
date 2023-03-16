package collections

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/loyalty-application/go-gin-backend/config"
	"github.com/loyalty-application/go-gin-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

var transactionCollection *mongo.Collection = config.OpenCollection(config.Client, "transactions")

func RetrieveAllTransactions(userId string, skip int64, slice int64) (transaction []models.Transaction, err error) {
	log.Println("Testing Retrieve")
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

func CreateTransactions(userId string, transactions models.TransactionList) (result interface{}, err error) {
	fmt.Println("Tesadhjskfdafj")
	log.Println("Tshjsfsh")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// convert from slice of struct to slice of interface
	t := make([]interface{}, len(transactions.Transactions))
	for i, v := range transactions.Transactions {
		v.UserId = userId
		t[i] = v
	}

	// Setting write permissions
	wc := writeconcern.New(writeconcern.WMajority())
	txnOpts := options.Transaction().SetWriteConcern(wc)

	// Start new session
	session, err := config.Client.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())

	// Start transaction
	if err = session.StartTransaction(txnOpts); err != nil {
		return nil, err
	}

	// Insert documents
	result, err = transactionCollection.InsertMany(ctx, t)
	if err != nil {
		log.Fatal(err)
		// Abort session if got error
		session.AbortTransaction(context.Background())
		return result, err
	}

	// Commit documents if no error
	err = session.CommitTransaction(context.Background())

	return result, err

}
