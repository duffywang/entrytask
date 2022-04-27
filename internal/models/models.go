package models

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
func NewDBEngine() (*gorm.DB, error) {
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
