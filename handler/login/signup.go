package login

import (
	"api-service/logger"
	"api-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminSignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func (h *Authenticationhandler) Signup(ctx *gin.Context) {
	var req AdminSignupRequest
	logService := logger.GetLogger()
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logService.Errorln("error in Signup API request body parsing reason:", err)
		utils.APIResponse(ctx, reqBodyParseErr, http.StatusBadRequest, nil)
		return
	}
}
