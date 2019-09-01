package conf

import (
	"github.com/kelseyhightower/envconfig"
)

type DBConf struct {
	MysqlHost string `default:"127.0.0.1" split_words:"true"`
	MysqlPort string `default:"3306" split_words:"true"`
	MysqlUser string `default:"root" split_words:"true"`
	MysqlPass string `default:"123456" split_words:"true"`
	MysqlCharSet string `default:"utf8" split_words:"true"`
	DBName string `default:"todos" split_words:"true"`
}

func InitDBConf() (config *DBConf) {
	config = new(DBConf)
	err := envconfig.Process("DB", config)

	if err != nil {
		panic(err.Error())
	}

	return config
}
