package routes

import (
	"one-to-one/internal/middleware"
	one_to_one "one-to-one/internal/services/one-to-one"

	"github.com/gin-gonic/gin"
)

// GROUP: /one-to-one
func OneToOneRoutes(group *gin.Engine) {
	oneToOneRepo := one_to_one.NewOneToOneRepository()
	oneToOneHandler := one_to_one.NewOneToOneHandler(oneToOneRepo)

	oneToOneGroup := group.Group("/one-to-one")

	// --- PROTECTED ROUTES ---
	oneToOneGroup.Use(middleware.JWTAuthMiddleware())
	{
		oneToOneGroup.POST("/create", func(c *gin.Context) {
			oneToOneHandler.CreateWeeklyReport(c)
		})

		// --- REPORTEE ROUTES ---

		oneToOneGroup.GET("/reportee/all", func(c *gin.Context) {
			oneToOneHandler.GetAllWeeklyReportsForReportee(c)
		})

		oneToOneGroup.GET("/reportee", func(c *gin.Context) {
			oneToOneHandler.GetWeeklyReportByWeekAndYearForReportee(c)
		})

		oneToOneGroup.PUT("/reportee/update", func(c *gin.Context) {
			oneToOneHandler.UpdateWeeklyReportForReportee(c)
		})

		// --- REPORT TO ROUTES ---
		
		oneToOneGroup.GET("/report-to/all", func(c *gin.Context) {
			oneToOneHandler.GetAllWeeklyReportsForReportTo(c)
		})

		oneToOneGroup.GET("/report-to", func(c *gin.Context) {
			oneToOneHandler.GetWeeklyReportByWeekAndYearForReportTo(c)
		})

		oneToOneGroup.PUT("/report-to/update", func(c *gin.Context) {
			oneToOneHandler.UpdateWeeklyReportForReportTo(c)
		})
	}
}
