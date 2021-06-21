package main

import (
	"fmt"
)

func runningSum(nums []int) []int {

	nn := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			nn = append(nn, nums[i])
			continue
		}

		v := nums[i] + nn[i-1]

		nn = append(nn, v)
	}

	return nn
}

func main() {

	nums := []int{1, 1, 1, 1, 1}

	fmt.Println(runningSum(nums))
}
