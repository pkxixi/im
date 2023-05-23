package router

import (
	"github.com/gin-gonic/gin"
	"im/middleware"
	"im/service"
)

func Router() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("v1")
	user := v1.Group("user")
	{
		user.GET("/list", middleware.JWY(), service.List)
	}

	return router
}
