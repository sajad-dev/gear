package developer

import (
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func (a *AutoCompile) RunTest() error {
	command := exec.Command("go", "test", a.TestPath)

	var out, stderr strings.Builder

	command.Stdout = &out
	command.Stderr = &stderr

	err := command.Run()

	if err != nil {
		color.Red(stderr.String())
	}
	if strings.Contains(out.String(), "FAIL") {
		color.Yellow(out.String())
		return err
	} else {
		color.Cyan(out.String())
		return err
	}

	return nil
}
