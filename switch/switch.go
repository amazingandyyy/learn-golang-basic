package main

import (
	"fmt"
)

func main() {
	var name string
	name = "andy3"
	switch name {
	case "andy":
		fmt.Println("Hey no.1")
	case "andy2":
		fmt.Println("Hey no.2")
		fallthrough
	case "andy3":
		fmt.Println("Hey no.3")
		fallthrough
	default:
		fmt.Println("Hey", name)
	}
}
