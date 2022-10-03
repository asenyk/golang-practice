package controllers

import (
	"golang-practice/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAddressById(context *gin.Context) {
	var address_postals models.AddressPostals
	if context.Query("id") == "" {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Missing id parameter"})
		return
	}
	if err := models.DB.Where("id = ?", context.Query("id")).First(&address_postals).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": address_postals})
}
