package main

import "strings"

func GenerateAscii(lines []string, banner map[rune][]string) string{
	var result strings.Builder
	for _, line := range lines {
		if line == "" {
			result.WriteString("\n")
			continue
		}
		for i := 0; i < 8; i++ {
			for _, ch := range line{
				if art, exist := banner[ch]; exist {
					result.WriteString(art[i])
				}
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}