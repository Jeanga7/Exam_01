package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isSpace(char rune) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
}

func isAlphanumeric(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func reverseStrCap(s string) string {
	var result string

	for i, char := range s {
		if i < len(s)-1 && !isSpace(rune(s[i+1])) {
			result += string(ToLower(char))
		} else {
			result += string(ToUpper(char))
		}
	}

	return result
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	for _, arg := range args {
		reversed := reverseStrCap(arg)
		for _, char := range reversed {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

func ToUpper(char rune) rune {
	if char >= 'a' && char <= 'z' {
		return char - ('a' - 'A')
	}
	return char
}

func ToLower(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char + ('a' - 'A')
	}
	return char
}
