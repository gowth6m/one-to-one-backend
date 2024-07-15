package routes

import (
	_ "one-to-one/docs"

	"github.com/gin-gonic/gin"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SwaggerRoutes(router *gin.Engine) {
	router.GET("/swagger/*any", gin.WrapH(httpSwagger.WrapHandler))
}
