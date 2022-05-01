package router

import (
	"github.com/gin-gonic/gin"
	"go-api/services"
)

func Router() *gin.Engine {
	app := gin.Default()
	app.Use(services.Cors())
	api := app.Group("/api")
	{
		api.POST("/upload", services.Upload)
	}

	app.Static("storage", "storage")
	return app
}
