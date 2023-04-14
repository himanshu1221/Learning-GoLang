package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Making request in APIs")
	PerformGetRequest()
}

func PerformGetRequest() {

	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)

	fmt.Println("Bytecount is", byteCount)

	fmt.Println(responseString.String())
	// fmt.Println("Content Length:", response.ContentLength)

	// fmt.Println(string(content))
}
