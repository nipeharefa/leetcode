package main

import "fmt"

func frequencySort(nums []int) []int {

	// counter
	m := make(map[int]int)

	keysM := []int{}

	for _, v := range nums {
		_, ok := m[v]
		if ok {
			m[v]++
		} else {
			m[v] = 1
			// keysM = append(keysM, v)
		}
	}

	for _, v := range m {
		keysM = append(keysM, v)
	}

	for i := 0; i < len(keysM); i++ {
		for j := 0; j < len(keysM); j++ {
			if keysM[i] < keysM[j] {
				keysM[i], keysM[j] = keysM[j], keysM[i]
			}
		}
	}

	// in := 1

	hasil := []int{}
	for _, qty := range keysM {

		a := []int{}
		for k, v := range m {
			if v == qty {
				a = append(a, k)
			}
		}

		for i := 0; i < len(a); i++ {
			for j := 0; j < len(a); j++ {
				if a[i] > a[j] {
					a[i], a[j] = a[j], a[i]
				}
			}
		}

		for _, aa := range a {

			for i := 0; i < qty; i++ {
				hasil = append(hasil, aa)
			}

			delete(m, aa)
		}
	}

	// for len(m) > in {
	// 	keysM := []int{}
	// 	for _, v := range m {
	// 		if v == in {

	// 		}
	// 		keysM = append(keysM, )
	// 	}
	// 	in++
	// }

	return hasil
}

func main() {
	nums := []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}

	fmt.Println(frequencySort(nums))
}
