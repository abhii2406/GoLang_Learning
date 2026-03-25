package main

import (
	"learning/rest-api_gorm/db"
	"learning/rest-api_gorm/models"
	"learning/rest-api_gorm/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Conect()
	db.DB.AutoMigrate(&models.User{}, &models.Event{})
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
