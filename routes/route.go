package routes

import (
	"example.com/event-booking/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    // GET, POST, PUT, DELETE
	server.GET("/events/:id", getEvent) //events/1 /events/1

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvent)       //Creation
	authenticated.PUT("/events/:id", updateEvent)    //Updation
	authenticated.DELETE("/events/:id", deleteEvent) //Deletion
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
