package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("learning about Switch and case ")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is : ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("dice value is one you can open")
	case 2:
		fmt.Println("you can move 2 spot")
	case 3:
		fmt.Println("you can move 3 spot")
		fallthrough
	case 4:
		fmt.Println("you can move 4 spot")
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("you can move 6 spot and can roll the dice again")
	default:
		fmt.Println("what was that")

	}
}
