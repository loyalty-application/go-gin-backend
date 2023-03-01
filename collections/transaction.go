package collections

import (
	"context"
	"fmt"
	"time"

	"github.com/loyalty-application/go-gin-backend/config"
	"github.com/loyalty-application/go-gin-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var c *mongo.Collection = config.OpenCollection(config.Client, "transaction")

func CreateTransactions(id string, transactions models.TransactionList) (success bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// insert into database
	result, err := c.InsertMany(ctx, transactions.Transactions)

	fmt.Println(result)

	if err != nil {
		return false, err
	}

	return true, err
}
