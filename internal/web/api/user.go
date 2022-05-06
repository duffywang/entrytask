package api

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	http_service "github.com/duffywang/entrytask/internal/service/http-service"
	"github.com/duffywang/entrytask/internal/status"
	"github.com/duffywang/entrytask/pkg/response"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

//API层 用户登录
func (u User) Login(c *gin.Context) {

	//返回结果和参数
	resp := response.NewResponse(c)

	//检查数据格式是否对应正确
	param := http_service.LoginRequest{}
	err := c.ShouldBind(&param)
	//log.Printf("Login param %v\n", param)
	if err != nil {
		resp.ResponseError(status.InvalidParamsError)
		return
	}
	//使用到服务，依赖倒置
	svc := http_service.NewService(c.Request.Context())
	loginResponse, err := svc.Login(&param)
	if err != nil {
		log.Println(err.Error())
		c.HTML(http.StatusOK, "login.html", nil)
		resp.ResponseError(status.UserLoginError)
		return
	}
	c.HTML(http.StatusOK, "profile.html", gin.H{
		"Username": loginResponse.Username,
		"Nickname": loginResponse.Nickname,
	})
	c.SetCookie("session_id", loginResponse.SessionID, 3600, "/", "", false, true)
	resp.ResponseOK("Login Success", loginResponse)

}

//API层 获取用户信息
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
		log.Println(err.Error())
		resp.ResponseError(status.UserGetError)
		return
	}

	//文件路径名是否有问题，没有使用Gin框架的模板渲染
	tmpl, err := template.ParseFiles("template/user.tmpl")
	if err != nil {
		fmt.Println("template.ParseFiles failed", err)
		return
	}

	err = tmpl.Execute(c.Writer, getUserResponse)
	if err != nil {
		fmt.Println("template.Execute failed", err)
		return
	}

	resp.ResponseOK("Get User Success", getUserResponse)

}

//API层 注册用户信息
func (u User) Register(c *gin.Context) {
	resp := response.NewResponse(c)
	param := http_service.RegisterUserReuqest{}
	//登录后具有sessionID信息，
	err := c.ShouldBind(&param)
	if err != nil {
		resp.ResponseError(status.InvalidParamsError)
		return
	}
	svc := http_service.NewService(c.Request.Context())
	//通过sessionID查询用户信息
	registerUserResponse, err := svc.RegisterUser(&param)
	if err != nil {
		log.Println(err.Error())
		resp.ResponseError(status.UserRegisterError)
		return
	}
	c.HTML(http.StatusOK, "login.html", nil)
	resp.ResponseOK("Register User Success", registerUserResponse)

}

//API层 编辑用户信息
func (u User) Edit(c *gin.Context) {
	resp := response.NewResponse(c)
	param := http_service.EditUserRequest{}

	err := c.ShouldBind(&param)
	if err != nil {
		resp.ResponseError(status.InvalidParamsError)
		return
	}
	svc := http_service.NewService(c.Request.Context())
	//登录后具有sessionID信息，请求中带有session_id，通过sessionID查询用户信息
	sessionID, _ := c.Get("session_id")
	//TODO：为什么要这样？
	param.SessionID = fmt.Sprintf("%v", sessionID)
	editUserResponse, err := svc.EditUser(&param)
	if err != nil {
		log.Println(err.Error())
		resp.ResponseError(status.UserEditError)
		return
	}

	resp.ResponseOK("Edit User Success", editUserResponse)
}
