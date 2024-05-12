package response

import "github.com/gin-gonic/gin"

type Response struct {
	success bool
	code    int
	message string
	data    any
}

func SuccessResp(c *gin.Context, code int, message string, data any) {
	resp := Response{
		success: true,
		code:    code,
		message: message,
		data:    data,
	}
	c.JSON(code, resp)
}
func ErrResp(c *gin.Context, code int, message string) {
	resp := Response{
		success: false,
		code:    code,
		message: message,
	}
	c.JSON(code, resp)
}
