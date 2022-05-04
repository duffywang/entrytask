package api

import (
	"strconv"

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
func (f File) Upload(c *gin.Context) {
	resp := response.NewResponse(c)
	//TODO : file := c.FormFile("file")
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		resp.ToErrorResponse(status.FileFormError)
		return
	}

	//思考需要type嘛
	fileType, _ := strconv.Atoi(c.PostForm("type"))
	//检验file、fileHeader、fileType是否合法

	param := http_service.UploadFileRequest{File: file, FileHeader: fileHeader, FileType: fileType}

	svc := http_service.NewService(c.Request.Context())
	uploadResponse, err := svc.Upload(&param)
	if err != nil {
		resp.ToErrorResponse(status.FileUploadError)
		return
	}

	resp.ToNormalResponse("Upload Profile Picture Success", uploadResponse)
}
