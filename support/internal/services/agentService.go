package services

import (
	"snapp_food_task/support/constants/commands"
	"snapp_food_task/support/internal/repositories"
	"snapp_food_task/support/internal/repositories/models"
)

type AgentService interface {
	GetByID(id uint) (*models.Agent, error)
	PushOrder(command commands.AssignOrder)
	AssignOrderToAgent(agent *models.Agent) error
	RemoveTask(agent *models.Agent)
}

type agentService struct {
	repository repositories.AgentRepository
}

func NewAgentService(repository repositories.AgentRepository) AgentService {
	return &agentService{repository: repository}
}

func (s agentService) PushOrder(command commands.AssignOrder) {
	s.repository.PushOrder(command)
}

func (s agentService) RemoveTask(agent *models.Agent) {
	s.repository.RemoveTask(agent)
}

func (s agentService) GetByID(id uint) (*models.Agent, error) {
	return s.repository.GetAgentByID(id)
}

func (s agentService) AssignOrderToAgent(agent *models.Agent) error {
	return s.repository.AssignOrderToAgent(agent)
}
