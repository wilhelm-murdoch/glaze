package actions

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/wilhelm-murdoch/glaze"
	"github.com/wilhelm-murdoch/glaze/models"
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

	profile := parser.Decode(glaze.PrimaryGlazeSpec)

	if parser.HasErrors() {
		parser.WriteDiags()
		return nil
	}

	client := tmux.NewClient()

	if client.SessionExists(profile.Name) {
		if err := client.KillSessionByName(profile.Name); err != nil {
			return err
		}
	}

	session, err := client.NewSession(profile.Name, profile.StartingDirectory)
	if err != nil {
		return err
	}

	var outerErr error
	profile.Windows.Each(func(i int, w *models.Window) bool {
		window, err := session.NewWindow(w.Name)
		if err != nil {
			outerErr = err
			return true
		}

		w.Panes.Each(func(i int, p *models.Pane) bool {
			pane, err := window.Split(p.Name, p.Split, p.StartingDirectory)
			if err != nil {
				outerErr = err
				return true
			}

			for _, cmd := range p.Commands {
				time.Sleep(time.Millisecond * time.Duration(100))
				pane.SendKeys(cmd)
			}

			return false
		})

		if err := window.SelectLayout(w.Layout); err != nil {
			outerErr = err
			return true
		}

		return outerErr != nil
	})

	windows, err := client.Windows(session)
	if err != nil {
		return err
	}

	if window, found := windows.At(0); found {
		if err := window.Kill(); err != nil {
			return err
		}

		windows.Shift()

		windows.Each(func(i int, w *tmux.Window) bool {
			panes, err := client.Panes(w)
			if err != nil {
				return true
			}

			if pane, found := panes.At(0); found {
				if err := pane.Kill(); err != nil {
					return true
				}

				panes.Shift()
			}
			return false
		})
	}

	if err := client.Attach(session); err != nil {
		return err
	}

	return nil
}