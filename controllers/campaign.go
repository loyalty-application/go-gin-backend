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

// GetCampaignId @Summary Retrieve Campaign based on campaignId
// @Description Retrieve Campaign based on campaignId
// @Tags    campaign
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Param   campaign_id path string true "capaign's id"
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
		c.JSON(http.StatusInternalServerError, models.HTTPError{http.StatusInternalServerError, "Failed to retrieve campaign"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetCampaignId @Summary Retrieve Campaigns of a merchant
// @Description Retrieve Campaigns of a merchant
// @Tags    campaign
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Param   campaign_id path string true "capaign's id"
// @Param   limit query int false "maximum records per page" minimum(0) default(100)
// @Param   page query int false "page of records, starts from 0" minimum(0) default(0)
// @Success 200 {object} []models.Campaigns
// @Failure 400 {object} models.HTTPError
// @Router  /campaign/{campaign_id} [get]
func (t CampaignController) GetCampaigns(c *gin.Context) {
	merchantId := c.Param("merchantId")
	if merchantId == "" {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Merchant Id"})
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

// UpdateCampaign @Summary Update Campaign based on campaignId
// @Description Update Campaign based on campaignId
// @Tags    campaign
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Param   campaign_id path string true "campaign's id"
// @Param   body body models.CampaignList true "Campaign object"
// @Success 200 {object} models.Campaign
// @Failure 400 {object} models.HTTPError
// @Router  /campaign/{campaign_id} [put]
func (t CampaignController) UpdateCampaign(c *gin.Context) {
	campaignId := c.Param("campaignId")
	if campaignId == "" {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Campaign Id"})
		return
	}

	data := new(models.CampaignList)
	err := c.BindJSON(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Campaign Object" + err.Error()})
		return
	}

	result, err := collections.UpdateCampaign(campaignId, *data)
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

// DeleteCampaign @Summary Delete Campaign based on campaignId
// @Description Delete Campaign based on campaignId
// @Tags    campaign
// @Accept  application/json
// @Produce application/json
// @Param   Authorization header string true "Bearer eyJhb..."
// @Param   campaign_id path string true "campaign's id"
// @Success 204
// @Failure 400 {object} models.HTTPError
// @Router  /campaign/{campaign_id} [delete]
func (t CampaignController) DeleteCampaign(c *gin.Context) {
	campaignId := c.Param("campaignId")
	if campaignId == "" {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Campaign Id"})
		return
	}

	err := collections.DeleteCampaign(campaignId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Campaign"})
		return
	}

	c.JSON(http.StatusNoContent)
}
