package routes

import (
	"events-service/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetEvents	godoc
// @Summary		Get all events
// @Description	Get all events from the DB
// @Tags         Events
// @Produce		application/json
// @Event		event
// @Success		200 {object}	[]models.Event
// @Router		/events	[get]
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get events", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

// GetEvent	godoc
// @Summary		Get an even by ID
// @Description	Get one event from the DB by its ID
// @Tags         Events
// @Param		id path	int true "Event ID"
// @Produce		application/json
// @Event		event
// @Success		200 {object}	models.Event
// @Router		/events/{id}	[get]
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Params.ByName("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

// CreateEvent	godoc
// @Summary		Create events
// @Description	Create new events and save them to the DB
// @Tags         Events
// @Param		event body models.Event true "Create Event"
// @Produce		application/json
// @Event		event
// @Success		201 {object}	models.Event
// @Router		/events	[post]
// @Security ApiKeyAuth
func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event", "error": err.Error()})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

// UpdateEvent	godoc
// @Summary		Update event
// @Description	Updates an event by ID
// @Tags         Events
// @Param		id path	int true "Event ID"
// @Param		event body models.Event true "Update Event"
// @Produce		application/json
// @Event		event
// @Success		200 {object}	models.Event
// @Router		/events/{id}	[put]
// @Security ApiKeyAuth
func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Params.ByName("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event", "error": err.Error()})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to update this event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event", "error": err.Error()})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated", "event": updatedEvent})
}

// DeleteEvent	godoc
// @Summary		Deletes an even by ID
// @Description	Deletes one event from the DB by its ID
// @Tags         Events
// @Param		id path	int true "Event ID"
// @Produce		application/json
// @Event		event
// @Success		200
// @Router		/events/{id}	[delete]
// @Security ApiKeyAuth
func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Params.ByName("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event", "error": err.Error()})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to delete this event"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
