package debug

import (
	"fmt"

	"github.com/kr/pretty"
)

// Prettier is a utility function that takes any given value and
// attempts to print it in a human-readable format to stdout. Used
// primarily for debugging.
func Prettier(values ...any) {
	for _, value := range values {
		fmt.Printf("%# v\n", pretty.Formatter(value))
	}
}
