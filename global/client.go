package global

import "google.golang.org/grpc"

//全局变量，数据库引擎
var (
	GRPCClient *grpc.ClientConn
)