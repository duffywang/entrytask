package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"

	"github.com/duffywang/entrytask/global"
	"github.com/duffywang/entrytask/internal/models"
	"github.com/duffywang/entrytask/internal/web"
	"github.com/duffywang/entrytask/pkg/setting"
)

var (
	port   string
	mode   string
	config string
)

func main() {
	r := web.NewRouter()
	//服务器配置
	s := &http.Server{
		Addr:         ":" + global.ServerSetting.HttpPort,
		Handler:      r,
		ReadTimeout:  global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
	}

	go func() {
		//监听
		log.Printf("Starting HTTP Server , Listening %s ... \n", s.Addr)
		//r.Run(":9090")可以修改端口
		err := r.Run()
		if err != nil {
			log.Fatalf("Server ListenAndServe Fail %v", err)
		}
	}()

	//Go1.8 内置Shutdown()方法优雅关机
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("ShutDown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown:%v", err)
	}

	log.Println("Server Existing")

}

//初始化
func init() {
	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuNum - 1)
	err := setupFlag()
	if err != nil {
		log.Fatalf("HTTP Set up Flag fail: %v\n", err)
	}

	err = setupSetting()
	if err != nil {
		log.Fatalf("HTTP Set up Setting fail: %v\n", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("HTTP Set up DBEngine fail %v\n", err)
	}

	err = setupCacheClient()
	if err != nil {
		log.Fatalf("HTTP Set up Cache Client fail: %v\n", err)
	}

	err = setupRPCClient()
	if err != nil {
		log.Fatalf("HTTP Set up RPC Client fail: %v\n", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("HTTP Set up Logger fail: %v\n", err)
	}
}

func setupFlag() error {
	//StringVar defines a string flag with specified name, default value, and usage string. The argument p points to a string variable in which to store the value of the flag.
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&mode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "/Users/wangyufei/entrytask/configs", "配置文件路径")
	// Parse parses the command-line flags from os.Args[1:]. Must be called after all flags are defined and before flags are accessed by the program.
	flag.Parse()
	return nil
}

//装载config文件中配置数据
func setupSetting() error {
	log.Printf("%v", config)
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

	if port != "" {
		global.ServerSetting.RPCPort = port
	}
	if mode != "" {
		global.ServerSetting.Mode = mode
	}

	return nil
}

//装载数据库客户端
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

//装载缓存客户端
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

//装载RPC客户端
func setupRPCClient() error {
	var err error
	global.GRPCClient, err = models.NewRPCClient(global.ClientSetting)
	if err != nil {
		log.Println("Set up RPC Client fail")
		return err
	}
	log.Println("Set up RPC Client Success")
	return nil

}

func setupLogger() error {
	logfile, err := os.OpenFile("log/httpserver.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open log file error :%v\n", err)
		return err
	}
	log.SetOutput(logfile)
	log.SetFlags(log.Llongfile | log.Ltime)
	return err
}
