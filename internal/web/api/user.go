package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/duffywang/entrytask/internal/constant"
	http_service "github.com/duffywang/entrytask/internal/service/http-service"
	"github.com/duffywang/entrytask/pkg/response"
	"github.com/gin-gonic/gin"
)

const (
	MaxWorker = 100
	MaxQueue  = 200
)

var JobQueue chan Job

func init() {
	JobQueue = make(chan Job, MaxQueue)
}

type Payload struct {
	C      *gin.Context
	Method string
}

type Job struct {
	PayLoad Payload
}

//二级具体工作
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				//time.Sleep(100 * time.Millisecond)
				switch job.PayLoad.Method {
				case "Login":
					fmt.Println("Worker Login")
				case "Register":
					fmt.Println("Worker Register")
				case "Get":
					fmt.Println("Worker Get")
				case "Edit":
					fmt.Println("Worker Edit")
				}
				fmt.Printf("处理成功：%v\n", job)
			case <-w.quit:
				return

			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

//一级分发器
type Dispatcher struct {
	WorkerPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool}
}

func (d *Dispatcher) Run() {
	for i := 0; i < MaxWorker; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			go func(job Job) {
				JobChannel := <-d.WorkerPool

				JobChannel <- job
			}(job)
		}
	}
}

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
	if err != nil || param.Username == "" || param.Password == "" {
		resp.ResponseError(constant.InvalidParamsError)
		return
	}
	// if i := strings.Index(param.Username, ";"); i != -1 {
	// 	resp.ResponseError(constant.InvalidParamsError)
	// 	return
	// }

	//使用到服务，依赖倒置
	svc := http_service.NewService(c.Request.Context())
	loginResponse, err := svc.Login(&param)
	if err != nil {
		log.Println(err.Error())
		c.HTML(http.StatusOK, "login.html", nil)
		resp.ResponseError(constant.UserLoginError)
		return
	}
	c.HTML(http.StatusOK, "profile.html", gin.H{
		"Username": loginResponse.Username,
		"Nickname": loginResponse.Nickname,
	})
	c.SetCookie(constant.SessionId, loginResponse.SessionID, 3600, "/", "", false, true)
	resp.ResponseOK("Login Success", loginResponse)

}

//API层 获取用户信息
func (u User) Get(c *gin.Context) {
	resp := response.NewResponse(c)
	param := http_service.GetUserRequest{}
	//登录后具有sessionID信息，
	sessionID, _ := c.Get(constant.SessionId)
	//sessionID.(string)
	param.SessionID = fmt.Sprintf("%v", sessionID)

	svc := http_service.NewService(c.Request.Context())
	//通过sessionID查询用户信息
	getUserResponse, err := svc.GetUserInfo(&param)
	if err != nil {
		log.Println(err.Error())
		resp.ResponseError(constant.UserGetError)
		return
	}

	// //文件路径名是否有问题，没有使用Gin框架的模板渲染
	// tmpl, err := template.ParseFiles("template/user.tmpl")
	// if err != nil {
	// 	fmt.Println("template.ParseFiles failed", err)
	// 	return
	// }

	// err = tmpl.Execute(c.Writer, getUserResponse)
	// if err != nil {
	// 	fmt.Println("template.Execute failed", err)
	// 	return
	// }

	resp.ResponseOK("Get User Success", getUserResponse)

}

//API层 注册用户信息
func (u User) Register(c *gin.Context) {
	resp := response.NewResponse(c)
	param := http_service.RegisterUserReuqest{}
	//登录后具有sessionID信息，
	err := c.ShouldBind(&param)
	if err != nil {
		resp.ResponseError(constant.InvalidParamsError)
		return
	}
	svc := http_service.NewService(c.Request.Context())
	//通过sessionID查询用户信息
	registerUserResponse, err := svc.RegisterUser(&param)
	if err != nil {
		log.Println(err.Error())
		resp.ResponseError(constant.UserRegisterError)
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
		resp.ResponseError(constant.InvalidParamsError)
		return
	}
	svc := http_service.NewService(c.Request.Context())
	//登录后具有sessionID信息，请求中带有session_id，通过sessionID查询用户信息
	sessionID, _ := c.Get(constant.SessionId)
	param.SessionID = fmt.Sprintf("%v", sessionID)
	editUserResponse, err := svc.EditUser(&param)
	if err != nil {
		log.Println(err.Error())
		resp.ResponseError(constant.UserEditError)
		return
	}

	resp.ResponseOK("Edit User Success", editUserResponse)
}
