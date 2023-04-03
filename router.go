package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
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
	card := new(controllers.CardController)

	// necessary for swagger
	docs.SwaggerInfo.BasePath = "/api/v1"

	router := gin.Default()
	// logging to stdout
	router.Use(gin.Logger())
	log.Println("GIN_MODE = ", os.Getenv("GIN_MODE"))

	// recover from panics and respond with internal server error
	router.Use(gin.Recovery())

	// enabling cors
	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

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

	// users
	userGroup := v1.Group("/user")
	userGroup.Use(middlewares.AuthMiddleware())
	
	userGroup.GET("", auth.GetAllUsers)
	userGroup.GET("/:userId", auth.GetSpecificUser)
	userGroup.POST("", auth.PostAccount)
	userGroup.PUT("/:email", auth.UpdateUser)

	// transaction
	transactionGroup := v1.Group("/transaction")
	transactionGroup.Use(middlewares.AuthMiddleware())

	transactionGroup.GET("", transaction.GetAllTransactions)
	transactionGroup.GET("/:userId", transaction.GetTransactionsForUser)
	transactionGroup.PUT("/:transactionId", transaction.UpdateTransaction)
	transactionGroup.POST("", transaction.PostTransactions)
	transactionGroup.DELETE("/:transactionId", transaction.DeleteTransaction)

	// campaign
	campaignGroup := v1.Group("/campaign")
	campaignGroup.Use(middlewares.AuthMiddleware())

	campaignGroup.GET("", campaign.GetCampaigns)
	campaignGroup.GET("/:campaignId", campaign.GetCampaignId)
	campaignGroup.GET("/active/:date", campaign.GetActiveCampaigns)
	campaignGroup.POST("", campaign.PostCampaign)
	campaignGroup.PUT("/:campaignId", campaign.UpdateCampaign)
	campaignGroup.PUT("/:campaignId/delete", campaign.DeleteCampaign)

	// card
	cardGroup := v1.Group("/card")
	cardGroup.Use(middlewares.AuthMiddleware())

	cardGroup.GET("", card.GetCards)
	cardGroup.GET("/:cardId", card.GetSpecificCard)
	cardGroup.GET("/user/:userId", card.GetCardsFromUser)
	cardGroup.POST("", card.PostCard)
	cardGroup.PUT("/:cardId", card.UpdateCard)

	router.Run(":" + PORT)
}
