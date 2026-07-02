package main

var banner = make(map[string]map[rune][]string)

func LoadAllBannners() error {
	files := map[string]string{
		"standard":   "standard.txt",
		"shadow":     "shadow.txt",
		"thinkertoy": "thinkertoy.txt",
	}
	for name, path := range files {
		m, err := LoadBanner(path)
		if err != nil {
			return err
		}
		banner[name] = m
	}
	return nil
}
