package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Making request in APIs")
	// PerformGetRequest()
	PerformPostReuest()
}

// func PerformGetRequest() {

// 	const myUrl = "http://localhost:8000/get"

// 	response, err := http.Get(myUrl)

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer response.Body.Close()

// 	fmt.Println("Status Code:", response.StatusCode)

// 	var responseString strings.Builder

// 	content, _ := ioutil.ReadAll(response.Body)

// 	byteCount, _ := responseString.Write(content)

// 	fmt.Println("Bytecount is", byteCount)

// 	fmt.Println(responseString.String())
// 	// fmt.Println("Content Length:", response.ContentLength)

// 	// fmt.Println(string(content))
// }

func PerformPostReuest() {
	const Url = "http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"Course Name":"Let's go with golang",
			"Price":0,
			"Platform":"himanshu.dev"
		}
	`)

	response, err := http.Post(Url, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	Content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(Content))

}
