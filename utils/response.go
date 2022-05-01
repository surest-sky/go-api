package utils

import "github.com/gin-gonic/gin"

type Response struct{}

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) Success(c *gin.Context, data map[string]interface{}) {
	data["code"] = 200
	data["data"] = data
	data["message"] = "Success"
	c.JSON(200, data)
}

func (r *Response) Error(c *gin.Context, code int, data map[string]interface{}) {
	data["code"] = code
	data["data"] = data
	data["message"] = ""
	c.JSON(200, data)
}

func RSuccess(c *gin.Context, data map[string]interface{})  {
	response := NewResponse()
	response.Success(c, data)
}
