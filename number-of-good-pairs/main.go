package main

import "fmt"

func numIdenticalPairs(nums []int) int {

	n := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				n++
			}
		}
	}
	return n
}

func main() {

	nums := []int{1, 2, 3}

	nu := numIdenticalPairs(nums)
	fmt.Println(nu)
}
