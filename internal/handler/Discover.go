package handler

import (
	"net/http"

	"github.com/SLANGERES/Service-Discovery/internal/models"
	storage "github.com/SLANGERES/Service-Discovery/internal/Storage"
	"github.com/gin-gonic/gin"
)

func Discovery(c *gin.Context) {
	serviceName := c.Param("service")

	returnData := storage.GetServicesByName(serviceName)

	if len(returnData) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message":  "No instances found",
			"Instance": []models.Service{},
		})
		return
	}
	responseData := make([]models.Service, 0, len(returnData))
	for _, v := range returnData {
		responseData = append(responseData, v)
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "Service Instance",
		"Instance": responseData,
	})
}
