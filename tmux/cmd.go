package tmux

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Command represents a command to run within a tmux session.
type Command struct {
	args []string
	cmd  *exec.Cmd
}

// CommandError represents an error that occurred while running a command.
type CommandError struct {
	args []string
	err  error
}

// NewCommandError returns a new command error.
func NewCommandError(args []string, err error) CommandError {
	return CommandError{
		args: args,
		err:  err,
	}
}

// Error returns the error message.
func (ce CommandError) Error() string {
	return fmt.Sprintf(`Error: "%s" Command: "%s"`, ce.err, strings.Join(ce.args, " "))
}

// NewCommand returns a new command with the given arguments.
func NewCommand(args ...string) (Command, error) {
	ok, tmux := IsInstalled()
	if !ok {
		return Command{}, fmt.Errorf("tmux is not installed")
	}

	args = append([]string{tmux}, args...)

	return Command{
		args: args,
		cmd:  exec.Command(args[0], args[1:]...),
	}, nil
}

// String returns the full command with arguments as a string.
func (c Command) String() string {
	return strings.Join(c.args, " ")
}

// Exec executes the command and returns an error if one occurred.
func (c Command) Exec() error {
	c.cmd.Stdin = os.Stdin
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr

	if err := c.cmd.Run(); err != nil {
		return NewCommandError(c.args, err)
	}

	return nil
}

// ExecWithOutput executes the command and returns the output as a string.
func (c Command) ExecWithOutput() (string, error) {
	output, err := c.cmd.CombinedOutput()
	if err != nil {
		return "", NewCommandError(c.args, err)
	}

	return strings.TrimSuffix(string(output), "\n"), nil
}
