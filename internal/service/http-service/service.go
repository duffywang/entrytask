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
	dao *dao.Dao
	client *grpc.ClientConn
}

func NewService(ctx context.Context) Service {
	service := Service{ctx: ctx}
	service.dao = dao.New(global.DBEngine)
	service.client = global.GRPCCLient
	return service
}
