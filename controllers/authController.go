package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/services"
)

type AuthController struct{}

func (a AuthController) Login(c *gin.Context) {

	auth := new(services.AuthService)
	token := auth.GenerateToken()

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func (a AuthController) SignUp(c *gin.Context) {

}
