package main

import (
	"os"

	"github.com/01-edu/z01"
)

func repeatAlpha(str string) {
	for _, char := range str {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			start := 'A'
			if char >= 'a' {
				start = 'a'
			}
			for i := start; i <= char; i++ {
				z01.PrintRune(char)
			}
		} else {
			z01.PrintRune(char)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	repeatAlpha(os.Args[1])
	z01.PrintRune('\n')
}
