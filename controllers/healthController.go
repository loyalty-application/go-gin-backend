package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

// @Description health check endpoint
// @Success 200 {string} string	"Success"
// @Failure 404 {string} string "Not Found"
// @Router /health [get]
func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Success")
}
