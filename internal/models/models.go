package models

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/duffywang/entrytask/pkg/setting"
	"github.com/go-redis/redis/v8"
)


//返回数据库客户端
func NewDBEngine(databaseSetting *setting.DBSetting) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.PassWord,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)))
	if err != nil {
		return nil, err
	}
	sqlDB, _ := db.DB()

	//连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	//设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(20)
	//设置连接池中空闲的最大数量
	sqlDB.SetMaxIdleConns(10)
	return db, nil
}

//返回缓存redis客户端
func NewCacheClient(cacheSetting *setting.CacheSetting) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: cacheSetting.Host,
		DB:   cacheSetting.DBIndex,
	})
	return rdb, nil
}

//返回grpc客户端
func NewRPCClient(clientSetting *setting.ClientSetting) (*grpc.ClientConn, error) {
	// Background returns a non-nil, empty Context. It is never canceled, has no
	// values, and has no deadline. It is typically used by the main function,
	// initialization, and tests, and as the top-level Context for incoming
	// requests.
	ctx := context.Background()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, clientSetting.RPCHost, opts...)
}
