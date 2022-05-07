package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Employee struct {
	ID         int    `json:"-"`
	Name       string `json:"name"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"project"`
}

// func (Employee) MarshalJSON() ([]byte, error)

var employees = []Employee{
	{1, "Gaurav", "LnD", 1001},
	{2, "Shoba", "SRE", 1002},
	{3, "Naveen", "Cloud", 10010},
}

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")

	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmployee Employee
	err := json.NewDecoder(req.Body).Decode(&newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	newEmployee.ID = len(employees) + 1
	employees = append(employees, newEmployee)

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	json.NewEncoder(w).Encode(newEmployee)
}

// func EmployeesHandler(w http.ResponseWriter, req *http.Request) {
// 	if req.Method == "POST" {
// 		EmployeeCreateHandler(w, req)
// 	} else if req.Method == "GET" {
// 		EmployeesIndexHandler(w, req)
// 	} else {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 	}
// }

func main() {
	// r := http.NewServeMux()
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!"

		// w.Write([]byte(msg))
		fmt.Fprintln(w, msg)
	})

	// r.HandleFunc("/employees", EmployeesHandler)
	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")

	http.ListenAndServe(":8000", r)
}
