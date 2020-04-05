package rest

import (
	"github.com/gin-gonic/gin"
)

// Respons types
type Respons struct {
	Status  int         `json:"status"`
	Meta    int         `json:"meta"`
	Results interface{} `json:"results"`
}

// Response params
// @context: *gin.Context
// status: int
// payload: interface{}
func ResponseList(context *gin.Context, status int, payload interface{}) {
	var res Respons
	res.Status = status
	res.Results = payload
	res.Meta = 0
	context.JSON(status, res)
	return
}

// ResponseData params
// @context: *gin.Context
// status: int
// payload: interface{}
func ResponseData(context *gin.Context, status int, payload interface{}) {
	context.JSON(status, payload)
	return
}

// ResponseMessages params
// @context: *gin.Context
// status: int
// msg: string
func ResponseMessages(context *gin.Context, status int, msg string) {
	context.JSON(status, gin.H{
		"messages": msg,
	})
	return
}
