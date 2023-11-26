package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"snapp_food_task/info/constants/envKeys"
	"snapp_food_task/info/handler/controllers"
	"snapp_food_task/info/internal/repositories"
	"snapp_food_task/info/internal/services"
)

type Router interface {
	HandleRoutes(router *gin.Engine)
}

func HandleRoutes(db *gorm.DB, envService services.EnvService) {
	router := gin.Default()

	//--------------Repositories---------------//

	delayRepo := repositories.NewDelayRepository(db)
	if err := delayRepo.Migrate(); err != nil {
		log.Fatal(err)
	}

	orderRepo := repositories.NewOrderRepository(db)
	if err := orderRepo.Migrate(); err != nil {
		log.Fatal(err)
	}

	tripRepo := repositories.NewTripRepository(db)
	if err := tripRepo.Migrate(); err != nil {
		log.Fatal(err)
	}

	vendorRepo := repositories.NewVendorRepository(db)
	if err := vendorRepo.Migrate(); err != nil {
		log.Fatal(err)
	}

	mockRepo := repositories.NewMockRepository(db)
	mockRepo.CreateDefaultRows()
	//--------------Services---------------//
	amqpService := services.NewAmqpService(envService.Get(envKeys.AMQP_SERVER))
	delaySrv := services.NewDelayService(delayRepo)
	orderSrv := services.NewOrderService(orderRepo, tripRepo)
	tripSrv := services.NewTripService(tripRepo, orderRepo)
	vendorSrv := services.NewVendorService(vendorRepo)

	//--------------Controllers---------------//
	delayCtrl := controllers.NewDelayController(delaySrv, orderSrv, tripSrv, amqpService)
	orderCtrl := controllers.NewOrderController(orderSrv, vendorSrv)
	tripCtrl := controllers.NewTripController(tripSrv)

	//--------------Routers---------------//
	delayR := NewDelayRouter(delayCtrl)
	orderR := NewOrderRouter(orderCtrl)
	tripR := NewTripRouter(tripCtrl)

	delayR.HandleRoutes(router)
	orderR.HandleRoutes(router)
	tripR.HandleRoutes(router)

	//--------------Server---------------//
	host := envService.Get(envKeys.SERVER_HOST)
	port := envService.Get(envKeys.SERVER_PORT)
	server := fmt.Sprintf("%s:%s", host, port)
	amqpService.Serve("info")
	if err := router.Run(server); err != nil {
		log.Fatal(err)
	}
}
