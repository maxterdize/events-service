package routes

import (
	_ "events-service/docs"
	"events-service/middlewares"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/attendees", registerForEvents)
	authenticated.DELETE("/events/:id/attendees", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
	// adding swagger
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
