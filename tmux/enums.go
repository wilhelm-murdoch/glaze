package tmux

import (
	"os"
	"strings"
)

const (
	SplitHorizontal = "horizontal"
	SplitVertical   = "vertical"

	LayoutEventHorizontal = "even-horizontal"
	LayoutEvenVertical    = "even-vertical"
	LayoutMainHorizontal  = "main-horizontal"
	LayoutMainVertical    = "main-vertical"
	LayoutTiled           = "tiled"
	LayoutUnknown         = "unknown"
)

var (
	SplitList = []string{
		SplitHorizontal,
		SplitVertical,
	}

	LayoutList = []string{
		LayoutEventHorizontal,
		LayoutEvenVertical,
		LayoutMainHorizontal,
		LayoutMainVertical,
		LayoutTiled,
	}
)

func Contains(list []string, value string) bool {
	for _, v := range list {
		if string(v) == value {
			return true
		}
	}

	return false
}

func ExpandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		userHome, err := os.UserHomeDir()
		if err != nil {
			return path
		}

		return strings.Replace(path, "~", userHome, 1)
	}

	return path
}
