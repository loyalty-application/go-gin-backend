package validators

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/models"
	"net/http"
	"time"
)

func ValidateStartDate(c *gin.Context, startDate time.Time) error {
	if startDate.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Start Date"})
		return fmt.Errorf("invalid start date")
	}
	return nil
}
