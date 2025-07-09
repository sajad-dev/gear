package developer

import (
	// "bufio"
	// "os/exec"

	// "github.com/fatih/color"
	"log"

	"github.com/sajad-dev/gear/config"
)

// func (a *AutoCompile) Run(close chan int) error {
// 	go func() {
// 		cmd := exec.Command("go", "run", a.ExecPath)

// 		stdout, err := cmd.StdoutPipe()
// 		if err != nil {
// 			color.Red("stdout error: %v", err)
// 		}
// 		stderr, err := cmd.StderrPipe()
// 		if err != nil {
// 			color.Red("stderr error: %v", err)
// 		}

// 		if err := cmd.Start(); err != nil {
// 			color.Blue("Start error: " + err.Error())
// 		}

// 		go func() {
// 			scanner := bufio.NewScanner(stdout)
// 			for scanner.Scan() {
// 				color.Green(scanner.Text())
// 			}
// 		}()
// 		go func() {
// 			scanner := bufio.NewScanner(stderr)
// 			for scanner.Scan() {
// 				color.Red(scanner.Text())
// 			}
// 		}()

// 		done := make(chan error)
// 		go func() {
// 			done <- cmd.Wait()
// 		}()

// 		select {
// 		case <-close:
// 			cmd.Process.Kill()
// 			<-done
// 		case <-done:
// 			color.Red("End Process")
// 		}

//		}()
//		return nil
//	}
func (a *AutoCompile) Run(close chan int, port int) error {
	stop := make(chan struct{})

	go func() {
		<-close
		stop <- struct{}{}
	}()
	err := config.Config.Http(port, config.Config.Db, stop)
	if err != nil {
		log.Println(err)
	}

	return nil
}
