package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ipfans/adrgo"
	"github.com/urfave/cli/v2"
)

var initCmd = &cli.Command{
	Name:  "init",
	Usage: "Initialize a new ADR configure and directory",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "verbose",
			Aliases: []string{"v"},
			Usage:   "Verbosity level outputs",
			Value:   false,
		},
		&cli.StringFlag{
			Name:    "language",
			Aliases: []string{"l", "lang"},
			Usage:   "Set languages (en, zh-cn)",
			Value:   "en",
		},
		&cli.StringFlag{
			Name:    "path",
			Aliases: []string{"p"},
			Usage:   "Set path to ADR directory.",
			Value:   filepath.Join("docs", "adr"),
		},
		&cli.IntFlag{
			Name:    "digits",
			Aliases: []string{"d"},
			Usage:   "Set path to ADR directory.",
			Value:   5,
		},
	},
	Action: func(c *cli.Context) error {
		if c.Bool("verbose") {
			log.Default().SetOutput(os.Stdout)
		}
		conf := adrgo.Config{
			Language: c.String("language"),
			Path:     c.String("path"),
		}
		err := adrgo.Init(conf)
		if err != nil {
			log.Println("Init failed:", err.Error())
		}
		return err
	},
}
