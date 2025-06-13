package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func parseCommand(command string) []string {
	command = strings.TrimSpace(command)
	return strings.Fields(command)
}

func exit(exitCodeString string) {
	exitCode, err := strconv.Atoi(exitCodeString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	os.Exit(exitCode)
}

func echo(inputString string) {
	fmt.Println(inputString)
}

func main() {
	builtins := map[string]func(){
		"exit": exit,
		"echo": echo,
	}

	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		commandArgs := parseCommand(command)

		if commandArgs[0] == "exit" {
			exit(commandArgs[1])
		}

		if fn, ok := builtins[args[0]]; ok {
			fn(commandArgs[1])
		} else {
			var commandNotFound string = command + ": command not found"
			fmt.Println(commandNotFound)
		}
	}
}
