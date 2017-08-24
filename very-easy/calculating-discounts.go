package main

import "fmt"

// Make an algorithm that reads the price of a product
// and show its new price, with 5% discount.

func main() {
	var price float64
	fmt.Print("Price of the product: $")
	fmt.Scanf("%f\n", &price)

	discount := price - (price * 5 / 100)
	fmt.Printf("The product with discount is: $%.2f", discount)
}
