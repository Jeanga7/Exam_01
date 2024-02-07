package main

import (
	"os"

	"github.com/01-edu/z01"
)

func switchCase(s string) string {
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			result += string(char - 32) 
		} else if char >= 'A' && char <= 'Z' {
			result += string(char + 32) 
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		return
	}
	result := switchCase(os.Args[1])

	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
