package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/models"
)

type TransactionController struct{}

func (t TransactionController) GetTransactions(c *gin.Context) {
	// get path param
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, "Invalid User Id")
	}

	// retrieve user's transactions from database

	c.String(http.StatusOK, "Working!")
}

func (t TransactionController) PostTransactions(c *gin.Context) {
	// bind json to model
	data := new(models.TransactionList)
	err := c.BindJSON(data)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	//result, err := .InsertMany(context.TODO(), docs)
	c.String(http.StatusOK, "Success")
}
