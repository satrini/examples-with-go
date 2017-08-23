package main

import "fmt"

func main() {
	var measures float64
	fmt.Print("Enter a distance in meters: ")
	fmt.Scanf("%f", &measures)
	fmt.Printf("%f km\n", measures/1000)
	fmt.Printf("%f hm\n", measures/100)
	fmt.Printf("%.0f cm\n", measures*100)
	fmt.Printf("%.0f mm\n", measures*1000)
}
