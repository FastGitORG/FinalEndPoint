package main

import "strings"

func parseFromString(conf string) map[string]string {
	arr := strings.Split(conf, "\n")
	return parseFromStringArray(arr)
}

func parseFromStringArray(conf []string) map[string]string {
	m := make(map[string]string)
	for _, v := range conf {
		n := strings.Split(v, ":")
		if len(n) != 2 {
			continue
		}
		m[n[0]] = n[1]
	}
	return m
}

func parseFromFile(path string) (map[string]string, bool) {
	s, ok := readAllText(path)
	if !ok {
		return nil, false
	}
	return parseFromString(s), true
}
