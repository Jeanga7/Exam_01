package main

import (
	"os"

	"github.com/01-edu/z01"
)

func union(s1, s2 string) {
	var order []rune
	mapChar := make(map[rune]bool)
	for _, char := range s1 {
		if exist := mapChar[char]; !exist {
			order = append(order, char)
			mapChar[char] = true
		}
	}
	for _, char := range s2 {
		if !mapChar[char] {
			order = append(order, char)
			mapChar[char] = true
		}
	}

	for _, char := range order {
		z01.PrintRune(char)
		
	}
	z01.PrintRune('\n')
}

func main() {
	if len(os.Args) == 3 {
		union(os.Args[1], os.Args[2])
	} else {
		z01.PrintRune('\n')
	}
}
