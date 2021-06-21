package main

import "fmt"

func helper(i int, s []byte) {
	// s[0] = 22
	l := len(s) / 2

	if i >= l {
		return
	}

	last := len(s) - i - 1

	s[i], s[last] = s[last], s[i]
	helper(i+1, s)

}
func reverseString(s []byte) {

	if len(s) == 2 {
		s[0], s[1] = s[1], s[0]
		return
	}
	helper(0, s)
}

func main() {

	s := "hello"

	sb := []byte(s)

	reverseString(sb)

	fmt.Println(string(sb))
}
