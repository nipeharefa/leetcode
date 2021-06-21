package main

import "fmt"

func containsDuplicate(nums []int) bool {

	e := make(map[int]int)

	for _, v := range nums {
		_, ok := e[v]

		if ok {
			e[v] += 1
			if e[v] >= 2 {
				return true
			}
		} else {
			e[v] = 1
		}
	}

	return false
}

func main() {

	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	fmt.Println(containsDuplicate(nums))
}
