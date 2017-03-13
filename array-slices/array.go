// package main

// import (
// 	"fmt"
// )

// const gCapability = 5

// var gGroceries [gCapability]string
// var gLenth = 0

// func addGrocery(item string) {
// 	if gLenth < gCapability {
// 		gGroceries[gLenth] = item
// 		gLenth++
// 	} else {
// 		fmt.Println("It can only be at most", gCapability, "items")
// 	}
// }

// func listGroceries() {
// 	fmt.Println("Grocery list is as follows:")
// 	for i := 0; i < gLenth; i++ {
// 		fmt.Println(gGroceries[i])
// 	}
// }

// func main() {
// 	listGroceries()
// 	addGrocery("Apple")
// 	listGroceries()
// 	addGrocery("Banana")
// 	listGroceries()
// }
