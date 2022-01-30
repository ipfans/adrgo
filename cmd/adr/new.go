package main

import (
	"fmt"

	"github.com/ipfans/adrgo"
	"github.com/urfave/cli/v2"
)

var newCmd = &cli.Command{
	Name:      "new",
	Usage:     "create new ADR records",
	ArgsUsage: "[title]",
	Action: func(c *cli.Context) error {
		conf, err := adrgo.ReadConfig()
		if err != nil {
			fmt.Println("\nWe can not locate .adr.yml file, please run `adr init` first.")
			return err
		}
		adrgo.LoadLanguage(conf.Language)
		if c.NArg() == 0 {
			fmt.Println("\nPlease enter the title of the new ADR record.")
			return nil
		}
		record := adrgo.ADRecord{
			Title: c.Args().First(),
		}
		return adrgo.New(c.Context, conf, record)
	},
}
