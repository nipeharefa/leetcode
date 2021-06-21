package main

import "fmt"

func maximumWealth(accounts [][]int) int {

	terkaya := 0

	for i := 0; i < len(accounts); i++ {
		temp := 0
		for j := 0; j < len(accounts[i]); j++ {
			temp += accounts[i][j]
		}

		if temp > terkaya {
			terkaya = temp
		}
	}

	return terkaya
}

func main() {

	a := [][]int{{1, 5}, {7, 3}, {3, 5}}

	fmt.Println(maximumWealth(a))
}
