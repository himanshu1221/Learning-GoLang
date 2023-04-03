package main

import "fmt"

func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("learning about Defer")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
