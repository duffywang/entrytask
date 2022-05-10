package grpc_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	_ "log"
	"time"

	"github.com/duffywang/entrytask/global"
	"github.com/duffywang/entrytask/internal/constant"
	"github.com/duffywang/entrytask/internal/dao"
	"github.com/duffywang/entrytask/pkg/utils/hashutils"
	"github.com/duffywang/entrytask/proto"
	uuid "github.com/satori/go.uuid"
	_ "golang.org/x/crypto/bcrypt"
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

//TODO:为啥http_service Service中用的指针，grpc_service Service没有用指针，自动处理了还是和调用有关系
//RPC服务端 用户登录方法
func (svc UserService) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginReply, error) {
	// pu, err := svc.GetUserProfileFromCache(request.Username)
	// if err == nil && pu.Password != "" {
	// 	log.Println("Login Cache ")
	// 	pwd := hashutils.Hash(request.Password)
	// 	if pu.Password != pwd {
	// 		return nil, errors.New("Login fail : pwd incorrect")
	// 	}
	// 	return &proto.LoginReply{Username: pu.Username, Nickname: pu.Nickname, ProfilePic: pu.ProfilePic}, nil
	// }

	//1.用户账户是否存在
	u, err := svc.dao.GetUserInfo(request.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Login Fail : User Not Exist")
		}
		return nil, err
	}

	//web(http-server)-service(grpc-server)-dao

	//2.验证用户密码是否正确
	//bcrypt耗时太高，性能较差，选择MD5+salt方法
	//err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(request.Password))
	pwd := hashutils.Hash(request.Password)
	if u.Password != pwd {
		return nil, errors.New("Login fail : pwd incorrect")
	}
	//log.Printf("user %v Login Password Valid Correct\n", u.Username)

	//3.session 存储
	//3.1 使用uuid生成sessionID
	sessionID := uuid.NewV4()
	//3.2 存储sessionID以及对应的username
	getUserResponse := &proto.GetUserReply{
		Username:   u.Username,
		Nickname:   u.Nickname,
		ProfilePic: u.ProfilePic,
		Password:   u.Password,
	}

	_ = svc.cache.Set(svc.ctx, constant.SessionIdWithColon+sessionID.String(), u.Username, time.Hour)
	_ = svc.UpdateUserProfileToCache(u.Username, getUserResponse)

	return &proto.LoginReply{Username: u.Username, Nickname: u.Nickname, ProfilePic: u.ProfilePic, SessionId: sessionID.String()}, nil

}

//RPC服务端 注册用户方法
func (svc UserService) RegisterUser(ctx context.Context, request *proto.RegisterUserReuqest) (*proto.RegisterUserReply, error) {
	//1.判断username是否已存在
	_, err := svc.dao.GetUserInfo(request.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		//说明已存在
		return nil, errors.New("RegisterUser Fail : Username Exist")
	}
	//数据库中不存在要注册的username，则继续执行注册逻辑
	//pwd, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	pwd := hashutils.Hash(request.Password)
	_, err = svc.dao.CreateUser(request.Username, request.Nickname, pwd, request.ProfilePic, 0)
	if err != nil {
		return nil, err
	}
	//RegisterUserResponse 没有定义字段
	return &proto.RegisterUserReply{}, nil
}

//RPC服务端 编辑用户方法
func (svc UserService) EditUser(ctx context.Context, request *proto.EditUserRequest) (*proto.EditUserReply, error) {
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

	//3.修改用户信息
	user, err := svc.dao.UpdateUser(u.ID, request.Nickname, request.ProfilePic)
	if err != nil {
		return nil, errors.New("userservice Edit Fail : Update User Information Fail")
	}

	//4.更新缓存
	getUserResponse := &proto.GetUserReply{
		Username:   user.Username,
		Nickname:   user.Nickname,
		ProfilePic: user.ProfilePic,
	}
	err = svc.UpdateUserProfileToCache(username, getUserResponse)
	if err != nil {
		return nil, errors.New("userservice Edit Fail : Cache User Information Fail")
	}

	return &proto.EditUserReply{Username: user.Username, Nickname: user.Nickname, ProfilePic: user.ProfilePic}, nil
}

//RPC服务端 获取用户信息方法
func (svc UserService) GetUser(ctx context.Context, request *proto.GetUserRequest) (*proto.GetUserReply, error) {
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

	//3.数据中获取用户信息
	if u == nil {
		mu, err := svc.dao.GetUserInfo(username)
		return &proto.GetUserReply{
			Username:   mu.Username,
			Nickname:   mu.Nickname,
			ProfilePic: mu.ProfilePic,
		}, err
	}

	return u, nil
}

//通过session_id从缓存中获取用户名
func (svc UserService) GetUsernameFromCache(sessionID string) (string, error) {
	username, err := svc.cache.Get(svc.ctx, constant.SessionIdWithColon+sessionID)
	if err != nil {
		return "", err
	}
	return username, nil
}

//更新缓存中用户信息
func (svc UserService) UpdateUserProfileToCache(key string, u *proto.GetUserReply) error {
	cacheKey := constant.ProfileWithColon + key

	cacheUser, err := json.Marshal(u)
	if err != nil {
		fmt.Printf("userservice GetUser UpdateUserProfile json Marchal Failed")
	}
	err = svc.cache.Set(svc.ctx, cacheKey, cacheUser, time.Minute*30)
	return err
}

//从缓存中获取用户信息
func (svc UserService) GetUserProfileFromCache(key string) (*proto.GetUserReply, error) {
	cacheKey := constant.ProfileWithColon + key

	value, err := svc.cache.Get(svc.ctx, cacheKey)
	if err != nil {
		return nil, err
	} else {
		getUserResponse := proto.GetUserReply{}
		json.Unmarshal([]byte(value), &getUserResponse)
		return &getUserResponse, nil
	}
}
