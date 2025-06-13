package main

import (
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/utils"
)

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
