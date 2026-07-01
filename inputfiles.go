package main

import (
	"strings"
)

func inputFiles(input string) []string {
	if input == "" {
		return []string{}
	}
	text := strings.ReplaceAll(input, "\\n", "\n")
	if text == "\n" {
		return []string{"\n"}
	}
	return strings.Split(text, "\n")
}
