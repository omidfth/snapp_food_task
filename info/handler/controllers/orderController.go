package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"snapp_food_task/info/internal/services"
)

type OrderController interface {
	Append(ctx *gin.Context)
}

type orderController struct {
	orderService  services.OrderService
	vendorService services.VendorService
}

func NewOrderController(
	orderService services.OrderService,
	vendorService services.VendorService,
) OrderController {
	return &orderController{
		orderService:  orderService,
		vendorService: vendorService,
	}
}

func (c orderController) Append(ctx *gin.Context) {
	vendorID := ctx.Param("vendor_id")
	vendor, err := c.vendorService.GetByID(vendorID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}
	order := c.orderService.AppendNewOrder(vendor)
	ctx.JSON(http.StatusOK, order)
}
