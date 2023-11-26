package repositories

import (
	"gorm.io/gorm"
	"snapp_food_task/support/internal/repositories/models"
	"strconv"
)

type MockRepository interface {
	CreateDefaultRows()
}

type mockRepository struct {
	db *gorm.DB
}

func NewMockRepository(db *gorm.DB) MockRepository {
	return &mockRepository{db: db}
}

func (r mockRepository) CreateDefaultRows() {
	for i := 0; i < 10; i++ {
		name := "agent_" + strconv.Itoa(i+1)
		agent := models.Agent{
			Name: name,
		}
		r.db.Where("name=?", name).FirstOrCreate(&agent)
	}
}
