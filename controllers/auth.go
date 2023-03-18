package controllers

import (
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

// @Summary Login
// @Description Users can login to the application and obtain a JWT token through this endpoint
// @Tags    authentication
// @Accept  application/json
// @Produce application/json
// @Param   request body models.AuthLoginRequest true "Login"
// @Success 200 {object} models.AuthLoginResponse
// @Failure 400 {object} models.HTTPError
// @Router  /auth/login [post]
func (a AuthController) Login(c *gin.Context) {

	var user models.User
	var dbUser models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Request Body"})
		return
	}

	dbUser, err := collections.RetrieveUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Login"})
		return
	}

	passwordIsValid := services.VerifyPassword(*user.Password, *dbUser.Password)
	if passwordIsValid != true {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Login"})
		return
	}

	token, refreshToken, _ := services.GenerateAllTokens(*dbUser.Email, *dbUser.FirstName, *dbUser.LastName, dbUser.UserID.Hex())
	services.UpdateAllTokens(token, refreshToken, dbUser.UserID.Hex())

	c.JSON(http.StatusOK, dbUser)

}

// @Summary Registration
// @Description Registration endpoint for user new users to register for an account, after registering for an account, the user will be able to login to the system and obtain a JWT Token
// @Tags    authentication
// @Accept  application/json
// @Produce application/json
// @Param request body models.AuthRegistrationRequest true "Registration"
// @Success 200 {object} models.AuthRegistrationResponse
// @Failure 400 {object} models.HTTPError
// @Failure 500 {object} models.HTTPError
// @Router  /auth/registration [post]
func (a AuthController) Registration(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Registration Request"})
		return
	}

	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Registration Request"})
		return
	}

	count, err := collections.CountUserEmail(*user.Email)
	if err != nil || count > 0 {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Email already exists"})
		return
	}

	count, err = collections.CountUserPhone(*user.Phone)
	if err != nil || count > 0 {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Phone number already exists"})
		return
	}

	password := services.HashPassword(*user.Password)
	user.Password = &password
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	// generate token for user
	user.UserID = primitive.NewObjectID()
	token, refreshToken, _ := services.GenerateAllTokens(*user.Email, *user.FirstName, *user.LastName, user.UserID.Hex())

	user.Token = &token
	user.RefreshToken = &refreshToken

	result, err := collections.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPError{http.StatusBadRequest, "User was not created"})
		return
	}

	// TODO: change to proper request instead of mongodb's successful insertion format
	c.JSON(http.StatusOK, result)

}
