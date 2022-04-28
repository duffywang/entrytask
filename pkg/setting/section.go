package setting

import "time"

type ServerSetting struct {
	HttpPort string
	RPCPort string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

type DatabaseSetting struct {
	DBType string
	UserName string
	PassWord string
	Host string
	DBName string
	TablePrefix string
	Charset string
	ParseTime bool
	MaxOpenConns int
	MaxIdleConns int
}

type CacheSetting struct {
	Host string
	DBIndex int 
}

type ClientSetting struct {
	RPCHost string
	ConnNum int
}



func (s *Setting)ReadSection(k string,v any)error {

	err := s.vp.UnmarshalKey(k,v)
	if err != nil {
		return err
	}
	return nil
}