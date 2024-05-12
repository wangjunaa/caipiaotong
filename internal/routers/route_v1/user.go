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
			middleware.UserRegisterParam,
			api_v1.UserRegister,
		)
		user.POST("/login",
			middleware.UserLoginParam,
			api_v1.UserLogin,
		)
	}
}
