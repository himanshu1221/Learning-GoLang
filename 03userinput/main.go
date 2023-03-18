package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welocme to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	// comma ok syntax || error ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type of rating is %T \n", input)

}
