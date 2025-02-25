package handlers

import (
	"example.com/mod/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"io"
	"encoding/json"
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
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": vendor})
}

func HandleUpdateVendor(context *gin.Context, id int) {
	var vendorUpdates models.Vendor

	body, err := io.ReadAll(context.Request.Body)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	json.Unmarshal([]byte(body), &vendorUpdates)

	vendor, err := models.UpdateVendor(id, vendorUpdates)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"data": vendor})
}
