package grpc_service

import (
	"context"
	"errors"

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
		dao:   dao.NewDBClient(global.DBEngine),
		cache: dao.NewRedisClient(global.RedisClient),
	}
}

//RPC服务端 用户上传图片方法
func (svc FileService) Upload(ctx context.Context, request *proto.UploadRequest) (*proto.UploadReply, error) {
	fileName := fileutils.GetFileName(request.FileName)
	savePath := fileutils.GetSavePath()
	dest := savePath + "/" + fileName

	if fileutils.CheckSavePathValid(savePath) {
		//如果存储路径不存在，创建一个
		err := fileutils.CreateSavePath(savePath)
		if err != nil {
			return nil, errors.New("svc.Upload CreateSavePath Failed")
		}
	}

	if fileutils.CheckPermisson(savePath) {
		return nil, errors.New("svc.Upload CheckPermisson Failed")
	}

	err := fileutils.SaveFileByte(&request.Contents, dest)
	if err != nil {
		return nil, errors.New("svc.Upload SaveFileByte Failed")
	}

	fileURL := "http://localhost:8080/static/" + fileName
	return &proto.UploadReply{FileUrl: fileURL, FileName: fileName}, nil
}
