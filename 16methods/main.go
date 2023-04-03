package main

import "fmt"

func main() {
	fmt.Println("Practice on Methods")

	Himanshu := User{"Himanshu Chhatwal", "himanshuchhatwal9295@gmail.com", true, 19}
	fmt.Println(Himanshu)
	fmt.Printf("Himanshu details are : %v\n", Himanshu)
	fmt.Printf("Name is %v and email is %v \n", Himanshu.Name, Himanshu.Email)
	Himanshu.GetStatus()
	Himanshu.newEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user is active:", u.Status)
}

func (u User) newEmail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this person is :", u.Email)
}
