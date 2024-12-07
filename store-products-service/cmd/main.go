package main

import (
	"log"
	"os"
	"product-service/internal/config"
	"product-service/internal/di"

	_ "product-service/docs"
)

// @title Product Service
// @version 1.0
// @description API Server for Product Service
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
