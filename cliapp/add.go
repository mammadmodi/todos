package cliapp

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
	"mods/todos/models"
)

func NewAddCommand(db *gorm.DB) cli.Command {
	command := cli.Command{
		Name:    "add",
		Aliases: []string{"a"},
		Usage:   "add a task in todo list",
		Action: func(c *cli.Context) error {
			taskName := c.Args().First()
			if taskName == "" {
				fmt.Print("task cannot be null!\n")
				return nil
			}
			taskId := models.CreateTask(taskName, db)
			if taskId > 0 {
				fmt.Printf("task with id:'%v' added to todo list.\n", taskId)
			} else {
				fmt.Printf("task can not be stored in database")
			}
			return nil
		},
	}

	return command
}
