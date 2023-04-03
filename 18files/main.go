package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("learning about files")

	content := "learning code Go from hitesh choudhary"

	files, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(files, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("the length is", length)

	defer files.Close()

	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("text data inside a file is ", databyte)
}
