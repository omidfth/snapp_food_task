package repositories

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"snapp_food_task/support/constants/commands"
	"snapp_food_task/support/constants/errorKeys"
	"snapp_food_task/support/constants/redisKeys"
	"snapp_food_task/support/internal/repositories/models"
	"snapp_food_task/support/producer"
	"time"
)

type AgentRepository interface {
	Migrate() error
	PushOrder(command commands.AssignOrder)
	RemoveTask(agent *models.Agent)
	AssignOrderToAgent(agent *models.Agent) error
	GetAgentByID(id uint) (*models.Agent, error)
}

type agentRepository struct {
	db            *gorm.DB
	redisProducer producer.RedisProducer
}

func NewAgentRepository(db *gorm.DB, redisProducer producer.RedisProducer) AgentRepository {
	return &agentRepository{db: db, redisProducer: redisProducer}
}

func (r agentRepository) Migrate() error {
	return r.db.AutoMigrate(models.Agent{})
}

func (r agentRepository) GetAgentByID(id uint) (*models.Agent, error) {
	var agent models.Agent
	if r.db.Where("id=?", id).First(&agent).RowsAffected < 1 {
		return nil, errors.New(errorKeys.AGENT_NOT_FOUND)
	}

	return &agent, nil
}

func (r agentRepository) RemoveTask(agent *models.Agent) {
	agent.OrderID = 0
	agent.HasOrder = false
	r.db.Save(&agent)
}

func (r agentRepository) PushOrder(command commands.AssignOrder) {
	r.saveOrders(command)
}

func (r agentRepository) AssignOrderToAgent(agent *models.Agent) error {
	if agent.HasOrder {
		return errors.New(errorKeys.FORBIDDEN_ASSIGN_NEW_ORDER)
	}
	firstOrder := r.getFirstOrder()
	if firstOrder == nil {
		return errors.New(errorKeys.ORDER_NOT_FOUND)
	}
	if r.db.Debug().Where("order_id=?", firstOrder.OrderID).Find(&models.Agent{}).RowsAffected > 0 {
		return errors.New(errorKeys.HAS_BEEN_ASSIGNED)
	}
	agent.OrderID = firstOrder.OrderID
	agent.HasOrder = true
	r.removeFirstOrder()
	r.db.Save(&agent)
	return nil
}

func (r agentRepository) removeFirstOrder() {
	ordersStr := r.redisProducer.Get(redisKeys.ORDERS)
	var orders []commands.AssignOrder
	err := json.Unmarshal([]byte(ordersStr), &orders)
	if err != nil {
		return
	}
	if len(orders) < 1 {
		orders = make([]commands.AssignOrder, 0)
	} else {
		orders = orders[1:]
	}
	j, _ := json.Marshal(orders)
	r.redisProducer.Set(redisKeys.ORDERS, string(j), time.Duration(168)*time.Hour)
}

func (r agentRepository) getFirstOrder() *commands.AssignOrder {
	ordersStr := r.redisProducer.Get(redisKeys.ORDERS)
	var orders []commands.AssignOrder
	json.Unmarshal([]byte(ordersStr), &orders)
	if len(orders) < 1 {
		return nil
	}
	return &orders[0]
}

// saveOrders returns first order
func (r agentRepository) saveOrders(command commands.AssignOrder) *commands.AssignOrder {
	if r.db.Debug().Where("order_id=?", command.OrderID).Find(&models.Agent{}).RowsAffected > 0 {
		return nil
	}
	ordersStr := r.redisProducer.Get(redisKeys.ORDERS)
	if ordersStr == "" {
		return r.addCommand(command)
	}
	var orders []commands.AssignOrder
	err := json.Unmarshal([]byte(ordersStr), &orders)
	if err != nil {
		return nil
	}

	if len(orders) < 1 {
		return r.addCommand(command)
	}

	if r.hasOrder(orders, command) {
		return nil
	}

	orders = append(orders, command)
	j, _ := json.Marshal(orders)
	r.redisProducer.Set(redisKeys.ORDERS, string(j), time.Duration(168)*time.Hour)
	return &orders[0]
}

func (r agentRepository) addCommand(command commands.AssignOrder) *commands.AssignOrder {
	orders := []commands.AssignOrder{command}
	j, _ := json.Marshal(orders)
	r.redisProducer.Set(redisKeys.ORDERS, string(j), time.Duration(168)*time.Hour)
	return &orders[0]
}

func (r agentRepository) hasOrder(orders []commands.AssignOrder, order commands.AssignOrder) bool {
	for _, o := range orders {
		if o.OrderID == order.OrderID {
			return true
		}
	}
	return false
}
