package main

import "fmt"

func rotate(x *[]int, k int) {

	for i := 0; i < k; i++ {
		nums := *x
		end := nums[len(nums)-1]
		temp := nums[0 : len(nums)-1]

		nums = []int{end}
		nums = append(nums, temp...)
		*x = nums
	}
}

func main() {

	nums := []int{-1, -100, 3, 99}

	rotate(&nums, 2)

	fmt.Println(nums)
}
