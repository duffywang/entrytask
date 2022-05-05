package global

import "github.com/duffywang/entrytask/pkg/setting"

//http服务端和RPC服务端、数据库、缓存、RPC客户端
var (
	ServerSetting *setting.ServerSetting
	DBSetting     *setting.DBSetting
	CacheSetting  *setting.CacheSetting
	ClientSetting *setting.ClientSetting
)
