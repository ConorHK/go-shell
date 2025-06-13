package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
var builtins = map[string]func([]string){
	"exit": exit,
	"echo": echo,
	"type": typeCommand,
}


func parseCommand(command string) []string {
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

func typeCommand(commands []string) {
	if _, ok := builtins[commands[0]]; ok {
		fmt.Println(commands[0] + " is a shell builtin")
	} else {
		invalidCommand(commands[0])
	}
}

func invalidCommand(command string) {
	var commandNotFound = command + ": command not found"
	fmt.Println(commandNotFound)
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}

		command = strings.TrimSpace(command)
		commandArgs := parseCommand(command)

		if fn, ok := builtins[commandArgs[0]]; ok {
			fn(commandArgs[1:])
		} else {
			invalidCommand(commandArgs[0])
		}
	}
}
