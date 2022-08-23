package schedule_service

import (
	"api-service/config"
	"api-service/service"
)

func StartServices(c *config.ConfigStruct) (*service.APIServices, error) {
	apiServices := &service.APIServices{}
	if mongodbClient, err := c.MongoDBConfig.NewConnection(); err != nil {
		return nil, err
	} else {
		apiServices.MongoDBServices.MongoDBClient = mongodbClient
	}
	return apiServices, nil
}
