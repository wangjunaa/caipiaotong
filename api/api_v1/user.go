package api_v1

import (
	"caipiaotong/configs/constant"
	"caipiaotong/internal/models"
	"caipiaotong/internal/service"
	"caipiaotong/internal/type/response"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	data := c.GetStringMapString(constant.CData)
	err := service.UserCreate(data["phone"], data["username"], data["password"])
	if err != nil {
		response.ErrResp(c, 400, err.Error())
		return
	}
	response.SuccessResp(c, 200, constant.MsgSuccess, nil)
}
func Login(c *gin.Context) {
	data := c.GetStringMapString(constant.CData)
	token, err := service.UserLogin(data["phone"], data["password"])
	if err != nil {
		response.ErrResp(c, 400, err.Error())
		return
	}
	response.SuccessResp(c, 200, constant.MsgSuccess, token)
}
func UserGet(c *gin.Context) {
	user := c.MustGet(constant.CUser).(models.User)
	response.SuccessResp(c, 200, constant.MsgSuccess, user)
}
func UserUpdate(c *gin.Context) {
	data := c.GetStringMapString(constant.CData)
	//更新用户
	err := service.UserUpdate(data["phone"], data["newUsername"], data["password"], data["newPassword"])
	if err != nil {
		response.ErrResp(c, 400, err.Error())
		return
	}
	response.SuccessResp(c, 200, constant.MsgSuccess, nil)
}
func UserDel(c *gin.Context) {
	data := c.GetStringMapString(constant.CData)
	err := service.UserDel(data["phone"], data["password"])
	if err != nil {
		response.ErrResp(c, 400, err.Error())
		return
	}
	response.SuccessResp(c, 200, constant.MsgSuccess, nil)
}
