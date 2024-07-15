package main

import (
	"log"
	"one-to-one/internal/config"
	"one-to-one/internal/db"
	"one-to-one/internal/pusher"
	"one-to-one/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title OneToOne API
// @version 1
// @description This is the REST API for OneToOne.
// @host one-to-one.backend.vercel.app
// @BasePath /
// @schemes https
// @securityDefinitions.apikey BasicAuth
// @in header
// @name Authorization
func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}

	db.ConnectToMongoDB()
	pusher.Init()
	defer db.DisconnectFromMongoDB()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		AllowWildcard:    true,
	}))
	routes.SetupRoutes(router)
	router.Run(config.AppConfig().App.Host + ":" + config.AppConfig().App.Port)
}
