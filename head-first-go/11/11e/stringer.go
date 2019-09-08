// B"H

package main

import "fmt"

// -- -------------------------------------
type Gallons float64

// Satisfies the fmt.Stringer interface:
func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

// -- -------------------------------------
type Liters float64

// Satisfies the fmt.Stringer interface:
func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

// -- -------------------------------------
type Milliliters float64

// Satisfies the fmt.Stringer interface:
func (m Milliliters) String() string {
	return fmt.Sprintf("%0.2f mL", m)
}

// -- -------------------------------------
func main() {

	fmt.Println(Gallons(12.09248342))

	fmt.Println(Liters(12.09248342))

	fmt.Println(Milliliters(12.09248342))
}
