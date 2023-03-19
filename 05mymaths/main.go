package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome  to maths in GOLang")

	// random number generator using math/rand

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
