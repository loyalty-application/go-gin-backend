package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/models"
)

type TransactionController struct{}

func (t TransactionController) GetTransaction(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

func (t TransactionController) CreateTransaction(c *gin.Context) {
	// bind json to model
	data := new(models.TransactionRowList)
	err := c.BindJSON(data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Bad Request,"+err.Error())
		return
	}

	// iterate over all transactions and output the result
	var totalTransactionAmountByA float32 = 0
	for _, transaction := range data.Transactions {
		if transaction.Merchant == "A" {
			totalTransactionAmountByA += transaction.Amount
		}
	}

	// return it
	c.JSON(http.StatusOK, gin.H{
		"total": totalTransactionAmountByA,
	})

}
