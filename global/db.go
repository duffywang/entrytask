package global

import "gorm.io/gorm"

//全局变量，数据库引擎
var (
	DBEngine *gorm.DB
)
