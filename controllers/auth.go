package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/loyalty-application/go-gin-backend/collections"
	"github.com/loyalty-application/go-gin-backend/models"
	"github.com/loyalty-application/go-gin-backend/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthController struct{}

var validate = validator.New()

func (a AuthController) Login(c *gin.Context) {

	var user models.User
	var dbUser models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbUser, err := collections.RetrieveUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Login"})
		return
	}

	passwordIsValid, msg := services.VerifyPassword(*user.Password, *dbUser.Password)
	if passwordIsValid != true {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	token, refreshToken, _ := services.GenerateAllTokens(*dbUser.Email, *dbUser.First_name, *dbUser.Last_name, dbUser.User_id)

	services.UpdateAllTokens(token, refreshToken, dbUser.User_id)

	c.JSON(http.StatusOK, dbUser)

}

// @Summary     Registration
// @Description Registration endpoint for user new users to register for an account
// @Accept      application/json
// @Produce     application/json
// @Success     200 {string} string	"OK"
// @Failure     400 {string} string "Bad Request"
// @Router      /auth/register [post]
func (a AuthController) Register(c *gin.Context) {
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

	count, err := collections.CountUserEmail(*user.Email)
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the email"})
		return
	}

	count, err = collections.CountUserPhone(*user.Phone)
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the phone number"})
		return
	}

	// hash password
	password := services.HashPassword(*user.Password)
	user.Password = &password

	if count > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "this email or phone number already exists"})
		return
	}

	user.ID = primitive.NewObjectID()
	user.User_id = user.ID.Hex()
	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	// generate token for user
	token, refreshToken, _ := services.GenerateAllTokens(*user.Email, *user.First_name, *user.Last_name, user.User_id)
	user.Token = &token
	user.Refresh_token = &refreshToken

	insertRes, insertErr := collections.CreateUser(user)

	if insertErr != nil {
		msg := fmt.Sprintf("User item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, insertRes)

}