package main

import (
	"log"

	"github.com/hokkyar/rest-bookapp/src/config"
	"github.com/hokkyar/rest-bookapp/src/models"
	"github.com/hokkyar/rest-bookapp/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	config.InitDB()

	// Initialize Redis
	config.InitRedis()

	// Auto Migrate Database
	db := config.GetDB()
	db.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.UserFavorite{},
		&models.UserComment{},
	)

	// Initialize Router
	r := gin.Default()

	routes.InitRoutes(r)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
