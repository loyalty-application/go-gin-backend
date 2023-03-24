package collections

import (
	"context"
	"log"
	"time"

	"github.com/imdario/mergo"
	"github.com/loyalty-application/go-gin-backend/config"
	"github.com/loyalty-application/go-gin-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection = config.OpenCollection(config.Client, "users")

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

func RetrieveAllUsers(skip int64, slice int64) (result []models.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetSort(bson.D{{Key: "_id", Value: 1}}).SetLimit(slice).SetSkip(skip)

	cursor, err := userCollection.Find(ctx, bson.D{}, opts)
	if err != nil {
		panic(err)
	}

	if err := cursor.All(ctx, &result); err != nil {
		panic(err)
	}

	output := make([]models.User, len(result))
	// Calculate total points / miles / cashback
	for _, user := range result {
		cardIdList := user.Card
		cardList, _ := RetrieveListOfCards(cardIdList)

		for _, card := range cardList {
			switch card.ValueType {
			case "Miles":
				user.Miles += card.Value
			case "Points":
				user.Points += card.Value
			case "Cashback":
				user.Cashback += card.Value
			default:
				log.Println("Invalid Card ValueType")
			}
		}
		output = append(output, user)
	}

	return output, err
}

func RetrieveSpecificUser(email string) (result models.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "email", Value: email}}

	err = userCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}

	// Calculate total points / miles / cashback
	cardIdList := result.Card
	cardList, err := RetrieveListOfCards(cardIdList)

	for _, card := range cardList {
		switch card.ValueType {
		case "Miles":
			result.Miles += card.Value
		case "Points":
			result.Points += card.Value
		case "Cashback":
			result.Cashback += card.Value
		default:
			log.Println("Invalid Card ValueType")
		}
	}

	return result, err
}

func UpdateUser(email string, user models.User) (result *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Println("Before Retrieve")
	// Get original data
	initialUser, err := RetrieveSpecificUser(email)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	log.Println("Before Merge")
	// Update original data with changed fields in transaction
	if err := mergo.Merge(&initialUser, user, mergo.WithOverride); err != nil {
		return nil, err
	}
	log.Println("After Merge")
	// Insert into db
	filter := bson.D{{Key: "email", Value: email}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "first_name", Value: initialUser.FirstName},
		{Key: "last_name", Value: initialUser.LastName},
		{Key: "password", Value: initialUser.Password},
		{Key: "email", Value: initialUser.Email},
		{Key: "cards", Value: initialUser.Card}}}}

	log.Println(initialUser)

	result, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		panic(err)
	}

	return result, err
}
