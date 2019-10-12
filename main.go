package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// Student represents a ...
type Student struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
}

// Course represents a ...
type Course struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// HomeRouterHandler represents a ...
func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, http.StatusOK)
	log.Printf("%v GET '%s'\n", http.StatusOK, r.URL.Path)
}

// APIRouterHandler represents a ...
func APIRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

// StudentsRouterHandler represents a ...
func StudentsRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, string(getStudents()))
}

// CoursesRouterHandler represents a ...
func CoursesRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, string(getCourses()))
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

func getCourses() []byte {
	connStr := "host=localhost user=postgres password=postgres dbname=playground sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM courses;")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	courses := []Course{}

	for rows.Next() {
		c := Course{}
		err := rows.Scan(&c.ID, &c.Title)
		if err != nil {
			log.Fatal("rows.Scan: ", err)
		}
		courses = append(courses, c)
	}

	response, err := json.Marshal(courses)
	if err != nil {
		log.Fatal("json.Marshal: ", err)
	}

	return response
}

func main() {
	http.HandleFunc("/", HomeRouterHandler)
	http.HandleFunc("/api/healthz", APIRouterHandler)
	http.HandleFunc("/api/students", StudentsRouterHandler)
	http.HandleFunc("/api/courses", CoursesRouterHandler)
	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
