package models

type Campaign struct {
	Merchant string `json:"merchant" bson:"merchant" example:"7-11"`
}
