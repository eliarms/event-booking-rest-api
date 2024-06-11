package routes

import (
	"eliarms.events.com/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch Event"})
		return
	}
	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could register user for the Event"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "The User has been successfully registered for the event!"})
}
func cancelRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	var event models.Event
	event.ID = eventId
	event.CancelRegistration(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "The User has been successfully deregistered from the event!"})
}
