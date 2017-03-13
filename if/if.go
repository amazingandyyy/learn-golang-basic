package main

import "fmt"

func main() {
	age := 1

	if age < 10 && age > 0 {
		fmt.Println("You are younger than 10")
	} else if age > 10 {
		fmt.Println("You are older than 10")
	} else {
		fmt.Println("?")
	}
}
