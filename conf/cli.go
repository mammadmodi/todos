package conf

import (
	"github.com/kelseyhightower/envconfig"
	"mods/todos/cliapp"
)

func InitCliAppConf() (config *cliapp.AppConf) {
	config = new(cliapp.AppConf)
	err := envconfig.Process("CLI", config)

	if err != nil {
		panic(err.Error())
	}

	return config
}
