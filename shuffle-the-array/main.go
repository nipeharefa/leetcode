package main

import "fmt"

func shuffle(nums []int, n int) []int {

	x := []int{}

	xx := nums[0:n]
	yy := nums[n:]
	for i := 0; i < n; i++ {
		x = append(x, xx[i], yy[i])
	}
	return x
}

func main() {

	nums := []int{1, 1, 2, 2}
	n := 2

	shuffled := shuffle(nums, n)

	fmt.Println(shuffled)
}
