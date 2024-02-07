package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune('0')
	}
	count := 0
	for range os.Args[1:] {
		count++
	}
	Println(PrintNbr(count))
	// fmt.Println(count)
}

func Println(str string) {
	for _, v := range str {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func PrintNbr(num int) string {
	res := ""
	dizaine := 0
	for num > 0 {
		dizaine = num % 10
		res = string('0'+dizaine) + res
		num /= 10
	}
	return res
}
