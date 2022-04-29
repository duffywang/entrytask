package middleware

import (
	"context"
	"log"
	"runtime/debug"

	"google.golang.org/grpc"
)

//实现grpc UnaryServerInterceptor 函数，rpc过程中捕捉异常处理器
func Recovery(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("method: %s, message: %s, stack: %s", info.FullMethod, err, string(debug.Stack()[:]))
		}
	}()
	return handler(ctx, req)

}
