package apiservice

import (
	"api-service/config"
	"api-service/handler"
	"api-service/service"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

type Api struct {
}

func (a Api) SetRouter(config *config.ConfigStruct, apiServices *service.APIServices) *gin.Engine {
	// Create a new Gin router with default middleware
	r := gin.New()

	gin.SetMode(gin.DebugMode)

	r.Use(helmet.Default())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowWildcard:    true,
	}))

	api := r.Group("/api/v1")

	apiInterface := handler.NewAPIInterface(apiServices, *config)

	// Middleware to set a key-value pair in the context
	api.Use(func(c *gin.Context) {
		c.Set("config", config)
		c.Set("services", apiServices)
		c.Next()
	})

	api.POST("/login", apiInterface.Auth.LoginHandler)
	api.POST("/signup", apiInterface.Auth.Signup)

	return r

}
