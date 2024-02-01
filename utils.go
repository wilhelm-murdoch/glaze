package glaze

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/kr/pretty"
)

func FileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil || errors.Is(err, fs.ErrNotExist) || fileInfo.IsDir() {
		return false
	}

	return true
}

func JoinWithOr(choices []string) string {
	return joinWith(choices, "or")
}

func JoinWithAnd(choices []string) string {
	return joinWith(choices, "and")
}

func joinWith(choices []string, conjunction string) string {
	length := len(choices)
	if length == 0 {
		return ""
	} else if length == 1 {
		return choices[0]
	} else if length == 2 {
		return fmt.Sprintf(`%s %s %s`, choices[0], conjunction, choices[1])
	}

	return fmt.Sprintf(`%s %s %s`, strings.Join(choices[:length-1], ", "), conjunction, choices[length-1])
}

func Prettier(values ...any) {
	for _, value := range values {
		fmt.Printf("%# v\n", pretty.Formatter(value))
	}
}
