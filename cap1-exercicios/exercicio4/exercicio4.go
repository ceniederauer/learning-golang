package main

import "fmt"

type hotdog int

var x hotdog

func main() {

	fmt.Printf("Value of x: %v\n", x)
	fmt.Printf("Type of x: %T\n", x)

	x = 42
	fmt.Printf("New value of x: %v\n", x)
}
