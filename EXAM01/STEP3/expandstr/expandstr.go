package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	s := os.Args[1]
	start, end := 0, len(s)-1

	for start < len(s) && (s[start] == ' ' || s[start] == '\t') {
		start++
	}

	for end >= 0 && (s[end] == ' ' || s[end] == '\t') {
		end--
	}

	for i := start; i <= end; i++ {
		if s[i] == ' ' || s[i] == '\t' {
			word := s[start:i]
			for _, char := range word {
				z01.PrintRune(char)
			}
			z01.PrintRune(' ')
			z01.PrintRune(' ')
			z01.PrintRune(' ')

			for i < end && (s[i+1] == ' ' || s[i+1] == '\t') {
				i++
			}
			start = i + 1
		}
	}

	lastWord := s[start : end+1]
	for _, char := range lastWord {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
