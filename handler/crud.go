package handler

import (
	"api-service/logger"
	"encoding/json"
	"net/http"
)

func GET(w http.ResponseWriter, r *http.Request) {
	loggerService := logger.GetLogger()
	config := getConfig(r.Context())
	service := getServices(r.Context())
	c, err := r.Cookie("jwt-Token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value
	if err := ValidateUser(sessionToken, config.Auth.Key); err != nil {
		loggerService.Errorln("error in validating user , Reason:", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		loggerService.Errorln("error in URL Query, Reason:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	mongoManager, err := service.MongoDBServices.MongoDBClient.GetManager(DB, testCollection)
	if err != nil {
		loggerService.Errorln("error in mongoManager, Reason:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	payload, err := mongoManager.GetDoc(id)
	if err != nil {
		loggerService.Errorln("error in get mongoManager, Reason:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	payloadByte, err := json.Marshal(&payload)
	if err != nil {
		loggerService.Errorln("error in json marsh , Reason:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(payloadByte)
	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	loggerService := logger.GetLogger()
	config := getConfig(r.Context())
	service := getServices(r.Context())
	c, err := r.Cookie("jwt-Token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value
	if err := ValidateUser(sessionToken, config.Auth.Key); err != nil {
		loggerService.Errorln("error in validating user , Reason:", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		loggerService.Errorln("error in URL Query, Reason:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	mongoManager, err := service.MongoDBServices.MongoDBClient.GetManager(DB, testCollection)
	if err != nil {
		loggerService.Errorln("error in mongoManager, Reason:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = mongoManager.DeleteDoc(id); err != nil {
		loggerService.Errorln("error in post mongoManager, Reason:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Successfully Deleted Data"))
	return
}
