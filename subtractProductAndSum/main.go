package main

import "fmt"

func subtractProductAndSum(n int) int {

	// a :=n % 10

	a := []int{}
	ax := 1
	aPlus := 0

	for n != 0 {
		lastDigit := n % 10
		a = append(a, lastDigit)
		n /= 10
	}

	for _, v := range a {
		ax *= v
		aPlus += v
	}

	return ax - aPlus
}

func main() {

	fmt.Println(subtractProductAndSum(4421))
}
