package main

import (
	"os"

	"github.com/01-edu/z01"
)

func alphamirror(str string) {
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			z01.PrintRune('a' + 'z' - char)
		} else if char >= 'A' && char <= 'Z' {
			z01.PrintRune('A' + 'Z' - char)
		} else {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) == 2 {
		alphamirror(os.Args[1])
	} else {
		z01.PrintRune('\n')
	}
}
