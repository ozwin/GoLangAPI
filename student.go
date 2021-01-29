package main

import (
	"errors"
	"math/rand"
	"strconv"
)

//Student ...
type Student struct {
	Name  string
	USN   string
	Batch string
}

//Students ...
type Students []Student

var studentList = Students{
	Student{
		"Jhon Legen",
		"NI12",
		"UG",
	},
	Student{
		"Jackson pawl",
		"N14",
		"PG",
	},
	Student{
		"Mark Rafelo",
		"N15",
		"UG",
	},
}

//RemoveStudent ...
func (students Students) RemoveStudent(removeUSN string) error {
	removeIndex := -1
	for index, student := range students {
		if student.USN == removeUSN {
			removeIndex = index
			break
		}
	}
	if removeIndex > -1 {
		students[removeIndex] = students[len(students)-1]
		studentList = students[:len(students)-1]
		return nil
	}
	return errors.New("Student does not exist")
}
func nextAvailableUSN(c chan string) {
	for {
		c <- "NI" + strconv.Itoa(rand.Intn(9999))
	}
}
