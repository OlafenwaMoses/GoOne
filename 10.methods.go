package main

import "fmt"

func main() {
	var1 := object{width: 40, length: 120, height: 200}

	fmt.Println(var1.computeVolume())
	fmt.Println(var1.computeSurfaceArea())

}

type object struct {
	width, length, height int
}

func (item object) computeVolume() int {
	volume := item.width * item.length * item.height
	return volume
}

func (item object) computeSurfaceArea() int {
	surface_area := 2 * ((item.width * item.length) + (item.width * item.height) + (item.height * item.length))
	return surface_area
}
