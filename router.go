package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/controllers"
	"github.com/loyalty-application/go-gin-backend/docs"
	"github.com/loyalty-application/go-gin-backend/middlewares"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRoutes() {
	PORT := os.Getenv("PORT")

	health := new(controllers.HealthController)
	auth := new(controllers.AuthController)
	transaction := new(controllers.TransactionController)
	campaign := new(controllers.CampaignController)

	// necessary for swagger
	docs.SwaggerInfo.BasePath = "/api/v1"

	router := gin.New()
	// logging to stdout
	router.Use(gin.Logger())
	// recover from panics and respond with internal server error
	router.Use(gin.Recovery())

	// swagger
	swaggerGroup := router.Group("/docs")
	swaggerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// v1 group
	v1 := router.Group("/api/v1")

	// healthcheck
	healthGroup := v1.Group("/health")
	healthGroup.GET("/", health.GetStatus)

	// authentication
	authGroup := v1.Group("/auth")
	authGroup.POST("/registration", auth.Registration)
	authGroup.POST("/login", auth.Login)

	// transaction
	transactionGroup := v1.Group("/transaction")
	transactionGroup.Use(middlewares.AuthMiddleware())

	transactionGroup.GET("/:userId", transaction.GetTransactions)
	transactionGroup.POST("/:userId", transaction.PostTransactions)

	// Create a campaign
	campaignGroup := v1.Group("/campaign")
	campaignGroup.Use(middlewares.AuthMiddleware())

	campaignGroup.GET("/", campaign.GetCampaigns)
	//campaignGroup.GET("/:id", campaign.GetCampaignId)
	//campaignGroup.POST("", campaign.PostCampaign)
	//campaignGroup.PUT(":id", campaign.UpdateCampaign)
	//campaignGroup.DELETE(":id", campaign.DeleteCampaign)

	router.Run(":" + PORT)
}
