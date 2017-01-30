package main

import (
	"fmt"
	"os"
	//    "os/exec"
)

func main() {
	if len(os.Args) < 2 {
		createfile("main.go")
	} else if len(os.Args) == 2 {
		createfile(os.Args[1] + ".go")
	} else {
		panic("Too many arguments")
	}
}

func createfile(fname string) {
	if _, err := os.Stat("./fname"); err == nil {
		fmt.Println("File does not exist")
	} else {
		fmt.Println("File exists")
	}
}
