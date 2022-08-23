package service

import "api-service/mongodb"

type APIServices struct {
	MongoDBServices MongoDBServices
}
type MongoDBServices struct {
	MongoDBClient *mongodb.DBConnector
}
