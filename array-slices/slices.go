package main

import (
	"fmt"
)

var gGroceries[]string

func addGrocery(item string) {
	fmt.Println("Capacity: ", cap(gGroceries))
	gGroceries=append(gGroceries, item)
}

func listGroceries() {
	fmt.Println("Grocery list is as follows:")
	for i := 0; i < len(gGroceries); i++ {
		fmt.Println(gGroceries[i])
	}
}

func main() {
	listGroceries()
	addGrocery("Apple")
	listGroceries()
	addGrocery("Banana")
	addGrocery("Banana")
	addGrocery("Banana")
	addGrocery("Banana")
	addGrocery("Banana")
	addGrocery("Banana")
	addGrocery("Banana")
	listGroceries()
}
