package login

import (
	"api-service/service"

	"github.com/gin-gonic/gin"
)

type AuthenticationInterface interface {
	LoginHandler(*gin.Context)
}

type Authenticationhandler struct {
	service service.APIServices
}

func NewHandlerLogin(service *service.APIServices) AuthenticationInterface {
	return &Authenticationhandler{service: *service}
}

func (h *Authenticationhandler) LoginHandler(ctx *gin.Context) {

}
