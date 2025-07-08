package developer

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Mkdir("sample", 0700)
	file, _ := os.Create("sample/main.go")
	defer file.Close()

	file.WriteString(`package main

import "fmt"

func main () {
	fmt.Println("Hello World")
}
	`)
	m.Run()
}
