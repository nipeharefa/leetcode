package main

import "fmt"

func plusOne(digits []int) []int {

	// digits[len(digits)-1] += 1

	val := digits[len(digits)-1]

	ten := val + 1

	if ten < 9 {
		digits[len(digits)-1] += 1
		return digits
	}

	digits[len(digits)-1] += 1

	for i := len(digits) - 1; i > 0; i-- {

		if digits[i]%10 == 0 && digits[i] != 0 {
			digits[i] = 0
			digits[i-1] += 1
		}
	}

	if digits[0]%10 == 0 {
		temp := []int{1}
		digits[0] = 0

		temp = append(temp, digits...)

		return temp
	}

	return digits
}

func main() {

	nums := []int{8, 9, 4, 5, 0, 0, 7, 9}

	fmt.Println(plusOne(nums))
}
