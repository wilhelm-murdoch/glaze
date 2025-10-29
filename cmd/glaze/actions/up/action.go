package up

import (
	"fmt"
	"time"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"

	"github.com/wilhelm-murdoch/glaze/cmd/glaze/actions"
	"github.com/wilhelm-murdoch/glaze/internal/parser"
	"github.com/wilhelm-murdoch/glaze/internal/schema"
	"github.com/wilhelm-murdoch/glaze/internal/schema/session"
	"github.com/wilhelm-murdoch/glaze/internal/tmux"
)

type Action struct {
	actions.BaseAction
	client *tmux.Client
}

func NewAction(ctx *cli.Context) (*Action, error) {
	base, err := actions.NewBaseAction(ctx)
	if err != nil {
		return nil, err
	}

	client := tmux.NewClient(
		ctx.String("socket-name"),
		ctx.String("socket-path"),
		ctx.Bool("debug"),
	)

	return &Action{
		BaseAction: *base,
		client:     &client,
	}, nil
}

func (a *Action) Run() error {
	variables, err := parser.CollectVariables(a.Context.StringSlice("var"))
	if err != nil {
		return fmt.Errorf("could not parse specified variables: %s", err)
	}

	profile, decodeDiags := a.Parser.Decode(
		schema.PrimaryGlazeSpec,
		parser.BuildEvalContext(variables),
	)

	if decodeDiags.HasErrors() {
		a.DiagnosticsManager.Extend(decodeDiags)
		return a.DiagnosticsManager.Write()
	}

	session, err := a.resolveSession(profile)
	if err != nil {
		return err
	}

	// Iterate through the windows and panes defined within the specified profile and create them within the tmux session.
	for _, ws := range profile.Windows.Items() {
		log.Info("creating new window", "window", ws.Name)
		wtmx, err := session.NewWindow(ws.Name)
		if err != nil {
			return fmt.Errorf("could not create new window `%s`: %s", ws.Name, err)
		}

		panes, err := a.client.Panes(wtmx)
		if err != nil {
			return fmt.Errorf("could not read panes for window `%s`: %s", wtmx.Name, err)
		}

		defaultPane := panes.Find(func(i int, item *tmux.Pane) bool {
			return item.IsFirst
		})

		if defaultPane == nil {
			return fmt.Errorf("could not locate default pane for window `%s`", wtmx.Name)
		}

		// Panes are originally parsed and created in the reverse order of how they are
		// defined within the glaze definition file. So, we'll just reverse them here to
		// set them back to the user-defined order.
		for _, ps := range ws.Panes.Reverse().Items() {
			log.Info("adding pane", "pane", ps.Name, "from", defaultPane.Target())
			ptmx, err := wtmx.Split(defaultPane.Target(), ps.Name, ps.StartingDirectory)
			if err != nil {
				return fmt.Errorf(
					"could not split pane `%d` for window `%s`: %s",
					defaultPane.Index,
					wtmx.Name,
					err,
				)
			}

			// Run any defined commands in order as defined within the
			// current profile. Add a small delay between each command
			// to ensure they are executed in order.
			for _, cmd := range ps.Commands {
				log.Info("sending command", "pane", ptmx.Name, "cmd", cmd)
				time.Sleep(time.Millisecond * time.Duration(100))
				if err := ptmx.SendKeys(cmd); err != nil {
					return fmt.Errorf(
						"could not execute command `%s` for pane `%s` in window `%s`: %s",
						cmd,
						ptmx.Name,
						wtmx.Name,
						err,
					)
				}
			}

			if ps.Focus {
				log.Info("setting focus", "pane", ptmx.Name)
				ptmx.Select()
			}
		}

		// Remove the default pane directly from the session.
		defaultPane.Kill()

		if err := wtmx.SelectLayout(ws.Layout); err != nil {
			return fmt.Errorf(
				"could not select layout `%s` for window `%s`: %s",
				ws.Layout,
				wtmx.Name,
				err,
			)
		}

		if ws.Focus {
			log.Info("setting focus", "window", wtmx.Name)
			wtmx.Select()
		}

	}

	if err := a.removeDefaultWindow(session); err != nil {
		return err
	}

	if !a.Context.Bool("detached") {
		if err := a.client.Attach(session); err != nil {
			return err
		}
	}

	return nil
}

func (a *Action) generateWindows() {}

// Tmux creates a default window with a default pane for every
// session. Remove the defaults so only windows and panes defined
// within the profile are left.
func (a *Action) removeDefaultWindow(session *tmux.Session) error {
	windows, err := a.client.Windows(session)
	if err != nil {
		return fmt.Errorf("could not read windows for session `%s`: %s", session.Name, err)
	}

	defaultWindow := windows.Find(func(i int, window *tmux.Window) bool {
		return window.IsFirst
	})

	if defaultWindow != nil {
		if err := defaultWindow.Kill(); err != nil {
			return err
		}
	}

	return nil
}

func (a *Action) generatePanes() {}

func (a *Action) removeDefaultPane() {}

func (a *Action) resolveSession(profile *session.Session) (*tmux.Session, error) {
	if a.Context.Bool("clear") {
		log.Info("clearing previous session", "session", profile.Name)
		a.client.KillSessionByName(profile.Name)
	}

	if a.client.SessionExists(profile.Name) {
		session, err := a.client.FindSessionByName(profile.Name)
		if err != nil {
			return nil, fmt.Errorf("could not find session `%s`: %s", profile.Name, err)
		}

		if !a.Context.Bool("detached") {
			log.Info("attaching to existing session", "session", profile.Name)
			if err := a.client.Attach(session); err != nil {
				return nil, fmt.Errorf("could not attach to session `%s`: %s", session.Name, err)
			}
		}

		return nil, nil
	}

	log.Info("creating new session", "session", profile.Name)
	session, err := a.client.NewSession(profile.Name, profile.StartingDirectory)
	if err != nil {
		return nil, fmt.Errorf("could not create new session `%s`: %s", session.Name, err)
	}

	return session, nil
}
