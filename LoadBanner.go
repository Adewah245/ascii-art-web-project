package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filesname string)(map[rune][]string, error) {
	data, err := os.ReadFile(filesname)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("Ops!, File Opened Is Empty..")
	}
	text := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(text, "\n")
	if len(lines) != 855 {
		return nil, errors.New("Ops!, File Opened Has Incomplete Banner Lines")
	}
	result := make(map[rune][]string)
	for i := 0; i < 95; i++ {
			start := i*9+1
			r := rune(i+32)
			result[r] = lines[start:start+8]
	}
	return result, nil
}