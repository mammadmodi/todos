//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/urfave/cli"
	"mods/todos/cliapp"
	"mods/todos/conf"
	"mods/todos/models"
)

var SuperSet = wire.NewSet(conf.InitDBConf, conf.InitCliAppConf, models.InitDB, cliapp.InitCliApp)

func initCliApp() (*cli.App, error) {
	wire.Build(SuperSet)

	return &cli.App{}, nil
}
