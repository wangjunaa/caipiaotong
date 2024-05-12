package middleware

import (
	"caipiaotong/configs/constant"
	"caipiaotong/internal/type/request"
	"caipiaotong/internal/type/response"
	"github.com/gin-gonic/gin"
)

// UserRegisterParam 解析数据
func UserRegisterParam(c *gin.Context) {
	data := &request.UserRegister{}
	err := c.ShouldBind(data)
	if err != nil {
		response.ErrResp(c, 400, constant.MsgBadReqs)
		c.Abort()
		return
	}
	c.Set(constant.CData, data)
	c.Next()
}
func UserLoginParam(c *gin.Context) {
	data := &request.UserLogin{}
	err := c.ShouldBind(data)
	if err != nil {
		response.ErrResp(c, 400, constant.MsgBadReqs)
		c.Abort()
		return
	}
	c.Set(constant.CData, data)
	c.Next()
}
