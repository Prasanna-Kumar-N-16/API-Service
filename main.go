package main

import (
	"api-service/config"
	apiservice "api-service/http_service"
	"api-service/logger"
	service "api-service/schedule_service"
	"os"
	"os/signal"
	"syscall"
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
	defer func() {
		if r := recover(); r != nil {
			// Recovered from a panic, handle the error
			loggerService.Errorln("Recovered from panic:", r)
			os.Exit(1)
		}
	}()
	loggerService.Infoln("Http services started at port :", config.HttpConfig.Host)

	// Channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	loggerService.Infoln("Server started successfully")

	// Wait for termination signal
	<-quit
	loggerService.Infoln("Shutting down server...")

}
