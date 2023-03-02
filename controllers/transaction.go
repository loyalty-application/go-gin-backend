package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/collections"
	"github.com/loyalty-application/go-gin-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionController struct{}

func (t TransactionController) GetTransactions(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User Id"})
		return
	}

	// required
	limit := c.Query("limit")
	if limit == "" {
		limit = "2"
	}

	// optional
	page := c.Query("page")
	if page == "" {
		page = "0"
	}

	pageInt, err := strconv.ParseInt(page, 10, 64)
	limitInt, err := strconv.ParseInt(limit, 10, 64)

	if pageInt < 0 || limitInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "param page should be >= 0 and limit should be > 0 "})
		return
	}

	skipInt := pageInt * limitInt
	result, err := collections.RetrieveAllTransactions(userId, skipInt, limitInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)

}

func (t TransactionController) PostTransactions(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User Id"})
		return
	}

	data := new(models.TransactionList)
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Transaction Object" + err.Error()})
		return
	}

	// TODO: make this operation atomic https://www.mongodb.com/docs/drivers/go/current/fundamentals/transactions/
	result, err := collections.CreateTransactions(userId, *data)
	if err != nil {
		msg := "Invalid Transactions"
		if mongo.IsDuplicateKeyError(err) {
			msg = "transaction_id already exists"
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, result)
}
