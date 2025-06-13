package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"errors"
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

		cmd := args[0]
		if _, ok := builtins[cmd]; ok {
			fmt.Printf("%s is a shell builtin\n", cmd)
		} else {
			absolutePath, err := searchDirectoriesForFile(pathDirectories(), cmd)
			if err == nil {
				fmt.Printf("%s is %s\n", cmd, absolutePath)
			} else {
				fmt.Printf("%s: not found\n", cmd)
			}
		}
	}
}

func pathDirectories() []string {
	pathStr, exists := os.LookupEnv("PATH")
	if !exists {
		return []string{}
	}
	return strings.Split(pathStr, ":")
}

func searchDirectoriesForFile(directories []string, fileName string) (string, error) {
	for _, directory := range directories {
		possibleFile := directory + "/" + fileName
		if ok := fileExists(possibleFile); ok {
			return possibleFile, nil
		}
	}
	return "", fmt.Errorf("not found")
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)

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
