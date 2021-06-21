package main

import "fmt"

func singleNumber(nums []int) int {

	e := make(map[int]int)

	total := 0
	for i := 0; i < len(nums); i++ {

		value := nums[i]

		_, ok := e[value]
		if !ok {
			e[value] = 1
		} else {
			e[value] += 1
		}
	}

	// fmt.Println(e)

	for k, v := range e {

		if v == 1 {
			return k
		}
	}

	return total

}

func main() {

	nums := []int{2, 2, 1}
	fmt.Println(singleNumber(nums))
}
