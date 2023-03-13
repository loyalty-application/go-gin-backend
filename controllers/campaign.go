package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/collections"
	"github.com/loyalty-application/go-gin-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strconv"
)

type CampaignController struct{}

func (t CampaignController) GetCampaigns(c *gin.Context) {

}

// GetCampaignId @Summary Retrieve Campaign of a campaignId
// @Description Retrieve Campaign of a campaignId
// @Tags    campaigns
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Param   campaign_id path string true "capaign's id"
// @Param   limit query int false "maximum records per page" minimum(0) default(100)
// @Param   page query int false "page of records, starts from 0" minimum(0) default(0)
// @Success 200 {object} []models.Campaigns
// @Failure 400 {object} models.HTTPError
// @Router  /campaign/{campaign_id} [get]
func (t CampaignController) GetCampaignId(c *gin.Context) {
	campaginId := c.Param("campaignId")
	if campaginId == "" {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Campaign Id"})
		return
	}

	// required
	limit := c.Query("limit")
	if limit == "" {
		limit = "100"
	}

	// optional
	page := c.Query("page")
	if page == "" {
		page = "0"
	}

	pageInt, err := strconv.ParseInt(page, 10, 64)
	limitInt, err := strconv.ParseInt(limit, 10, 64)

	if pageInt < 0 || limitInt <= 0 {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Param page should be >= 0 and limit should be > 0 "})
		return
	}

	skipInt := pageInt * limitInt
	result, err := collections.RetrieveAllCampaigns(campaginId, skipInt, limitInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPError{http.StatusInternalServerError, "Failed to retrieve campaigns"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (t CampaignController) PostCampaign(c *gin.Context) {
	// TODO: should post campaign on merchantId
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid User Id"})
		return
	}

	data := new(models.CampaignList)
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Campaign Object" + err.Error()})
		return
	}

	result, err := collections.CreateCampaign(userId, *data)
	if err != nil {
		msg := "Invalid Campaign"
		if mongo.IsDuplicateKeyError(err) {
			msg = "campaign_id already exists"
		}
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, msg})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (t CampaignController) UpdateCampaign(c *gin.Context) {

}

func (t CampaignController) DeleteCampaign(c *gin.Context) {

}
