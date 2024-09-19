package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/wilhelm-murdoch/glaze/cmd/glaze/actions"
)

const (
	TmuxBinaryName = "tmux"
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
		Before: func(ctx *cli.Context) error {
			_, err := exec.LookPath(TmuxBinaryName)
			if err != nil {
				return err
			}
			return nil
		},
		Commands: []*cli.Command{{
			Name:  "up",
			Usage: "apply the specified glaze profile",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "detached",
					Usage: "start a tmux session using glaze in detached mode",
				},
				&cli.BoolFlag{
					Name:  "clear",
					Usage: "if it exists, clear the current glaze session before starting",
				},
				&cli.BoolFlag{
					Name:  "debug",
					Usage: "prints a list of all commands sent to the specified tmux socket",
				},
				&cli.StringFlag{
					Name:  "socket-path",
					Value: "",
					Usage: "optional path to the tmux socket",
				},
				&cli.StringFlag{
					Name:  "socket-name",
					Value: "",
					Usage: "optional name for the tmux socket",
				},
			},
			Action: actions.ActionUp,
		}, {
			Name:  "fmt",
			Usage: "rewrites the target glaze profile file to a canonical format",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "stdout",
					Usage: "writes the formatted glaze output to your terminal",
				},
			},
			Action: actions.ActionFmt,
		}, {
			Name:  "save",
			Usage: "running this within a tmux session will save its current state to the specified glaze profile",
			Action: func(ctx *cli.Context) error {
				return actions.ActionSave(ctx.Args().First())
			},
		}},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
