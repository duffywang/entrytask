package models

import (
	"fmt"
)

type CommonModel struct {
	ID int32 `json:"id"`
	CreateTime int32 `json:"createTime,omitempty" `
	UpdateTime int32 `json:"updateTime,omitempty"`
}


func