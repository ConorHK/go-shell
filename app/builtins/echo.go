package main

import (
	"fmt"
	"strings"
)

func echo(inputStrings []string) {
	outputString := strings.Join(inputStrings, " ")
	fmt.Println(outputString)
}
