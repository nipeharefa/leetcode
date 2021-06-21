package main

import "fmt"

func getHint(secret string, guess string) string {

	var a string

	buls := make(map[int]byte)
	cows := make(map[int]byte)

	secretByte := []byte(secret)
	guessByte := []byte(guess)

	for i, sb := range secretByte {

		if guessByte[i] == sb {
			buls[i] = sb
			continue
		}

		cows[i] = sb
	}

	cowsCounter := 0
	for i, sb := range secretByte {
		if buls[i] == sb {
			continue
		}
		for kv, v := range cows {
			if v == sb {
				cowsCounter++
				delete(cows, kv)
			}
		}
	}

	fmt.Println(buls, cows, cowsCounter)
	// a = fmt.Sprintf("%dA%dB", buls, cows)
	return a
}

func main() {

	fmt.Println(getHint("1123", "0111"))
}
