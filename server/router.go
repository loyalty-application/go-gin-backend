package server

import (
	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/controllers"
	"github.com/loyalty-application/go-gin-backend/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// health check controller
	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	// use auth middleware
	router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("v1")
	transactionGroup := v1.Group("transaction")

	// transaction controller
	transaction := new(controllers.TransactionController)
	transactionGroup.GET("/", transaction.GetTransaction)
	transactionGroup.POST("/:id", transaction.CreateTransaction)

	return router

}
