package cliapp

import (
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
	"log"
	"mods/todos/models"
)

func NewMigrateCommand(db *gorm.DB) cli.Command {
	command := cli.Command{
		Name:    "migrate",
		Aliases: []string{"m"},
		Usage:   "updates migrate table",
		Action: func(c *cli.Context) error {
			db.AutoMigrate(&models.Task{})
			log.Println("migration done.")
			return nil
		},
	}

	return command
}
