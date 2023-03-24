package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in GoLang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "tomato"
	fruitList[3] = "banana"

	fmt.Println("fruitlist is :", fruitList)
	fmt.Println("Length of array is:", len(fruitList))

	var vegList = [3]string{"potato", "onion", "beans"}
	fmt.Println("veggies list:", vegList)
	fmt.Println("length of veg list:", len(vegList))

}
