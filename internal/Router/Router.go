package router

import (
	"github.com/SLANGERES/Service-Discovery/internal/handler"
	"github.com/SLANGERES/Service-Discovery/internal/handler/Register"
	"github.com/gin-gonic/gin"
)

func RegisterRouter()*gin.Engine{
	Router:=gin.Default()

	
	Router.POST("/register",Register.RegisterService)
	Router.GET("/discover/:service",handler.Discovery)
	Router.POST("/unregister",Register.UnRegisterRoute)
	return  Router
}