package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/collections"
	"github.com/loyalty-application/go-gin-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionController struct{}

// @Summary Retrieve Transactions of User
// @Description Retrieve transaction records of a user
// @Tags    transaction
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Param   user_id path string true "user's id"
// @Param   limit query int false "maximum records per page" minimum(0) default(100)
// @Param   page query int false "page of records, starts from 0" minimum(0) default(0)
// @Success 200 {object} []models.Transaction
// @Failure 400 {object} models.HTTPError
// @Router  /transaction/{user_id} [get]
func (t TransactionController) GetTransactions(c *gin.Context) {
	fmt.Println("Print Test")
	log.Println("Log Test")
	os.Stdout.Sync()
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid User Id"})
		return
	}

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
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Param page should be >= 0 and limit should be > 0 "})
		return
	}

	skipInt := pageInt * limitInt
	result, err := collections.RetrieveAllTransactions(userId, skipInt, limitInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPError{http.StatusInternalServerError, "Failed to retrieve transactions"})
		return
	}

	c.JSON(http.StatusOK, result)

}

// @Summary Create Transactions for User
// @Description Create transaction records
// @Tags    transaction
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Param   user_id path string true "user's id"
// @Param   request body models.TransactionList true "transactions"
// @Success 200 {object} []models.Transaction
// @Failure 400 {object} models.HTTPError
// @Router  /transaction/{user_id} [post]
func (t TransactionController) PostTransactions(c *gin.Context) {
	userId := c.Param("userId")
	fmt.Printf("SDhuajkhdfsjkhfs")
	log.Println("sahdgas")
	if userId == "" {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid User Id"})
		return
	}

	data := new(models.TransactionList)
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Transaction Object" + err.Error()})
		return
	}

	// TODO: make this operation atomic https://www.mongodb.com/docs/drivers/go/current/fundamentals/transactions/
	result, err := collections.CreateTransactions(userId, *data)
	if err != nil {
		msg := "Invalid Transactions"
		if mongo.IsDuplicateKeyError(err) {
			msg = "transaction_id already exists"
		}
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, msg})
		return
	}

	c.JSON(http.StatusOK, result)
}
