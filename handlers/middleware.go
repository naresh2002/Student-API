package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"student-api/data"
)

type KeyStudent struct{}

func (s *Students) MiddlewareValidateStudent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		student := data.Student{}
		err := json.NewDecoder(req.Body).Decode(&student)
		if err != nil {
			s.lg.Println("[ERROR] deserializing student", err)
			http.Error(rw, "Error reading student", http.StatusBadRequest)
			return
		}

		student.Name = strings.Join(strings.Fields(student.Name), " ")
		// validate the struct Student
		err = student.Validate()
		if err != nil {
			s.lg.Println("[ERROR] validating student", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating student: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		// add the student to the context
		ctx := context.WithValue(req.Context(), KeyStudent{}, student)
		req = req.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, req)
	})
}
