package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primeFactors(n int) []int {
	factors := make([]int, 0)
	for i := 2; i <= n; i++ {
		for n%i == 0 && isPrime(i) {
			factors = append(factors, i)
			n /= i
		}
	}
	return factors
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 1 {
		return
	}

	factors := primeFactors(num)
	if len(factors) == 0 {
		return
	}

	for i, factor := range factors {
		fmt.Print(factor)
		if i < len(factors)-1 {
			fmt.Print("*")
		}
	}
	fmt.Println()
}
