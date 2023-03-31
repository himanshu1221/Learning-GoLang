package main

import "fmt"

func main() {
	fmt.Println("Learning about loops")
	days := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and the value is %v\n", index, day)
	// }

	rougueValue := 1

	for rougueValue < 11 {

		if rougueValue == 2 {
			goto lco

		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println("value is ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("welcome to my github profile")

}
