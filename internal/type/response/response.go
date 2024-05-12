package response

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SuccessResp(c *gin.Context, code int, message string, data any) {
	resp := Response{
		Success: true,
		Code:    code,
		Message: message,
		Data:    data,
	}
	c.JSON(code, resp)
}
func ErrResp(c *gin.Context, code int, message string) {
	resp := Response{
		Success: false,
		Code:    code,
		Message: message,
		Data:    nil,
	}
	c.JSON(code, resp)
}
