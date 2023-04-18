package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Making request in APIs")
	// PerformGetRequest()
	// PerformPostJsonReuest()
	PerformPostFormReuest()
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

func PerformPostJsonReuest() {
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

func PerformPostFormReuest() {
	const Url = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "himanshu")
	data.Add("lastname", "chhatwal")
	data.Add("email", "himanshu@go.dev")

	response, err := http.PostForm(Url, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
