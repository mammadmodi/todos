package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"mods/todos/commands"
	"mods/todos/conf"
	"mods/todos/models"
	"os"

	"github.com/urfave/cli"
)

var (
	dbConf = conf.InitDBConf()
	db *gorm.DB
)

func init() {
	db = models.Init(dbConf)
}

func main() {
	app := cli.NewApp()
	app.Name = "TODO"
	app.Usage = "manage daily tasks ..."
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello manager!")
		return nil
	}
	app.Commands = commands.Init(db)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
