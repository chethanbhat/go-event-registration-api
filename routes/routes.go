package routes

import (
	"github.com/chethanbhat/go-rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Events
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEventByID)
	authenticated.DELETE("/events/:id", deleteEventByID)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/cancel", cancelRegistration)

	// Users
	server.POST("/signup", signup)
	server.POST("/login", login)
}
