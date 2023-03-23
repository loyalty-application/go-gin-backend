package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/collections"
	"github.com/loyalty-application/go-gin-backend/models"
	"github.com/loyalty-application/go-gin-backend/services"
	"github.com/loyalty-application/go-gin-backend/validators"
	"go.mongodb.org/mongo-driver/mongo"
)

type CardController struct{}

// @Summary Retrieve Cards
// @Description Retrieve all available cards
// @Tags    card
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Param   limit query int false "maximum records per page" minimum(0) default(100)
// @Param   page query int false "page of records, starts from 0" minimum(0) default(0)
// @Success 200 {object} []models.Card
// @Failure 400 {object} models.HTTPError
// @Router  /card [get]
func (t CardController) GetCards(c *gin.Context) {

	// required
	limit := c.Query("limit")
	if limit == "" {
		limit = "100"
	}

	// optional
	page := c.Query("page")
	if page == "" {
		page = "0"
	}

	pageInt, err := strconv.ParseInt(page, 10, 64)
	limitInt, err := strconv.ParseInt(limit, 10, 64)

	if pageInt < 0 || limitInt <= 0 {
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: "Param page should be >= 0 and limit should be > 0 "})
		return
	}

	skipInt := pageInt * limitInt
	result, err := collections.RetrieveAllCards(skipInt, limitInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPError{Code: http.StatusInternalServerError, Message: "Failed to retrieve cards"})
		return
	}

	c.JSON(http.StatusOK, result)

}

// @Summary Retrieve specific Card
// @Description Retrieve card based on its card_id
// @Tags    card
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Param   card_id path string true "card's id"
// @Success 200 {object} models.Card
// @Failure 400 {object} models.HTTPError
// @Router  /card/{card_id} [get]
func (t CardController) GetSpecificCard(c *gin.Context) {
	cardId := c.Param("cardId")
	if cardId == "" {
		c.JSON(http.StatusInternalServerError, models.HTTPError{Code: http.StatusInternalServerError, Message: "cardId cannot be blank"})
		return
	}
	
	result, err := collections.RetrieveSpecificCard(cardId)
	if err != nil {
		msg := "Failed to retrieve card"
		if err == mongo.ErrNoDocuments {
			msg = "No card found with given card id"
		}
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: msg})
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Create Card
// @Description Create new Card
// @Tags    card
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Param   request body models.Card true "card"
// @Success 200 {object} []models.Card
// @Failure 400 {object} models.HTTPError
// @Router  /card [post]
func (t CardController) PostCard(c *gin.Context) {

	data := new(models.Card)
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: "Invalid Card Object" + err.Error()})
		return
	}

	// validating card type
	if err = validators.ValidateCardType(c, data.CardType); err != nil {
		return
	}

	// setting card valueType (points / miles / cashback)
	data.ValueType = services.ProcessCardType(*data)

	// set card value to 0.0
	data.Value = 0.0

	result, err := collections.CreateCard(*data)
	if err != nil {
		msg := "Failed to insert card" + err.Error()
		if mongo.IsDuplicateKeyError(err) {
			msg = "CardId already exists"
		}
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: msg})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// @Summary Update a Card
// @Description Update specific Card
// @Tags    card
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Param   card_id path string true "card's id"
// @Param   request body models.Card true "card"
// @Success 200 {object} models.CardUpdateRequest
// @Failure 400 {object} models.HTTPError
// @Router  /card/{card_id} [put]
func (t CardController) UpdateCard(c *gin.Context) {
	cardId := c.Param("cardId")
	if cardId == "" {
		c.JSON(http.StatusInternalServerError, models.HTTPError{Code: http.StatusBadRequest, Message: "cardId cannot be blank"})
		return
	}

	data := new(models.CardUpdateRequest)
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: "Invalid Card Object" + err.Error()})
		return
	}

	result, err := collections.UpdateCardPoints(cardId, *data)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: "Card Id doesn't exist"})
		return
	}

	c.JSON(http.StatusOK, result)
}