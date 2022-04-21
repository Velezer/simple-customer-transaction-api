package routes

import (
	"sctrans/middlewares"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var r *gin.Engine = gin.Default()

func SetupRouter() *gin.Engine {
	r.Use(middlewares.ErrorMiddleware())

	r.POST("/customers")

	r.GET("/orders")
	r.POST("/orders")

	r.GET("/orders")
	r.POST("/orders")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
