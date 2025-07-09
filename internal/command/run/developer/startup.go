package developer

import (
	"bytes"
	"os/exec"
)

type AutoCompile struct {
	ProjectPath string
	ExecPath    string
	TestPath    string
	Test        bool
	Port        int
}

func (a *AutoCompile) GenerateFile() error {
	init := exec.Command("git", "init", a.ProjectPath)

	var stderr bytes.Buffer
	init.Stderr = &stderr

	if err := init.Run(); err != nil {
		printError(stderr.String())
		return err
	}

	// 	if err := os.MkdirAll(a.BuildPath, 0755); err != nil {
	// 		printError(err.Error())
	// 		return err
	// 	}

	return nil
}
