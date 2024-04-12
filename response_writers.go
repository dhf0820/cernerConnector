package main

import (
	"encoding/json"
	"fmt"
	"strings"

	//"io/ioutil"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	log "github.com/dhf0820/vslog"

	//"github.com/sirupsen/logrus"

	//"github.com/davecgh/go-spew/spew"
	//"github.com/dhf0820/fhir4"
	//cm "github.com/dhf0820/baseConnector/common"
	//"github.com/davecgh/go-spew/spew"
	"github.com/dhf0820/fhir4"
	fhir "github.com/dhf0820/fhir4"
	common "github.com/dhf0820/uc_core/common"
	//common "github.com/dhf0820/uc_core/common"
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
	//log.Debug3("resp: " + spew.Sdump(resp))
	resResp := common.ResourceResponse{}
	resResp.Status = status
	resResp.ResourceType = "OperationOutcome"
	resResp.Message = *resp.Issue[0].Details.Text
	resResp.OpOutcome = resp
	//resResp.OperationOutcome = resp
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
	log.Debug3("resResp: " + spew.Sdump(resResp))
	err := json.NewEncoder(w).Encode(resResp)
	if err != nil {
		log.Error("WriteFhirOperationOutcome  Error marshaling JSON:" + err.Error())
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
	log.Debug3(fmt.Sprint("--  Status: ", status))
	//log.Debug3(fmt.Sprintf("WriteFhirResponse:170  --  Status: %d\n", status))
	//log.Debug3("Data:  " + spew.Sdump(resp))
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

func WriteInitialFhirResponse(w http.ResponseWriter, status int, resp *common.InitialResourceResponse) error {
	w.Header().Set("Content-Type", "application/json")
	log.Debug3(fmt.Sprint("--  Status: ", status))
	//log.Debug3(fmt.Sprintf("WriteFhirResponse:170  --  Status: %d\n", status))
	//log.Debug3("Data:  " + spew.Sdump(resp))
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

		fmt.Println("WriteInitialFhirResponse:184  --  Error marshaling JSON:", err)
		return log.Errorf("--  Error marshaling JSON: " + err.Error())
	}
	return nil
}

//################################### FHIR Responses ####################################

func CreateOperationOutcome(code fhir.IssueType, severity fhir.IssueSeverity, details *string) *fhir.OperationOutcome {
	log.Debug3(fmt.Sprintf("Code: %s   Error Message : %s", code, *details))
	log.ErrMsg(fmt.Sprintf("Code: %s   Error Message : %s", code, *details))
	parts := strings.Split(*details, "--")
	fmt.Printf("\n\n")
	log.Debug3("Parts : " + spew.Sdump(parts))
	fmt.Printf("\n\n")
	t := parts[len(parts)-1]

	s := fmt.Sprintf("Query: %s -- %s", QueryString, *details)
	outcome := fhir.OperationOutcome{}
	issue := fhir.OperationOutcomeIssue{}
	issue.Code = code
	issue.Severity = severity
	issue.Details = &fhir.CodeableConcept{}
	issue.Details.Text = &t
	issue.Diagnostics = &s
	outcome.Issue = append(outcome.Issue, issue)
	return &outcome
}

func CreateOpOutcome(srcIssues []fhir.OperationOutcomeIssue) *fhir.OperationOutcome {
	log.Debug3("--  Error Message : %s" + *srcIssues[0].Details.Text)
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
