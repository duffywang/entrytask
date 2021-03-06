package middleware

import (
	"log"
	"strings"
	"time"

	"github.com/duffywang/entrytask/internal/constant"
	http_service "github.com/duffywang/entrytask/internal/service/http-service"
	"github.com/duffywang/entrytask/pkg/response"
	"github.com/gin-gonic/gin"
)

//客户端请求携带session_id校验
func SessionRequired(c *gin.Context) {
	res := response.NewResponse(c)
	sessionID, err := c.Cookie(constant.SessionId)
	if err != nil {
		res.ResponseError(constant.SessionError)
		return
	}
	c.Set(constant.SessionId, sessionID)
	c.Next()
}

//登录校验
func LoginRequired(c *gin.Context) {
	res := response.NewResponse(c)
	sessionID, _ := c.Cookie(constant.SessionId)

	svc := http_service.NewService(c.Request.Context())
	username, err := svc.AuthUser(sessionID)
	if err != nil {
		res.ResponseError(constant.UserLoginError)
		return
	}
	c.Set("username", username)
	c.Next()

}

//请求耗时统计
func TimeMonitor(c *gin.Context) {
	log.Printf("Before Request: %v\n",c.Request)
	log.Printf("Request.Body : %v\n",c.Request.Body)
	if strings.HasSuffix(c.Request.URL.String(),"js") || strings.HasSuffix(c.Request.URL.String(),"ico") {
		return
	}
	start := time.Now()
	c.Next()
	cost := time.Since(start)
	log.Printf("After Request: %v\n",c.Request)
	log.Printf("Process Cost Time : %v\n", cost)
}
