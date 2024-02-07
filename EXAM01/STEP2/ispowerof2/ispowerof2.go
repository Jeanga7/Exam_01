package main

import (
	"os"

	"github.com/01-edu/z01"
)

const (
	true  = "true"
	false = "false"
)

func main() {
	// num, err := strconv.Atoi(os.Args[1])
	num := Atoi(os.Args[1])
	if len(os.Args) != 2 /* || err != nil */ {
		return
	}
	IsPowerOf2(num)
}

func IsPowerOf2(num int) {
	if num <= 0 {
		return
	}
	for num > 1 {
		if num%2 != 0 {
			println(false)
			return
		}
		num /= 2
	}
	println(true)
}

func Println(str string) {
	for _, v := range str {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func Atoi(str string) int {
	res := 0
	for _, v := range str {
		if v >= '0' && v <= '9' {
			res = res*10 + int(v-'0')
		} else {
			return 0
		}
	}
	return res
}
