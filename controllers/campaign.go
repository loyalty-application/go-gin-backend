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