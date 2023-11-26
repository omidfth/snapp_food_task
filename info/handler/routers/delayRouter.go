package routers

import (
	"github.com/gin-gonic/gin"
	"snapp_food_task/info/handler/controllers"
)

type delayRouter struct {
	controller controllers.DelayController
}

func NewDelayRouter(controller controllers.DelayController) Router {
	return &delayRouter{
		controller: controller,
	}
}

func (r delayRouter) HandleRoutes(router *gin.Engine) {
	ro := router.Group("delay")
	ro.GET("report/:order_id", r.controller.Report)
	ro.GET("fetch", r.controller.FetchReports)
}
