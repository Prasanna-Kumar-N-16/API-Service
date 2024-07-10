package utils

import "github.com/gin-gonic/gin"

type Responses struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Data interface{}) {
	jsonResponse := Responses{
		StatusCode: StatusCode,
		Message:    Message,
		Data:       Data,
	}
	ctx.JSON(StatusCode, jsonResponse)

}

func AbortAPIResponse(ctx *gin.Context, Message string, StatusCode int, Data interface{}) {
	jsonResponse := Responses{
		StatusCode: StatusCode,
		Message:    Message,
		Data:       Data,
	}
	ctx.AbortWithStatusJSON(StatusCode, jsonResponse)

}
