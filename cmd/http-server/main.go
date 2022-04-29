package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

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
		Addr:         global.ServerSetting.HttpPort,
		Handler:      r,
		ReadTimeout:  global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
	}

	go func() {
		//监听
		err := s.ListenAndServe()
		if err != nil {
			log.Fatalf("Server ListenAndServe Fail %v", err)
		}
	}()

	//TODO：服务器退出逻辑

}

func init() {
	err := setupFlag()
	if err != nil {
		log.Fatalf("HTTP Set up Flag fail: %v", err)
	}
	fmt.Println("HTTP server Setup Flag success")

	err = setupSetting()
	if err != nil {
		//Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
		log.Fatalf("HTTP Set up Setting fail: %v", err)
	}
	fmt.Println("HTTP server Setup Setting success")

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("HTTP Set up DBEngine fail %v", err)
	}
	fmt.Println("HTTP server Setup DB success")
	err = setupCacheClient()
	if err != nil {
		log.Fatalf("HTTP Set up Cache Client fail: %v", err)
	}
	fmt.Println("HTTP server Setup Cache success")
	err = setupRPCClient()
	if err != nil {
		log.Fatalf("HTTP Set up RPC Client fail: %v", err)
	}
	fmt.Println("HTTP server Setup RPC Clieny success")

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
	err = s.ReadSection("Database", &global.DatabaseSetting)
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

	return nil
}

func setupFlag() error {
	//StringVar defines a string flag with specified name, default value, and usage string. The argument p points to a string variable in which to store the value of the flag.
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&mode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "配置文件路径")
	// Parse parses the command-line flags from os.Args[1:]. Must be called after all flags are defined and before flags are accessed by the program.
	flag.Parse()
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = models.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		log.Println("Set up DBEngine fail")
		return err
	}
	log.Println("Set up DBEngine Success")
	return nil
}

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
