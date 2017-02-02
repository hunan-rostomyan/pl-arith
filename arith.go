package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Main compiles a program written in pl-arith into Go.
// The input is a path to the program. The output will
// be produced in the same path as the program with a
// special suffix.
func main() {

	// Try to get the program name.
	args := os.Args[1:]
	if len(args) == 0 {
		Err("Please specify the name of the program")
	}
	programName := args[0]

	// Read the program text.
	programContents, err := ioutil.ReadFile(programName)
	if err != nil {
		Err(fmt.Sprintf("Unable to read program named '%s'", programName))
	}

	// Display its contents
	fmt.Print(fmt.Sprintf("Program contents: %s", programContents))
}
