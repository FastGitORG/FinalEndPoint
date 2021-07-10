package main

import (
	"fmt"
	"os"
)

func readAllText(path string) (string, bool) {
	file, err := os.Open(path)
	if err != nil {
		return "", false
	}
	return fmt.Sprintf("%v", file), true
}
