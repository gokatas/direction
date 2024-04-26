// Package direction shows idiomatic way to implement an enumerated type:
//
//  1. Create a new integer type.
//  2. List its values using iota.
//  3. Give the type a String function.
//
// Based on https://yourbasic.org/golang/iota.
package direction

type Cardinal int

// Cardinal directions.
const (
	North Cardinal = iota
	South
	East
	West
)

func (d Cardinal) String() string {
	return [...]string{"North", "South", "East", "West"}[d]
}
