package main

import (
	"fmt"
	"math/rand"

	"github.com/gokatas/direction"
)

func main() {
	cardinal := direction.Cardinal(rand.Intn(3))
	fmt.Print(cardinal)
	switch cardinal {
	case direction.North:
		fmt.Println(" goes up.")
	case direction.South:
		fmt.Println(" goes down.")
	default:
		fmt.Println(" stays put.")
	}
}
