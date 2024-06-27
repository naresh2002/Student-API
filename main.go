package main

import (
	"log"
	"net/http"
	"os"
	"student-api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	lg := log.New(os.Stdout, "student-api ", log.LstdFlags)
	studentHandler := handlers.NewStudent(lg)

	router := mux.NewRouter()

	router.HandleFunc("/student/all", studentHandler.GetStudents).Methods("GET")
	router.HandleFunc("/student/{id}", studentHandler.GetStudentByID).Methods("GET")
	router.HandleFunc("/student/add", studentHandler.CreateStudent).Methods("POST")
	router.HandleFunc("/student/update/{id}", studentHandler.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student/delete/{id}", studentHandler.DeleteStudent).Methods("DELETE")
	router.HandleFunc("/student/summary/{id}", studentHandler.GetStudentSummary).Methods("GET")

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	log.Printf("Server starting on port 8000\n")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
