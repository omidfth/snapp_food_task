package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"snapp_food_task/info/commands"
	"snapp_food_task/info/constants/amqpKeys"
	"snapp_food_task/info/constants/errorKeys"
	"snapp_food_task/info/internal/repositories/models"
	"snapp_food_task/info/internal/services"
)

type DelayController interface {
	Report(ctx *gin.Context)
	FetchReports(ctx *gin.Context)
}

type delayController struct {
	delayService services.DelayService
	orderService services.OrderService
	tripService  services.TripService
	amqpService  services.AmqpService
}

func NewDelayController(
	delayService services.DelayService,
	orderService services.OrderService,
	tripService services.TripService,
	amqpService services.AmqpService,
) DelayController {
	return &delayController{
		delayService: delayService,
		orderService: orderService,
		tripService:  tripService,
		amqpService:  amqpService,
	}
}

func (c delayController) Report(ctx *gin.Context) {
	orderID := ctx.Param("order_id")
	order, err := c.orderService.GetByID(orderID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	trip, _ := c.tripService.GetByOrderID(order.ID)
	delayReport, estimateTime := c.delayService.Report(order, trip)

	req := commands.AssignOrder{OrderID: delayReport.OrderID}
	amqpModel := models.AmqpModel{
		Type: amqpKeys.ASSING_ORDER_TO_AGENT,
		Body: req,
	}
	j, _ := json.Marshal(amqpModel)
	c.amqpService.Publish("support", j)

	if estimateTime != nil {
		ctx.JSON(http.StatusOK, estimateTime)
		return
	} else if delayReport != nil {
		ctx.JSON(http.StatusOK, delayReport)
		return
	}
	ctx.JSON(http.StatusForbidden, errorKeys.FORBIDDEN_REPORT_DELAY)
}

func (c delayController) FetchReports(ctx *gin.Context) {
	reports := c.delayService.FetchReports()
	ctx.JSON(http.StatusOK, reports)
}
