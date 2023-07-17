package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/wilhelm-murdoch/glaze/cmd/glaze/actions"
)

var (
	// Version describes the version of the current build.
	Version = "dev"

	// Commit describes the commit of the current build.
	Commit = "none"

	// Date describes the date of the current build.
	Date = "unknown"

	// Release describes the stage of the current build, eg; development, production, etc...
	Stage = "unknown"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			os.Exit(1)
		}
	}()

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("Version: %s, Stage: %s, Commit: %s, Date: %s\n", Version, Stage, Commit, Date)
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Usage:   "print only the version",
		Aliases: []string{"v"},
	}

	currentYear, _, _ := time.Now().Date()

	app := &cli.App{
		Name:     "glaze",
		Usage:    "easily manage tmux windows and panes",
		Version:  Version,
		Compiled: time.Now(),
		Authors: []*cli.Author{{
			Name:  "Wilhelm Murdoch",
			Email: "wilhelm@devilmayco.de",
		}},
		Copyright: fmt.Sprintf(`(c) %d Wilhelm Codes ( https://wilhelm.codes )`, currentYear),
		Commands: []*cli.Command{{
			Name:  "apply",
			Usage: "apply the specified glaze profile",
			Action: func(c *cli.Context) error {
				return actions.ActionApply(c.Args().First())
			},
		}, {
			Name:  "fmt",
			Usage: "rewrites the target glaze profile file to a canonical format",
			Action: func(c *cli.Context) error {
				return actions.ActionFmt(c.Args().First())
			},
		}, {
			Name:  "save",
			Usage: "running this within a tmux session will save its current state to the specified glaze profile",
			Action: func(c *cli.Context) error {
				return actions.ActionSave(c.Args().First())
			},
		}},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
