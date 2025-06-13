package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)


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

func main() {
	reader := bufio.NewReader(os.Stdin)

	builtins := make(map[string]func([]string))
	builtins["exit"] = exit
	builtins["echo"] = echo
	builtins["type"] = typeCommand(builtins)

	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, _ := reader.ReadString('\n')
		args := strings.Fields(strings.TrimSpace(input))

		cmd := args[0]
		if fn, ok := builtins[cmd]; ok {
			fn(args[1:])
		} else {
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}
