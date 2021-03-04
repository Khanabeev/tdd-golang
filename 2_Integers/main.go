package main

import (
	"fmt"
)

func main() {
	fmt.Println(Add(2,2))
}
// Add takes two integers and returns the sum of them.
func Add(x int, y int) int {
	return x + y
}