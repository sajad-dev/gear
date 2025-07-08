package run

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/sajad-dev/gear/internal/config"
	"github.com/spf13/cobra"
)

func runDeveloper(port int) {
	// ou := exec.Command("pwd")
	// addr, _ := ou.Output()

	// com := Compiler{Addr: strings.TrimSpace(string(addr)), Port: 8000, BuildPath: "/home/sajad/Documents/Programming/gingo-framework/gear/build/gear", ExecuteFilePath: "/home/sajad/Documents/Programming/gingo-framework/gear/cmd/main.go"}

	// // color.Blue("Adress Local : http://127.0.0.1:8000")

	// // fmt.Println(com)

	// com.Run()
	// com.RunTest()

	// com.Compile()

}
func runProduction(port int) error {
	_, err := config.Config.Http(port, config.Config.Db)
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
