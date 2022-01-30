package main

import (
	"fmt"
	"os"

	"github.com/ipfans/adrgo"
	"github.com/urfave/cli/v2"
)

var listCmd = &cli.Command{
	Name:  "list",
	Usage: "List all ADR records",
	Action: func(c *cli.Context) error {
		conf, err := adrgo.ReadConfig()
		if err != nil {
			fmt.Println("\nWe can not locate .adr.yml file, please run `adr init` first.")
			return err
		}
		adrgo.LoadLanguage(conf.Language)
		return adrgo.List(c.Context, conf, os.Stdout)
	},
}
