// Iota shows idiomatic way to implement an enumerated type:
//  1. Create a new integer type.
//  2. List its values using iota.
//  3. Give the type a String function.
//
// Based on https://yourbasic.org/golang/iota.
package main

import (
	"fmt"
	"math/rand"
)

type Direction int

// Cardinal directions.
const (
	North Direction = iota
	South
	East
	West
)

func (d Direction) String() string {
	return [...]string{"North", "South", "East", "West"}[d]
}

func main() {
	d := Direction(rand.Intn(3))
	fmt.Print(d)
	switch d {
	case North:
		fmt.Println(" goes up.")
	case South:
		fmt.Println(" goes down.")
	default:
		fmt.Println(" stays put.")
	}
}
