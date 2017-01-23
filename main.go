package main

import (
	"cmd"

	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Smolgu"
	app.Usage = "Smolgu service"
	app.Commands = []cli.Command{
		cmd.CmdWeb,
		cmd.CmdMigrate,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
