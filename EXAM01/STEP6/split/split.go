package main

import (
	"fmt"
)

func Split(s, sep string) []string {
	var result []string

	if s == "" || sep == "" {
		return []string{s}
	}

	sepLen := len(sep)
	strLen := len(s)
	start := 0

	for i := 0; i <= strLen-sepLen; i++ {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			start = i + sepLen
			i += sepLen - 1
		}
	}

	result = append(result, s[start:])
	return result
}

func main() {
	s := "HAhelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
