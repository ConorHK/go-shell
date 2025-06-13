package main

import (
	"fmt"
	"os"
	"strings"
	"errors"
)

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
}
