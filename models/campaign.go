package models

import "time"

type Campaign struct {
	Id          int       `json:"id" bson:"id" example:"1"`
	Merchant    string    `json:"merchant" bson:"merchant" example:"7-11"`
	CardType    string    `json:"card_type" bson:"card_type" example"super_miles_card"`
	Description string    `json:"description" bson:"description"`
	StartDate   time.Time `json:"start_date" example:"2023-03-02T13:10:23Z"'`
	EndDate     time.Time `json:"end_date" example:"2023-03-03T13:10:23Z"'`
}
