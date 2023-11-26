package routers

import (
	"github.com/gin-gonic/gin"
	"snapp_food_task/info/handler/controllers"
)

type orderRouter struct {
	controller controllers.OrderController
}

func NewOrderRouter(controller controllers.OrderController) Router {
	return &orderRouter{
		controller: controller,
	}
}

func (r orderRouter) HandleRoutes(router *gin.Engine) {
	ro := router.Group("order")
	ro.GET("append/:vendor_id", r.controller.Append)
}
