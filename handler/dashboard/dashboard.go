package dashboard

import (
	"api-service/config"
	"api-service/logger"
	"api-service/service"
	"api-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	reqBodyParseErr = "request body parsing error"
)

type ReportsInterface interface {
	DashboardHandler(*gin.Context)
}

type Reportshandler struct {
	service service.APIServices
	c       config.ConfigStruct
}

func NewReportsInt(service *service.APIServices, c config.ConfigStruct) ReportsInterface {
	return &Reportshandler{service: *service, c: c}
}

func (h *Reportshandler) DashboardHandler(c *gin.Context) {
	logService := logger.GetLogger()
	id := c.Query("id")
	if id == "" {
		logService.Errorln("empty id")
		utils.APIResponse(c, reqBodyParseErr, http.StatusBadRequest, nil)
		return
	}
}
