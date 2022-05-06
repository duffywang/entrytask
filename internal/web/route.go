package web

import (
	"net/http"

	"github.com/duffywang/entrytask/internal/web/api"
	"github.com/duffywang/entrytask/pkg/middleware"
	"github.com/gin-gonic/gin"
)

//建立路由关系
func NewRouter() *gin.Engine {
	//gin.Default 默认使用了Logger 和 Recovery中间件，Logger将日志写入gin.DefaultWriter,Recovery中间件会recover任何panic,返回500状态码
	r := gin.New()
	
	ping := api.NewPing()
	user := api.NewUser()
	file := api.NewFile()
	//加载其他
	r.LoadHTMLGlob("view/*")

	//r.Use(middleware.TimeMonitor)
	pingGroup := r.Group("api")
	//TODO:Login(c *gin.Context) 没有带参数
	pingGroup.GET("/ping", ping.Ping)

	r.GET("/api/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	loginGroup := r.Group("api")
	//1.习惯性用一对{ }包裹同组的路由，只是说为了看着清晰，和不用{ }功能上没有区别
	//2.路由组支持嵌套的
	{
		//注册和登录POST请求
		loginGroup.POST("/user/login", user.Login)
	}
	registerGroup := r.Group("api")
	{
		registerGroup.GET("/register", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "register.html", nil)
		})
		registerGroup.POST("/user/register", user.Register)
	}

	//获取用户信息和编辑用户信息
	sessionGroup := r.Group("api")
	sessionGroup.Use(middleware.SessionRequired)
	{
		sessionGroup.GET("/user/get", user.Get)
		sessionGroup.POST("/user/edit", user.Edit)
	}

	uploadGroup := r.Group("api")
	//uploadGroup.Use(middleware.LoginRequired)
	{
		uploadGroup.POST("/file/upload", file.Upload)
	}

	//路由兜底逻辑
	r.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", nil)
	})

	return r
}
