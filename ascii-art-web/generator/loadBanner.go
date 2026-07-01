package main

import (
	"errors"
	"os"
	"strings"
)

var banner = make(map[string]map[rune][]string)

func LoadAllBanner() error {
	file := map[string]string{
		"standard":   "banners/standard.txt",
		"shadow":     "banners/shadow.txt",
		"thinkertoy": "banners/thinkertoy.txt",
	}
	for name, r := range file {
		m, err := LoadBanner(r)
		if err != nil {
			return err
		}
		banner[name] = m
	}
	return nil
}
func LoadBanner(filname string) (map[rune][]string, error) {
	data, err := os.ReadFile(filname)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("Ops! This Is An Empty File")
	}
	text := strings.ReplaceAll(string(data), "\r\n", "\n")
	Lines := strings.Split(text, "\n")
	if len(Lines) != 855 {
		return nil, errors.New("Ops! This Banner Files Is Incomplete Banner")
	}
	var Result = make(map[rune][]string)
	for i := 0; i < 95; i++ {
		start := i*9 + 1
		r := rune(i + 32)
		Result[r] = Lines[start : start+8]
	}
	return Result, nil
}
