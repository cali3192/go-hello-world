package main // executable type

import (
	"fmt"
)

// using interface allows multiple types of inputs
func print(message interface{}) {
	switch message.(type) {
	case string:
		fmt.Println("\ntest message", message)
		return
	case int:
		fmt.Println("\nI'm an int", message)
		return
	default:
		fmt.Println("\nDefault")
		return

	}
}

func main() {
	// code goes here
	var card string = "Ace of Spades"

	var piVar int = 3

	print(card)

	print(piVar)

}
