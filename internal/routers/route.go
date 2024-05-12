package routers

import (
	"caipiaotong/internal/routers/route_v1"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		route_v1.BindUserRouter(v1)
	}
	return r
}
