package main

import (
	"os"

	"github.com/01-edu/z01"
)

func searchReplace(str, search, replace string) {
	for _, char := range str {
		if string(char) == search {
			z01.PrintRune(rune(replace[0]))
		} else {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	str := os.Args[1]
	search := os.Args[2]
	replace := os.Args[3]

	searchReplace(str, search, replace)
}
