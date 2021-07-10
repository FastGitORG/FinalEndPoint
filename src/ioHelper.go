package main

import (
	"fmt"
	"os"
)

func readAllText(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v", file)
}
