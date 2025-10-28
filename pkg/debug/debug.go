package debug

import (
	"fmt"

	"github.com/kr/pretty"
)

func Prettier(values ...any) {
	for _, value := range values {
		fmt.Printf("%# v\n", pretty.Formatter(value))
	}
}
