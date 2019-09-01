package commands

import (
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
)

func Init(db *gorm.DB) []cli.Command {
	var registeredCmds []cli.Command

	registeredCmds = append(registeredCmds, NewAddCommand(db))
	registeredCmds = append(registeredCmds, NewListCommand(db))
	registeredCmds = append(registeredCmds, NewDoCommand(db))

	return registeredCmds
}
