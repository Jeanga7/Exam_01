package main

import (
	"os"

	"github.com/01-edu/z01"
)

func hiddenp(s1, s2 string) bool {
	i, j := 0, 0

	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}

	return i == len(s1)
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1, s2 := os.Args[1], os.Args[2]

	if hiddenp(s1, s2) {
		z01.PrintRune('1')
	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')
}
