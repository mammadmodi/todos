package commands

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
)

func NewListCommand(db *gorm.DB) cli.Command {
	command := cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "list of tasks",
		Action:  func(c *cli.Context) error {
			fmt.Print("1:) task number 1!\n")
			fmt.Print("2:) task number 2!\n")
			fmt.Print("3:) task number 3!\n")
			return nil
		},
	}

	return command
}
