package run

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/sajad-dev/gear/config"
	"github.com/sajad-dev/gear/internal/command/run/developer"
	"github.com/spf13/cobra"
)

func runDeveloper(port int) error {
	ou := exec.Command("pwd")
	addr, _ := ou.Output()
	addrString := strings.TrimSpace(string(addr))
	color.Red(string(addr))

	st := developer.AutoCompile{
		Port:        port,
		ProjectPath: addrString,
		ExecPath:    fmt.Sprintf("%s/cmd/main.go", addrString),
	}
	err := st.GenerateFile()
	if err != nil {
		return err
	}

	err = st.Loop()
	if err != nil {
		return err
	}

	return nil

}
func runProduction(port int) error {
	stop := make(chan struct{})
	err := config.Config.Http(port, config.Config.Db,stop)
	if err != nil {
		return err
	}
	return nil
}

func Http(cmd *cobra.Command, args []string) error {
	// port, err := cmd.Flags().GetInt("port")
	// if err != nil {
	// 	return err
	// }
	port := 8000

	switch args[0] {
	case "dev":
		runDeveloper(port)
	case "production":
		runProduction(port)
	default:
		color.Red(fmt.Sprintf("Not exsis %s", args[0]))
	}

	return nil
}
