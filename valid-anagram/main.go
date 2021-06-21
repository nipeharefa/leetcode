package main

import "fmt"

func isAnagram(s string, t string) bool {

	sourceMap := make(map[byte]int)
	tMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		_, ok := sourceMap[s[i]]
		if ok {
			sourceMap[s[i]]++
		} else {
			sourceMap[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		_, ok := sourceMap[t[i]]
		if ok {
			tMap[t[i]]++
		} else {
			tMap[t[i]] = 1
		}
	}

	if len(sourceMap) != len(tMap) {
		return false
	}

	for k, v := range sourceMap {
		if tMap[k] != v {
			return false
		}
		delete(tMap, k)
	}

	if len(tMap) > 0 {
		return false
	}

	return true
}

func main() {

	fmt.Println(isAnagram("rat", "car"))
}
