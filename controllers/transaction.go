package controllers

import (
	"net/http"
	"strconv"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/collections"
	"github.com/loyalty-application/go-gin-backend/kafka"
	"github.com/loyalty-application/go-gin-backend/models"
)

type TransactionController struct{}

// @Summary Update Transaction
// @Description Update a Specific Transaction
// @Tags    transaction
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Param   transaction_id path string true "transaction's id"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} models.HTTPError
// @Router  /transaction/{transaction_id} [put]
func (t TransactionController) UpdateTransaction(c *gin.Context) {
	transactionId := c.Param("transactionId")
	if transactionId == "" {
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: "Invalid Transaction Id"})
		return
	}

	data := new(models.Transaction)
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: "Invalid Transaction Object " + err.Error()})
		return
	}

	// transaction id should not be modified
	if data.TransactionId != "" {
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: "Transaction Id should not be modified"})
	}

	result, err := collections.UpdateTransaction(transactionId, *data)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: "Invalid Transaction " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Retrieve Transactions of all Users
// @Description Retrieve all transaction records
// @Tags    transaction
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Param   limit query int false "maximum records per page" minimum(0) default(100)
// @Param   page query int false "page of records, starts from 0" minimum(0) default(0)
// @Success 200 {object} []models.Transaction
// @Failure 400 {object} models.HTTPError
// @Router  /transaction [get]
func (t TransactionController) GetAllTransactions(c *gin.Context) {

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
	result, err := collections.RetrieveAllTransactions(skipInt, limitInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPError{Code: http.StatusInternalServerError, Message: "Failed to retrieve transactions"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Counts the total records in the db
// @Description Counts the total records in the db
// @Tags    transaction
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Success 200 {object} int64
// @Failure 400 {object} models.HTTPError
// @Router  /transaction/count [get]
func (t TransactionController) CountAllTransactions(c *gin.Context) {

	count, err := collections.CountTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPError{Code: http.StatusInternalServerError, Message: "Failed to count transactions"})
		return
	}

	c.JSON(http.StatusOK, count)
}

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
func (t TransactionController) GetTransactionsForUser(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: "Invalid User Id"})
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
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: "Param page should be >= 0 and limit should be > 0 "})
		return
	}

	skipInt := pageInt * limitInt
	cardList, err := collections.RetrieveCardsByUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPError{Code: http.StatusInternalServerError, Message: "User doesn't have any cards"})
		return
	}

	cardIdList := make([]string, len(cardList))
	for i, card := range cardList {
		cardIdList[i] = card.CardId
	}

	result, err := collections.RetrieveAllTransactionsForUser(cardIdList, skipInt, limitInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPError{Code: http.StatusInternalServerError, Message: "Failed to retrieve transactions"})
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
	data := new(models.TransactionList)
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: "Invalid Transaction Object" + err.Error()})
		return
	}
	log.Println("Hello World")
	result, err := collections.CreateTransactions(*data)
	if err != nil {
		// msg := "Invalid Transactions"
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	for _, transaction := range data.Transactions {
		// index is the index where we are
		// element is the element from someSlice for where we are
		kafka.ProduceMessage(transaction)
	}

	c.JSON(http.StatusCreated, result)
}

func (t TransactionController) DeleteTransaction(c *gin.Context) {
	transactionId := c.Param("transactionId")
	if transactionId == "" {
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: "Invalid Transaction Id"})
		return
	}

	result, err := collections.DeleteTransaction(transactionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{Code: http.StatusBadRequest, Message: "Transaction doesn't exist"})
		return
	}

	c.JSON(http.StatusOK, result)
}
