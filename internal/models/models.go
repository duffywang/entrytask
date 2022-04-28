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

type CommonModel struct {
	ID         uint32 `json:"id"`
	CreateTime uint32 `json:"createTime,omitempty" `
	UpdateTime uint32 `json:"updateTime,omitempty"`
}

/*
Full-Featured ORM
Associations (Has One, Has Many, Belongs To, Many To Many, Polymorphism, Single-table inheritance)
Hooks (Before/After Create/Save/Update/Delete/Find)
Eager loading with Preload, Joins
Transactions, Nested Transactions, Save Point, RollbackTo to Saved Point
Context, Prepared Statement Mode, DryRun Mode
Batch Insert, FindInBatches, Find To Map
SQL Builder, Upsert, Locking, Optimizer/Index/Comment Hints, NamedArg, Search/Update/Create with SQL Expr
Composite Primary Key
Auto Migrations
Logger
Extendable, flexible plugin API: Database Resolver (Multiple Databases, Read/Write Splitting) / Prometheus…
Every feature comes with tests
Developer Friendly
*/
func NewDBEngine(databaseSetting *setting.DatabaseSetting) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		"username",
		"password",
		"host",
		"DBname",
		"Charset",
		true,
	)))
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()

	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	return db, nil
}

func NewCacheClient(cacheSetting *setting.CacheSetting) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: cacheSetting.Host,
		DB:   cacheSetting.DBIndex,
	})
	return rdb, nil
}

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
