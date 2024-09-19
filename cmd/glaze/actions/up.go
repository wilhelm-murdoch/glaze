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
			return fmt.Errorf("could not read current working directory: %s", err)
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
		return parser.WriteDiags()
	}

	profile := parser.Decode(glaze.PrimaryGlazeSpec)

	if parser.HasErrors() {
		return parser.WriteDiags()
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
			return fmt.Errorf("could not find session `%s`: %s", profile.Name, err)
		}

		if !ctx.Bool("detached") {
			if err := client.Attach(session); err != nil {
				return fmt.Errorf("could not attach to session `%s`: %s", session.Name, err)
			}
		}

		return nil
	}

	session, err := client.NewSession(profile.Name, profile.StartingDirectory)
	if err != nil {
		return fmt.Errorf("could create new session `%s`: %s", session.Name, err)
	}

	// Iterate through the windows and panes defined within the specified profile and create them within the tmux session.
	for _, wm := range profile.Windows.Items() {
		wc, err := session.NewWindow(wm.Name)
		if err != nil {
			return fmt.Errorf("could not create new window `%s`: %s", wm.Name, err)
		}

		panes, err := client.Panes(wc)
		if err != nil {
			return fmt.Errorf("could not read panes for window `%s`: %s", wc.Name, err)
		}

		defaultPane := panes.Find(func(i int, item *tmux.Pane) bool {
			return item.IsFirst
		})

		if defaultPane == nil {
			return fmt.Errorf("could not locate default pane for window `%s`", wc.Name)
		}

		for _, pm := range wm.Panes.Items() {
			pc, err := wc.Split(defaultPane.Target(), pm.Name, pm.StartingDirectory)
			if err != nil {
				return fmt.Errorf("could not split pane `%s` for window `%s`: %s", defaultPane.Name, wc.Name, err)
			}

			// Run any defined commands in order as defined within the
			// current the profile. Add a small delay between each command
			// to ensure they are executed in order.
			for _, cmd := range pm.Commands {
				time.Sleep(time.Millisecond * time.Duration(100))
				if err := pc.SendKeys(cmd); err != nil {
					return fmt.Errorf("could not execute command `%s` for pane `%s` in window `%s`: %s", cmd, pc.Name, wc.Name, err)
				}
			}

			if pm.Focus {
				pc.Select()
			}
		}

		if err := wc.SelectLayout(wm.Layout); err != nil {
			return fmt.Errorf("could not select layout `%s` for window `%s`: %s", wm.Layout, wc.Name, err)
		}

		if wm.Focus {
			wc.Select()
		}

		// Remove the default pane directly from the session.
		defaultPane.Kill()
	}

	// Tmux creates a default window with a default pane for every
	// session. Remove the defaults so only windows and panes defined
	// within the profile are left.
	windows, err := client.Windows(session)
	if err != nil {
		return fmt.Errorf("could not read windows for session `%s`: %s", session.Name, err)
	}

	defaultWindow := windows.Find(func(i int, window *tmux.Window) bool {
		return window.IsFirst
	})

	if defaultWindow != nil {
		defaultWindow.Kill()
	}

	if !ctx.Bool("detached") {
		if err := client.Attach(session); err != nil {
			return err
		}
	}

	return nil
}
