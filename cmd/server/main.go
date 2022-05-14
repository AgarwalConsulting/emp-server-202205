package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"algogrit.com/emp-server/entities"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
)

var empRepo = repository.NewInMem()
var empSvc = service.NewV1(empRepo)

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")

	employees, err := empSvc.Index()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmployee entities.Employee
	err := json.NewDecoder(req.Body).Decode(&newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	createdEmp, err := empSvc.Create(newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf8")
	json.NewEncoder(w).Encode(createdEmp)
}

func LoggingMiddleware(h http.Handler) http.Handler {
	middleware := func(w http.ResponseWriter, req *http.Request) {
		begin := time.Now()

		h.ServeHTTP(w, req)

		log.Printf("%s %s | duration: %s\n", req.Method, req.URL, time.Since(begin))
	}

	return http.HandlerFunc(middleware)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!"

		fmt.Fprintln(w, msg)
	})

	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")

	log.Println("Starting server on port: 8000...")
	// http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))
	http.ListenAndServe(":8000", LoggingMiddleware(r))
}
