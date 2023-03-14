package collections

import (
	"context"
	"github.com/loyalty-application/go-gin-backend/config"
	"github.com/loyalty-application/go-gin-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"errors"
)

var campaignCollection *mongo.Collection = config.OpenCollection(config.Client, "campaigns")


func RetrieveCampaign(campaignID string) (campaign models.Campaign, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "campaign_id", Value: campaignID}}

	err = campaignCollection.FindOne(ctx, filter).Decode(&campaign)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Campaign{}, errors.New("campaign not found")
		}
		return models.Campaign{}, err
	}

	return campaign, nil
}


func RetrieveAllCampaigns(campaignId string, skip int64, slice int64) (campaign []models.Campaign, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "campaign_id", Value: campaignId}}
	opts := options.Find().SetSort(bson.D{{"start_date", 1}}).SetLimit(slice).SetSkip(skip) // unsure if can sort by time.Time class

	cursor, err := campaignCollection.Find(ctx, filter, opts)

	if err != nil {
		panic(err)
	}

	err = cursor.All(ctx, &campaign)

	return campaign, err
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
