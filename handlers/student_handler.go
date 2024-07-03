package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"student-api/data"

	"github.com/gorilla/mux"
)

type Students struct {
	lg *log.Logger
}

func NewStudent(lg *log.Logger) *Students {
	return &Students{lg}
}

var mutex sync.Mutex

func (s *Students) GetStudents(rw http.ResponseWriter, req *http.Request) {
	s.lg.Println("Handle Get All Students")
	ls := data.GetStudents()
	d, err := json.Marshal(ls)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	rw.Write(d)
}

func (s *Students) GetStudentByID(rw http.ResponseWriter, req *http.Request) {
	s.lg.Println("Handle Get Student By ID")
	vars := mux.Vars(req)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(rw, "Invalid student ID", http.StatusBadRequest)
		return
	}
	student, exists := data.GetStudents()[id]
	if !exists {
		http.Error(rw, "Student not found", http.StatusNotFound)
		return
	}
	d, err := json.Marshal(student)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	rw.Write(d)
}

func (s *Students) CreateStudent(rw http.ResponseWriter, req *http.Request) {
	s.lg.Println("Handle Create Student")

	student := req.Context().Value(KeyStudent{}).(data.Student)

	mutex.Lock()
	defer mutex.Unlock()

	student.ID = data.GetNextID()
	student.CreatedAt = time.Now().String()
	student.UpdatedAt = time.Now().String()
	data.StudentsList[student.ID] = student

	rw.WriteHeader(http.StatusCreated)
	d, err := json.Marshal(student)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	rw.Write(d)
}

func (s *Students) UpdateStudent(rw http.ResponseWriter, req *http.Request) {
	s.lg.Println("Handle Update Student")
	vars := mux.Vars(req)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(rw, "Invalid student ID", http.StatusBadRequest)
		return
	}
	student, exists := data.StudentsList[id]
	if !exists {
		http.Error(rw, "Student not found", http.StatusNotFound)
		return
	}

	updatedStudent := req.Context().Value(KeyStudent{}).(data.Student)

	mutex.Lock()
	defer mutex.Unlock()

	updatedStudent.ID = id
	updatedStudent.CreatedAt = student.CreatedAt
	updatedStudent.UpdatedAt = time.Now().String()
	data.StudentsList[id] = updatedStudent

	d, err := json.Marshal(updatedStudent)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	rw.Write(d)
}

func (s *Students) DeleteStudent(rw http.ResponseWriter, req *http.Request) {
	s.lg.Println("Handle Delete Student")
	vars := mux.Vars(req)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(rw, "Invalid student ID", http.StatusBadRequest)
		return
	}
	_, exists := data.StudentsList[id]
	if !exists {
		http.Error(rw, "Student not found", http.StatusNotFound)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	delete(data.StudentsList, id)
	rw.WriteHeader(http.StatusNoContent)
}

func (s *Students) GetStudentSummary(rw http.ResponseWriter, req *http.Request) {
	s.lg.Println("Handle Get Student Summary")
	vars := mux.Vars(req)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(rw, "Invalid student ID", http.StatusBadRequest)
		return
	}
	student, exists := data.GetStudents()[id]
	if !exists {
		http.Error(rw, "Student not found", http.StatusNotFound)
		return
	}

	type OllamaRequest struct {
		Model  string `json:"model"`
		Prompt string `json:"prompt"`
		Stream bool   `json:"stream"`
	}
	studentMarshal, err := json.Marshal(student)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	promptData := strings.ReplaceAll(fmt.Sprintf("Summarise the student details given in data into a sentence. data : %s", string(studentMarshal)), "\"", "")

	payload := OllamaRequest{
		Model:  "llama3",
		Prompt: promptData,
		Stream: false,
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

	url := "http://localhost:11434/api/generate"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		http.Error(rw, "Unable to unmarshal response", http.StatusInternalServerError)
		return
	}

	// Removing unnecessary fields
	delete(response, "context")
	delete(response, "done_reason")
	delete(response, "eval_count")
	delete(response, "eval_duration")
	delete(response, "load_duration")
	delete(response, "prompt_eval_count")
	delete(response, "prompt_eval_duration")

	// Marshal the response back to JSON
	cleanedBody, err := json.Marshal(response)
	if err != nil {
		http.Error(rw, "Unable to marshal cleaned response", http.StatusInternalServerError)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(rw, string(cleanedBody), resp.StatusCode)
		return
	}

	rw.Write(cleanedBody)
}
