package dashboard

import (
	"api-service/config"
	"api-service/service"
)

type ReportsInterface interface {
}

type Reportshandler struct {
	service service.APIServices
	c       config.ConfigStruct
}

func NewReportsInt(service *service.APIServices, c config.ConfigStruct) ReportsInterface {
	return &Reportshandler{service: *service, c: c}
}
