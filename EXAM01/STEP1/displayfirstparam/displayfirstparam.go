package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) >= 2 {
		Println(os.Args[1])
	} else {
		return
	}

}

func Println(str string) {
	for _, v := range str {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
