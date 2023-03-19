package collections

import (
	"context"
	"log"
	"time"

	"github.com/loyalty-application/go-gin-backend/config"
	"github.com/loyalty-application/go-gin-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var transactionCollection *mongo.Collection = config.OpenCollection(config.Client, "transactions")

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

func RetrieveAllTransactionsForUser(userId string, skip int64, slice int64) (transaction []models.Transaction, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "user_id", Value: userId}}
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
		v.UserId = userId
		log.Println(v)
		t[i] = v
		// // To test 100k records in 1 transaction
		// for j, count := 0, 0; j < 100000; j++ {
		// 	v.UserId = userId
		// 	v.TransactionId = strconv.Itoa(count)
		// 	count++
		// 	t[j] = v
		// }
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
	result, err = transactionCollection.BulkWrite(ctx, models, bulkWriteOptions)
    if err != nil {
        log.Println(err.Error())
		return result, err
    }

	// Please don't delete the code below in case we need to reuse transactions in the future - Gabriel


	// // Setting write permissions
	// wc := writeconcern.New(writeconcern.WMajority())
	// txnOpts := options.Transaction().SetWriteConcern(wc)

	// // Start new session
	// session, err := config.Client.StartSession()
	// if err != nil {
	// 	return nil, err
	// }
	// defer session.EndSession(context.Background())

	// // Start transaction
	// if err = session.StartTransaction(txnOpts); err != nil {
	// 	return nil, err
	// }
	// log.Println("Transaction Start without errors")

	// // Insert documents in the current session
	// log.Println("Before Insert")
	// result, err = transactionCollection.InsertMany(mongo.NewSessionContext(context.Background(),session), t)
	// log.Println("After Insert")
	// defer cancel()

	// if err != nil {
	// 	log.Println("Insert Many Error = ", err.Error())
	// 	// Abort session if got error
	// 	session.AbortTransaction(context.Background())
	// 	// log.Println("Aborted Transaction")
	// 	return result, err
	// }

	// // Commit documents if no error
	// err = session.CommitTransaction(context.Background())

	return result, err
}
