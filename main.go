package main

import (
	"api-service/config"
	apiservice "api-service/http_service"
	"api-service/logger"
	service "api-service/schedule_service"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	loggerService := logger.StartLogger()
	config, err := config.Parse()
	if err != nil {
		loggerService.Panic("Unable to parse config . Reason: ", err)
		return
	}
	apiService, err := service.StartServices(config)
	if err != nil {
		loggerService.Panic("Unable to start the services. Reason: ", err)
		return
	}
	loggerService.Infoln("services started successfully")
	defer func() {
		if r := recover(); r != nil {
			// Recovered from a panic, handle the error
			loggerService.Errorln("Recovered from panic:", r)
			os.Exit(1)
		}
	}()

	app := apiservice.Api{}
	r := app.SetRouter(config, apiService)
	// Create an HTTP server
	server := &http.Server{
		Addr:    config.HttpConfig.Host,
		Handler: r,
	}

	// Channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// Start server in a separate goroutine
	go func() {
		var err error
		if config.HttpConfig.ISSecureConnection {
			err = server.ListenAndServeTLS(config.HttpConfig.SSLConfig.CrtFile, config.HttpConfig.SSLConfig.PrivateKey)
		} else {
			err = server.ListenAndServe()
		}

		if err != nil && err != http.ErrServerClosed {
			loggerService.Errorln("Error in HTTP server:", err)
		}
	}()
	loggerService.Infoln("Server started successfully")

	// Wait for termination signal
	<-quit
	loggerService.Infoln("Shutting down server...")

	// Create a context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		loggerService.Errorln("Error during server shutdown:", err)
	} else {
		loggerService.Infoln("Server shut down gracefully")
	}

}
