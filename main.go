package main

import (
	"log"

	"github.com/Akaame/golang-action-test/dao"
	"go.uber.org/zap"
)

func main() {
	config := zap.NewProductionConfig()
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("%v", err)
	}
	logger.Info("Logger setup successful")

	client, err := dao.NewClient("cassandra")
	if err != nil {
		log.Fatalf("Error during cassandra client initialization: %v", err)
	}
	defer client.Close()

}
