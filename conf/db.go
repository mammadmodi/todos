package conf

import (
	"github.com/kelseyhightower/envconfig"
	"mods/todos/models"
)

func InitDBConf() (config *models.DBConf) {
	config = new(models.DBConf)
	err := envconfig.Process("DB", config)

	if err != nil {
		panic(err.Error())
	}

	return config
}
