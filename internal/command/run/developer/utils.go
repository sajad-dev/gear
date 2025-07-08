package developer

import (
	"github.com/fatih/color"
)

func printError(stderr string) {
	color.Red(stderr)
}
