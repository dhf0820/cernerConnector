package main

import (
	"encoding/json"
	"fmt"

	//m "github.com/dhf0820/ROIPrint/pkg/model"
	"github.com/gorilla/mux"
	//"gopkg.in/mgo.v2/bson"
	"net/http"
	"os"
	//"strconv"
	//"strings"
	//"time"
	//"github.com/davecgh/go-spew/spew"
	//"github.com/oleiade/reflections"
)

type HealthResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func WriteHealthResponse(w http.ResponseWriter, status int, message string) error {
	w.Header().Set("Content-Type", "application/json")
	var resp HealthResponse
	resp.Status = status
	resp.Message = message

	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err

	}
	return nil
}

// Routes processes
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	version := fmt.Sprintf("OK: Version %s-%s  Environment: %s", "CernerConnector", os.Getenv("CodeVersion"), Env)
	WriteHealthResponse(w, 200, version)
	fmt.Println(version)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GotHere")
	params := mux.Vars(r)
	fmt.Printf("params:59 - %v\n", params)
	// fhirSystem, uri, err := determineFhirSystem(r)
	// if err != nil {
	// 	WriteHealthResponse(w, 400, "Not Responding")
	// 	return
	// }
	curVersion := os.Getenv("VERSION")
	version := fmt.Sprintf("OK: cernerConn Version %s  Environment: %s", curVersion, Env)
	// version := fmt.Sprintf("OK: Version %s-%s Facility: %s at %s using FhirVersion: %s  URI: %s", "uc_Cache", os.Getenv("CodeVersion"),
	// 		fhirSystem.DisplayName, fhirSystem.FacilityName,fhirSystem.FhirVersion, uri)
	WriteHealthResponse(w, 200, version)
	fmt.Println(version)
}
