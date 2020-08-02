package main

import (
	"github.com/gg-tools/hashid/commands"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "HashID"
	app.Compiled = time.Now()
	app.Usage = "HashID - encode & decode id"
	app.UsageText = `hashid [command] [args...]`
	app.Commands = []cli.Command{commands.Encode, commands.Decode}
	app.Action = func(c *cli.Context) error {
		cli.ShowAppHelpAndExit(c, 1)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
