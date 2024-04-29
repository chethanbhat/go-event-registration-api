package routes

import (
	"net/http"
	"strconv"

	"github.com/chethanbhat/go-rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {

	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	err = event.Register(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "successfully registered for the event"})

}

func cancelRegistration(context *gin.Context) {

	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	err = event.CancelRegistration(userID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel event registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "successfully cancelled the registration"})
}
