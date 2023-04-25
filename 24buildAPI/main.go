package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB

var courses []Course

// middleware/helper file

func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("API-Build my first api using GoLang")
	r := mux.NewRouter()

	//seeding

	courses = append(courses, Course{CourseId: "1", CourseName: "ViteJS", CoursePrice: 0, Author: &Author{Fullname: "himanshu chhatwal", Website: "himashu.go.dev"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "JavaScrpt", CoursePrice: 0, Author: &Author{Fullname: "himanshu chhatwal", Website: "himashu.go.dev"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "NodeJS", CoursePrice: 0, Author: &Author{Fullname: "himanshu chhatwal", Website: "himashu.go.dev"}})

	//Routing

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")

	//listen to a port

	log.Fatal(http.ListenAndServe(":3000", r))

}

// controller

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome To API by Himanshu Chhatwal</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	// Grab Id from request

	params := mux.Vars(r)

	//looping thru course to find the id and returning the reponse

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id ")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post One Course")
	w.Header().Set("Content-Type", "application/json")

	// what if  body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about the data which is sent like this "{}"

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		json.NewEncoder(w).Encode("No Data inside the json")
		return
	}

	// Generate the unique id , converting it to a string
	// append course into  courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode("Added the data")
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	// loop,id,remove(idex,index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}
