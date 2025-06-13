package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


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
