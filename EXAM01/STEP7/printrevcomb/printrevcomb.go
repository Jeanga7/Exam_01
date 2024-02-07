package main

import "github.com/01-edu/z01"

func printRevComb() {
	for i := 9; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			for k := j - 1; k >= 0; k-- {
				z01.PrintRune(rune('0' + i))
				z01.PrintRune(rune('0' + j))
				z01.PrintRune(rune('0' + k))

				if i != 0 || j != 1 || k != 0 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	printRevComb()
}
