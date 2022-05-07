package http_service

import (
	"bytes"
	"io"
	"mime/multipart"

	proto "github.com/duffywang/entrytask/proto"
)

type UploadFileRequest struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
}

type UploadFileResponse struct {
	FileName string `json:"filename"`
	FileUrl  string `json:"fileurl"`
}

//RPC客户端 上传图片方法
func (svc *Service) Upload(request *UploadFileRequest) (*UploadFileResponse, error) {
	//上传图片解析，转化为字节类型
	src, err := request.FileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		return nil, err
	}
	content := buf.Bytes()

	
	fileClient := svc.GetFileClient()
	resp, err := fileClient.Upload(svc.ctx, &proto.UploadRequest{
		FileName: request.FileHeader.Filename,
		Contents: content,
	})
	return &UploadFileResponse{FileName: resp.FileName, FileUrl: resp.FileUrl}, nil
}

var fileClient proto.FileServiceClient

//获取图片上传服务RPC客户端
func (svc *Service) GetFileClient() proto.FileServiceClient {
	if fileClient == nil {
		fileClient = proto.NewFileServiceClient(svc.client)
	}
	return fileClient
}
