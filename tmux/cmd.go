package tmux

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Command struct {
	args []string
	cmd  *exec.Cmd
}

type CommandError struct {
	args []string
	err  error
}

func NewCommandError(args []string, err error) CommandError {
	return CommandError{
		args: args,
		err:  err,
	}
}

func (ce CommandError) Error() string {
	return fmt.Sprintf(`Error: "%s" Command: "%s"`, ce.err, strings.Join(ce.args, " "))
}

func NewCommand(args ...string) (Command, error) {
	tmux, ok := IsInstalled()
	if !ok {
		return Command{}, fmt.Errorf("tmux is not installed")
	}

	args = append([]string{tmux}, args...)

	return Command{
		args: args,
		cmd:  exec.Command(args[0], args[1:]...),
	}, nil
}

func (c Command) String() string {
	return strings.Join(c.args, " ")
}

func (c Command) Exec() error {
	c.cmd.Stdin = os.Stdin
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr

	if err := c.cmd.Run(); err != nil {
		return NewCommandError(c.args, err)
	}

	return nil
}

func (c Command) ExecWithOutput() (string, error) {
	output, err := c.cmd.CombinedOutput()
	if err != nil {
		return "", NewCommandError(c.args, err)
	}

	return strings.TrimSuffix(string(output), "\n"), nil
}
