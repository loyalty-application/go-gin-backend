package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (a AuthController) UserLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"token": "TOKEN HERE",
	})
}
