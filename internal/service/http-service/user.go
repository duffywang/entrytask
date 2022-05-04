package http_service

import (
	"errors"

	proto "github.com/duffywang/entrytask/proto"
)

//定义各种请求结构体
//请求使用form格式
type LoginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type RegisterUserReuqest struct {
	Username   string `form:"username" binding:"required,min=3,max=20"`
	Password   string `form:"password" binding:"required,min=3,max=20"`
	Nickname   string `form:"nickname" binding:"required,min=3,max=20"`
	ProfilePic string `form:"profilepic" binding:"-"` //跳过校验，否则取不到profile_pic
}

type EditUserRequest struct {
	SessionID  string `form:"session_id"`
	Nickname   string `form:"nickname" binding:"min=3,max=20"`
	ProfilePic string `form:"profilepic" binding:"-"`
}

type GetUserRequest struct {
	SessionID string `form:"session_id"`
}

type LoginResponse struct {
	Username   string `json:"username" `
	Nickname   string `json:"nickname" `
	ProfilePic string `json:"profile_pic" `
	SessionID  string `json:"session_id"`
}

//返回值为json格式
type GetUserResponse struct {
	Username   string `json:"username" `
	Nickname   string `json:"nickname" `
	ProfilePic string `json:"profile_pic" `
}

type RegisterUserResponse struct {
}

type EditUserResponse struct {
}

func (svc *Service) Login(request *LoginRequest) (*LoginResponse, error) {
	res, err := svc.GetUserClient().Login(svc.ctx, &proto.LoginRequest{
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}
	return &LoginResponse{Username: res.Username, Nickname: res.Nickname, ProfilePic: res.ProfilePic, SessionID: res.SessionId}, nil

}

func (svc *Service) GetUserInfo(request *GetUserRequest) (*GetUserResponse, error) {
	res, err := svc.GetUserClient().GetUser(svc.ctx, &proto.GetUserRequest{
		SessionId: request.SessionID,
	})
	if err != nil {
		return nil, err
	}
	return &GetUserResponse{Username: res.Username, Nickname: res.Nickname, ProfilePic: res.ProfilePic}, nil
}

func (svc *Service) RegisterUser(request *RegisterUserReuqest) (*RegisterUserResponse, error) {
	//TODO:生成参数快捷键
	_, err := svc.GetUserClient().RegisterUser(svc.ctx, &proto.RegisterUserReuqest{
		Username: request.Username,
		Password: request.Password,
		Nickname: request.Nickname,
	})

	if err != nil {
		return nil, err
	}
	return &RegisterUserResponse{}, nil
}

func (svc *Service) EditUser(request *EditUserRequest) (*EditUserResponse, error) {
	_, err := svc.GetUserClient().EditUser(svc.ctx, &proto.EditUserRequest{
		SessionId:  request.SessionID,
		Nickname:   request.Nickname,
		ProfilePic: request.ProfilePic,
	})
	if err != nil {
		return nil, err
	}
	return &EditUserResponse{}, nil
}

func (svc *Service) AuthUser(sessionID string) (string, error) {
	username, err := svc.cache.Get(svc.ctx, "session_id"+sessionID)
	if err != nil {
		return "", errors.New("login authuser fail")
	}
	return username, err
}

var userClient proto.UserServiceClient

func (svc *Service) GetUserClient() proto.UserServiceClient {
	if userClient == nil {
		userClient = proto.NewUserServiceClient(svc.client)
	}
	return userClient
}
