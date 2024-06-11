package routes

import (
	"eliarms.events.com/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getAllEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch All Events. Try again later."})
		return
	}
	if events == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No records Found"})
		return
	} else {
		c.JSON(200, events)
	}

}
func createEvent(c *gin.Context) {

	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Could not passed request Data": err.Error()})
		return
	}
	userId := c.GetInt64("userId")
	event.UserID = userId
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Create Events. Try again later."})
		return
	}
	c.JSON(201, gin.H{"message": "Event created!", "event": event})

}
func getEventById(c *gin.Context) {

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	if event == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No records Found"})
		return
	} else {
		c.JSON(200, event)
	}

}
func updateEvent(c *gin.Context) {

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	userId := c.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	if event.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
		return

	}
	var updateEvent models.Event
	err = c.ShouldBindJSON(&updateEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Could not passed request Data": err.Error()})
		return
	}
	updateEvent.ID = eventId
	err = updateEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event Updated Successfully!"})
}
func deleteEvent(c *gin.Context) {

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	userId := c.GetInt64("userId")
	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	if event.UserID != userId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event"})
		return

	}
	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event Deleted Successfully!"})
}
