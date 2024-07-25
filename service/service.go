package service

import "api-service/models"

type APIServices struct {
	MongoDBServices MongoDBServices
	PostgesQL       *models.Service
}
type MongoDBServices struct {
	MongoDBClient *models.DBConnector
}
