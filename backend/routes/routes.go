package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"vehicle-telemetry-system/backend/handlers"
)

func SetupRouter(handler *handlers.TelemetryHandler) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.POST("/telemetry", handler.PostTelemetry)
	router.GET("/telemetry/latest", handler.GetLatest)
	router.GET("/telemetry/history", handler.GetHistory)

	return router
}
