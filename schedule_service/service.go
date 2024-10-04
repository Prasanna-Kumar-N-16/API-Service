package schedule_service

import (
	"api-service/config"
	k "api-service/kafka_service"
	"api-service/service"
)

func StartServices(c *config.ConfigStruct) (*service.APIServices, error) {
	apiServices := &service.APIServices{}
	if mongodbClient, err := c.MongoDBConfig.NewConnection(); err != nil {
		return nil, err
	} else {
		apiServices.MongoDBServices.MongoDBClient = mongodbClient
	}
	if postService, err := c.PostgresQL.NewService(); err != nil {
		return nil, err
	} else {
		apiServices.PostgesQL = postService
	}
	ks := k.NewKService()
	//TODO add config
	if err := ks.NewConsumer("", "", ""); err != nil {
		return nil, err
	}
	if err := ks.NewProducer(""); err != nil {
		return nil, err
	}
	return apiServices, nil
}
