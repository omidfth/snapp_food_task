package routers

import (
	"github.com/gin-gonic/gin"
	"snapp_food_task/info/handler/controllers"
)

type tripRouter struct {
	controller controllers.TripController
}

func NewTripRouter(controller controllers.TripController) Router {
	return &tripRouter{
		controller: controller,
	}
}

func (r tripRouter) HandleRoutes(router *gin.Engine) {
	ro := router.Group("trip")
	ro.POST("status", r.controller.ChangeStatus)
}
