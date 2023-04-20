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
	// EncodeJSON()
	DecodeJSON()
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

func DecodeJSON() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "MERN",
		"Price": 299,
		"Platform": "Youtube.com",
		"tags": ["full-stack","JS"]
	}
	
	`)

	var MERNCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &MERNCourse)
		fmt.Printf("%#v\n", MERNCourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and the type is %T\n", k, v, v)
	}
}
