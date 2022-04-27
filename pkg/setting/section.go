package setting

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