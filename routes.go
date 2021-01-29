package main

import "net/http"

//Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes ...
type Routes []Route

var routes = Routes{
	Route{
		"GetAll",
		"GET",
		"/Student/GetAll",
		GetStudentList,
	},
	Route{
		"DeleteStudent",
		"DELETE",
		"/Student/Delete/{id}",
		RemoveByID,
	},
	Route{
		"AddNewStudent",
		"POST",
		"/Student/Create",
		AddNewStudent,
	},
	Route{
		"GetStudentDetails",
		"GET",
		"/Student/{id}",
		GetStudentDetails,
	},
	Route{
		"BulkInsert",
		"POST",
		"/Student/BulkInsert",
		BulkInsert,
	},
	Route{
		"UpdateDetails",
		"PATCH",
		"/Student/Update",
		UpdateStudentDetails,
	},
}
