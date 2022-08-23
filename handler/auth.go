package handler

import (
	"api-service/config"
	"api-service/service"
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Login(w http.ResponseWriter, r *http.Request) {
	config := getConfig(r.Context())
	//Can improvize by string login data in DB and authenticate
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userName":   "API Service",
		"expiryTime": time.Now().Add(24 * time.Hour).UnixMilli(),
		"IssuedAt":   time.Now().UnixMilli(),
		"Subject":    "API DESC",
	})
	tokenString, err := token.SignedString([]byte(config.Auth.Key))
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	cookie := &http.Cookie{
		Name:   "jwt-Token",
		Value:  tokenString,
		MaxAge: 1,
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Successfully authenticated"))
	return
}

func getConfig(ctx context.Context) *config.ConfigStruct {
	value := ctx.Value("config")
	if value == nil {
		return nil
	}
	return value.(*config.ConfigStruct)
}

func getServices(ctx context.Context) *service.APIServices {
	value := ctx.Value("services")
	if value == nil {
		return nil
	}
	return value.(*service.APIServices)
}
