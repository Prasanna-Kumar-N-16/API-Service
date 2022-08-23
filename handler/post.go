package handler

import (
	"api-service/logger"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

const (
	DB             = "testDB"
	testCollection = "testCollection"
)

type Payload struct {
	ID           string `json:"_id" bson:"_id"`
	CompanyName  string `json:"companyName" bson:"companyName"`
	EmployeeName string `json:"employeeName" bson:"employeeName"`
}

func Post(w http.ResponseWriter, r *http.Request) {
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
	reqData := Payload{}
	err = json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		loggerService.Errorln("error in decoding data from body, Reason:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	mongoManager, err := service.MongoDBServices.MongoDBClient.GetManager(DB, testCollection)
	if err != nil {
		loggerService.Errorln("error in mongoManager, Reason:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = mongoManager.PostDoc(reqData); err != nil {
		loggerService.Errorln("error in post mongoManager, Reason:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Successfully Posted Data"))
	return
}

func PUT(w http.ResponseWriter, r *http.Request) {
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
	reqData := Payload{}
	err = json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		loggerService.Errorln("error in decoding data from body, Reason:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	mongoManager, err := service.MongoDBServices.MongoDBClient.GetManager(DB, testCollection)
	if err != nil {
		loggerService.Errorln("error in mongoManager, Reason:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = mongoManager.PutDoc(reqData, reqData.ID); err != nil {
		loggerService.Errorln("error in post mongoManager, Reason:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Successfully Updated/Created Data"))
	return
}

func ValidateUser(tokenString, secretKey string) error {
	key := []byte(secretKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error in signing method")
		}
		return key, nil
	})
	if err != nil {
		return errors.New("error in Parsing JWT")
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("error in token validation")
}
