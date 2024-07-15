package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Swagger routes for API documentation
	SwaggerRoutes(router)

	// Default routes for the root path
	DefaultRoutes(router)

	// User routes for the /user path
	UserRoutes(router)

	// One-to-one routes for the /one-to-one path
	OneToOneRoutes(router)
}
