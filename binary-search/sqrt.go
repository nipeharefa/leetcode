package main

import "fmt"

func sqrt(input int) int {

	a := []int{}

	for input != 0 {
		h := input / 2

		if h == 2 {
			a = append(a, 2)
		} else {
			a = append(a, h)
		}
		input = h
	}
	fmt.Println(a)

	return 0
}

func main() {
	fmt.Println(sqrt(8))
}
