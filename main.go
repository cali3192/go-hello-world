package main // package named main for executable files

import (
	"fmt"
	"math/rand"
)

var cuteLevel = 10

// basic pring
func helloWorld(john string) {

	fmt.Println("\nsecond function: helloWorld: ", john, cuteLevel)
}

// referencing another package
func printRandom() {
	var num = rand.Int()

	fmt.Println("\nthird function: printRan", num)
}

// taking the pointer to integer
func swap(x, y *int) int {

	// taking a temporary variable
	var tmp int

	// * operator modifies values
	tmp = *x
	*x = *y
	*y = tmp

	return tmp
}

func main() {
	fmt.Println("Hi, John!")

	// calling the hello world function
	helloWorld("Babe")

	// calling math ran function
	printRandom()

	var first = 700
	var second = 900

	fmt.Printf("\nfourth function: swap \nbefore function call\n")
	fmt.Printf("f = %d and s = %d\n", first, second)

	// & operator to get address of data
	// by passing the address
	// of the variables
	swap(&first, &second)
	// swap(first, second)

	fmt.Printf("-> after function call\n")
	fmt.Printf("f = %d and s = %d\n", first, second)
}
