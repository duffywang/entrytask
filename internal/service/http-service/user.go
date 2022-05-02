package http_service

import (
	proto "github.com/duffywang/entrytask/proto"
)

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
	ProfilePic string `form:"profilepic" binding:"-"` //跳过校验，否则取不到profile_pic
}

type EditUserRequest struct {
	SessionID  string `form:"session_id"`
	NickName   string `form:"nickname" binding:"min=3,max=20"`
	ProfilePic string `form:"profilepic" binding:"-"`
}

type GetUserRequest struct {
	SessionID string `form:"session_id"`
}

type LoginResponse struct {
	SessionID string `json:"session_id"`
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

func (svc *Service) Login(request *LoginRequest) (*LoginResponse, error) {
	res, err := svc.GetUserClient().Login(svc.ctx, &proto.LoginRequest{
		Username: request.UserName,
		Password: request.PassWord,
	})
	if err != nil {
		return nil, err
	}
	return &LoginResponse{SessionID: res.SessionId}, nil

}

func (svc *Service) GetUserInfo(request *GetUserRequest) (*GetUserResponse, error) {
	res, err := svc.GetUserClient().GetUser(svc.ctx, &proto.GetUserRequest{
		SessionId: request.SessionID,
	})
	if err != nil {
		return nil, err
	}
	return &GetUserResponse{UserName: res.Username, NickName: res.Nickname, ProfilePic: res.ProfilePic}, nil
}

func (svc *Service) RegisterUser(request *RegisterUserReuqest) (*RegisterUserResponse, error) {
	//TODO:生成参数快捷键
	_, err := svc.GetUserClient().RegisterUser(svc.ctx, &proto.RegisterUserReuqest{
		Username: request.UserName,
		Password: request.PassWord,
		Nickname: request.NickName,
	})

	if err != nil {
		return nil, err
	}
	return &RegisterUserResponse{}, nil
}

func (svc *Service) EditUser(request *EditUserRequest) (*EditUserResponse, error) {
	_, err := svc.GetUserClient().EditUser(svc.ctx, &proto.EditUserRequest{
		SessionId:  request.SessionID,
		Nickname:   request.NickName,
		ProfilePic: request.ProfilePic,
	})
	if err != nil {
		return nil, err
	}
	return &EditUserResponse{}, nil
}

var userClient proto.UserServiceClient

//方法名小写是私有的吗？
func (svc *Service) GetUserClient() proto.UserServiceClient {
	if userClient == nil {
		userClient = proto.NewUserServiceClient(svc.client)
	}
	return userClient
}
