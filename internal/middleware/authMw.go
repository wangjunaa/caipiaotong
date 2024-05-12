package middleware

import (
	"caipiaotong/configs/constant"
	"caipiaotong/internal/service"
	"caipiaotong/internal/type/response"
	"context"
	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

// MwAuthentication 验证令牌并将用户phone存入context
func MwAuthentication(c *gin.Context) {
	//验证令牌
	token := c.GetHeader(constant.CToken)
	phone, err := service.CheckToken(token)
	if err != nil {
		response.ErrResp(c, 401, err.Error())
		c.Abort()
		return
	}
	//存入用户信息
	c.Set(constant.CData, phone)
	c.Next()
}
