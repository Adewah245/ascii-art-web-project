package main

import (
	"strings"
)
func generate(lines []string, banner map[rune][]string) string {
	var result strings.Builder
	for _, line := range lines{
		if line == ""{
			result.WriteString("\n")
			continue
		}
		for row := 0; row < 8; row++ {
			for _, ch := range line{
				if art, ok := banner[ch]; ok {
				result.WriteString(art[row])
				}	
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}