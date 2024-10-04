package service

import (
	k "api-service/kafka_service"
	"api-service/models"
)

type APIServices struct {
	MongoDBServices MongoDBServices
	PostgesQL       *models.Service
	KafkaService    *k.KService
}
type MongoDBServices struct {
	MongoDBClient *models.DBConnector
}
