package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		res := wdMatch(os.Args[1], os.Args[2])
		if res == os.Args[1]{
			printStr(res)
		}
	}
}

func wdMatch(s1, s2 string) string {
	var result string
	i, j := 0, 0

	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			result += string(s1[i])
			i++
		}
		j++
	}

	return result
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
