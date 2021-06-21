package main

import "fmt"

type ParkingSystem struct {
	a map[int]int
	b map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {

	p := ParkingSystem{}

	a := make(map[int]int)
	a[1] = big
	a[2] = medium
	a[3] = small

	p.a = a

	b := make(map[int]int)
	b[1] = 0
	b[2] = 0
	b[3] = 0

	p.b = b

	return p
}

func (this *ParkingSystem) AddCar(carType int) bool {

	n := this.a[carType]
	m := this.b[carType]

	can := (m + 1) <= n
	if can {
		this.b[carType] += 1
		return true
	}

	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */

func main() {
	obj := Constructor(1, 1, 0)

	fmt.Println(obj.AddCar(1))
	fmt.Println(obj.AddCar(2))
	fmt.Println(obj.AddCar(3))
}
