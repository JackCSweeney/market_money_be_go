package handlers

import (
	"example.com/mod/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func HandleGetAllVendors(context *gin.Context) {
	vendors, err := models.GetAllVendors()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": vendors})
}

func HandleGetVendorById(context *gin.Context, id int) {
	vendor, err := models.GetVendorById(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"data": vendor})
}
