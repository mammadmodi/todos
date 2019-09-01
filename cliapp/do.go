package cliapp

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
	"mods/todos/models"
)

func NewDoCommand(db *gorm.DB) cli.Command {
	command := cli.Command{
		Name:    "do",
		Aliases: []string{"d"},
		Usage:   "doing a task",
		Action: func(c *cli.Context) error {
			taskId := c.Args().First()
			var task models.Task
			db.Where("id = ?", taskId).First(&task)
			if task.ID != 0 {
				if task.State != "waiting" {
					fmt.Print("task can not be done or is done before!\n")
					return nil
				}
				task.State = "done"
				db.Save(&task)
				fmt.Printf("task with '%v' added to done list.\n", taskId)
			} else {
				fmt.Print("task not found!\n")
			}
			return nil
		},
	}

	return command
}
