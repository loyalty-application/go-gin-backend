package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/services"
)

// look in the header for the token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.String(http.StatusForbidden, "No Authorization header provided")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")
		if tokenString == auth {
			c.String(http.StatusForbidden, "Could not find bearer token in Authorization header")
			c.Abort()
			return
		}

		authService := new(services.AuthService)
		tokenValid := authService.ValidateToken(tokenString)

		if !tokenValid {
			c.String(http.StatusInternalServerError, "Invalid Token")
			c.Abort()
			return
		}

		c.Next()
	}
}
