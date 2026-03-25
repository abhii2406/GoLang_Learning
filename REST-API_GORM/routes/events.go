package routes

import (
	"learning/rest-api_gorm/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEventById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse to int"})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not get by id"})
		return
	}
	ctx.JSON(http.StatusOK, event)
}

func getEvent(ctx *gin.Context) {
	events, err := models.GetAllevents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {

	var event models.Event

	err := ctx.ShouldBindJSON(&event)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	uid := ctx.GetInt64("userId")
	// event.Id = 1
	event.UserId = uid
	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to save in db","error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"Event created :": event})

}

func updateEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse to int"})
		return
	}
	_, err = models.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch"})
		return
	}
	var updatedEvent models.Event

	err = ctx.ShouldBindJSON(&updatedEvent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed"})
		return
	}
	
	updatedEvent.Id = id
	err = updatedEvent.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedEvent)
}

func deleteEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse to int"})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch"})
		return
	}

	err = event.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"messege": "delete successfully"})

}
