package main

import (
	"log"
	"os"
	"users-service/internal/config"
	"users-service/internal/di"

	_ "users-service/docs"
)

// @title Users Service API
// @version 1.0
// @description API Server for Users Service
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
