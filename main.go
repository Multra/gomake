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
	if _, err := os.Stat("./"+fname); os.IsNotExist(err) {
		fmt.Println(fname +" created")
        file, err := os.Create(fname)
        if err != nil {
            panic(err)
        }
        fmt.Fprintf(file,"package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n}")
        file.Close()
	} else {
		fmt.Println(fname+" already exists, exiting")
	}
}
