package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/collections"
	"github.com/loyalty-application/go-gin-backend/models"
	"github.com/loyalty-application/go-gin-backend/validators"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type CampaignController struct{}

// @Summary Retrieve Campaign based on campaignId
// @Description Retrieve Campaign based on campaignId
// @Tags    campaign
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Param   campaign_id path string true "campaign's id"
// @Success 200 {object} models.Campaign
// @Failure 400 {object} models.HTTPError
// @Router  /campaign/{campaign_id} [get]
func (t CampaignController) GetCampaignId(c *gin.Context) {
	campaignId := c.Param("campaignId")

	if campaignId == "" {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Campaign Id"})
		return
	}

	result, err := collections.RetrieveCampaign(campaignId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPError{http.StatusInternalServerError, "Failed to retrieve campaign on campaignId"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Retrieve all Campaigns
// @Description Retrieve all campaigns, sorted by start date
// @Tags    campaign
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Success 200 {object} models.CampaignList
// @Failure 400 {object} models.HTTPError
// @Router  /campaign [get]
func (t CampaignController) GetCampaigns(c *gin.Context) {
	result, err := collections.RetrieveAllCampaigns()

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.HTTPError{http.StatusInternalServerError, "Failed to retrieve campaign"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// @Summary Create Campaigns for Merchants
// @Description Create campaigns
// @Tags    campaign
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Param   user_id path string true "user's id"
// @Param   request body models.CampaignList true "campaigns"
// @Success 200 {object} []models.Campaign
// @Failure 400 {object} models.HTTPError
// @Router  /campaign/{user_id} [post]
func (t CampaignController) PostCampaign(c *gin.Context) {
	// TODO: should post campaign on merchantId
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid User Id"})
		return
	}

	data := new(models.CampaignList)
	err := c.BindJSON(data)
	startDate := data.Campaigns[0].StartDate
	endDate := data.Campaigns[0].EndDate
	cardType := data.Campaigns[0].CardType

	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Campaign Object" + err.Error()})
		return
	}

	if err := validators.ValidateStartDate(c, startDate); err != nil {
		return
	}

	if err := validators.ValidateEndDate(c, startDate, endDate); err != nil {
		return
	}

	if err := validators.ValidateCardType(c, cardType); err != nil {
		return
	}

	_, err = collections.CreateCampaign(userId, *data)
	if err != nil {
		msg := "Invalid Campaign"
		if mongo.IsDuplicateKeyError(err) {
			msg = "campaign_id already exists"
		}
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, msg})
		return
	}

	c.JSON(http.StatusOK, "Campaign with campaign ID: "+data.Campaigns[0].CampaignId+" is created")
}

// @Summary Update Campaign based on campaignId
// @Description Update Campaign based on campaignId
// @Tags    campaign
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Param   campaign_id path string true "campaign's id"
// @Param   body body models.CampaignList true "campaign"
// @Success 200 {object} models.Campaign
// @Failure 400 {object} models.HTTPError
// @Router  /campaign/{campaign_id} [put]
func (t CampaignController) UpdateCampaign(c *gin.Context) {
	campaignId := c.Param("campaignId")
	if campaignId == "" {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Campaign Id"})
		return
	}

	data := new(models.Campaign)
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Campaign Object" + err.Error()})
		return
	}

	_, err = collections.UpdateCampaign(campaignId, *data)
	if err != nil {
		msg := "Invalid Campaign"
		if mongo.IsDuplicateKeyError(err) {
			msg = "campaign_id already exists"
		}
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, msg})
		return
	}

	c.JSON(http.StatusOK, "Campaign with campaign ID: "+data.CampaignId+" is updated")
}

// @Summary Delete Campaign based on campaignId
// @Description Delete Campaign based on campaignId
// @Tags    campaign
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Param   campaign_id path string true "campaign's id"
// @Param   body body models.CampaignList true "campaign"
// @Success 200 {object} models.Campaign
// @Failure 400 {object} models.HTTPError
// @Router  /campaign/{campaign_id}/delete [put]
func (t CampaignController) DeleteCampaign(c *gin.Context) {
	campaignId := c.Param("campaignId")
	if campaignId == "" {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Campaign Id"})
		return
	}

	data := new(models.Campaign)
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Campaign Object" + err.Error()})
		return
	}

	_, err = collections.DeleteCampaign(campaignId, *data)
	if err != nil {
		msg := "Invalid Campaign"
		if mongo.IsDuplicateKeyError(err) {
			msg = "campaign_id already exists"
		}
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, msg})
		return
	}

	c.JSON(http.StatusOK, "Campaign with campaign ID: "+campaignId+" is deleted")
}