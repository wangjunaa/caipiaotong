package route_v1

import (
	"caipiaotong/internal/api/api_v1"
	"caipiaotong/internal/middleware"
	"github.com/gin-gonic/gin"
)

func BindUserRouter(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/register",
			middleware.GetRegisterParam,
			api_v1.Register,
		)
		user.POST("/login",
			middleware.GetLoginParam,
			api_v1.Login,
		)

		user.Use(middleware.MwAuthentication)
		user.GET("/get", api_v1.UserGet)
		user.POST("/update",
			middleware.GetUpdateParam,
			api_v1.UserUpdate,
		)
		user.POST("/del",
			middleware.GetDelParam,
			api_v1.UserDel,
		)
	}
}
