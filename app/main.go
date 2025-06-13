package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)


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

func typeCommand(builtins map[string]func([]string)) func([]string) {
	return func(args []string) {
		if len(args) == 0 {
			fmt.Println("type: missing argument")
			return
		}
		if _, ok := builtins[args[0]]; ok {
			fmt.Printf("%s is a shell builtin\n", args[0])
		} else {
			fmt.Printf("%s: not found\n", args[0])
		}
	}
}

func invalidCommand(command string) {
	var commandNotFound = command + ": command not found"
	fmt.Println(commandNotFound)
}

func main() {
	builtins := make(map[string]func([]string))
	builtins["exit"] = exit
	builtins["echo"] = echo
	builtins["type"] = typeCmd(builtins)

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
