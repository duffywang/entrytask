package dao

import (
	"database/sql"
	_ "mysql"
)

type MySqlDao struct {
	
}

func main() {
	db,err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}