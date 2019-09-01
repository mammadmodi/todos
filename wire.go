//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/urfave/cli"
	"mods/todos/cliapp"
	"mods/todos/conf"
	"mods/todos/models"
)

func initCliApp() (*cli.App, error) {
	wire.Build(wire.NewSet(conf.InitDBConf, conf.InitCliAppConf, models.InitDB, cliapp.InitCliApp))

	return &cli.App{}, nil
}
