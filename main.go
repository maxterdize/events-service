package main

import (
	"events-service/db"
	"events-service/routes"

	"github.com/gin-gonic/gin"
)

// @title	Events Service API
// @version 1.0
// @description A GoLang demo API to manage Events
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	db.InitDB()

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
