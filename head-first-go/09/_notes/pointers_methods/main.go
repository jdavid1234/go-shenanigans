// B"H

package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

// -- ------------------------------------------
// Notice the Receiver Parameters
func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

// -- ------------------------------------------
func (l Liters) DoubleWrong() {
	l *= 2
}

// -- ------------------------------------------
// Notice how receiver paramter is a pointer:
func (l *Liters) Double() {
	// Update value at pointer:
	*l *= 2
}

// -- ------------------------------------------
func main() {

	// -- --------------------------------------
	var soda Liters
	var water Milliliters

	soda = 2
	water = 500

	// -- --------------------------------------
	fmt.Println("Soda:", soda)
	fmt.Println("Water:", water)

	fmt.Println("Soda + float64:", soda+2.5)

	// -- --------------------------------------
	// Notice the Method Receiver
	var sodaGallons Gallons
	sodaGallons = soda.ToGallons()

	fmt.Println("sodaGallons", sodaGallons)

	// -- --------------------------------------
	var waterGallons Gallons
	waterGallons = water.ToGallons()

	fmt.Println("waterGallons", waterGallons)

	// -- --------------------------------------
	soda.DoubleWrong()
	fmt.Println("Am I double?", soda)

	// -- --------------------------------------
	// Go will automatically convert the
	// method receiver to a pointer for you:
	soda.Double()
	fmt.Println("Am I double?", soda)

}
