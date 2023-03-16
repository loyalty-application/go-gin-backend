package collections

import (
	"context"
	"errors"
	"github.com/loyalty-application/go-gin-backend/config"
	"github.com/loyalty-application/go-gin-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var campaignCollection *mongo.Collection = config.OpenCollection(config.Client, "campaigns")

func RetrieveCampaign(campaignId string) (campaign models.Campaign, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "campaign_id", Value: campaignId}}
	err = campaignCollection.FindOne(ctx, filter).Decode(&campaign)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Campaign{}, errors.New("campaign not found")
		}
		return models.Campaign{}, err
	}

	return campaign, nil
}

func RetrieveAllCampaigns() (campaigns []models.Campaign, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	options := options.Find().SetSort(bson.M{"start_date": 1})
	cursor, err := campaignCollection.Find(ctx, bson.M{}, options)
	if err != nil {
		panic(err)
	}

	defer cursor.Close(ctx)
	err = cursor.All(ctx, &campaigns)

	return campaigns, err
}

func CreateCampaign(userId string, campaigns models.CampaignList) (result *mongo.InsertManyResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// convert from slice of struct to slice of interface
	t := make([]interface{}, len(campaigns.Campaigns))
	for i, v := range campaigns.Campaigns {
		v.UserId = userId
		t[i] = v
	}

	result, err = campaignCollection.InsertMany(ctx, t)
	return result, err
}

func UpdateCampaign(campaignId string, updateData models.Campaign) (result *mongo.UpdateResult, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "campaign_id", Value: campaignId}}
	update := bson.D{{Key: "$set", Value: updateData}}

	result, err = campaignCollection.UpdateOne(ctx, filter, update)
	return result, err
}

func DeleteCampaign(campaignId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "campaign_id", Value: campaignId}}

	_, err := campaignCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
