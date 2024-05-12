package middleware

import (
	"caipiaotong/configs/constant"
	"caipiaotong/internal/type/response"
	"github.com/gin-gonic/gin"
)

// GetRegisterParam 解析数据
func GetRegisterParam(c *gin.Context) {
	username, b1 := c.GetPostForm("username")
	password, b2 := c.GetPostForm("password")
	phone, b3 := c.GetPostForm("phone")
	if !b1 || !b2 || !b3 {
		response.ErrResp(c, 400, constant.MsgBadReqs)
		c.Abort()
		return
	}

	c.Set(constant.CData, map[string]string{
		"username": username,
		"password": password,
		"phone":    phone,
	})
	c.Next()
}
func GetLoginParam(c *gin.Context) {
	password, b1 := c.GetPostForm("password")
	phone, b2 := c.GetPostForm("phone")
	if !b1 || !b2 {
		response.ErrResp(c, 400, constant.MsgBadReqs)
		c.Abort()
		return
	}
	c.Set(constant.CData, map[string]string{
		"password": password,
		"phone":    phone,
	})
	c.Next()
}
func GetDelParam(c *gin.Context) {
	password, exist := c.GetPostForm("password")
	if !exist {
		response.ErrResp(c, 400, constant.MsgBadReqs)
		c.Abort()
		return
	}
	//读取手机号
	phone := c.GetString(constant.CData)
	c.Set(constant.CData, map[string]string{
		"phone":    phone,
		"password": password,
	})
	c.Next()
}
func GetUpdateParam(c *gin.Context) {
	password, exist := c.GetPostForm("password")
	newPassword := c.PostForm("newPassword")
	newUsername := c.PostForm("newUsername")
	if !exist || (newPassword == "" && newUsername == "") {
		response.ErrResp(c, 400, constant.MsgBadReqs)
		c.Abort()
		return
	}
	//读取手机号
	phone := c.GetString(constant.CData)
	c.Set(constant.CData, map[string]string{
		"phone":       phone,
		"password":    password,
		"newPassword": newPassword,
		"newUsername": newUsername,
	})
	c.Next()
}
