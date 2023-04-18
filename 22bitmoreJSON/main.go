package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Learning about JSON")
	EncodeJSON()
}

func EncodeJSON() {
	YTCourses := []course{
		{"ReactJS", 299, "Youtube.com", "abc123", []string{"web-dev", "JS"}},
		{"MERN", 299, "Youtube.com", "sdc123", []string{"full-stack", "JS"}},
		{"MEVN", 299, "Youtube.com", "dvsc123", nil},
	}

	finalJSON, err := json.MarshalIndent(YTCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJSON)

}
