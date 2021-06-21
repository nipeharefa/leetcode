package main

import "fmt"

func numberOfSteps(num int) int {

	step := 0

	for num != 0 {

		isEven := num%2 == 0

		if isEven {

			num /= 2
		} else {
			num--
		}

		step++
	}

	return step
}

func main() {

	fmt.Println(numberOfSteps(123))

}
