package cliapp

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
)

func NewAddCommand(db *gorm.DB) cli.Command {
	command := cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add a task in todo list",
		Action: func(c *cli.Context) error {
			task := c.Args().First()
			if task == "" {
				fmt.Print("task cannot be null!\n")
			} else {
				fmt.Printf("task '%v' added to todo list.\n", task)
			}
			return nil
		},
	}

	return command
}
