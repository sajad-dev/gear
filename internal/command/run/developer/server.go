package developer

import (
	"bytes"
	"os/exec"
)

func (a *AutoCompile) Run() error {
	runServer := exec.Command("go", "run", a.ExecPath)

	var stderr, stdout bytes.Buffer
	runServer.Stderr = &stderr
	runServer.Stdout = &stdout

	if err := runServer.Run(); err != nil {
		printError(stderr.String())
		return err
	}

	return nil
}
