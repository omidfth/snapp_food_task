package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"snapp_food_task/support/constants/commands"
	"snapp_food_task/support/internal/repositories/models"
	"snapp_food_task/support/internal/services"
	"strconv"
)

type AgentController interface {
	PushOrder(command models.AmqpModel)
	RemoveTask(ctx *gin.Context)
	AddTask(ctx *gin.Context)
}

type agentController struct {
	agentService services.AgentService
}

func NewAgentController(agentService services.AgentService) AgentController {
	return &agentController{agentService: agentService}
}

func (c agentController) PushOrder(data models.AmqpModel) {
	var command commands.AssignOrder
	d, _ := json.Marshal(data.Body)
	err := json.Unmarshal(d, &command)
	if err != nil {
		log.Println(err)
		return
	}
	c.agentService.PushOrder(command)
}

func (c agentController) AddTask(ctx *gin.Context) {
	agentIdString := ctx.Param("agent_id")
	agentID, err := strconv.Atoi(agentIdString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}
	agent, err := c.agentService.GetByID(uint(agentID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}
	err = c.agentService.AssignOrderToAgent(agent)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "succeed"})
}

func (c agentController) RemoveTask(ctx *gin.Context) {
	agentIdString := ctx.Param("agent_id")
	agentID, err := strconv.Atoi(agentIdString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}
	agent, err := c.agentService.GetByID(uint(agentID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}
	c.agentService.RemoveTask(agent)
	ctx.JSON(http.StatusOK, gin.H{"message": "succeed"})
}
