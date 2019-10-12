package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// Global represents a ...
type Global struct {
	Students []Student
	Courses  []Course
}

// Student represents a ...
type Student struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
}

// Course represents a ...
type Course struct {
	ID    int `json:"id"`
	Title string `json:"title"`
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
	fmt.Fprintln(w, string(1))
}

// StudentsRouterHandler represents a ...
func StudentsRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, string(getStudents()))
}

// CoursesRouterHandler represents a ...
func CoursesRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, string(3))
}

func getStudents() []byte {
	connStr := "host=localhost user=postgres password=postgres dbname=playground sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM students;")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	students := []Student{}

	for rows.Next() {
		s := Student{}
		err := rows.Scan(&s.ID, &s.FirstName, &s.SecondName, &s.Email)
		if err != nil {
			log.Fatal("rows.Scan: ", err)
		}
		students = append(students, s)
	}

	response, err := json.Marshal(students)
	if err != nil {
		log.Fatal("json.Marshal: ", err)
	}

	return response
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
