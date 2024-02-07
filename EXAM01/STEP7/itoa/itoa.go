package main

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	var result string
	for n != 0 {
		digit := n % 10
		result = string('0'+digit) + result
		n /= 10
	}

	if n < 0 {
		result = "-" + result
	}

	return result
}

