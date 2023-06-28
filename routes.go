package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"HealthCheck",
		"GET",
		"/api/rest/v1/healthcheck",
		HealthHandler,
	},
	/////////////////////////////////////////////////////////////////////////////////
	//                                FindResources                                //
	/////////////////////////////////////////////////////////////////////////////////

	Route{
		"FindPatient",
		"GET",
		"/api/rest/v1/Patient",
		findResource,
		//findPatient,
	},
	Route{
		"GetPatient",
		"GET",
		"/api/rest/v1/Patient/{patientId}",
		getResource,
		//getPatient,
	},
	Route{
		"FindPatientsDocRefs",
		"GET",
		"/api/rest/v1/DocumentReference",
		findResource,
	},
	// Route{
	// 	"SavePatient",
	// 	"POST",
	// 	"/api/rest/v1/Patient",
	// 	savePatient,
	// },
	// Route{
	// 	"DebbieTest",
	// 	"GET",
	// 	"/api/rest/v1/Patient/{patientId}/document?type=diagnosticReport",
	// 	DebbieTest,
	// },
	// Route{
	// 	"DebbieTest",
	// 	"GET",
	// 	"/facility/{facilityId}/patient/{patientId}",
	// 	getPatient,
	// },
}
