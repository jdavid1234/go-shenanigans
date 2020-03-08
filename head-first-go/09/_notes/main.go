// B"H

package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64
type Kiloliters float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (k Kiloliters) ToGallons() Gallons {
	return Gallons(k * 264)
}

func main() {

	// -- --------------------------------------
	var soda Liters
	var water Milliliters
	var oil Kiloliters

	soda = Liters(1.0)
	water = Milliliters(1.0)
	oil = Kiloliters(1.0)

	// -- --------------------------------------
	fmt.Println("Soda:", soda)
	fmt.Println("Water:", water)

	// -- --------------------------------------
	var sodaGallons Gallons
	sodaGallons = soda.ToGallons()

	fmt.Println("sodaGallons", sodaGallons)

	// -- --------------------------------------
	var waterGallons Gallons
	waterGallons = water.ToGallons()

	fmt.Println("waterGallons", waterGallons)

	// -- --------------------------------------
	var oilGallons Gallons
	oilGallons = oil.ToGallons()

	fmt.Println("oilGallons", oilGallons)

	fmt.Println(oilGallons*2 - waterGallons*2)

}
