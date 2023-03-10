package models

type Campaign struct {
	Id int `json:"id" bson:"id" example:"1"`
	Merchant string `json:"merchant" bson:"merchant" example:"7-11"`
	CardType string  `json:"card_type" bson:"card_type" example"super_miles_card"`
	Description string `json:"description" bson:"description"`
}