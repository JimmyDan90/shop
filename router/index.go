package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop/config"
)

func InitRouter() *gin.Engine  {
	route := gin.New()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	gin.SetMode(config.RunMode)
	apiv1 := route.Group("/api/v1")
	{
		apiv1.GET("/shop", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "successful",
			})
		})
	}
	return route
}
