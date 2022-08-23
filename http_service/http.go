package apiservice

import (
	"context"
	"net/http"

	"api-service/config"
	"api-service/handler"
	"api-service/service"

	"github.com/gorilla/mux"
)

type Api struct {
	router *mux.Router
}

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (apimux Api) Initialize(config *config.ConfigStruct, apiServices *service.APIServices) error {
	apimux.router = mux.NewRouter()
	apimux.SetRouter(config, apiServices)
	if err := apimux.Run(*config.HttpConfig); err != nil {
		return err
	}
	return nil
}

func (apimux Api) SetRouter(config *config.ConfigStruct, apiServices *service.APIServices) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "config", config)
	ctx = context.WithValue(ctx, "services", apiServices)
	apimux.Post("/login", apimux.handleRequest(handler.Login, ctx))
	apimux.Post("/post", apimux.handleRequest(handler.Post, ctx))
	apimux.GET("/get", apimux.handleRequest(handler.GET, ctx))
	apimux.PUT("/put", apimux.handleRequest(handler.PUT, ctx))
	apimux.DELETE("/delete", apimux.handleRequest(handler.Delete, ctx))
}

func (apimux Api) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	apimux.router.HandleFunc(path, f).Methods("POST")
}
func (apimux Api) GET(path string, f func(w http.ResponseWriter, r *http.Request)) {
	apimux.router.HandleFunc(path, f).Methods("GET")
}
func (apimux Api) PUT(path string, f func(w http.ResponseWriter, r *http.Request)) {
	apimux.router.HandleFunc(path, f).Methods("PUT")
}
func (apimux Api) DELETE(path string, f func(w http.ResponseWriter, r *http.Request)) {
	apimux.router.HandleFunc(path, f).Methods("DELETE")
}

func (a *Api) handleRequest(handler RequestHandlerFunction, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r.WithContext(ctx))
	}
}
func (a *Api) Run(apiConfig config.Service) error {
	if apiConfig.ISSecureConnection {
		if err := http.ListenAndServeTLS(apiConfig.Host, apiConfig.SSLConfig.CrtFile, apiConfig.SSLConfig.PrivateKey, a.router); err != nil {
			return err
		}
	} else {
		if err := http.ListenAndServe(apiConfig.Host, a.router); err != nil {
			return err
		}
	}
	return nil
}
