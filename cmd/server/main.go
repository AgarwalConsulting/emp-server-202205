package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	empHTTP "algogrit.com/emp-server/employees/http"
	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
)

var (
	port  = envOrDefault("PORT", "8000")
	dbURL = envOrDefault("DB_URL", "postgres://localhost:5432/emp-demo?sslmode=disable")
)

func envOrDefault(key string, dflt string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		return dflt
	}

	return val
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
	var empRepo = repository.NewInMem()
	// var empRepo = repository.NewSQL(dbURL)
	var empSvc = service.NewV1(empRepo)
	var empHandler = empHTTP.New(empSvc)

	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!"

		fmt.Fprintln(w, msg)
	})

	empHandler.SetupRoutes(r)

	log.Println("Starting server on port: " + port + "...")
	// http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))
	http.ListenAndServe(":"+port, LoggingMiddleware(r))
}
