package collections

import (
	"context"
	"time"

	"github.com/loyalty-application/go-gin-backend/config"
	"github.com/loyalty-application/go-gin-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.OpenCollection(config.Client, "user")

func RetrieveUser(user models.User) (dbUser models.User, err error) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	err = userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)

	return dbUser, err

}

func CountUserPhone(phone string) (count int64, err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	count, err = userCollection.CountDocuments(ctx, bson.M{"phone": phone})

	return count, err
}

func CountUserEmail(email string) (count int64, err error) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	count, err = userCollection.CountDocuments(ctx, bson.M{"email": email})

	return count, err
}

func CreateUser(user models.User) (insertionNo *mongo.InsertOneResult, err error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	insertionNo, err = userCollection.InsertOne(ctx, user)

	return insertionNo, err

}
