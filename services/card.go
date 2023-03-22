package services

import "github.com/loyalty-application/go-gin-backend/models"

// return type of card
func ProcessCardType(card models.Card) string {

	switch str := card.CardType; str {
	case "scis_freedom":
		return "Points"
	case "scis_premiummiles", "scis_platinummiles":
		return "Miles"
	case "scis_shopping":
		return "Cashback"
	default:
		return "Error"
	}

}