package routes

import (
	"eventapi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// registerForEvent registers a user for an event. The userId is required to be a valid user id in the event
//
// @param context - The Gin context to
func registerForEvent(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id."})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered"})

}

// CancelRegistration is a handler for cancellation of registration for an event.
//
// @param context - The Gin context to
func cancelRegistration(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, _ := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event

	event.ID = eventId

	err := event.CancelRegistration(userId)

	// Cancel registration for event.
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration for event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled"})

}
