package grpc_service

import (
	"context"

	"github.com/duffywang/entrytask/global"
	"github.com/duffywang/entrytask/internal/dao"
	"github.com/duffywang/entrytask/pkg/utils/fileutils"
	"github.com/duffywang/entrytask/proto"
)

type FileService struct {
	ctx   context.Context
	dao   *dao.Dao
	cache *dao.RedisClient
	proto.UnimplementedFileServiceServer
}

func NewFileService(ctx context.Context) FileService {
	return FileService{
		ctx:   ctx,
		dao:   dao.New(global.DBEngine),
		cache: dao.NewRedisClient(global.RedisClient),
	}
}

func (svc FileService) Upload(ctx context.Context, request *proto.UploadRequest) (*proto.UploadResponse, error) {
	fileName := fileutils.GetFileName(request.FileName)
	fileURL := ""
	//
	return &proto.UploadResponse{FileUrl: fileURL, FileName: fileName}, nil
}
