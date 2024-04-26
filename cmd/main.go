package main

import (
	"fmt"
	"math/rand"

	"direction"
)

func main() {
	d := direction.Cardinal(rand.Intn(4))
	fmt.Print(d)
	switch d {
	case direction.North:
		fmt.Println(" goes up.")
	case direction.South:
		fmt.Println(" goes down.")
	default:
		fmt.Println(" stays put.")
	}
}
