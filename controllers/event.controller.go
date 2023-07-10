package controllers

import (
	"jejetickets/models"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EventController struct {
	DB *gorm.DB
}

func NewEventController(DB *gorm.DB) EventController {
	return EventController{DB}
}

func (ec *EventController) CreateEvent(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.EventInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newEvent := models.Event{
		Title:           payload.Title,
		Description:     payload.Description,
		Location:        payload.Location,
		CoverPhotoURL:   payload.CoverPhotoURL,
		Tickets:         payload.Tickets,
		TicketDetails:   payload.TicketDetails,
		PrivacySettings: payload.PrivacySettings,
		GuestInfo:       payload.GuestInfo,
		User:            currentUser.ID,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	result := ec.DB.Create(&newEvent)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "Duplicate Key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Event Name matches another!"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newEvent})
}

// Get Single Event Handler
func (ec *EventController) FindEventById(ctx *gin.Context) {
	eventId := ctx.Param("eventId")

	var event models.Event
	result := ec.DB.First(&event, "id = ?", eventId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No Event With That Title Exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": event})
}

// Get All Events Handler
func (ec *EventController) FindEvents(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var events []models.Event
	results := ec.DB.Limit(intLimit).Offset(offset).Find(&events)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(events), "data": events})
}

// Delete Event Handler
func (ec *EventController) DeleteEvent(ctx *gin.Context) {
	eventId := ctx.Param("eventId")

	result := ec.DB.Delete(&models.Event{}, "id = ?", eventId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No Event With That Title Exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
