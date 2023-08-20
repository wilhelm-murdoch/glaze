package actions

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

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

	// Iterate through the windows and panes defined within the specified
	// profile and create them within the tmux session.
	for _, wm := range profile.Windows.Items() {
		wc, err := session.NewWindow(wm.Name)
		if err != nil {
			return err
		}

		for _, pm := range wm.Panes.Items() {
			pc, err := wc.Split(pm.Name, pm.Split, pm.StartingDirectory, pm.Size, pm.Placement, pm.Full)
			if err != nil {
				return err
			}

			// Run any defined commands in order as defined within the
			// current the profile. Add a small delay between each command
			// to ensure they are executed in order.
			for _, cmd := range pm.Commands {
				time.Sleep(time.Millisecond * time.Duration(100))
				pc.SendKeys(cmd)
			}
		}

		if err := wc.SelectLayout(wm.Layout); err != nil {
			return err
		}
	}

	// Tmux creates a default window with a default pane for every
	// session. Remove the defaults so only windows and panes defined
	// within the profile are left.
	windows, err := client.Windows(session)
	if err != nil {
		return err
	}

	if window, found := windows.At(0); found {
		if err := window.Kill(); err != nil {
			return err
		}

		windows.Shift()

		for _, wc := range windows.Items() {
			pc, err := client.Panes(wc)
			if err != nil {
				return err
			}

			if pane, found := pc.At(0); found {
				if err := pane.Kill(); err != nil {
					return err
				}

				pc.Shift()
			}
		}
	}

	// if err := client.Attach(session); err != nil {
	// 	return err
	// }

	return nil
}
