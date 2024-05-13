package routers

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"os"
	"time"
)

func Router() *gin.Engine {
	path := "./logs/ "
	writer, _ := rotatelogs.New(
		path+"%Y%m%d%H.log",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(1*time.Hour),
	)
	gin.DefaultWriter = io.MultiWriter(writer, os.Stdout)

	r := gin.Default()
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		BindUserRouter(v1)
	}
	return r
}
