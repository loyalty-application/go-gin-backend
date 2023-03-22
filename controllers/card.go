package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/collections"
	"github.com/loyalty-application/go-gin-backend/models"
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
// @Router  /transaction [get]
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
	result, err := collections.GetAllCards(skipInt, limitInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPError{Code: http.StatusInternalServerError, Message: "Failed to retrieve cards"})
		return
	}

	c.JSON(http.StatusOK, result)

}