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

	// GET Subrouter
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/student/all", studentHandler.GetStudents)
	getRouter.HandleFunc("/student/{id}", studentHandler.GetStudentByID)
	getRouter.HandleFunc("/student/summary/{id}", studentHandler.GetStudentSummary)

	// POST Subrouter with middleware
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.Use(studentHandler.MiddlewareValidateStudent)
	postRouter.HandleFunc("/student/add", studentHandler.CreateStudent)

	// PUT Subrouter with middleware
	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.Use(studentHandler.MiddlewareValidateStudent)
	putRouter.HandleFunc("/student/update/{id}", studentHandler.UpdateStudent)

	// DELETE Subrouter
	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/student/delete/{id}", studentHandler.DeleteStudent)

	// // OR :
	// router.HandleFunc("/student/all", studentHandler.GetStudents).Methods(http.MethodGet)
	// router.HandleFunc("/student/{id}", studentHandler.GetStudentByID).Methods(http.MethodGet)
	// router.HandleFunc("/student/add", studentHandler.CreateStudent).Methods(http.MethodPost)
	// router.HandleFunc("/student/update/{id}", studentHandler.UpdateStudent).Methods(http.MethodPut)
	// router.HandleFunc("/student/delete/{id}", studentHandler.DeleteStudent).Methods(http.MethodDelete)
	// router.HandleFunc("/student/summary/{id}", studentHandler.GetStudentSummary).Methods(http.MethodGet)

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	log.Printf("Server starting on port 8000\n")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
