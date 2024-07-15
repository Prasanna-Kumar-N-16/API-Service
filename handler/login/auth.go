package login

import (
	"api-service/logger"
	"api-service/service"
	"api-service/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthenticationInterface interface {
	LoginHandler(*gin.Context)
	Signup(*gin.Context)
}

type Authenticationhandler struct {
	service service.APIServices
}

func NewHandlerLogin(service *service.APIServices) AuthenticationInterface {
	return &Authenticationhandler{service: *service}
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

const (
	reqBodyParseErr = "request body parsing error"
)

func (h *Authenticationhandler) LoginHandler(ctx *gin.Context) {
	var reqBody UserLogin
	logService := logger.GetLogger()
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		logService.Errorln("error in API request body parsing reason:", err)
		utils.APIResponse(ctx, reqBodyParseErr, http.StatusBadRequest, nil)
		return
	}
	reqBody.Password = strings.ReplaceAll(reqBody.Password, " ", "")
	if reqBody.Email == "" || reqBody.Password == "" {
		logService.Errorln("empty request body param")
		utils.APIResponse(ctx, reqBodyParseErr, http.StatusBadRequest, nil)
		return
	}
	if !strings.Contains(reqBody.Email, "@") && !strings.Contains(reqBody.Email, ".") {
		logService.Errorln("Email format Invalid")
		utils.APIResponse(ctx, reqBodyParseErr, http.StatusBadRequest, nil)
		return
	}
}

func (h *Authenticationhandler) Signup(*gin.Context) {

}
