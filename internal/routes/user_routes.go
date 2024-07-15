package routes

import (
	"github.com/gin-gonic/gin"
	"one-to-one/internal/middleware"
	"one-to-one/internal/services/user"
)

// GROUP: /user
func UserRoutes(group *gin.Engine) {
	userRepo := user.NewUserRepository()
	userHandler := user.NewUserHandler(userRepo)

	userGroup := group.Group("/user")

	// --- PUBLIC ROUTES ---
	userGroup.POST("/create", func(c *gin.Context) {
		userHandler.CreateUser(c)
	})

	userGroup.GET("/all", func(c *gin.Context) {
		userHandler.GetAllUsers(c)
	})

	userGroup.GET("/email/:email", func(c *gin.Context) {
		userHandler.GetUserByEmail(c)
	})

	userGroup.POST("/login", func(c *gin.Context) {
		userHandler.LoginUser(c)
	})

	// --- PROTECTED ROUTES ---
	userGroup.Use(middleware.JWTAuthMiddleware())
	{
		userGroup.GET("/current", func(c *gin.Context) {
			userHandler.GetCurrentUser(c)
		})
	}
}
