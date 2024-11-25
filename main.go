package main

import (
	"fmt"

	"github.com/thechicken82/car"
)

func main() {

	s1 := car.Honk()
	s2 := car.Honks()
	fmt.Println(s1)
	fmt.Println(s2)

	// also like this
	fmt.Println(car.Honk())

	fmt.Println(car.The_Car_Base())
}
