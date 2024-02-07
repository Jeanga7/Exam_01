package main

import (
	"fmt"
	"os"
	"strconv"
)

func Range(num1, num2 string) {
	n1, err := strconv.Atoi(num1)
	if err != nil {
		fmt.Println(err)
	}
	n2, err := strconv.Atoi(num2)
	if err != nil {
		fmt.Println(err)
		return
	}
	if n1 <= n2 {
		for i := n1; i <= n2; i++ {
			fmt.Print(i)
			fmt.Print(" ")
		}
	}
fmt.Println()
}

func main() {
	if len(os.Args) == 3 {
		Range(os.Args[1], os.Args[2])
	} else {
		return
	}
}
