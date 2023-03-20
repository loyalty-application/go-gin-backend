package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/controllers"
	"github.com/loyalty-application/go-gin-backend/docs"
	"github.com/loyalty-application/go-gin-backend/middlewares"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes() {
	PORT := os.Getenv("SERVER_PORT")
	gin.SetMode(os.Getenv("GIN_MODE"))

	health := new(controllers.HealthController)
	auth := new(controllers.AuthController)
	transaction := new(controllers.TransactionController)
	campaign := new(controllers.CampaignController)

	// necessary for swagger
	docs.SwaggerInfo.BasePath = "/api/v1"

	router := gin.Default()
	// logging to stdout
	router.Use(gin.Logger())
	log.Println("GIN_MODE = ", os.Getenv("GIN_MODE"))

	// recover from panics and respond with internal server error
	router.Use(gin.Recovery())

	// swagger
	swaggerGroup := router.Group("/docs")
	swaggerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// v1 group
	v1 := router.Group("/api/v1")

	// healthcheck
	healthGroup := v1.Group("/health")
	healthGroup.GET("", health.GetStatus)

	// authentication
	authGroup := v1.Group("/auth")
	authGroup.POST("/registration", auth.Registration)
	authGroup.POST("/login", auth.Login)

	// transaction
	transactionGroup := v1.Group("/transaction")
	transactionGroup.Use(middlewares.AuthMiddleware())

	transactionGroup.GET("/", transaction.GetAllTransactions)
	transactionGroup.GET("/:userId", transaction.GetTransactionsForUser)
	transactionGroup.PUT("/:transactionId", transaction.UpdateTransaction)
	transactionGroup.POST("/:userId", transaction.PostTransactions)
	transactionGroup.DELETE("/:transactionId", transaction.DeleteTransaction)

	// Create a campaign
	campaignGroup := v1.Group("/campaign")
	campaignGroup.Use(middlewares.AuthMiddleware())

	campaignGroup.GET("/", campaign.GetCampaigns)
	campaignGroup.GET("/:campaignId", campaign.GetCampaignId)
	campaignGroup.POST("/:userId", campaign.PostCampaign)
	campaignGroup.PUT("/:campaignId", campaign.UpdateCampaign)
	campaignGroup.DELETE("/:campaignId", campaign.DeleteCampaign)

	router.Run(":" + PORT)
}
