package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/controllers"
	"github.com/loyalty-application/go-gin-backend/middlewares"
)

func InitRoutes() {
	PORT := os.Getenv("SERVER_PORT")
	HOST := "0.0.0.0"

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// /health
	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	// /auth
	auth := new(controllers.AuthController)
	router.GET("/auth", auth.Login)

	// apply middleware to all routes after
	router.Use(middlewares.AuthMiddleware())

	// /v1 route
	v1 := router.Group("v1")

	// /v1/transaction
	transactionGroup := v1.Group("transaction")
	transaction := new(controllers.TransactionController)
	transactionGroup.GET("/", transaction.GetTransaction)
	transactionGroup.POST("/:id", transaction.CreateTransaction)

	router.Run(HOST + ":" + PORT)
}
