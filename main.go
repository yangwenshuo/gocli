package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Action = func(c *cli.Context) error {
		fmt.Println("BOOM!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
