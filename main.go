package main

import "github.com/sajad-dev/gear/internal/command/run/developer"

func main() {
	st := developer.AutoCompile{
		ExecPath: "/home/sajad/Documents/Programming/gingo-framework/gear/internal/command/run/developer/sample/main.go",
		ProjectPath: "/home/sajad/Documents/Programming/gingo-framework/gear/internal/command/run/developer/sample",
		Test: true,
		TestPath: "/home/sajad/Documents/Programming/gingo-framework/gear/internal/command/run/developer/sample/...",
	}
	st.GenerateFile()
	st.Loop()

}
