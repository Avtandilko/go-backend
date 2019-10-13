package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// StudentID represents a ...
// swagger:parameters getStudent
type StudentID struct {
	// The ID of the student
	//
	// in: path
	// required: true
	ID string `json:"id"`
}

// CourseID represents a ...
// swagger:parameters getCourse
type CourseID struct {
	// The ID of the course
	//
	// in: path
	// required: true
	ID string `json:"id"`
}

// Student represents a ...
// swagger:response StudentResponse
type Student struct {
    // in: path
    // required: true
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

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
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

// getStudents represents a ...
func getStudents(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /api/v2/students students getStudents
	// responses:
	//   200: StudentResponse
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(readStudents())
}

// getStudent represents a ...
func getStudent(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /api/v2/students/{id} students getStudent
	// responses:
	//   200: StudentResponse
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

// getCourses represents a ...
func getCourses(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /api/v2/courses courses getCourses
	// responses:
	//   200: CourseResponse
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(readCourses())
}

// getCourse represents a ...
func getCourse(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /api/v2/courses/{id} courses getCourse
	// responses:
	//   200: CourseResponse
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

func createStudent(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    s := Student{}
    _ = json.NewDecoder(r.Body).Decode(&s)
    fmt.Println(s)
	fmt.Println(json.NewEncoder(w).Encode(s))
	
	conn := NewDbConn()

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", conn.DbHost, conn.DbUsername, conn.DbPassword, conn.DbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(fmt.Sprintf("INSERT INTO students (firstname, lastname, email) VALUES ('%s', '%s', '%s');", s.Firstname, s.Lastname, s.Email))

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
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

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeRouterHandler)
	r.HandleFunc("/api/v2/health", APIRouterHealthHandler).Methods("GET")
	r.HandleFunc("/api/v2/students", getStudents).Methods("GET")
	r.HandleFunc("/api/v2/students/{id}", getStudent).Methods("GET")
	r.HandleFunc("/api/v2/students", createStudent).Methods("POST")
	r.HandleFunc("/api/v2/courses", getCourses).Methods("GET")
	r.HandleFunc("/api/v2/courses/{id}", getCourse).Methods("GET")
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
