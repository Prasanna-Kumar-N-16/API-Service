package login

import (
	"api-service/config"
	"api-service/encryption"
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
	c       config.ConfigStruct
}

func NewHandlerLogin(service *service.APIServices, c config.ConfigStruct) AuthenticationInterface {
	return &Authenticationhandler{service: *service, c: c}
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
	dbOut, exist := h.service.PostgesQL.GetRecord(reqBody.Email)
	if !exist {
		logService.Errorln("error : Failed to get record for email :", reqBody.Email)
		utils.APIResponse(ctx, "Failed to get record for email :"+reqBody.Email, http.StatusInternalServerError, nil)
		return
	}
	outst := Admin{}
	if err := utils.MarshalUnmarshal(dbOut, &outst); err != nil {
		logService.Errorln("error : Failed to marshal and unmarshal struct :", reqBody.Email)
		utils.APIResponse(ctx, "Failed to marshal and unmarshal struct :"+reqBody.Email, http.StatusInternalServerError, nil)
		return
	}
	validUser, err := encryption.VerifyPassword(reqBody.Password, outst.Password, h.c.EncryptKey)
	if !validUser || err != nil {
		logService.Errorln("Email / Password entered is wrong")
		utils.APIResponse(ctx, "Email / Password entered is wrong", http.StatusUnauthorized, nil)
		return
	}
}
