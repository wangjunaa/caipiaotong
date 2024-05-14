package v1

import (
	"caipiaotong/internal/constant"
	"caipiaotong/internal/models"
	"caipiaotong/internal/response"
	"caipiaotong/internal/service"
	"caipiaotong/internal/utils/encrypt"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
	Del(c *gin.Context)
}
type userHandler struct {
	service service.UserService
	resp    response.Resp
}

func NewUserHandler() UserHandler {
	return &userHandler{
		service: service.NewUserService(),
		resp:    response.NewResp(),
	}
}

func (h *userHandler) Register(c *gin.Context) {
	phone := c.PostForm(constant.DPhone)
	username := c.PostForm(constant.DUsername)
	password := c.PostForm(constant.DPassword)

	err := h.service.Register(phone, username, password)
	if err != nil {
		h.resp.Error(c, 400, err)
		return
	}
	h.resp.Success(c, constant.MsgSuccess, nil)
}
func (h *userHandler) Login(c *gin.Context) {
	phone := c.PostForm(constant.DPhone)
	password := c.PostForm(constant.DPassword)
	token, err := h.service.Login(phone, password)
	if err != nil {
		h.resp.Error(c, 400, err)
		return
	}
	h.resp.Success(c, constant.MsgSuccess, token)
}
func (h *userHandler) Get(c *gin.Context) {
	user := c.MustGet(constant.DUser).(models.User)
	user.Password = ""
	h.resp.Success(c, constant.MsgSuccess, user)
}
func (h *userHandler) Update(c *gin.Context) {
	user := c.MustGet(constant.DUser).(models.User)
	password := c.PostForm(constant.DPassword)
	newUsername := c.PostForm(constant.DNewUsername)
	newPassword := c.PostForm(constant.DNewPassword)
	//验证密码
	err := encrypt.Check(password, user.Password)
	if err != nil {
		h.resp.Error(c, 400, err)
		return
	}
	//更新用户
	if newPassword != "" {
		newPassword, err = encrypt.Encode(newPassword)
		if err != nil {
			h.resp.Error(c, 400, err)
		}
		user.Password = newPassword
	}
	if newUsername != "" {
		user.Username = newUsername
	}
	err = h.service.Update(&user)
	if err != nil {
		h.resp.Error(c, 400, err)
		return
	}
	h.resp.Success(c, constant.MsgSuccess, nil)
}
func (h *userHandler) Del(c *gin.Context) {
	user := c.MustGet(constant.DUser).(models.User)
	password := c.PostForm(constant.DPassword)
	//验证密码
	err := encrypt.Check(password, user.Password)
	if err != nil {
		h.resp.Error(c, 400, err)
		return
	}
	err = h.service.Del(user.Phone, password)
	if err != nil {
		h.resp.Error(c, 400, err)
		return
	}
	h.resp.Success(c, constant.MsgSuccess, nil)
}
