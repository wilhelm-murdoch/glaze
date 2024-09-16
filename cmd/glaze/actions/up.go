package actions

import (
	"errors"
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

	if ctx.String("socket-name") != "" && ctx.String("socket-path") != "" {
		return errors.New("cannot specify both --socket-name and --socket-path flags")
	}

	if ctx.String("socket-path") != "" {
		if !glaze.FileExists(ctx.String("socket-path")) {
			return fmt.Errorf("specified --socket-path of %s does not exist", ctx.String("socket-path"))
		}
	}

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
		return fmt.Errorf("profile `%s` not found on the specified path, the current directory, or the GLAZE_PATH environment variable", profilePath)
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

	client := tmux.NewClient(
		ctx.String("socket-name"),
		ctx.String("socket-path"),
		ctx.Bool("debug"),
	)

	if ctx.Bool("clear") {
		client.KillSessionByName(profile.Name)
	}

	if client.SessionExists(profile.Name) {
		session, err := client.FindSessionByName(profile.Name)
		if err != nil {
			return err
		}

		if !ctx.Bool("detached") {
			if err := client.Attach(session); err != nil {
				return err
			}
		}

		return nil
	}

	session, err := client.NewSession(profile.Name, profile.StartingDirectory)
	if err != nil {
		return err
	}

	glaze.Prettier(profile)

	// Iterate through the windows and panes defined within the specified profile and create them within the tmux session.
	for _, wm := range profile.Windows.Items() {
		wc, err := session.NewWindow(wm.Name)
		if err != nil {
			return err
		}

		panes, _ := client.Panes(wc)
		defaultPane := panes.Find(func(i int, item *tmux.Pane) bool {
			return item.IsFirst
		})

		for _, pm := range wm.Panes.Items() {
			pc, err := wc.Split(pm.Split, pm.Placement, pm.Full, defaultPane.Name, pm.Name, pm.StartingDirectory, pm.Size)
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
	}

	// if err := wc.SelectLayout(wm.Layout); err != nil {
	// 	return err
	// }
	// }

	// Tmux creates a default window with a default pane for every
	// session. Remove the defaults so only windows and panes defined
	// within the profile are left.
	// windows, err := client.Windows(session)
	// if err != nil {
	// 	return err
	// }

	// defaultWindow := windows.Find(func(i int, window *tmux.Window) bool {
	// 	return window.Index == 0
	// })

	// if defaultWindow != nil {
	// 	defaultWindow.Kill()
	// 	windows.Shift()
	// }

	// for _, window := range windows.Items() {
	// 	panes, err := client.Panes(window)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	if panes.Length() > 1 {
	// 		defaultPane := panes.Find(func(i int, pane *tmux.Pane) bool {
	// 			return pane.Index == 0
	// 		})

	// 		if defaultPane != nil {
	// 			defaultPane.Kill()
	// 			panes.Shift()
	// 		}
	// 	}
	// }

	if !ctx.Bool("detached") {
		if err := client.Attach(session); err != nil {
			return err
		}
	}

	return nil
}
