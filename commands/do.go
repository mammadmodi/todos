package commands

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
)

func NewDoCommand(db *gorm.DB) cli.Command {
	command := cli.Command{
		Name:    "do",
		Aliases: []string{"d"},
		Usage:   "doing a task",
		Action:  func(c *cli.Context) error {
			taskId := c.Args().First()
			if taskId == ""{
				fmt.Print("task id cannot be null!\n")
			} else {
				fmt.Printf("task with '%v' added to todo list.\n", taskId)
			}
			return nil
		},
	}

	return command
}
