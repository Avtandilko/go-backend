package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

// swagger:route GET /api/v2/students students idOfStudentsEndpoint
// responses:
//   200: StudentResponse

// Student represents a ...
// swagger:response StudentResponse
type Student struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"secondname"`
	Email     string `json:"email"`
}

// swagger:route GET /api/v2/courses courses idOfCoursesEndpoint
// responses:
//   200: CourseResponse

// Course represents a ...
// swagger:response CourseResponse
type Course struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// HomeRouterHandler represents a ...
func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, http.StatusOK)
	log.Printf("%v GET '%s'\n", http.StatusOK, r.URL.Path)
}

// APIRouterHealthHandler represents a ...
func APIRouterHealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

// getStudents represents a ...
func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(readStudents())
}

// getStudent represents a ...
func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range readStudents() {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Student{})
}

// getCourse represents a ...
func getCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range readCourses() {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Course{})
}

// getCourses represents a ...
func getCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(readCourses())
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func readStudents() []Student {
	conn := NewDbConn()

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", conn.DbHost, conn.DbUsername, conn.DbPassword, conn.DbName)
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
		err := rows.Scan(&s.ID, &s.Firstname, &s.Lastname, &s.Email)
		if err != nil {
			log.Fatal("rows.Scan: ", err)
		}
		students = append(students, s)
	}

	return students
}

func readCourses() []Course {
	conn := NewDbConn()
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", conn.DbHost, conn.DbUsername, conn.DbPassword, conn.DbName)

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

	return courses
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeRouterHandler)
	r.HandleFunc("/api/v2/healthz", APIRouterHealthHandler)
	r.HandleFunc("/api/v2/students", getStudents)
	r.HandleFunc("/api/v2/students/{id}", getStudent)
	r.HandleFunc("/api/v2/courses", getCourses)
	r.HandleFunc("/api/v2/courses/{id}", getCourse)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
