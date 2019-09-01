package cliapp

import (
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
)

type AppConf struct {
	Name  string
	Usage string
}

func InitCliApp(conf *AppConf, db *gorm.DB) *cli.App {
	app := cli.NewApp()
	app.Name = conf.Name
	app.Usage = conf.Usage

	var registeredCmds []cli.Command

	registeredCmds = append(registeredCmds, NewAddCommand(db))
	registeredCmds = append(registeredCmds, NewListCommand(db))
	registeredCmds = append(registeredCmds, NewDoCommand(db))

	app.Commands = registeredCmds

	return app
}
