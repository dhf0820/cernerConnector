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
		//"/system/{SystemId}/Patient",
		//findResource,
		findPatient,
	},
	// Route{
	// 	"GetPatient",
	// 	"GET",
	// 	"/api/rest/v1/{Resource}/{resourceId}",
	// 	getResource,
	// 	//getPatient,
	// },
	Route{
		"FindDiagnosticReports",
		"GET",
		"/api/rest/v1/DiagnosticReport",
		findDiagnosticRept,
	},
	Route{
		"GetDiagnosticReport",
		"GET",
		"/api/rest/v1/DiagnosticReport/{resourceId}",
		getDiagnosticRept,
	},
	Route{
		"FindObservation",
		"GET",
		"/api/rest/v1/Observation",
		findObservation,
	},
	Route{
		"GetObservation",
		"GET",
		"/api/rest/v1/Observation/{resourceId}",
		getObservation,
	},
	Route{
		"FindConditions",
		"GET",
		"/api/rest/v1/Condition",
		findCondition,
	},
	Route{
		"GetCondition",
		"GET",
		"/api/rest/v1/Condition/{resourceId}",
		getCondition,
	},

	Route{
		"GetNutritionOrder",
		"GET",
		"/api/rest/v1/NutritionOrder/{resourceId}",
		getResource,
	},
	Route{
		"FindAllergyIntolerance",
		"GET",
		"/api/rest/v1/AllergyIntolerance",
		findResource,
	},
	Route{
		"FindConditions",
		"GET",
		"/api/rest/v1/Condition",
		findResource,
	},
	Route{
		"GetCondition",
		"GET",
		"/api/rest/v1/Condition/{resourceId}",
		getResource,
	},
	Route{
		"FindDocumentReferences",
		"GET",
		"/api/rest/v1/DocumentReference",
		findDocumentReference,
	},
	Route{
		"GetDocumentReferences",
		"GET",
		"/api/rest/v1/DocumentReference/{resourceId}",
		findDocumentReference,
	},
	Route{
		"FindCoverage",
		"GET",
		"/api/rest/v1/Coverage",
		findResource,
	},
	Route{
		"FindCoverage",
		"GET",
		"/api/rest/v1/Coverage/resourceId",
		getResource,
	},
	// Route{
	// 	"GetDocumentReference",
	// 	"GET",
	// 	"/api/rest/v1/{Resource}/{resourceId}",
	// 	getResource,
	// },
	Route{
		"FindEncounters",
		"GET",
		"/api/rest/v1/Encounter",
		findResource,
	},

	// Route{
	// 	"FindResource",
	// 	"GET",
	// 	"/api/rest/v1/resource/{resource}",
	// 	findResource,
	// },
	// Route{
	// 	"GetResource",
	// 	"GET",
	// 	"/api/rest/v1/resource/{Resource}/{resourceId}",
	// 	getResource,
	// },
	Route{
		"FindGoal",
		"GET",
		"/api/rest/v1/Goal",
		findResource,
	},
	Route{
		"GetGoal",
		"GET",
		"/api/rest/v1/Goal/{resourceId}",
		getResource,
	},
	// Route{
	// 	"FindMetadata",
	// 	"GET",
	// 	"/api/rest/v1/resource/{resource}",
	// 	findResource,
	// },
	// Route{
	// 	"GetMetadata",
	// 	"GET",
	// 	"/api/rest/v1/resource/{resource}",
	// 	findResource,
	// },
	Route{
		"FindObservations",
		"GET",
		"/api/rest/v1/Observation",
		findObservation,
	},
	Route{
		"GetObservation",
		"GET",
		"/api/rest/v1/Observation/{ressourceId}",
		getResource,
	},
	Route{
		"GetBinary",
		"GET",
		"/api/rest/v1/Binary/{resourceId}",
		getResource,
	},
	Route{
		"FindProcedure",
		"GET",
		"/api/rest/v1/Procedure",
		findResource,
	},
	// Route{
	// 	"GetProcedure",
	// 	"GET",
	// 	"/api/rest/v1/{Resource}/{resourceId}",
	// 	getResource,
	// },
	Route{
		"FindQuestionaireResponse",
		"GET",
		"/api/rest/v1/QuestionaireResponse",
		findResource,
	},
	Route{
		"FindQuestionaireResponse",
		"GET",
		"/api/rest/v1/QuestionaireResponse{resourceId}",
		getResource,
	},
	// Route{
	// 	"GetResource",
	// 	"GET",
	// 	"/api/rest/v1/{Resource}/{resourceId}",
	// 	getResource,
	// },
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
