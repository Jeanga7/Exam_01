package main

import (
	"fmt"
	"os"
)

/* func rostring(s string) string {
	words := strings.Fields(s)
	if len(words) <= 1 {
		return s
	}
	rotated := append(words[1:], words[0])
	return strings.Join(rotated, " ")
} */


func isSpace(r byte) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

func trimSpaces(s string) string {
	start, end := 0, len(s)-1
	for start <= end && isSpace(s[start]) {
		start++
	}
	for end >= start && isSpace(s[end]) {
		end--
	}
	return s[start : end+1]
}

func reverse(str []byte) {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
}

func rostring(s string) string {
	s = trimSpaces(s)
	if len(s) == 0 {
		return ""
	}

	str := []byte(s)
	reverse(str)

	wordStart := 0
	for i := 0; i < len(str); i++ {
		if isSpace(str[i]) {
			reverse(str[wordStart:i])
			wordStart = i + 1
		}
	}

	reverse(str[wordStart:])

	return string(str)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	result := rostring(os.Args[1])
	fmt.Println(result)
}
