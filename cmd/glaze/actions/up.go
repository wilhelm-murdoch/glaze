package actions

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/k0kubun/pp"
	"github.com/urfave/cli/v2"
	"github.com/wilhelm-murdoch/glaze"
	"github.com/wilhelm-murdoch/glaze/tmux"
)

func ActionUp(ctx *cli.Context) error {
	profilePath := ctx.Args().First()

	if profilePath == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		profilePath = filepath.Join(cwd, ".glaze")

		if !glaze.FileExists(profilePath) && os.Getenv("GLAZE_PATH") != "" {
			profilePath = filepath.Join(os.Getenv("GLAZE_PATH"), ".glaze")
		}
	}

	if !glaze.FileExists(profilePath) {
		return fmt.Errorf("default glaze profile not found on the specified path, the current directory, or the GLAZE_PATH environment variable")
	}

	parser := glaze.NewParser(profilePath)

	if parser.HasErrors() {
		parser.WriteDiags()
		return nil
	}

	// profile := parser.Decode(glaze.PrimaryGlazeSpec)

	if parser.HasErrors() {
		parser.WriteDiags()
		return nil
	}

	// pp.Print(profile)

	var t tmux.Tmux

	sessions, _ := t.Sessions()

	for _, session := range sessions {
		pp.Print(t.Windows(session))
	}

	// var err error
	// var server *tmux.Server
	// if ctx.String("socket-path") != "" {
	// 	if !glaze.FileExists(ctx.String("socket-path")) {
	// 		return fmt.Errorf("the specified socket path could not be found")
	// 	}

	// 	server, err = tmux.NewServerWithSocket(ctx.String("socket-path"), ctx.String("socket-name"))

	// } else {
	// 	server, err = tmux.NewServer()
	// }

	// if err != nil {
	// 	return err
	// }

	// if err := server.Apply(profile); err != nil {
	// 	return err
	// }

	return nil
}
