package utils

import "github.com/tsuru/config"

var parameterConfig *Config

type Config struct {
	ListenAddress string
	SqlConfig     SqlConfig
}

type SqlConfig struct {
	DataSource string
	Debug      bool
}

func LoadConfig(path string) {
	parameterConfig = &Config{
		SqlConfig: SqlConfig{},
	}
	err := config.ReadConfigFile(path)
	if err != nil {
		panic(err)
	}
	if parameterConfig.ListenAddress, err = config.GetString("listen_address"); err != nil {
		panic(err)
	}

	if parameterConfig.SqlConfig.DataSource, err = config.GetString("mysql_setting:data_source"); err != nil {
		panic(err)
	}

	if parameterConfig.SqlConfig.Debug, err = config.GetBool("mysql_setting:enable_debug"); err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	return parameterConfig
}
