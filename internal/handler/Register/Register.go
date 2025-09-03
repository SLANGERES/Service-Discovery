package Register

import (
	"fmt"
	"log/slog"
	"net/http"

	storage "github.com/SLANGERES/Service-Discovery/internal/Storage"
	models "github.com/SLANGERES/Service-Discovery/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)
func validation(service *models.Service)bool{
	err:=validator.New().Struct(service)
	if err!=nil{
		return false
	}
	return true
}
func RegisterService(c *gin.Context) {
	var service models.Service
	if err := c.ShouldBindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to parse the request body",
		})
		return
	}

	service.TTl = 60
	ok:=validation(&service);if !ok{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Validation error",
		})
		return
	}
	fmt.Print(service.Name,service.Host,service.Port,service.TTl)

	storage.AddIntoStorage(service)

	c.JSON(http.StatusOK, gin.H{
		"message": "Service update sucessfully",
	})
}
func UnRegisterRoute(c *gin.Context) {
	var data models.Service
	err:=c.BindJSON(&data);if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to parse the request body",
		})
		return
	}
	storage.RemoveFromStorage(data)
	slog.String(data.Name, ": Unregisterd sucessfully")
}
