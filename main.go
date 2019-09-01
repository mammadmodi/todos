package main

import (
	"log"
	"os"
)

func main() {
	app, _ := initCliApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
