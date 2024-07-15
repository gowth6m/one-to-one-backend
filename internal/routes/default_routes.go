package routes

import (
	"net/http"
	"one-to-one/internal/api"
	"one-to-one/internal/config"

	"github.com/gin-gonic/gin"
)

// For /
func DefaultRoutes(group *gin.Engine) {
	group.GET("/", func(c *gin.Context) {
		api.Success(c, http.StatusOK, "Online OneToOne", gin.H{
			"message": "Welcome to the Online OneToOne REST API v0",
			"version": config.AppConfig().App.AppVersion,
		})
	})
}
