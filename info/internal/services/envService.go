package services

import (
	"github.com/joho/godotenv"
	"log"
)

type envService struct {
	data map[string]string
}

type EnvService interface {
	Get(key string) string
}

func NewEnvService(path string) EnvService {
	err := godotenv.Load(path)
	if err != nil {
		log.Panicf("env not found! path:%s", path)
		return nil
	}
	appEnv, _ := godotenv.Read()
	return &envService{data: appEnv}
}

func (e *envService) Get(key string) string {
	val, ok := e.data[key]
	if ok {
		return val
	}
	return ""
}
