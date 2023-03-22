package validators

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/models"
	"net/http"
	"time"
)

// ValidateStartDate Validate StartDate is not before current time
func ValidateStartDate(c *gin.Context, startDate time.Time) error {
	if startDate.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest,
			"Invalid Start Date. Start Date cannot be before current date and time."})
		return fmt.Errorf("invalid start date")
	}
	return nil
}

// ValidateEndDate Validate EndDate is not before StartDate
func ValidateEndDate(c *gin.Context, startDate time.Time, endDate time.Time) error {
	if endDate.Before(startDate) {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest,
			"Invalid End Date. End Date cannot be before Start Date."})
		return fmt.Errorf("invalid end date")
	}
	return nil
}

// ValidateCardType Validate CardType is of the valid types
func ValidateCardType(c *gin.Context, cardType string) error {
	ValidCardType := [4]string{"scis_platinummiles", "scis_premiummiles", "scis_shopping", "scis_freedom"}
	for _, elem := range ValidCardType {
		if elem == cardType {
			return nil
		}
	}

	c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest,
		"Invalid Card Type."})
	return fmt.Errorf("invalid card type")
}