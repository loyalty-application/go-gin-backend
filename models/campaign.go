package models

import "time"

type CampaignList struct {
	Campaigns []Campaign `json:"campaigns" bson:",inline"`
}

type Campaign struct {
	// TODO: should have on merchantId
	UserId      string    `json:"-" bson:"user_id" example:"6400a..."`
	CampaignId  string    `json:"campaign_id" bson:"campaign_id" example:"cmp00001"`
	Merchant    string    `json:"merchant" bson:"merchant" example:"7-11"`
	CardType    string    `json:"card_type" bson:"card_type" example:"super_miles_card"`
	Description string    `json:"description" bson:"description"`
	StartDate   time.Time `json:"start_date" bson:"start_date" example:"2023-03-02T13:10:23Z"`
	EndDate     time.Time `json:"end_date" bson:"end_date" example:"2023-03-03T13:10:23Z"`
}
