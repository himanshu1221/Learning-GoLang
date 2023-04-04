package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Learning about Golang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	databytes, error := ioutil.ReadAll(response.Body)

	if error != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}
