package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"snapp_food_task/support/constants/amqpKeys"
	"snapp_food_task/support/constants/envKeys"
	"snapp_food_task/support/handler/controllers"
	"snapp_food_task/support/internal/repositories"
	"snapp_food_task/support/internal/services"
	"snapp_food_task/support/producer"
)

type Router interface {
	HandleRoutes(router *gin.Engine)
}

func HandleRoutes(db *gorm.DB, envService services.EnvService) {
	router := gin.Default()
	//--------------Producer---------------//
	redisProducer := producer.NewRedisProducer(envService.Get(envKeys.REDIS_HOST), envService.Get(envKeys.REDIS_PORT))

	//--------------Repositories---------------//
	agentRepo := repositories.NewAgentRepository(db, redisProducer)
	if err := agentRepo.Migrate(); err != nil {
		log.Fatal(err)
	}

	mockRepo := repositories.NewMockRepository(db)
	mockRepo.CreateDefaultRows()
	//--------------Services---------------//
	amqpService := services.NewAmqpService(envService.Get(envKeys.AMQP_SERVER))
	agentSrv := services.NewAgentService(agentRepo)

	//--------------Controllers---------------//
	agentCtrl := controllers.NewAgentController(agentSrv)

	amqpService.On(amqpKeys.ASSING_ORDER_TO_AGENT, agentCtrl.PushOrder)

	//--------------Routers---------------//
	agentR := NewAgentRouter(agentCtrl)
	agentR.HandleRoutes(router)

	//--------------Server---------------//
	host := envService.Get(envKeys.SERVER_HOST)
	port := envService.Get(envKeys.SERVER_PORT)
	server := fmt.Sprintf("%s:%s", host, port)
	amqpService.Serve("support")
	if err := router.Run(server); err != nil {
		log.Fatal(err)
	}
}
