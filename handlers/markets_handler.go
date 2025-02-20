package handlers

import (
	"example.com/mod/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func HandleGetAllMarkets(context *gin.Context) {
	markets, err := models.GetAllMarkets()
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": markets})
}

func HandleGetMarketById(context *gin.Context, id int) {
	market, err := models.GetMarketById(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"data": market})
}