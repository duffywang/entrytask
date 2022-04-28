package web

import (
	"github.com/duffywang/entrytask/internal/web/api"
	"github.com/duffywang/entrytask/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	user := api.NewUser()
	file := api.NewFile()

	loginGroup := r.Group("api")
	{
		//注册和登录POST请求
		loginGroup.POST("/user/login", user.Login)
	}

	registerGroup := r.Group("api")
	{
		registerGroup.POST("/user/register", user.Register)
	}

	//获取用户信息和编辑用户信息
	sessionGroup := r.Group("api")
	sessionGroup.Use(middleware.SessionRequired)
	{
		sessionGroup.GET("/user/get", user.Get)
		sessionGroup.GET("/user/edit", user.Edit)
	}

	uploadGroup := r.Group("api")
	uploadGroup.Use(middleware.LoginRequired)
	{
		uploadGroup.POST("user/upload", file.Upload)
	}

	return r
}
