package pkg

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int
	Message string
	Data    *interface{}
}

func NewResponse(status int, message string) *Response {
	return &Response{
		Status:  status,
		Message: message,
	}
}

func (response *Response) Body(data interface{}) *Response {
	response.Data = &data

	return response
}

func (response *Response) Build() (status int, obj any) {
	return response.Status, gin.H{
		"message": response.Message,
		"success": response.Status >= 200 && response.Status < 300,
		"data":    response.Data,
	}
}
