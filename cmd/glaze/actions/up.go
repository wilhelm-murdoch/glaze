package actions

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/k0kubun/pp"
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

		if err := window.SelectLayout(w.Layout); err != nil {
			outerErr = err
			return true
		}

		w.Panes.Each(func(i int, p *models.Pane) bool {
			pane, err := window.Split(p.Name, p.Split, p.StartingDirectory)
			pp.Print(pane, err)
			return false
		})

		return false
	})

	if outerErr != nil {
		return outerErr
	}

	windows, err := client.Windows(session)
	if err != nil {
		return err
	}

	// Kill the first window and associated pane.
	windows.Each(func(i int, w tmux.Window) bool {
		if w.Index == 0 {
			w.Kill()
			return true
		}

		return false
	})

	if err := client.Attach(session); err != nil {
		return err
	}

	// if t.SessionExists(profile.Name) && !profile.ReattachOnStart {
	// 	if _, err := t.KillSession(profile.Name); err != nil {
	// 		return err
	// 	}

	// 	if _, err := t.NewSession(profile.Name, profile.StartingDirectory); err != nil {
	// 		return err
	// 	}
	// } else if !t.SessionExists(profile.Name) {
	// 	if _, err := t.NewSession(profile.Name, profile.StartingDirectory); err != nil {
	// 		return err
	// 	}
	// }

	// every new window starts with a single pane
	//  - find new pane index
	//  - rename pane
	//
	// create new panes by splitting the target window
	// var err error
	// profile.Windows.Each(func(i int, w *models.Window) bool {
	// 	o, _ := t.NewWindow(w.Name)
	// 	fmt.Println(o)
	// 	return false
	// })

	// if err != nil {
	// 	return err
	// }

	// if _, err := t.AttachToSession(profile.Name); err != nil {
	// 	return err
	// }
	// sessions, _ := t.Sessions()

	// for _, session := range sessions {
	// 	fmt.Println("Session Name:", session.Name)
	// 	windows, _ := t.Windows(session)
	// 	for _, window := range windows {
	// 		fmt.Println("... Window Name:", window.Name)
	// 		panes, _ := t.Panes(window)
	// 		for _, pane := range panes {
	// 			fmt.Println("... ... Pane Name:", pane.Id, pane.Name)
	// 		}
	// 	}
	// }

	return nil
}
