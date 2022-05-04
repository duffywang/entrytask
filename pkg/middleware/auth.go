package middleware

import (
	"log"
	"strings"
	"time"

	http_service "github.com/duffywang/entrytask/internal/service/http-service"
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
	sessionID, err := c.Cookie("session_id")

	svc := http_service.NewService(c.Request.Context())
	username, err := svc.AuthUser(sessionID)
	if err != nil {
		res.ToErrorResponse(status.UserLoginError)
		return
	}
	c.Set("username", username)
	c.Next()

}

func TimeMonitor(c *gin.Context) {
	if strings.HasSuffix(c.Request.URL.String(),"js") {
		return
	}
	start := time.Now()
	c.Next()
	cost := time.Since(start)
	log.Printf("%v\n",c.Request)
	log.Printf("Process Cost Time : %v\n", cost)
}
