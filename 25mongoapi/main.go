package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/himanshu1221/Learning-GoLang/25mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening to the port 4000")
}
