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

func exit(exitCodeStrings []string) {
	exitCode, err := strconv.Atoi(exitCodeStrings[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	os.Exit(exitCode)
}

func echo(inputStrings []string) {
	outputString := strings.Join(inputStrings, " ")
	fmt.Println(outputString)
}

func main() {
	builtins := map[string]func(string){
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
			exit(commandArgs[1:])
		}

		if fn, ok := builtins[commandArgs[0]]; ok {
			fn(commandArgs[1:])
		} else {
			var commandNotFound string = command + ": command not found"
			fmt.Println(commandNotFound)
		}
	}
}
