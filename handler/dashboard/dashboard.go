package dashboard

import (
	"api-service/config"
	"api-service/service"

	"github.com/gin-gonic/gin"
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

}
