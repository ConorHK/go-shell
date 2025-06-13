package main

import (
	"fmt"
	"os"
)

func exit(exitCodeStrings []string) {
	exitCode, err := strconv.Atoi(exitCodeStrings[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	os.Exit(exitCode)
}
