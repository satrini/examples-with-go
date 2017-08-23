package main

import "fmt"

func main() {
	var value int
	fmt.Print("Enter value: ")
	fmt.Scanf("%d", &value)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", i, value, i*value)
	}
}
