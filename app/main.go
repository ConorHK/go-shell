package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		command = strings.TrimSpace(command)
		args := strings.Fields(command)
		if args[0] == "exit" {
			os.Exit(args[1])
		}
		var output = command[:len(command)-1] + ": command not found"
		fmt.Println(output)
	}
}
