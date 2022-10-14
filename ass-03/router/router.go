package router

import (
	"ass-03/controller"
	"ass-03/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	// router.POST("/statuswater", controller.CreateStatus)
	// router.GET("/statuswater", controller.GetNewStatus)

	router.Use(middleware.CORSMiddleware())
	router.GET("statuswater", controller.GetNewStatus)

	return router
}
