package stringutils

import (
	"fmt"
	"strings"
)

func JoinWithOr(choices []string) string {
	return joinWith(choices, "or")
}

func JoinWithAnd(choices []string) string {
	return joinWith(choices, "and")
}

func joinWith(choices []string, conjunction string) string {
	length := len(choices)
	switch length {
	case 0:
		return ""
	case 1:
		return choices[0]
	case 2:
		return fmt.Sprintf(`%s %s %s`, choices[0], conjunction, choices[1])
	}

	return fmt.Sprintf(`%s %s %s`, strings.Join(choices[:length-1], ", "), conjunction, choices[length-1])
}
