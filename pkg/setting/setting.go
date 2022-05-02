package setting

import (
	"github.com/spf13/viper"
)

/*
Viper
setting defaults
reading from JSON, TOML, YAML, HCL, envfile and Java properties config files
live watching and re-reading of config files (optional)
reading from environment variables
reading from remote config systems (etcd or Consul), and watching changes
reading from command line flags
reading from buffer
setting explicit values
*/

type Setting struct {
	vp *viper.Viper
}

//读取 config.yanm.default配置文件
func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}
