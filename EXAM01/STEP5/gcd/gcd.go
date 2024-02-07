package main

import (
	"fmt"
	"os"
	"strconv"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	num1, err1 := strconv.Atoi(os.Args[1])
	num2, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil || num1 <= 0 || num2 <= 0 {
		return
	}

	result := gcd(num1, num2)
	fmt.Println(result)
}
