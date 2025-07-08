package developer

import (
	"os/exec"
	"strings"
)

func (a *AutoCompile) HandelChange() (bool, error) {
	command := exec.Command("git", "status", a.ProjectPath)
	statusOutput, err := command.Output()
	if err != nil {
		printError(err.Error())
		return false, err
	}

	strStatusOutput := string(statusOutput)

	if strings.Contains(strStatusOutput, "Changes not staged for commit:") || strings.Contains(strStatusOutput, "Untracked files:") {
		commandGit := exec.Command("git", "add", ".", a.ProjectPath)

		if err := commandGit.Run(); err != nil {
			printError(err.Error())
			return false, err
		}
		return true, nil
	}
	return false, nil
}
