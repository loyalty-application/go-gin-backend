package models

type TransactionList struct {
	Transactions []Transaction `json:"transactions" bson:",inline"`
}

type Transaction struct {
	UserId          string  `json:"-" bson:"user_id"`
	Id              int     `json:"id" bson:"id"`
	TransactionId   string  `json:"transaction_id" bson:"transaction_id"`
	Merchant        string  `json:"merchant" bson:"merchant"`
	MCC             string  `json:"mcc" bson:"mcc"`
	Currency        string  `json:"currency" bson:"currency"`
	Amount          float64 `json:"amount" bson:"amount"`
	TransactionDate string  `json:"transaction_date" bson:"transaction_date"`
	CardId          string  `json:"card_id" bson:"card_id"`
	CardPan         string  `json:"card_pan" bson:"card_pan"`
	CardType        string  `json:"card_type" bson:"card_type"`
}
