package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Println("1. address of t is", &t)
	// & - Reference
	// * - De-Rederence
	add(3, 5, &t)
	fmt.Println(t)
}

func add(i int, j int, t *int) int {
	fmt.Println("2. address of t is", t)
	fmt.Println("3. valut of t of this address is", *t)
	*t = i + j
	return *t
}
