package main

import (
	"api-service/config"
	apiservice "api-service/http_service"
	"api-service/logger"
	service "api-service/schedule_service"
)

func main() {
	loggerService := logger.StartLogger()
	config, err := config.Parse()
	if err != nil {
		return
	}
	ediService, err := service.StartServices(config)
	if err != nil {
		loggerService.Panic("Unable to start the services. Reason: ", err)
	}
	loggerService.Infoln("services started successfully")
	app := apiservice.Api{}
	if err = app.Initialize(config, ediService); err != nil {
		loggerService.Errorln("unable to start http service at port :", config.HttpConfig.Host)
		return
	}
	loggerService.Infoln("Http services started at port :", config.HttpConfig.Host)
}
