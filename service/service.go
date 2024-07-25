package service

import "api-service/models"

type APIServices struct {
	MongoDBServices MongoDBServices
}
type MongoDBServices struct {
	MongoDBClient *models.DBConnector
}
