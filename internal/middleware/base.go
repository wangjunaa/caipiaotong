package middleware

import (
	"caipiaotong/internal/constant"
	"caipiaotong/internal/response"
	"caipiaotong/internal/service"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	// Authentication 验证令牌并将用户phone存入context
	Authentication(c *gin.Context)
}
type middleware struct {
	userService  service.UserService
	tokenService service.TokenService
	resp         response.Resp
}

func NewMiddleware() Middleware {
	return &middleware{
		userService:  service.NewUserService(),
		tokenService: service.NewTokenService(),
		resp:         response.NewResp(),
	}
}

func (m *middleware) Authentication(c *gin.Context) {
	//验证令牌
	token := c.GetHeader(constant.DToken)
	phone, err := m.tokenService.Check(token)
	if err != nil {
		m.resp.Error(c, 401, err)
		c.Abort()
		return
	}
	user, err := m.userService.Get(phone)
	if err != nil {
		m.resp.Error(c, 401, constant.ErrUserNotExist)
		c.Abort()
		return
	}
	//存入用户信息
	c.Set(constant.DUser, *user)
	c.Next()
}
