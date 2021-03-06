package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"strings"

	"github.com/duffywang/entrytask/global"
	"github.com/duffywang/entrytask/internal/models"
	grpc_service "github.com/duffywang/entrytask/internal/service/grpc-service"
	"github.com/duffywang/entrytask/pkg/middleware"
	"github.com/duffywang/entrytask/pkg/setting"
	"github.com/duffywang/entrytask/proto"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

var (
	port   string
	mode   string
	config string
)

//RPC服务端初始化配置
func init() {
	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum - 1)
	err := setupFlag()
	if err != nil {
		log.Fatalf("GRPC Set up Flag fail: %v\n", err)
	}
	err = setupSetting()
	if err != nil {
		//Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
		log.Fatalf("GRPC Set up Setting fail: %v\n", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("GRPC Set up DBEngine fail %v\n", err)
	}
	err = setupCacheClient()
	if err != nil {
		log.Fatalf("GRPC Set up Cache Client fail: %v\n", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("GRPC Set up Logger fail: %v\n", err)
	}

}

func setupSetting() error {
	s, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DBSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Redis", &global.CacheSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Client", &global.ClientSetting)
	if err != nil {
		return err
	}

	if port != ""{
		global.ServerSetting.RPCPort = port
	}
	if mode != ""{
		global.ServerSetting.Mode = mode
	}

	return nil
}

//Flag数据绑定
func setupFlag() error {
	//StringVar defines a string flag with specified name, default value, and usage string. The argument p points to a string variable in which to store the value of the flag.
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&mode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "/Users/wangyufei/entrytask/configs", "配置文件路径")
	// Parse parses the command-line flags from os.Args[1:]. Must be called after all flags are defined and before flags are accessed by the program.
	flag.Parse()
	return nil
}

//装载数据库配置
func setupDBEngine() error {
	var err error
	global.DBEngine, err = models.NewDBEngine(global.DBSetting)
	if err != nil {
		log.Println("Set up DBEngine fail")
		return err
	}
	log.Println("Set up DBEngine Success")
	return nil
}

//装载缓存配置
func setupCacheClient() error {
	var err error
	global.RedisClient, err = models.NewCacheClient(global.CacheSetting)
	if err != nil {
		log.Println("Set up Redis Client fail")
		return err
	}
	log.Println("Set up Redis Client Success")
	return nil
}

//设置log配置
// Ldate         = 1 << iota     // 日期：2009/01/23
// Ltime                         // 时间：01:23:23
// Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
// Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
// Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
// LUTC                          // 使用UTC时间
// LstdFlags     = Ldate | Ltime // 标准logger的初始值
func setupLogger() error {
	logfile, err := os.OpenFile("log/rpcserver.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open log file error :%v\n", err)
		return err
	}
	log.SetOutput(logfile)
	log.SetFlags(log.Lshortfile | log.Ldate)
	return err
}

//RPC服务端主函数，提供RPC服务
func main() {
	s := grpc.NewServer(grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
		middleware.Recovery,
	)))
	ctx := context.Background()
	//作为服务方注册图片上传服务和用户服务
	proto.RegisterFileServiceServer(s, grpc_service.NewFileService(ctx))
	proto.RegisterUserServiceServer(s, grpc_service.NewUserService(ctx))
	fmt.Println("Rpc-Server Main Func Success")

	lis, err := net.Listen("tcp", ":"+global.ServerSetting.RPCPort)
	if err != nil {
		log.Fatalf("GRPC Listen Fail: %v\n", err)
	}
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("GRPC Serve Fail: %v\n", err)
	}
}


