package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/collections"
	"github.com/loyalty-application/go-gin-backend/models"
)

type TransactionController struct{}

func (t TransactionController) GetTransactions(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User Id"})
		return
	}

	result, err := collections.RetrieveAllTransactions(userId)
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
	data.UserId = userId
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Transactions"})
		return
	}

	result, err := collections.CreateTransactions(*data)
	fmt.Println(result)

	c.String(http.StatusOK, "Success")
}
