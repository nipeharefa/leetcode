package main

import "fmt"

const MaxInt = 2147483647
const MinInt = -2147483647

func reverse(n int) int {

	if n > MaxInt || n < MinInt {
		return 0
	}
	result := 0

	for n != 0 {

		lastDigit := n % 10
		n /= 10
		result = (result * 10) + lastDigit

	}

	if result > MaxInt || result < MinInt {
		return 0
	}

	return result
}

func main() {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)

	fmt.Println(MaxInt)

	// 2147483648
	// 9646324351
	fmt.Println(reverse(1534236469))
}
