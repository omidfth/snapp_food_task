package routers

import (
	"github.com/gin-gonic/gin"
	"snapp_food_task/support/handler/controllers"
)

type agentRouter struct {
	controller controllers.AgentController
}

func NewAgentRouter(controller controllers.AgentController) Router {
	return &agentRouter{
		controller: controller,
	}
}

func (r agentRouter) HandleRoutes(router *gin.Engine) {
	ro := router.Group("agent")
	ro.GET("/add/:agent_id", r.controller.AddTask)
	ro.GET("/remove/:agent_id", r.controller.RemoveTask)
}
