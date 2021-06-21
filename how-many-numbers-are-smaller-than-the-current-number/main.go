package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {

	a := []int{}

	for i := 0; i < len(nums); i++ {
		total := 0
		my := nums[i]
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}

			v := nums[j]
			if v < my {
				total += 1
			}
		}

		a = append(a, total)
	}

	return a
}

func main() {

	a := []int{7, 7, 7, 7}

	fmt.Println(smallerNumbersThanCurrent(a))
}
