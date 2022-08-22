package schedule_service

import (
	"api-service/config"
	"api-service/service"
)

func StartServices(c *config.ConfigStruct) (*service.APIServices, error) {
	apiServices := &service.APIServices{}
	return apiServices, nil
}
