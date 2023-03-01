package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

// @Summary Health Check
// @Description Health Check Endpoint that doesn't require authentication
// @Success 200 {string} string	"Success"
// @Failure 404 {string} string "Not Found"
// @Router /health [get]
func (h HealthController) GetStatus(c *gin.Context) {
	c.String(http.StatusOK, "Success")
}
