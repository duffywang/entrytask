package status

import (
	"fmt"
	"net/http"
)

//定义异常结构体，封装异常

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []string `json:"data"`
}

func NewError(code int, msg string) *Error {
	//避免code重复逻辑，codes怎么添加key-value呢
	return &Error{Code: code, Msg: msg}
}

func NewErrorWithData(code int, msg string, data []string) *Error {
	return &Error{Code: code, Msg: msg, Data: data}
}

func (e *Error) GetError() string {
	return fmt.Sprintf("Error Code : %d, Error Msg : %s", e.Code, e.Msg)
}

func (e *Error) GetCode() int {
	return e.Code
}

func (e *Error) GetMsg() string {
	return e.Msg
}

func (e *Error) GetData() []string {
	return e.Data
}

func (e *Error) GetStatusCode() int {
	switch c := e.GetCode(); {
	case c == ServerError.GetCode():
		return http.StatusInternalServerError
	case c == InvalidParamsError.GetCode():
		return http.StatusBadRequest
	case c == NotFoundError.GetCode():
		return http.StatusNotFound
	case c >= 300000:
		return http.StatusBadGateway
	}
	return http.StatusInternalServerError
}
