package routes

import (
	"eliarms.events.com/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// API v1
	v1 := server.Group("/api/v1")
	{
		v1.GET("events", getAllEvents)
		v1.GET("events/:id", getEventById)

		authenticated := server.Group("/api/v1")
		authenticated.Use(middlewares.Authenticate)
		authenticated.POST("events", createEvent)
		authenticated.PUT("event/:id", updateEvent)
		authenticated.DELETE("event/:id", deleteEvent)
		authenticated.POST("events/:id/register", registerForEvent)
		authenticated.DELETE("events/:id/register", cancelRegistration)

		v1.POST("/signup", signup)
		v1.POST("/login", login)

	}
}
