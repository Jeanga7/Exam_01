package main

import (
	"fmt"
	"os"
	"strconv"
)

func toRoman(num int) (string, error) {
	if num <= 0 || num >= 4000 {
		return "", fmt.Errorf("ERROR: cannot convert to roman digit")
	}

	roman := ""
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			roman += romans[i]
			num -= values[i]
		}
	}

	return roman, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <number>")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	roman, err := toRoman(num)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(roman)
}
