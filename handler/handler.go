package handler

import (
	"api-service/config"
	"api-service/handler/login"
	"api-service/service"
)

type APIInterface struct {
	Auth login.AuthenticationInterface
}

func NewAPIInterface(s *service.APIServices, c config.ConfigStruct) APIInterface {
	return APIInterface{
		Auth: login.NewHandlerLogin(s, c),
	}
}
