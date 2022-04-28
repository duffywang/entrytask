package main

import (
	"flag"
	"log"
	"strings"

	"github.com/duffywang/entrytask/global"
	"github.com/duffywang/entrytask/internal/models"
	"github.com/duffywang/entrytask/pkg/setting"
)

var (
	port   string
	mode   string
	config string
)

func init() {
	err := setupSetting()
	if err != nil {
		//Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
		log.Fatalf("Set up Setting fail: %v", err)
	}

	err = setupFlag()
	if err != nil {
		log.Fatalf("Set up Flag fail: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("Set up DBEngine fail %v", err)
	}
	err = setupCacheClient()
	if err != nil {
		log.Fatalf("Set up Cache Client fail: %v", err)
	}

	err = setupRPCClient()
	if err != nil {
		log.Fatalf("Set up RPC Client fail: %v", err)
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
	flag.StringVar(&config, "config", "/configs", "配置文件路径")
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
