## Make a standard go template ##
Running "gomake" will create a main.go file in the current directory which contains:

	package main

	import (
    "fmt"
	)

	func main() {
	}
Running "gomake filename" will create filename.go in the current directory with the above output.
The file will not be created if the name already exists or lack of permissions.
