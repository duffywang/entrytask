package middleware

import (
	"github.com/duffywang/entrytask/internal/status"
	"github.com/duffywang/entrytask/pkg/response"
	"github.com/gin-gonic/gin"
)

func SessionRequired(c *gin.Context) {
	res := response.NewResponse(c)
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		res.ToErrorResponse(status.SessionError)
		return
	}
	c.Set("session_id", sessionID)
	c.Next()
}

func LoginRequired(c *gin.Context) {
	res := response.NewResponse(c)
	_, err := c.Cookie("session_id")

	//通过sessionID获取用户信息
	//svc := http_service.NewService(c)
	//svc.GetAuth(sessionId) //获取用户信息，使用redis存储
	username := "test"

	if err != nil {
		res.ToErrorResponse(status.UserLoginError)
		return
	}
	c.Set("username", username)
	c.Next()

}
