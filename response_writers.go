package main

import (
	"encoding/json"
	"fmt"

	//"io/ioutil"
	"net/http"

	//log "github.com/sirupsen/logrus"

	//"github.com/davecgh/go-spew/spew"
	//"github.com/dhf0820/fhir4"
	//cm "github.com/dhf0820/baseConnector/common"
	"github.com/davecgh/go-spew/spew"
	"github.com/dhf0820/fhir4"
	fhir "github.com/dhf0820/fhir4"
	common "github.com/dhf0820/uc_common"
	//"github.com/gorilla/mux"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"os"
	//"strconv"
	//"strings"
	//"time"
	//token "github.com/dhf0820/token"
)

// ####################################### Response Writers Functions #######################################
func WriteFhirOperationOutcome(w http.ResponseWriter, status int, resp *fhir.OperationOutcome) error {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	default:
		w.WriteHeader(status)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("WriteFhirOperationOutcome:42  --  Error marshaling JSON:", err)
		return err
	}
	return nil
}

func WriteHttpResponse(w http.ResponseWriter, status int, body []byte) error {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	case 409:
		w.WriteHeader(http.StatusConflict)
	default:
		w.WriteHeader(status)
	}

	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		fmt.Println("WriteHttpResponse:71  --  Error marshaling JSON:", err)
		return err
	}
	return nil
}
func WriteSaveResponse(w http.ResponseWriter, status int, resp *common.SaveResponse) error {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	default:
		w.WriteHeader(status)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("WriteSaveResponse:83  --  Error marshaling JSON:", err)
		return err
	}
	return nil
}

func WriteFhirResource(w http.ResponseWriter, status int, resp *common.ResourceResponse) error {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	default:
		w.WriteHeader(status)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("WriteFhirResource:108  --  Error marshaling JSON:", err)
		return err
	}
	return nil
}
func WriteFhirResourceBundle(w http.ResponseWriter, status int, resp *common.ResourceResponse) error {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	default:
		w.WriteHeader(status)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("WriteFhirResourceBundle:128  --  Error marshaling JSON:", err)
		return err
	}
	return nil
}

func WriteFhirBundle(w http.ResponseWriter, status int, resp *fhir4.Bundle) error {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	default:
		w.WriteHeader(status)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("WriteFhirBundle:149  --  Error marshaling JSON:", err)
		return err
	}
	return nil
}

func WriteFhirResponse(w http.ResponseWriter, status int, resp *common.ResourceResponse) error {
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("WriteFhirResponse:170  --  Status: %d\n", status)
	fmt.Printf("WriteFhirResponse:171  --  Data:  %s\n", spew.Sdump(resp.Resource.Binary))
	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	default:
		w.WriteHeader(status)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("WriteFhirResponse:184  --  Error marshaling JSON:", err)
		return err
	}
	return nil
}

//################################### FHIR Responses ####################################

func CreateOperationOutcome(code fhir.IssueType, severity fhir.IssueSeverity, details *string) *fhir.OperationOutcome {
	fmt.Printf("CreateOperationOutcome:179  --  Code: %s   Error Message : %s\n", code, *details)
	s := *details
	outcome := fhir.OperationOutcome{}
	issue := fhir.OperationOutcomeIssue{}
	issue.Code = code
	issue.Severity = severity
	issue.Details = &fhir.CodeableConcept{}
	issue.Details.Text = &s
	outcome.Issue = append(outcome.Issue, issue)
	return &outcome
}

func CreateOpOutcome(srcIssues []fhir.OperationOutcomeIssue) *fhir.OperationOutcome {
	fmt.Printf("CreateOpOutcome:192  --  Error Message : %s\n", *srcIssues[0].Details.Text)
	//s := *details
	outcome := fhir.OperationOutcome{}
	outcome.Issue = srcIssues
	//issue := srcIssue
	// issue.Code = code
	// issue.Severity = severity
	// issue.Details = &fhir.CodeableConcept{}
	// issue.Details.Text = &s
	//outcome.Issue = append(outcome.Issue, issue)
	return &outcome
}
