package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snapp_food_task/info/commands"
	"snapp_food_task/info/internal/repositories/models"
	"snapp_food_task/info/internal/services"
)

type TripController interface {
	ChangeStatus(ctx *gin.Context)
}

type tripController struct {
	tripService  services.TripService
	orderService services.OrderService
}

func NewTripController(tripService services.TripService) TripController {
	return &tripController{tripService: tripService}
}

func (c tripController) ChangeStatus(ctx *gin.Context) {
	var command commands.ChangeStatus
	if err := ctx.ShouldBindJSON(&command); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}
	trip, err := c.tripService.GetByID(command.TripID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}
	status := models.TripStatus(command.Status)
	data, err := c.tripService.ChangeStatus(trip, status)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, data)
}
