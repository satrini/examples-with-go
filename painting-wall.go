package main

import "fmt"

// Make a program that reads the width and height of a wall
// in meters, calculates its area and the amount of paint
// needed to paint it, knowing that each liter of paint paints
// an area of ​​2 square meters.

func main() {
	var with, height float64
	fmt.Print("Wall width: ")
	fmt.Scanf("%f\n", &with)
	fmt.Print("Wall height: ")
	fmt.Scanf("%f\n", &height)

	area := with * height
	paint := area / 2

	fmt.Printf("Wall dimension: %.1f x %.1f\nArea: %.1fm²\n", with, height, area)
	fmt.Printf("Needs: %.1f liters of paint wall\n", paint)
}
