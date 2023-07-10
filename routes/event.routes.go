package routes

import (
	"jejetickets/controllers"
	"jejetickets/middleware"

	"github.com/gin-gonic/gin"
)

type EventRouteController struct {
	eventController controllers.EventController
}

func NewRouteEventController(eventController controllers.EventController) EventRouteController {
	return EventRouteController{eventController}
}

func (ec *EventRouteController) EventRoute(rg *gin.RouterGroup) {

	router := rg.Group("events")
	router.Use(middleware.DeserializeUser())
	router.POST("/", ec.eventController.CreateEvent)
	router.GET("/", ec.eventController.FindEvents)
	router.PUT("/:eventId", ec.eventController.FindEventById)
	router.DELETE("/:eventId", ec.eventController.DeleteEvent)
}
