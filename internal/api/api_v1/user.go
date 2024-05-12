package api_v1

import (
	"caipiaotong/configs/constant"
	"caipiaotong/internal/service"
	"caipiaotong/internal/type/request"
	"caipiaotong/internal/type/response"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	val, exists := c.Get(constant.CData)
	if !exists {
		response.ErrResp(c, 500, constant.MsgMiddleErr)
		return
	}
	data := val.(request.UserRegister)
	err := service.CreateUser(data)
	if err != nil {
		response.ErrResp(c, 500, err.Error())
		return
	}
	response.SuccessResp(c, 200, constant.MsgSuccess, nil)
}

func UserLogin(c *gin.Context) {
	val, exists := c.Get(constant.CData)
	if !exists {
		response.ErrResp(c, 500, constant.MsgMiddleErr)
		return
	}
	data := val.(request.UserLogin)
	token, err := service.Login(data)
	if err != nil {
		response.ErrResp(c, 500, err.Error())
		return
	}
	response.SuccessResp(c, 200, constant.MsgSuccess, token)
}
