package http_service

import "github.com/golang/protobuf/proto"

//定义各种请求结构体
//请求使用form格式
type LoginRequest struct {
	UserName string `form:"username" binding:"required"`
	PassWord string `form:"password" binding:"required"`
}

type RegisterUserReuqest struct {
	UserName   string `form:"username" binding:"required,min=3,max=20"`
	PassWord   string `form:"password" binding:"required,min=3,max=20"`
	NickName   string `form:"nickname" binding:"required,min=3,max=20"`
	ProfilePic string `form:"profile_pic" `
}

type EditUserRequest struct {
	SessionID  string `form:session_id`
	NickName   string `form:"nickname" binding:"min=3,max=20"`
	ProfilePic string `form:"profile_pic" binding:"max = 1024"`
}

type GetUserRequest struct {
	SessionID string `form:session_id`
}

type LoginResponse struct {
	SessionID string `json:session_id`
}

//返回值为json格式
type GetUserResponse struct {
	UserName   string `json:"username" `
	NickName   string `json:"nickname" `
	ProfilePic string `json:"profile_pic" `
}

type RegisterUserResponse struct {
}

type EditUserResponse struct {
}

func (svc *Service) Login(request *LoginRequest) *LoginResponse {
	svc.GetUserClient()

}

func (svc *Service) GetUserInfo(request *GetUserRequest) *GetUserResponse {

}

func (svc *Service) RegisterUser(request *RegisterUserReuqest) *RegisterUserResponse {

}

func (svc *Service) EditUser(request *EditUserRequest) *EditUserResponse {

}

var userClient proto.UserServiceClient

func (svc *Service) GetUserClient() proto.UserServiceClient {
	if userClient == nil {
		userClient = proto.NewUserServiceClient(svc.client)
	}
	return userClient
}
