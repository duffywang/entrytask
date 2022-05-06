package response

import (
	"net/http"

	"github.com/duffywang/entrytask/internal/status"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

//提示信息、数据，正常请求状态码一定为业务自定义0 ，http.StatusOK = 200，r.Ctx.JSON(http.StatusOK, response)封装
func (r *Response) ResponseOK(msg string, data any) {
	response := gin.H{"code": 0, "msg": "success", "data": data}
	r.Ctx.JSON(http.StatusOK, response)
}

func (r *Response) ResponseError(err *status.Error) {
	response := gin.H{"code": err.GetCode(), "msg": err.GetMsg()}
	details := err.GetData()
	if len(details) > 0 {
		response["data"] = details[0]
	}
	r.Ctx.JSON(err.GetStatusCode(), response)
}
