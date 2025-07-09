package developer

import (
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func (a *AutoCompile) Chenges() (bool, error) {
	command := exec.Command("git", "status")
	// command.Dir = a.ProjectPath
	color.Red(a.ProjectPath)
	

	statusOutput, err := command.Output()
	if err != nil {
		printError(err.Error())
		return false, err
	}

	strStatusOutput := string(statusOutput)

	if strings.Contains(strStatusOutput, "Changes not staged for commit:") || strings.Contains(strStatusOutput, "Untracked files:") {
		commandGit := exec.Command("git", "add", ".")
		commandGit.Dir = a.ProjectPath

		if err := commandGit.Run(); err != nil {
			printError(err.Error())
			return false, err
		}

		return true, nil
	}
	return false, nil
}
