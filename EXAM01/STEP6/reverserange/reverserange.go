package main

import (
	"fmt"
	"os"
	"strconv"
)

func reverseRange(start, end int) {
	if start > end {
		for i := start; i >= end; i-- {
			fmt.Print(i, " ")
		}
		fmt.Println()
	} else {
		for i := start; i <= end; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		return
	}

	start, err1 := strconv.Atoi(args[0])
	end, err2 := strconv.Atoi(args[1])

	if err1 != nil || err2 != nil {
		fmt.Println(err1)
		return
	}

	reverseRange(start, end)
}
/* package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func reverseRange(start, end int) {
	if start > end {
		for i := start; i >= end; i-- {
			printNumber(i)
		}
		z01.PrintRune('\n')
	} else {
		for i := start; i <= end; i++ {
			printNumber(i)
		}
		z01.PrintRune('\n')
	}
}

func printNumber(num int) {
	if num < 0 {
		z01.PrintRune('-')
		num = -num
	}

	if num == 0 {
		z01.PrintRune('0')
		return
	}

	var digits []int
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(rune(digits[i] + '0'))
	}
	z01.PrintRune(' ')
}

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		return
	}

	start, err1 := strconv.Atoi(args[0])
	end, err2 := strconv.Atoi(args[1])

	if err1 != nil || err2 != nil {
		fmt.Println(err1)
		return
	}

	reverseRange(start, end)
}
 */