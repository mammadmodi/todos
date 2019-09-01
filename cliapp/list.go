package cliapp

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
	"mods/todos/models"
)

func NewListCommand(db *gorm.DB) cli.Command {
	var flags []cli.Flag
	var waiting bool
	flags = append(flags, cli.BoolFlag{
		Name:        "waiting",
		Usage:       "if set then shows waiting only.",
		Destination: &waiting,
	})
	command := cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "list of tasks",
		Action: func(c *cli.Context) error {
			var tasks []models.Task
			if waiting {
				db.Where("state = ?", "waiting").Find(&tasks)
			} else {
				db.Find(&tasks)
			}
			if db.Error != nil {
				fmt.Print(db.Error)
			}
			for _, task := range tasks {
				fmt.Printf("%v:) %v, state: %v.\n", task.ID, task.Name, task.State)
			}
			return nil
		},
		Flags: flags,
	}

	return command
}
