package handler

import (
	"api-service/handler/login"
	"api-service/service"
)

type APIInterface struct {
	Auth login.AuthenticationInterface
}

func NewAPIInterface(s *service.APIServices) APIInterface {
	return APIInterface{
		Auth: login.NewHandlerLogin(s),
	}
}
