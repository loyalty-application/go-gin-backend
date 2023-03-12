package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/loyalty-application/go-gin-backend/collections"
	"github.com/loyalty-application/go-gin-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type CampaignController struct{}

func (t CampaignController) GetCampaigns(c *gin.Context){

}

func (t CampaignController) GetCampaignId(c *gin.Context){

}

func (t CampaignController) CreateCampaign(c *gin.Context){
    var campaign Campaign

    if err := c.BindJSON(&campaign); err != nil {
        c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
        return
    }

    if err := ctrl.service.CreateCampaign(&campaign); err != nil {
        c.JSON(http.StatusBadRequest, models.HTTPError{http.StatusBadRequest, "Invalid Transaction Object" + err.Error()})
        return
    }

    c.JSON(http.StatusOK, campaign)
}

func (t CampaignController) UpdateCampaign(c *gin.Context){

}

func (t CampaignController) DeleteCampaign(c *gin.Context){

}