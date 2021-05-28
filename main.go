package main

import (
	"log"

	"go.uber.org/zap"
)

func main() {
	config := zap.NewProductionConfig()
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("%v", err)
	}
	logger.Info("Logger setup successful")
}
