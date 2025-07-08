package developer

import (
	"bytes"
	"os/exec"
)

type AutoCompile struct {
	ProjectPath string
}

func (a *AutoCompile) GenerateFile() {
	init := exec.Command("git", "init")

	var stderr bytes.Buffer
	init.Stderr = &stderr

	err := init.Run()
	errorHandel(err, stderr.String())
}
