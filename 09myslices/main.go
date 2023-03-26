package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to  slices tutorial")

	var fruitList = []string{"Apple", "Tomato", "Banana"}
	fmt.Println("Old fruit List", fruitList)

	fmt.Printf("type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "peach")

	fmt.Println("Updatded fruit list", fruitList)

	fruitList = append(fruitList[2:4])

	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 123
	highScore[1] = 567
	highScore[2] = 744
	highScore[3] = 987

	highScore = append(highScore, 222, 555, 666)

	fmt.Println(highScore)

	sort.Ints(highScore)

	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))

	// how to remove a value from slice based on index

	var courses = []string{"ReactJS", "JavaScript", "Python", "Gatsby"}

	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
}
