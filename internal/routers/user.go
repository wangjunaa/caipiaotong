package routers

import (
	v1 "caipiaotong/internal/handler/v1"
	"caipiaotong/internal/middleware"
	"github.com/gin-gonic/gin"
)

func BindUserRouter(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/register", v1.Register)
		user.POST("/login", v1.Login)

		user.Use(middleware.MwAuthentication)
		user.GET("/get", v1.UserGet)
		user.POST("/update", v1.UserUpdate)
		user.POST("/del", v1.UserDel)
	}
}
