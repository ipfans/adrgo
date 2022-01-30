package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	logger := log.Default()
	logger.SetFlags(0)

	app := cli.NewApp()
	app.Name = "adrgo"
	app.Version = Version
	app.Description = "Architecture Decision Records in Go with Reporter."
	app.Authors = []*cli.Author{
		{
			Name: "Kevin Lee",
		},
	}
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		initCmd,
		listCmd,
		newCmd,
		{
			Name:  "toc",
			Usage: "generate TOC file",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "graph",
			Usage: "generate DOT graph file",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "logs",
			Usage: "show logs for given ADR record",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "update",
			Usage: "update ADR record filename",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "export",
			Usage: "export ADR record to HTML/CSV file",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "search",
			Usage: "search ADR records with given keyword",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "version",
			Usage: "Print the version",
			Action: func(c *cli.Context) error {
				fmt.Println(Version)
				return nil
			},
		},
	}
	cli.HandleExitCoder(app.Run(os.Args))
}
