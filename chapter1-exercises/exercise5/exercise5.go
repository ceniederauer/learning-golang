package main

import "fmt"

type hotdog int

var x hotdog
var y int
var z int

func main() {

	fmt.Printf("Value of x: %v\n", x)
	fmt.Printf("Type of x: %T\n", x)

	x = 42
	fmt.Printf("New value of x: %v\n", x)

	y = int(x)
	fmt.Printf("Value of y: %v\n", y)
	fmt.Printf("Type of y: %T\n", y)
}
