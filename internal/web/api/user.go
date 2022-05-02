package api

import (
	"fmt"

	http_service "github.com/duffywang/entrytask/internal/service/http-service"
	"github.com/duffywang/entrytask/internal/status"
	"github.com/duffywang/entrytask/pkg/response"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

//为啥不是指针类型？
func (u User) Login(c *gin.Context) {
	//返回结果和参数
	resp := response.NewResponse(c)
	param := http_service.LoginRequest{}
	//检查数据格式是否对应正确
	err := c.ShouldBind(&param)
	if err != nil {
		resp.ToErrorResponse(status.InvalidParamsError)
		return
	}
	//使用到服务，依赖倒置
	svc := http_service.NewService(c.Request.Context())
	loginResponse, err := svc.Login(&param)
	if err != nil {
		resp.ToErrorResponse(status.UserLoginError)
		return
	}
	c.SetCookie("session_id", loginResponse.SessionID, 3600, "/", "", false, true)
	resp.ToNormalResponse("Login Success", loginResponse)

}

func (u User) Get(c *gin.Context) {
	resp := response.NewResponse(c)
	param := http_service.GetUserRequest{}
	//登录后具有sessionID信息，
	sessionID, _ := c.Get("session_id")
	//sessionID.(string) 
	param.SessionID = fmt.Sprintf("%v", sessionID)

	svc := http_service.NewService(c.Request.Context())
	//通过sessionID查询用户信息
	getUserResponse, err := svc.GetUserInfo(&param)
	if err != nil {
		//日志
		resp.ToErrorResponse(status.UserGetError)
		return
	}

	resp.ToNormalResponse("Get User Success", getUserResponse)

}

func (u User) Register(c *gin.Context) {
	resp := response.NewResponse(c)
	param := http_service.RegisterUserReuqest{}
	//登录后具有sessionID信息，
	err := c.ShouldBind(&param)
	if err != nil {
		resp.ToErrorResponse(status.InvalidParamsError)
		return
	}
	svc := http_service.NewService(c.Request.Context())
	//通过sessionID查询用户信息
	registerUserResponse, err := svc.RegisterUser(&param)
	if err != nil {
		//日志
		resp.ToErrorResponse(status.UserRegisterError)
		return
	}

	resp.ToNormalResponse("Get User Success", registerUserResponse)
}

func (u User) Edit(c *gin.Context) {
	resp := response.NewResponse(c)
	param := http_service.EditUserRequest{}

	err := c.ShouldBind(&param)
	if err != nil {
		resp.ToErrorResponse(status.InvalidParamsError)
		return
	}
	svc := http_service.NewService(c.Request.Context())
	//登录后具有sessionID信息，请求中带有session_id，通过sessionID查询用户信息
	sessionID, _ := c.Get("session_id")
	//TODO：为什么要这样？ 
	param.SessionID = fmt.Sprintf("%v", sessionID)
	editUserResponse, err := svc.EditUser(&param)
	if err != nil {
		//日志
		resp.ToErrorResponse(status.UserEditError)
		return
	}

	resp.ToNormalResponse("Get User Success", editUserResponse)
}
