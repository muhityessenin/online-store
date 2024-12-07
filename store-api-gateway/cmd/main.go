package main

import (
	"api-gateway-service/internal/config"
	"api-gateway-service/internal/di"
	"log"
	"os"

	_ "api-gateway-service/docs"
)

// @title API Gateway Service
// @version 1.0
// @description API Server for Online Store
// @BasePath /api
func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal(configErr)
	}

	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal(diErr)
	} else {
		server.Run(infoLog, errorLog)
	}
}
