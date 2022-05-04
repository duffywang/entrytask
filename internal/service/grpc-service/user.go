package grpc_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/duffywang/entrytask/global"
	"github.com/duffywang/entrytask/internal/dao"
	"github.com/duffywang/entrytask/pkg/utils/hashutils"
	"github.com/duffywang/entrytask/proto"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

//rpc服务端逻辑，提供服务

type UserService struct {
	ctx   context.Context
	dao   *dao.Dao
	cache *dao.RedisClient
	proto.UnimplementedUserServiceServer
}

func NewUserService(ctx context.Context) UserService {
	return UserService{
		ctx:   ctx,
		dao:   dao.NewDBClient(global.DBEngine), //dao : global.DBEngine 不行，因为是小写的，需要通过New方法注入
		cache: dao.NewRedisClient(global.RedisClient),
	}
}

//TODO:grpc不使用XxxResponse，使用XxxReply
//QUESTION:为啥http_service Service中用的指针，grpc_service Service没有用指针，自动处理了还是和调用有关系
func (svc UserService) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	//web(http-server)-service(grpc-server)-dao
	//1.用户账户是否存在
	u, err := svc.dao.GetUserInfo(request.Username)
	if err != nil {
		return nil, err
	}
	//2.用户密码是否正确
	pwd := hashutils.Hash(request.Password)
	if u.Password != pwd {
		//2.1 密码错误
		return nil, errors.New("userservice Login fail : pwd incorrect")
	} else if u.Status != 0 {
		//2.2 用户被删除
		return nil, errors.New("userservice Login fail : user status disabled")
	}
	fmt.Println("Login Password Valid Correct")

	//3.session 存储
	//3.1 使用uuid生成sessionID
	sessionID := uuid.NewV4()
	//3.2 存储sessionID以及对应的username
	getUserResponse := &proto.GetUserResponse{
		Username:   u.Username,
		Nickname:   u.Nickname,
		ProfilePic: u.ProfilePic,
	}

	_ = svc.cache.Set(svc.ctx, "session_id:"+sessionID.String(), u.Username, time.Hour)
	_ = svc.UpdateUserProfile(u.Username, getUserResponse)

	return &proto.LoginResponse{Username: u.Username, Nickname: u.Nickname, ProfilePic: u.ProfilePic, SessionId: sessionID.String()}, nil

}

func (svc UserService) RegisterUser(ctx context.Context, request *proto.RegisterUserReuqest) (*proto.RegisterUserResponse, error) {
	//1.判断username是否已存在
	_, err := svc.dao.GetUserInfo(request.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		//说明已存在
		return nil, errors.New(" RegisterUser Fail : Username exist")
	}
	//数据库中不存在要注册的username，则继续执行注册逻辑

	pwd := hashutils.Hash(request.Password)
	//TODO:过期时间设为0是什么意思
	_, err = svc.dao.CreateUser(request.Username, request.Nickname, pwd, request.ProfilePic, 0)
	if err != nil {
		return nil, err
	}
	//RegisterUserResponse 没有定义字段
	return &proto.RegisterUserResponse{}, nil

}

func (svc UserService) EditUser(ctx context.Context, request *proto.EditUserRequest) (*proto.EditUserResponse, error) {

	//1.通过sessionID获取username
	username, err := svc.GetUsernameFromCache(request.SessionId)
	if err != nil {
		return nil, errors.New("userservice Edit fail : User Is Not Login in")
	}
	//2.根据username查询用户信息
	u, err := svc.dao.GetUserInfo(username)
	if err != nil {
		return nil, errors.New("userservice Edit fail : Get User Infoemation Fail")
	}
	//3.用户状态合法，0和1放在常量里面
	if u.Status == 1 {
		return nil, errors.New("userservice Edit fail : User Status Disabled")
	}

	//4.修改用户信息
	err = svc.dao.UpdateUser(u.ID, request.Nickname, request.ProfilePic)
	if err != nil {
		return nil, errors.New("userservice Edit Fail : Update User Information Fail")
	}

	//5.更新缓存
	getUserResponse := &proto.GetUserResponse{
		Username:   u.Username,
		Nickname:   u.Nickname,
		ProfilePic: u.ProfilePic,
	}
	err = svc.UpdateUserProfile(username, getUserResponse)
	if err != nil {
		return nil, errors.New("userservice Edit Fail : Cache User Information Fail")
	}

	return &proto.EditUserResponse{}, nil
}

func (svc UserService) GetUser(ctx context.Context, request *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	//1.通过sessionID获取username
	username, err := svc.GetUsernameFromCache(request.SessionId)
	if err != nil {
		return nil, errors.New("userservice Get fail : User Is Not Login in")
	}

	//2.缓存中获取用户信息
	u, err := svc.GetUserProfileFromCache(username)
	if err != nil {
		return nil, errors.New("userservice Get Fail : Get User Profile From Cache Fail")
	}
	return u, nil

}

func (svc UserService) GetUsernameFromCache(sessionID string) (string, error) {
	username, err := svc.cache.Get(svc.ctx, "session_id:"+sessionID)
	if err != nil {
		return "", err
	}
	return username, nil
}

func (svc UserService) UpdateUserProfile(key string, u *proto.GetUserResponse) error {
	cacheKey := "user_profile" + key

	cacheUser, err := json.Marshal(u)
	if err != nil {
		fmt.Printf("userservice GetUser UpdateUserProfile json Marchal Failed")
	}
	err = svc.cache.Set(svc.ctx, cacheKey, cacheUser, time.Hour*24)
	return err
}

func (svc UserService) GetUserProfileFromCache(key string) (*proto.GetUserResponse, error) {
	cacheKey := "user_profile" + key

	value, err := svc.cache.Get(svc.ctx, cacheKey)
	if err != nil {
		return nil, err
	} else {
		getUserResponse := proto.GetUserResponse{}
		//TODO：反序列化时传入指针
		json.Unmarshal([]byte(value), &getUserResponse)
		return &getUserResponse, nil
	}
}
