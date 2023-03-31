package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")
	greeter()

	result := adder(3, 5)

	fmt.Println("Result is ", result)

	proResult, myMessage := proAdder(5, 6, 7, 8, 63, 43423, 232, 3, 2, 23)

	fmt.Println("pro result is ", proResult)
	fmt.Println("pro message is ", myMessage)

}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "hi pro result function"
}

func greeter() {
	fmt.Println("Namastey from golang")
}
