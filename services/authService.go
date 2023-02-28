package services

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct{}

var JWT_SECRET = os.Getenv("JWT_SECRET")

func (a AuthService) GetJWTSecret() []byte {
	JWT_SECRET_BYTES := []byte(JWT_SECRET)
	return JWT_SECRET_BYTES

}

// generates a token with the JWT_SECRET environment variable
func (a AuthService) GenerateToken() string {

	JWT_SECRET_BYTES := a.GetJWTSecret()

	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(24))),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(JWT_SECRET_BYTES)
	if err != nil {

	}
	return ss
}

// validates the given token, returns false if the token is invalid
func (a AuthService) ValidateToken(tokenString string) bool {

	JWT_SECRET_BYTE := a.GetJWTSecret()

	// attempt to parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return JWT_SECRET_BYTE, nil
	})

	// if there was an error parsing
	if err != nil {
		return false
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// token is valid, return true
		return true
	}

	// if it reaches here, its invalid
	return false

}
