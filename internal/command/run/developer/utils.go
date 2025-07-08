package developer

import (
	"os"

	"github.com/fatih/color"
)

func errorHandel(err error, stderr string) {
	if err != nil {
		color.Red(stderr)
		os.Exit(1)
	}
}
