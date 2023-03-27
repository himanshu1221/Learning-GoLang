package main

import "fmt"

func main() {
	fmt.Println("Practice on structs")

	Himanshu := User{"Himanshu Chhatwal", "himanshuchhatwal9295@gmail.com", true, 19}
	fmt.Println(Himanshu)
	fmt.Printf("Himanshu details are : %v\n", Himanshu)
	fmt.Printf("Name is %v and email is %v \n", Himanshu.Name, Himanshu.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
