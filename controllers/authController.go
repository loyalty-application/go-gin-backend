package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/loyalty-application/go-gin-backend/config"
	"github.com/loyalty-application/go-gin-backend/models"
	"github.com/loyalty-application/go-gin-backend/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct{}

var userCollection *mongo.Collection = config.OpenCollection(config.Client, "user")
var validate = validator.New()

func (a AuthController) Login(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user models.User
	var foundUser models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "login or passowrd is incorrect"})
		return
	}

	passwordIsValid, msg := services.VerifyPassword(*user.Password, *foundUser.Password)
	defer cancel()
	if passwordIsValid != true {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	token, refreshToken, _ := services.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, foundUser.User_id)

	services.UpdateAllTokens(token, refreshToken, foundUser.User_id)

	c.JSON(http.StatusOK, foundUser)

}

// @Summary     Registration
// @Description Registration endpoint for user new users to register for an account
// @Accept      application/json
// @Produce     application/json
// @Success     200 {string} string	"OK"
// @Failure     400 {string} string "Bad Request"
// @Router      /auth/register [post]
func (a AuthController) Register(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
		return
	}

	password := services.HashPassword(*user.Password)
	user.Password = &password

	count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
	defer cancel()
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the phone number"})
		return
	}

	if count > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "this email or phone number already exists"})
		return
	}

	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	token, refreshToken, _ := services.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, user.User_id)
	user.Token = &token
	user.Refresh_token = &refreshToken

	resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil {
		msg := fmt.Sprintf("User item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, resultInsertionNumber)

}
