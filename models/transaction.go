package models

type TransactionList struct {
	UserId       string        `bson:"_id"`
	Transactions []interface{} `json:"transactions"`
}

type TransactionRow struct {
	Id              int     `json:"id"`
	TransactionId   string  `json:"transaction_id"`
	Merchant        string  `json:"merchant"`
	MCC             string  `json:"mcc"`
	Currency        string  `json:"currency"`
	Amount          float32 `json:"amount"`
	TransactionDate string  `json:"transaction_date"`
	CardId          string  `json:"card_id"`
	CardPan         string  `json:"card_pan"`
	CardType        string  `json:"card_type"`
}
