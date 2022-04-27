package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct{}

func  NewUser() User{
	return User{}
}

func (u User)Login(c *gin.Context){
	
}