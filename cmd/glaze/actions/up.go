package actions

import (
	"fmt"
	"time"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
	"github.com/wilhelm-murdoch/glaze"
	"github.com/wilhelm-murdoch/glaze/tmux"
)

func Up(ctx *cli.Context) error {
	profilePath, err := glaze.ResolveProfilePath(ctx.Args().First())
	if err != nil {
		return err
	}

	diagsManager := glaze.NewDiagnosticsManager(profilePath)
	if diagsManager.HasErrors() {
		return diagsManager.Write()
	}

	parser, parserDiags := glaze.NewParser(profilePath)
	if parserDiags.HasErrors() {
		diagsManager.Extend(parserDiags)
		return diagsManager.Write()
	}

	variables, err := glaze.CollectVariables(ctx.StringSlice("var"))
	if err != nil {
		return fmt.Errorf("could not parse specified variables: %s", err)
	}

	profile, decodeDiags := parser.Decode(glaze.PrimaryGlazeSpec, glaze.BuildEvalContext(variables))
	if decodeDiags.HasErrors() {
		diagsManager.Extend(decodeDiags)
		return diagsManager.Write()
	}

	client := tmux.NewClient(
		ctx.String("socket-name"),
		ctx.String("socket-path"),
		ctx.Bool("debug"),
	)

	if ctx.Bool("clear") {
		log.Info("clearing previous session", "session", profile.Name)
		client.KillSessionByName(profile.Name)
	}

	if client.SessionExists(profile.Name) {
		session, err := client.FindSessionByName(profile.Name)
		if err != nil {
			return fmt.Errorf("could not find session `%s`: %s", profile.Name, err)
		}

		if !ctx.Bool("detached") {
			log.Info("attaching to existing session", "session", profile.Name)
			if err := client.Attach(session); err != nil {
				return fmt.Errorf("could not attach to session `%s`: %s", session.Name, err)
			}
		}

		return nil
	}

	log.Info("creating new session", "session", profile.Name)
	session, err := client.NewSession(profile.Name, profile.StartingDirectory)
	if err != nil {
		return fmt.Errorf("could not create new session `%s`: %s", session.Name, err)
	}

	for option, value := range profile.Options {
		log.Info("... setting option", "session", session.Name, option, value)
		if err := session.SetOption(option, value); err != nil {
			return fmt.Errorf("could not set option `%s` with value `%s` for session `%s`: %s", option, value, session.Name, err)
		}
	}

	// Iterate through the windows and panes defined within the specified profile and create them within the tmux session.
	for _, wm := range profile.Windows.Items() {
		log.Info("... creating new window", "window", wm.Name)
		wc, err := session.NewWindow(wm.Name)
		if err != nil {
			return fmt.Errorf("could not create new window `%s`: %s", wm.Name, err)
		}

		for option, value := range wm.Options {
			log.Info("... setting option", "window", wm.Name, option, value)
			if err := wc.SetOption(option, value); err != nil {
				return fmt.Errorf("could not set option `%s` with value `%s` for window `%s`: %s", option, value, session.Name, err)
			}
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
			log.Info("... ... adding pane", "pane", pm.Name, "from", defaultPane.Target())
			pc, err := wc.Split(defaultPane.Target(), pm.Name, pm.StartingDirectory)
			if err != nil {
				return fmt.Errorf("could not split pane `%d` for window `%s`: %s", defaultPane.Index, wc.Name, err)
			}

			// Run any defined commands in order as defined within the
			// current the profile. Add a small delay between each command
			// to ensure they are executed in order.
			for _, cmd := range pm.Commands {
				log.Info("... ... sending command", "pane", pc.Name, "cmd", cmd)
				time.Sleep(time.Millisecond * time.Duration(100))
				if err := pc.SendKeys(cmd); err != nil {
					return fmt.Errorf("could not execute command `%s` for pane `%s` in window `%s`: %s", cmd, pc.Name, wc.Name, err)
				}
			}

			for option, value := range pm.Options {
				log.Info("... ... setting option", "pane", pc.Name, option, value)
				if err := pc.SetOption(option, value); err != nil {
					return fmt.Errorf("could not set option `%s` with value `%s` for pane `%s` in window `%s`: %s", option, value, pc.Name, wc.Name, err)
				}
			}

			if pm.Focus {
				log.Info("... ... setting focus", "pane", pc.Name)
				pc.Select()
			}
		}

		if err := wc.SelectLayout(wm.Layout); err != nil {
			return fmt.Errorf("could not select layout `%s` for window `%s`: %s", wm.Layout, wc.Name, err)
		}

		if wm.Focus {
			log.Info("... setting focus", "window", wc.Name)
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
