package main

import (
	"os"

	"github.com/01-edu/z01"
)

/* func isSpace(char rune) bool {
	return char == ' ' || char == '\t' || char == '\n' || char == '\r'
} */

func firstWord(s string) string {
	result := ""
	for _, v := range s {
		if v != ' '{
			result+=string(v)
		}else{
			break
		}
	}

	return result
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	result := firstWord(os.Args[1])

	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
