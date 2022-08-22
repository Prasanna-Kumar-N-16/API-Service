package main

import (
	"api-service/config"
	"api-service/logger"
)

func main() {
	loggerService := logger.StartLogger()
	config, err := config.Parse()
	if err != nil {
		return
	}
	loggerService.Infoln(config)
}
