package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers section")

	// var ptr *int

	// fmt.Println("value of pointer is :", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("vaue of actual ptr is ", ptr)
	fmt.Println("vaue of actual ptr is ", *ptr)

	*ptr = *ptr + 2

	fmt.Println("new value is :", myNumber)

}
