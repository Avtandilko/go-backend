package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Global represents a ...
type Global struct {
	Students []Students
	Courses  []Courses
}

// Students represents a ...
type Students struct {
	Name, Surname string
}

// Courses represents a ...
type Courses struct {
	Name string
	ID   int
}

var studentsOne = Students{
	Name:    "Student",
	Surname: "One",
}

var studentsTwo = Students{
	Name:    "Student",
	Surname: "Two",
}

var coursesGo = Courses{
	Name: "Go",
	ID:   1,
}

var global = Global{
	Students: []Students{studentsOne, studentsTwo},
	Courses:  []Courses{coursesGo},
}

func getGlobal() []byte {

	response, err := json.Marshal(global)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return response
}

func getStudents() []byte {

	response, err := json.Marshal([]Students{studentsOne, studentsTwo})
	if err != nil {
		fmt.Println("Error:", err)
	}

	return response
}

func getCourses() []byte {

	response, err := json.Marshal([]Courses{coursesGo})
	if err != nil {
		fmt.Println("Error:", err)
	}

	return response
}

// HomeRouterHandler represents a ...
func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
	r.ParseForm()
	log.Println("path", r.URL.Path)
}

// APIRouterHandler represents a ...
func APIRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

// GlobalRouterHandler represents a ...
func GlobalRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, string(getGlobal()))
}

// StudentsRouterHandler represents a ...
func StudentsRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, string(getStudents()))
}

// CoursesRouterHandler represents a ...
func CoursesRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, string(getCourses()))
}

func main() {

	http.HandleFunc("/", HomeRouterHandler)
	http.HandleFunc("/api/healthz", APIRouterHandler)
	http.HandleFunc("/students", StudentsRouterHandler)
	http.HandleFunc("/courses", CoursesRouterHandler)
	http.HandleFunc("/global", GlobalRouterHandler)
	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
