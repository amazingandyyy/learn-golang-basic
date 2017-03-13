package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 10; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second * 1 / 10)
	}
	fmt.Println("Happy Birthday")

	a := 10
	for a > 0 {
		fmt.Println(a)
		time.Sleep(time.Second * 1 / 10)
		a--
	}
	fmt.Println("Happy Birthday")
}
