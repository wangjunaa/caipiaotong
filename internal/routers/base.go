package routers

import (
	v1 "caipiaotong/internal/controller/v1"
	"caipiaotong/internal/middleware"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"os"
	"time"
)

type Router interface {
	Run(port string)
	initLogs()
	bindRoute()
}
type router struct {
	userHandler v1.UserHandler
	billHandler v1.BillHandler
	middleware  middleware.Middleware
	engine      *gin.Engine
}

func NewRouter() Router {
	r := router{
		userHandler: v1.NewUserHandler(),
		billHandler: v1.NewBillHandler(),
		middleware:  middleware.NewMiddleware(),
		engine:      gin.Default(),
	}
	r.initLogs()
	r.bindRoute()
	return &r

}

func (r *router) Run(port string) {
	err := r.engine.Run(port)
	if err != nil {
		panic(err)
	}
}

func (r *router) initLogs() {
	path := "./logs/ "
	writer, _ := rotatelogs.New(
		path+"%Y %m %d %H.log",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(1*time.Hour),
	)
	gin.DefaultWriter = io.MultiWriter(writer, os.Stdout)
}
func (r *router) bindRoute() {
	group := r.engine.Group("/api/v1")
	{
		r.bindUserRouter(group)
		r.bindBillRouter(group)
	}
}

func (r *router) bindUserRouter(g *gin.RouterGroup) {
	user := g.Group("/user")
	{
		user.POST("/register", r.userHandler.Register)
		user.POST("/login", r.userHandler.Login)

		user.Use(r.middleware.Authentication)
		user.GET("/get", r.userHandler.Get)
		user.POST("/update", r.userHandler.Update)
		user.POST("/del", r.userHandler.Del)
	}
}
func (r *router) bindBillRouter(g *gin.RouterGroup) {
	bill := g.Group("/bill", r.middleware.Authentication)
	{
		bill.POST("/orc", r.billHandler.OCR)
		bill.POST("/upload", r.billHandler.Upload)
		bill.GET("/getBills", r.billHandler.GetBills)
		bill.POST("/revocation", r.billHandler.Revocation)
		bill.GET("/summarize", r.billHandler.Summarize)
	}
}
