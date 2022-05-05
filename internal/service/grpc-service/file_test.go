package grpc_service

import (
	"context"
	"reflect"
	"testing"

	"github.com/duffywang/entrytask/internal/dao"
	"github.com/duffywang/entrytask/proto"
)


func TestFileService_Upload(t *testing.T) {
	type fields struct {
		ctx                            context.Context
		dao                            *dao.Dao
		cache                          *dao.RedisClient
		UnimplementedFileServiceServer proto.UnimplementedFileServiceServer
	}
	type args struct {
		ctx     context.Context
		request *proto.UploadRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.UploadReply
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := FileService{
				ctx:                            tt.fields.ctx,
				dao:                            tt.fields.dao,
				cache:                          tt.fields.cache,
				UnimplementedFileServiceServer: tt.fields.UnimplementedFileServiceServer,
			}
			got, err := svc.Upload(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileService.Upload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileService.Upload() = %v, want %v", got, tt.want)
			}
		})
	}
}
