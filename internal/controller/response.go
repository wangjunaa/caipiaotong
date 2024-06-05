package controller

import "github.com/gin-gonic/gin"

type respData struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
type Resp interface {
	Success(c *gin.Context, msg string, data any)
	Error(c *gin.Context, code int, err error)
}
type resp struct{}

func NewResp() Resp {
	return &resp{}
}

func (r *resp) Success(c *gin.Context, msg string, data any) {
	rd := respData{
		Success: true,
		Code:    200,
		Message: msg,
		Data:    data,
	}
	c.JSON(200, rd)
}

func (r *resp) Error(c *gin.Context, code int, err error) {
	rd := respData{
		Success: false,
		Code:    code,
		Message: err.Error(),
		Data:    nil,
	}
	c.JSON(code, rd)
}
