package routes

import (
	"events-service/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// RegisterForEvent	godoc
// @Summary		Register a user for an event
// @Description	Create new user registration for events and save them to the DB
// @Tags         Registrations
// @Param		id path	int true "Event ID"
// @Produce		application/json
// @Success		201
// @Router		/events/{id}/attendees	[post]
// @Security ApiKeyAuth
func registerForEvents(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event", "error": err.Error()})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered for event successfully"})
}

// CancelRegistration	godoc
// @Summary		Cancel event registration
// @Description	Cancel an event registration and delete them from the DB
// @Tags         Registrations
// @Param		id path	int true "Event ID"
// @Produce		application/json
// @Success		201
// @Router		/events/{id}/attendees	[delete]
// @Security ApiKeyAuth
func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled successfully"})
}
