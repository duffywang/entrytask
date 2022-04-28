package status

import "fmt"

//定义异常结构体，封装异常

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	//为何定义为切片
	Data []string `json:"data"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	//避免code重复逻辑，codes怎么添加key-value呢
	return &Error{Code: code, Msg: msg}
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

