package main

import (
	"app-cinema/cinema"
	"fmt"
	"log"
)

func main() {
	user1 := cinema.SetUser("John Doe", 19)
	result, err := user1.BuyTicket(16.0)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Printf("change: R$%.2f\n", result)
	fmt.Println("Thank you, goodbye!")
}
