package main

import "fmt"

func restoreString(s string, indices []int) string {

	x := make(map[int]byte)
	temp := []byte(s)
	for i := 0; i < len(s); i++ {
		x[indices[i]] = temp[i]
	}

	newB := make([]byte, 0)
	for i := 0; i < len(s); i++ {

		newB = append(newB, x[i])
	}

	return string(newB)
}

func main() {
	s := "codeleet"

	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}

	str := restoreString(s, indices)
	fmt.Println(str)
}
