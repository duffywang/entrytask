package http_service

import (
	"context"

	"github.com/duffywang/entrytask/global"
	"github.com/duffywang/entrytask/internal/dao"
	"google.golang.org/grpc"
)

type Service struct {
	ctx context.Context
	//全局唯一
	dao    *dao.Dao
	cache  *dao.RedisClient
	client *grpc.ClientConn
}

func NewService(ctx context.Context) Service {
	service := Service{ctx: ctx}
	service.dao = dao.NewDBClient(global.DBEngine)
	service.cache = dao.NewRedisClient(global.RedisClient)
	service.client = global.GRPCClient
	return service
}
