package cmd

import (
	"github.com/urfave/cli"
	"models"
)

var (
	CmdMigrate = cli.Command{
		Name:        "migrate",
		Usage:       "Migrate",
		Description: `Migrate`,
		Action:      migrate,
	}
)

func migrate(ctx *cli.Context) error {
	err := models.NewContext()
	if err != nil {
		return err
	}

	err = models.Migrate()
	if err != nil {
		return err
	}
	return nil
}
