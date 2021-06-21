package main

import (
	"fmt"
	"strconv"
)

func myAtoi(s string) int {

	dest := []byte{}
	sb := []byte(s)

	for _, v := range sb {

		if len(dest) == 0 && v == 45 {
			dest = append(dest, v)
			continue
		}
		if v >= 48 && v <= 57 {
			dest = append(dest, v)
		} else {
			dest = []byte(nil)
			break
		}
	}
	ss, _ := strconv.Atoi(string(dest))

	return ss
}

func main() {

	fmt.Println(myAtoi(" -.     42"))
}
