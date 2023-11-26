package main

import (
	"flag"
	"log"
	"snapp_food_task/support/constants/envKeys"
	"snapp_food_task/support/handler/routers"
	"snapp_food_task/support/internal/database"
	"snapp_food_task/support/internal/services"
)

func main() {
	envPath := flag.String("env", envKeys.PATH, "env file path")
	envService := services.NewEnvService(*envPath)
	db, err := database.Connect(
		envService.Get(envKeys.DB_HOST),
		envService.Get(envKeys.DB_PORT),
		envService.Get(envKeys.DB_USERNAME),
		envService.Get(envKeys.DB_PASSWORD),
		envService.Get(envKeys.DB_NAME))
	if err != nil {
		log.Fatalf("%v", err)
	}
	routers.HandleRoutes(db, envService)
}
