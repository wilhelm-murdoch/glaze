package tmux

import (
	"fmt"
	"strings"

	"github.com/wilhelm-murdoch/glaze/tmux/enums"
	"github.com/wilhelm-murdoch/go-collection"
)

type Option[OT enums.OptionTyper[OT]] struct {
	Name  OT
	Value string
}

func (o Option[OT]) String() string {
	return fmt.Sprintf("%s %s", o.Name, o.Value)
}

func showOptions[OT enums.OptionTyper[OT]](client Client, args ...string) (collection.Collection[Option[OT]], error) {
	var out collection.Collection[Option[OT]]

	cmd, err := NewCommand(client, args...)
	if err != nil {
		return out, err
	}

	output, err := cmd.ExecWithOutput()
	if err != nil {
		return out, err
	}

	for _, line := range strings.Split(strings.TrimSpace(output), "\n") {
		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			return out, fmt.Errorf("could not parse output of option `%s`", line)
		}

		var name OT

		out.Push(Option[OT]{
			Name:  name.FromString(parts[0]),
			Value: parts[1],
		})
	}

	return out, nil
}

func getOption[OT enums.OptionTyper[OT]](client Client, args ...string) (Option[OT], error) {
	var out Option[OT]

	cmd, err := NewCommand(client, args...)
	if err != nil {
		return out, err
	}

	output, err := cmd.ExecWithOutput()
	if err != nil {
		return out, err
	}

	parts := strings.SplitN(output, " ", 2)
	if len(parts) != 2 {
		return out, fmt.Errorf("could not parse output of option `%s`", output)
	}

	var name OT

	out.Name = name.FromString(parts[0])
	out.Value = parts[1]

	return out, nil
}

func setOption[OT enums.OptionTyper[OT]](client Client, args ...string) error {
	cmd, err := NewCommand(client, args...)
	if err != nil {
		return err
	}

	return cmd.Exec()
}
