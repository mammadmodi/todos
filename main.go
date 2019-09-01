package main

import (
	"log"
	"os"
)

func main() {
	app, err := initCliApp()

	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
