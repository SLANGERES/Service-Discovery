package Register

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	storage "github.com/SLANGERES/Service-Discovery/internal/Storage"
	models "github.com/SLANGERES/Service-Discovery/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func validation(service *models.Service) bool {
	return validator.New().Struct(service) == nil
}
func RegisterService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to parse the request body",
		})
		return
	}

	service.TTl = 180
	service.Expires = time.Now().Add(time.Duration(service.TTl) * time.Second)
	ok := validation(&service)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation error",
		})
		return
	}
	fmt.Print(service.Name, service.Host, service.Port, service.TTl)

	storage.AddIntoStorage(service)

	c.JSON(http.StatusOK, gin.H{
		"message": "Service update sucessfully",
		
	})
}
func UnRegisterRoute(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to parse the request body",
		})
		return
	}
	storage.RemoveFromStorage(id)
	c.JSON(http.StatusOK,gin.H{
		"message":"service is unregistered Sucessfully",
	})
	slog.Info(fmt.Sprintf("%s: Unregistered successfully", id))
}
