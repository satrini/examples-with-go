package main

import (
	"fmt"
	"math"
)

func main() {
	var input int
	fmt.Print("Enter value: ")
	fmt.Scanf("%d", &input)
	fmt.Println("Double: ", input*2)
	fmt.Println("Triple: ", input*3)
	fmt.Println("Sqrt: ", math.Sqrt(float64(input)))
}
