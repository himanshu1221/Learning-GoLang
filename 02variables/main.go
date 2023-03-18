package main

import "fmt"

// public
const LoginToken string = "sknvk"

func main() {

	// Data Types

	var username string = "himanshu"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type : %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type : %T \n", smallVal)

	var smallFloat float64 = 255.22342342344342323523523
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type : %T \n", smallFloat)

	// default value and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type : %T \n", anotherVariable)

	// impicit type

	var website = "Himanshuchhatwal.netlify.app"
	fmt.Println(website)

	// no var style

	numberOfUser := 3000
	fmt.Println(numberOfUser)

	// calling public constant

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type : %T \n", LoginToken)

}
