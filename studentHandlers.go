package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//AddNewStudent ...
func AddNewStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, "Requested body is not in correct format", http.StatusBadRequest)
	}
	usnChannel := make(chan string)
	go nextAvailableUSN(usnChannel)

	student.USN = <-usnChannel
	studentList = append(studentList, student)
	JSONResponseWriter(w, nil, http.StatusCreated)
}

//BulkInsert ...
func BulkInsert(w http.ResponseWriter, r *http.Request) {
	var students Students
	err := json.NewDecoder(r.Body).Decode(&students)
	if err != nil {
		http.Error(w, "Requested body is not in correct format", http.StatusBadRequest)
	}
	usnChannel := make(chan string)
	go nextAvailableUSN(usnChannel)
	for _, student := range students {
		student.USN = <-usnChannel
		studentList = append(studentList, student)
	}
	JSONResponseWriter(w, nil, http.StatusCreated)
}

//GetStudentList ...
func GetStudentList(w http.ResponseWriter, r *http.Request) {
	JSONResponseWriter(w, studentList, http.StatusOK)
}

//RemoveByID ...
func RemoveByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usn, _ := vars["id"]
	err := studentList.RemoveStudent(usn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

//GetStudentDetails ...
func GetStudentDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usn, _ := vars["id"]
	var result Student
	for _, student := range studentList {
		if student.USN == usn {
			result = student
			break
		}
	}
	if len(result.USN) > 0 {
		JSONResponseWriter(w, result, http.StatusOK)
	} else {
		http.NotFound(w, r)
	}
}

//UpdateStudentDetails ...
func UpdateStudentDetails(w http.ResponseWriter, r *http.Request) {
	var updateStudent Student
	json.NewDecoder(r.Body).Decode(&updateStudent)
	for index, student := range studentList {
		if student.USN == updateStudent.USN {
			studentList[index] = updateStudent
			break
		}
	}
	w.WriteHeader(http.StatusOK)
}
