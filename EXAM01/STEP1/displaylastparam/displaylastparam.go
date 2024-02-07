package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) >= 2 {
		for _, v := range os.Args[len(os.Args)-1] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
