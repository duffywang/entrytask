package api

import (
	http_service "github.com/duffywang/entrytask/internal/service/http-service"
	"github.com/duffywang/entrytask/internal/status"
	"github.com/duffywang/entrytask/pkg/response"
	"github.com/gin-gonic/gin"
)

type File struct {
}

func NewFile() File {
	return File{}
}

//API层 上传用户图片
func (f File) Upload(c *gin.Context) {
	resp := response.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		resp.ResponseError(status.FileFormError)
		return
	}

	param := http_service.UploadFileRequest{File: file, FileHeader: fileHeader}

	svc := http_service.NewService(c.Request.Context())
	uploadResponse, err := svc.Upload(&param)
	if err != nil {
		resp.ResponseError(status.FileUploadError)
		return
	}

	resp.ResponseOK("Upload Profile Picture Success", uploadResponse)
}
