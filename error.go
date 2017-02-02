package main

import (
	"fmt"
	"os"
)

func Err(message string) {
	fmt.Fprintf(os.Stderr, message+".\n")
	os.Exit(1)
}
