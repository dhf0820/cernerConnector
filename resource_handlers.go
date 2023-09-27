package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/dhf0820/vslog"

	"github.com/davecgh/go-spew/spew"
	fhir "github.com/dhf0820/fhir4"
	common "github.com/dhf0820/uc_core/common"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// //"os"
	// //"strconv"
	"strings"
	"time"

	jw_token "github.com/dhf0820/jwToken"
)

//####################################### Response Writers Functions #######################################
// func WriteFhirOperationOutcome(w http.ResponseWriter, status int, resp *fhir.OperationOutcome) error {
// 	w.Header().Set("Content-Type", "application/json")

// 	switch status {
// 	case 200:
// 		w.WriteHeader(http.StatusOK)
// 	case 400:
// 		w.WriteHeader(http.StatusBadRequest)
// 	case 401:
// 		w.WriteHeader(http.StatusUnauthorized)
// 	case 403:
// 		w.WriteHeader(http.StatusForbidden)
// 	default:
// 		w.WriteHeader(status)
// 	}
// 	err := json.NewEncoder(w).Encode(resp)
// 	if err != nil {
// 		fmt.Println("WriteFhirOperationOutcome:42  --  Error marshaling JSON:", err)
// 		return err
// 	}
// 	return nil
// }

// func WriteHttpResponse(w http.ResponseWriter, status int, body []byte) error {
// 	w.Header().Set("Content-Type", "application/json")

// 	switch status {
// 	case 200:
// 		w.WriteHeader(http.StatusOK)
// 	case 400:
// 		w.WriteHeader(http.StatusBadRequest)
// 	case 401:
// 		w.WriteHeader(http.StatusUnauthorized)
// 	case 403:
// 		w.WriteHeader(http.StatusForbidden)
// 	case 409:
// 		w.WriteHeader(http.StatusConflict)
// 	default:
// 		w.WriteHeader(status)
// 	}

// 	err := json.NewEncoder(w).Encode(body)
// 	if err != nil {
// 		fmt.Println("WriteHttpResponse:71  --  Error marshaling JSON:", err)
// 		return err
// 	}
// 	return nil
// }
// func WriteSaveResponse(w http.ResponseWriter, status int, resp *common.SaveResponse) error {
// 	w.Header().Set("Content-Type", "application/json")

// 	switch status {
// 	case 200:
// 		w.WriteHeader(http.StatusOK)
// 	case 400:
// 		w.WriteHeader(http.StatusBadRequest)
// 	case 401:
// 		w.WriteHeader(http.StatusUnauthorized)
// 	case 403:
// 		w.WriteHeader(http.StatusForbidden)
// 	default:
// 		w.WriteHeader(status)
// 	}
// 	err := json.NewEncoder(w).Encode(resp)
// 	if err != nil {
// 		fmt.Println("WriteSaveResponse:83  --  Error marshaling JSON:", err)
// 		return err
// 	}
// 	return nil
// }

// func WriteFhirResource(w http.ResponseWriter, status int, resp *common.ResourceResponse) error {
// 	w.Header().Set("Content-Type", "application/json")

// 	switch status {
// 	case 200:
// 		w.WriteHeader(http.StatusOK)
// 	case 400:
// 		w.WriteHeader(http.StatusBadRequest)
// 	case 401:
// 		w.WriteHeader(http.StatusUnauthorized)
// 	case 403:
// 		w.WriteHeader(http.StatusForbidden)
// 	default:
// 		w.WriteHeader(status)
// 	}
// 	err := json.NewEncoder(w).Encode(resp)
// 	if err != nil {
// 		fmt.Println("WriteFhirResource:108  --  Error marshaling JSON:", err)
// 		return err
// 	}
// 	return nil
// }
// func WriteFhirResourceBundle(w http.ResponseWriter, status int, resp *common.ResourceResponse) error {
// 	w.Header().Set("Content-Type", "application/json")

// 	switch status {
// 	case 200:
// 		w.WriteHeader(http.StatusOK)
// 	case 400:
// 		w.WriteHeader(http.StatusBadRequest)
// 	case 401:
// 		w.WriteHeader(http.StatusUnauthorized)
// 	case 403:
// 		w.WriteHeader(http.StatusForbidden)
// 	default:
// 		w.WriteHeader(status)
// 	}
// 	err := json.NewEncoder(w).Encode(resp)
// 	if err != nil {
// 		fmt.Println("WriteFhirResourceBundle:128  --  Error marshaling JSON:", err)
// 		return err
// 	}
// 	return nil
// }

// func WriteFhirBundle(w http.ResponseWriter, status int, resp *fhir4.Bundle) error {
// 	w.Header().Set("Content-Type", "application/json")

// 	switch status {
// 	case 200:
// 		w.WriteHeader(http.StatusOK)
// 	case 400:
// 		w.WriteHeader(http.StatusBadRequest)
// 	case 401:
// 		w.WriteHeader(http.StatusUnauthorized)
// 	case 403:
// 		w.WriteHeader(http.StatusForbidden)
// 	default:
// 		w.WriteHeader(status)
// 	}
// 	err := json.NewEncoder(w).Encode(resp)
// 	if err != nil {
// 		fmt.Println("WriteFhirBundle:149  --  Error marshaling JSON:", err)
// 		return err
// 	}
// 	return nil
// }

// func WriteFhirResponse(w http.ResponseWriter, status int, resp *common.ResourceResponse) error {
// 	w.Header().Set("Content-Type", "application/json")

// 	switch status {
// 	case 200:
// 		w.WriteHeader(http.StatusOK)
// 	case 400:
// 		w.WriteHeader(http.StatusBadRequest)
// 	case 401:
// 		w.WriteHeader(http.StatusUnauthorized)
// 	case 403:
// 		w.WriteHeader(http.StatusForbidden)
// 	default:
// 		w.WriteHeader(status)
// 	}
// 	err := json.NewEncoder(w).Encode(resp)
// 	if err != nil {
// 		fmt.Println("WriteFhirResponse:170  --  Error marshaling JSON:", err)
// 		return err
// 	}
// 	return nil
// }

// //################################### FHIR Responses ####################################

// func CreateOperationOutcome(code fhir.IssueType, severity fhir.IssueSeverity, details *string) *fhir.OperationOutcome {
// 	fmt.Printf("CreateOperationOutcome:179  --  Code: %s   Error Message : %s\n", code, *details)
// 	s := *details
// 	outcome := fhir.OperationOutcome{}
// 	issue := fhir.OperationOutcomeIssue{}
// 	issue.Code = code
// 	issue.Severity = severity
// 	issue.Details = &fhir.CodeableConcept{}
// 	issue.Details.Text = &s
// 	outcome.Issue = append(outcome.Issue, issue)
// 	return &outcome
// }

// func CreateOpOutcome(srcIssues []fhir.OperationOutcomeIssue) *fhir.OperationOutcome {
// 	fmt.Printf("CreateOpOutcome:192  --  Error Message : %s\n", *srcIssues[0].Details.Text)
// 	//s := *details
// 	outcome := fhir.OperationOutcome{}
// 	outcome.Issue = srcIssues
// 	//issue := srcIssue
// 	// issue.Code = code
// 	// issue.Severity = severity
// 	// issue.Details = &fhir.CodeableConcept{}
// 	// issue.Details.Text = &s
// 	//outcome.Issue = append(outcome.Issue, issue)
// 	return &outcome
// }

// //####################################### Route Handlers #######################################

// // // Header will have the fhir services token
// // // Header may or url will have the id of the FhirConnector
// // func searchPatient(w http.ResponseWriter, r *http.Request) {
// // 	// var pspTags map[string]string
// // 	// tagFields := make(map[string]string)
// // 	// var Limit int
// // 	// var Skip int
// // 	fmt.Printf("searchPatient:86 - Request: %s \n", spew.Sdump(r))
// // 	//buildFieldsByTagMap("schema", *psp)
// // 	//facility = "demo"
// // 	userId := r.Header.Get("UserId")
// // 	uri := r.RequestURI
// // 	parts := strings.Split(uri, "v1/")
// // 	uri = parts[1]
// // 	fmt.Printf("searchPatient:93 - URI = %s\n", uri)
// // 	resource := GetFHIRResource(r)

// // 	fhirVersion := r.Header.Get("FhirVersion")
// // 	if fhirVersion == "" {
// // 		fhirVersion = "R4"
// // 	}
// // 	// if resource == "Patient" {
// // 	// 	urlA, err := r.URL.Parse(r.RequestURI)
// // 	// 	if err != nil {
// // 	// 		err = fmt.Errorf("error parsing patient URI: %s", err.Error())
// // 	// 		errMsg := err.Error()
// // 	// 		fmt.Printf("findResource:102 - r.URL.Parse error = %s\n", errMsg)
// // 	// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// // 	// 		return
// // 	// 	}
// // 	// 	fmt.Printf("findResource:106 - r.URL.Parse = %v\n", urlA)
// // 	// 	urlB := *urlA
// // 	// 	uriValues := urlB.Query()
// // 	// 	fmt.Printf("findResource:109 - uriValues= %v\n", uriValues)
// // 	// 	ident := uriValues.Get("identifier")
// // 	// 	if ident != "" { // There is identifier Search, use it
// // 	// 		fmt.Printf("findResource:110 - using Identifier: %s to search\n", ident)
// // 	// 	} else {
// // 	// 		fmt.Printf("findResource:110 - using other search params: %v\n", uriValues)
// // 	// 	}

// // 	// }
// // 	//fhirVersion := GetFHIRVersion(r)
// // 	//cacheBaseURL := fmt.Sprintf("%s/%s/v1/", r.Host, parts[0])
// // 	if err := r.ParseForm(); err != nil {
// // 		err = fmt.Errorf("error parsing query: %s", err.Error())
// // 		errMsg := err.Error()
// // 		fmt.Printf("searchPatient:126 - %s\n", errMsg)
// // 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// // 		return
// // 	}
// // 	FhirId := GetFhirId(r)
// // 	fmt.Printf("findResource:132 - FhirKey - [%s]\n", FhirId)
// // 	fhirSystem, err := GetFhirSystem(FhirId)

// // 	if err != nil {
// // 		fmt.Printf("GetFhirSystem failed with : %s\n", err.Error())
// // 		err = fmt.Errorf("fhirSystem error:  %s", err.Error())
// // 		errMsg := err.Error()
// // 		fmt.Printf("searchPatient:138 - %s\n", errMsg)
// // 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityFatal, &errMsg))
// // 		return
// // 	}
// // 	fmt.Printf("searchPatient:142 -  %s/n", spew.Sdump(fhirSystem))

// // 	if resource == "Patient" {
// // 		urlA, err := r.URL.Parse(r.RequestURI)
// // 		if err != nil {
// // 			err = fmt.Errorf("error parsing patient URI: %s", err.Error())
// // 			errMsg := err.Error()
// // 			fmt.Printf("searchPatient:149 - r.URL.Parse error = %s\n", errMsg)
// // 			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// // 			return
// // 		}
// // 		//fmt.Printf("findResource:106 - r.URL.Parse = %v\n", urlA)
// // 		urlB := *urlA
// // 		uriValues := urlB.Query()
// // 		fmt.Printf("searchPatient:156 - uriValues= %v\n", uriValues)
// // 		idSearch := uriValues.Get("identifier")
// // 		idValue := ""
// // 		if idSearch != "" { // There is identifier Search, use it
// // 			fmt.Printf("searchPatient:160 - using Identifier: %s to search\n", idSearch)
// // 			ids := strings.Split(idSearch, "|")
// // 			if len(ids) != 2 {
// // 				err = fmt.Errorf("invalid identifier: %s", idSearch)
// // 				errMsg := err.Error()
// // 				fmt.Printf("searchPatient:165 - r.URL.Parse error = %s\n", errMsg)
// // 				WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// // 				return
// // 			}
// // 			idName := ids[0]
// // 			idSearchValue := ids[1]
// // 			idents := fhirSystem.Identifiers
// // 			for _, id := range idents {
// // 				if id.Name == idName {
// // 					idValue = id.Value
// // 					break
// // 				}
// // 			}
// // 			if idValue == "" { //Not configured identifier
// // 				err = fmt.Errorf("identifier type: %s is not configured", idName)
// // 				errMsg := err.Error()
// // 				fmt.Printf("searchPatient:181 - Identifiers = %s\n", errMsg)
// // 				WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// // 				return
// // 			}
// // 			uri = fmt.Sprintf("Patient?identifier=%s", idValue+idSearchValue)
// // 			fmt.Printf("searchPatient:186 - New Identifier search Value: %s\n", uri)
// // 			// uriValues.Del("identifier")
// // 			// uriValues.Set("identifier", id)
// // 			// //urlB.RawQuery = uriValues.Encode()
// // 			// r.URL.RawQuery = uriValues.Encode()
// // 			// //curUri := r.RequestURI
// // 			// //urUriParts := strings.Split(curUri, "?")
// // 			// r.RequestURI = uriValues.Encode()
// // 			// fmt.Printf("\n\n$$$ searchResources: 188 - Updated request: %s\n\n", spew.Sdump(r))
// // 		} else {
// // 			fmt.Printf("searchPatient:196 - using other search params: %v\n", uriValues)
// // 		}

// // 	}
// // 	var bundle *fhir.Bundle
// // 	var header *common.CacheHeader
// // 	//resourceId := r.Header.Get("Fhir-System")
// // 	// params := mux.Vars(r)
// // 	// fmt.Printf("findResource params:115 %v\n", params)

// // 	//resource := strings.Split(uri, "?")[0]
// // 	fmt.Printf("\nsearchPatient:207 - resource = %s  uri = %s\n", resource, uri)
// // 	url := fhirSystem.FhirUrl + "/" + uri
// // 	fmt.Printf("searchPatient:209 - calling %s \n", url)
// // 	var totalPages int64
// // 	// if resourceId != "" {
// // 	// 	fmt.Printf("findResource:128 - Get %s with [%s]\n", resource, resourceId)
// // 	// } else {
// // 	fmt.Printf("searchPatient:214 Search %s with %s\n", url, r.RequestURI)
// // 	totalPages, bundle, header, err = SearchPatient(fhirSystem, url, resource, userId, r.RequestURI)
// // 	if err != nil {
// // 		err = fmt.Errorf("fhirSearch url: %s error:  %s", url, err.Error())
// // 		errMsg := err.Error()
// // 		fmt.Printf("searchPatient:219 - %s\n", errMsg)
// // 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityInformation, &errMsg))
// // 		return
// // 	}
// // 	//}
// // 	fmt.Printf("searchPatient:224 - Get %s bundle successful\n", resource)
// // 	fmt.Printf("searchPatient:225 - Number in page: %d\n", len(bundle.Entry))
// // 	fmt.Printf("searchPatient:226 - PageNumber: %d\n", header.PageId)
// // 	resp := ResourceResponse{}
// // 	//hostParts := strings.Split(r.Host, ":")

// // 	host := common.GetKVData(GetConfig().Data, "cacheHost")
// // 	//host := os.Getenv("CORE_ADDRESS")
// // 	fmt.Printf("searchPatient:232 - ##host: %s\n\n\n", host)
// // 	header.CacheUrl = fmt.Sprintf("%s%sv1/Cache/%s/", host, parts[0], header.QueryId)

// // 	resp.Bundle = bundle
// // 	resp.Resource = header.ResourceType
// // 	resp.BundleId = *bundle.Id
// // 	resp.ResourceType = resource
// // 	resp.Status = 200
// // 	resp.QueryId = header.QueryId
// // 	resp.PageNumber = header.PageId
// // 	resp.CountInPage = len(bundle.Entry)
// // 	resp.TotalPages = totalPages
// // 	resp.Header = header
// // 	resp.Message = "Ok"
// // 	fmt.Printf("searchPatient:246 - returning a resource bundle\n")
// // 	WriteFhirResourceBundle(w, resp.Status, &resp)

// // }
// var Resource string
// var JWToken string
// var Payload *token.Payload

// func DebbieTest(w http.ResponseWriter, r *http.Request) {
// 	_, err := cm.FhirPatientSearch(r)
// 	if err != nil {
// 		msg := err.Error()
// 		WriteFhirOperationOutcome(w, 200, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &msg))
// 		return
// 	}
// 	params := mux.Vars(r)
// 	fmt.Printf("DebbieTest324  --  Prmd: %v\n", params)
// 	Resource := ""
// 	uri := r.URL.RequestURI()
// 	fmt.Printf("DebbieTest:327  --  uri: %s\n", uri)
// 	values := r.URL.Query()
// 	for k, v := range values {
// 		fmt.Println(k, " => ", v)
// 	}

// 	return
// 	parts := strings.Split(uri, "api/")
// 	p0 := parts[1]
// 	fmt.Printf("searchPatient:93 - URI-p0 = %s\n", p0)
// 	findParts := strings.Split(p0, "?")
// 	fmt.Printf("uriParts length = %d\n\n", len(findParts))
// 	if len(findParts) == 1 { // is a getResource
// 		getParts := strings.Split(p0, "/")
// 		fmt.Printf("getParts = %v\n", getParts)
// 		Resource = getParts[0]
// 	} else {
// 		Resource = findParts[0]
// 	}
// 	msg := fmt.Sprintf("Resource = %s\n", Resource)
// 	WriteFhirOperationOutcome(w, 200, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &msg))
// }

func findResource(w http.ResponseWriter, r *http.Request) {
	JWToken = r.Header.Get("Authorization")
	Payload, status, err := jw_token.ValidateToken(JWToken, "")
	if err != nil {
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	userId := Payload.UserId
	resourceType := ""
	params := mux.Vars(r)
	log.Debug3(fmt.Sprintf("Params: %v\n", params))
	if params["resource"] != "" {
		log.Debug3(" --  Using Resource in params")
		resourceType = params["resource"]
	} else {
		// url := r.URL
		log.Debug2(fmt.Sprintf("--  url = %s", r.URL.Path))
		log.Debug2(fmt.Sprintf("findResource:442  --  uri = %s", r.URL.RequestURI()))
		resourceType = DetermineResource(r.URL.Path, "/api/rest/v1/")
		if resourceType == "" {
			errMsg := log.ErrMsg("Resource not found in URL")
			WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
			return
		}
	}
	log.Debug1(" --  Resource = " + resourceType)

	//log.Debug1(" -- being called for resource: [%s]\n", Resource)
	log.Debug2("--  Reading Body")
	//fmt.Printf("findResource:453  --  r = %s\n", spew.Sdump(r))
	body, err := io.ReadAll(r.Body) // Should be ConnectorPayload
	if err != nil {
		fmt.Printf("findResource:456  --  ReadAll FhirSystem error %s\n", err.Error())
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	connectorPayload := common.ConnectorPayload{}

	//fmt.Printf("findResource:470  --  body = %s\n", string(body))
	err = json.Unmarshal(body, &connectorPayload)
	if err != nil {
		errMsg := log.ErrMsg("--  unmarshal ConnectorPayload err = " + err.Error())
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}

	connectorConfig := connectorPayload.ConnectorConfig
	log.Debug5("-- ConnectorPayload = " + spew.Sdump(connectorPayload))
	uri := r.URL.RequestURI()
	fmt.Printf("findResource:475  --  uri: %s\n", uri)
	fmt.Printf("findResource:476  --  URL.Path() = %s\n", r.URL.Path)
	fmt.Printf("findResource:477  --  query = %s\n", r.URL.RawQuery)
	//r.URL.RawQuery
	// uri = parts[1]
	// // p0 := parts[1]
	// fmt.Printf("findResource:430  --  part0 = %s\n", parts[0])
	// fmt.Printf("findResource:431  --  part1 = %s\n", parts[1])
	// findParts := strings.Split(uri, "?")
	// fmt.Printf("findResource:433  --  uriParts length = %d p0 = %s,  p1 = %s\n\n", len(findParts), findParts[0], findParts[1])
	// if len(findParts) == 1 { // is a getResource
	// 	getParts := strings.Split(uri, "/")
	// 	fmt.Printf("findResource:436  --  getParts = %v\n", getParts)
	// 	Resource = getParts[0]
	// 	// } else {
	// 	// 	Query = p0
	// 	// 	Resource = findParts[0]
	// } else {
	// 	Resource = findParts[0]
	// }

	//Resource = r.Header.Get("Resource")

	//buildFieldsByTagMap("schema", *psp)
	//facility = "demo"
	//userId := r.Header.Get("UserId")
	//fhirId := r.Header.Get("Fhir-System")
	// FhirId := GetFhirId(r)
	// fmt.Printf("findResource:275 - FhirKey - [%s]\n", FhirId)
	// fhirSystem, err := GetFhirSystem(FhirId)

	// if err != nil {
	// 	fmt.Printf("GetFhirSystem failed with : %s\n", err.Error())
	// 	err = fmt.Errorf("fhirSystem error:  %s", err.Error())
	// 	errMsg := err.Error()
	// 	fmt.Printf("findResource:282 - %s\n", errMsg)
	// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }

	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	log.Errorf("findResource:374  --  ReadAll error %s\n", err.Error())
	// 	errMsg := err.Error()
	// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// fmt.Printf("FindResource:379  --  bodyString = %s\n\n", string(body))
	// fhirSystem := common.FhirSystem{}
	// err = json.Unmarshal(body, &fhirSystem)
	// if err != nil {
	// 	fmt.Printf("\nFindResource:383  --  unmarshal err = %s\n", err.Error())
	// 	errMsg := err.Error()
	// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// fmt.Printf("FindResource:388  -- FhirSystem = %s\n", spew.Sdump(fhirSystem))

	//TODO: Split on the Resource keeping the actual url variables and query params
	//uri := r.RequestURI
	// fmt.Printf("findResource:341  --  uri = %s\n", r.RequestURI)
	// u := r.URL
	// fmt.Printf("findResource:342  --  %s\n", u.RequestURI())
	// fmt.Printf("findResource:343  --  URL = %s\n", spew.Sdump(u))
	// uri := r.URL.RequestURI()
	// uriParts := strings.Split(uri, "/")
	// queryString := ""
	// i := 1
	// for _, part := range uriParts {
	// 	log.Printf("uri part: %d = %s\n", i, part)
	// 	i++
	// }
	// parts := strings.Split(uriParts[len(uriParts)-1], "?")
	// if len(parts) > 1 {
	// 	//url query and keep the query element[1]
	// 	//element[0] is the resource
	// 	queryString = parts[len(parts)-1]
	// } else {
	// 	// last element is the id of the query
	// 	queryString = parts[len(uriParts)-1]
	// }

	// log.Printf("QueryString = %s\n", queryString)

	// fmt.Printf("findResource:362 - URI = %s\n", uri)
	// resource := Resource
	//resource := GetFHIRResource(r)

	//fhirVersion := "R4"
	//fhirVersion := r.Header.Get("FhirVersion")
	// if fhirVersion == "" {
	// 	fhirVersion = "R4"
	// }
	// if resource == "Patient" {
	// 	urlA, err := r.URL.Parse(r.RequestURI)
	// 	if err != nil {
	// 		err = fmt.Errorf("error parsing patient URI: %s", err.Error())
	// 		errMsg := err.Error()
	// 		fmt.Printf("findResource:102 - r.URL.Parse error = %s\n", errMsg)
	// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
	// 		return
	// 	}
	// 	fmt.Printf("findResource:106 - r.URL.Parse = %v\n", urlA)
	// 	urlB := *urlA
	// 	uriValues := urlB.Query()
	// 	fmt.Printf("findResource:109 - uriValues= %v\n", uriValues)
	// 	ident := uriValues.Get("identifier")
	// 	if ident != "" { // There is identifier Search, use it
	// 		fmt.Printf("findResource:110 - using Identifier: %s to search\n", ident)
	// 	} else {
	// 		fmt.Printf("findResource:110 - using other search params: %v\n", uriValues)
	// 	}
	// }
	//fhirVersion := GetFHIRVersion(r)
	//cacheBaseURL := fmt.Sprintf("%s/%s/v1/", r.Host, parts[0])
	if err := r.ParseForm(); err != nil {
		errMsg := log.ErrMsg("error parsing query: " + err.Error())
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
		return
	}

	queryStr := ""
	log.Debug3("-- Resource: " + resourceType)
	switch strings.ToLower(resourceType) {
	case "patient":
		log.Debug3("-  case: " + resourceType)
		queryStr = fmt.Sprintf("%s?%s", resourceType, r.URL.RawQuery) //BuildPatientQuery(r)
	case "documentreference":
		log.Debug3("-  case: " + resourceType)
	default:
		log.Debug3("- default: " + resourceType)
	}

	log.Debug3("  --  queryStr = " + queryStr)
	if strings.ToLower(resourceType) == "patient" {
		urlA, err := r.URL.Parse(r.URL.RequestURI())
		if err != nil {
			errMsg := log.ErrMsg("error parsing patient URI: " + err.Error())
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
			return
		}
		//fmt.Printf("findResource:106 - r.URL.Parse = %v\n", urlA)
		urlB := *urlA
		uriValues := urlB.Query()
		//fmt.Printf("\n\n\nfindResource:624 - uriValues= %v\n", uriValues)
		idSearch := uriValues.Get("identifier")
		idValue := ""
		if idSearch != "" { // There is identifier Search, use it
			log.Debug3("- using Identifier: " + idSearch + " to search")
			ids := strings.Split(idSearch, "|")
			if len(ids) != 2 {
				errMsg := log.ErrMsg("invalid identifier: " + idSearch)
				WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
				return
			}
			idName := ids[0]
			idSearchValue := ids[1]
			idents := connectorPayload.ConnectorConfig.Identifiers
			for _, id := range idents {
				if id.Name == idName {
					idValue = id.Value
					break
				}
			}
			if idValue == "" { //Not configured identifier
				errMsg := log.ErrMsg("identifier type: " + idName + " is not configured")
				WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
				return
			}
			uri = fmt.Sprintf("Patient?identifier=%s", idValue+idSearchValue)
			log.Debug3(" - New Identifier search Value: " + uri)
			// uriValues.Del("identifier")
			// uriValues.Set("identifier", id)
			// //urlB.RawQuery = uriValues.Encode()
			// r.URL.RawQuery = uriValues.Encode()
			// //curUri := r.RequestURI
			// //urUriParts := strings.Split(curUri, "?")
			// r.RequestURI = uriValues.Encode()
			// fmt.Printf("\n\n$$$ searchResources: 188 - Updated request: %s\n\n", spew.Sdump(r))
		} else {
			log.Debug3(fmt.Sprintf(" - using other search params: %v", uriValues))
		}
	}
	var bundle *fhir.Bundle
	var header *common.CacheHeader
	//resourceId := r.Header.Get("Fhir-System")
	log.Debug3(" - connectorPayload = " + spew.Sdump(connectorPayload))
	qryStr := r.URL.RawQuery

	log.Debug3(fmt.Sprintf(" - resource = %s  uri = %s", resourceType, qryStr))
	url := connectorPayload.ConnectorConfig.HostUrl + resourceType + "?" + qryStr
	//url := connectorPayload.System.Url + Resource + "?" + qryStr
	log.Debug3(" - calling " + url)
	var totalPages int64
	log.Debug3(fmt.Sprintf(" --  Search %s with %s", resourceType, qryStr))
	startTime := time.Now()
	totalPages, bundle, header, err = FindResource(&connectorPayload, resourceType, userId, qryStr, JWToken)
	if err != nil {
		errMsg := log.ErrMsg(fmt.Sprintf("--  fhirSearch url: %s  --error:  %s", url, err.Error()))
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityInformation, &errMsg))
		return
	}
	//}
	log.Debug3(fmt.Sprintf(" - Get %s bundle successful in %s", resourceType, time.Since(startTime)))
	log.Debug3(fmt.Sprintf(" - Total Pages: %d", totalPages))
	log.Debug3(fmt.Sprintf(" - Number in page: %d", len(bundle.Entry)))
	log.Debug3(fmt.Sprintf(" - PageNumber: %d", header.PageId))
	log.Debug3(fmt.Sprintf(" - QueryId: %s", header.QueryId))
	resp := common.ResourceResponse{}
	//fmt.Printf("findResource:628 - Header: %s\n", spew.Sdump(header))
	host := connectorPayload.ConnectorConfig.HostUrl
	//host := common.GetKVData(GetConfig().Data, "cacheHost")
	log.Debug3(" --  host: " + host)
	cacheBundleUrl := fmt.Sprintf("%s/%s/BundleTransaction", connectorConfig.CacheUrl, header.QueryId)
	//header.CacheUrl = fmt.Sprintf("%s%sv1/Cache/%s/", host, parts[0], header.QueryId)
	log.Debug3("  --  CacheUrl = " + cacheBundleUrl)
	//resp.Resource = header.ResourceType
	resp.BundleId = *bundle.Id
	resp.ResourceType = resourceType
	resp.Status = 200
	resp.QueryId = header.QueryId
	resp.PageNumber = header.PageId
	resp.CountInPage = len(bundle.Entry)
	resp.TotalPages = totalPages
	resp.Header = header
	resp.Message = "Ok"
	logTime := time.Now()
	log.Debug3(fmt.Sprintf("--  resp without bundle: " + spew.Sdump(resp)))
	log.Debug3(fmt.Sprintf("--  Time to log = %s", time.Since(logTime)))
	resp.Bundle = bundle
	log.Debug3(fmt.Sprintf("--  Number of entries in buldle: %d", len(bundle.Entry)))
	log.Debug3(fmt.Sprintf("--  QueryId: " + header.QueryId))
	FillResourceResponse(&resp, resourceType)
	//fmt.Printf("findResource:614  --  Returning Bundle: %s\n", spew.Sdump(bundle))
	//WriteFhirResourceBundle(w, resp.Status, &resp)
	WriteFhirResponse(w, resp.Status, &resp)
}

func getResource(w http.ResponseWriter, r *http.Request) {
	log.Debug3("getResource  Start")
	JWToken = r.Header.Get("Authorization")
	_, status, err := jw_token.ValidateToken(JWToken, "")
	if err != nil {
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	params := mux.Vars(r)
	Resource := DetermineGetResource(r.URL.Path, "/")
	resourceId := params["resourceId"]
	resourceType := params["Resource"]
	log.Debug3(fmt.Sprintf("params = %s", params))
	log.Debug3(fmt.Sprintf("param resource : %s - %s  ResId: %s", Resource, resourceType, resourceId))

	// fmt.Printf("getResource:674 - Request - ")
	// spew.Dump(r)
	// Resource = r.Header.Get("Resource")
	//url := r.URL
	//x := *url
	//fmt.Printf("getResource:679  --  url = %s\n", spew.Sdump(x))
	//Resource := DetermineResource(x.Path, "/api/rest/v1/")
	//fmt.Printf("getResource:744  --  Resource = %s", Resource)
	if err := r.ParseForm(); err != nil {
		errMsg := log.ErrMsg("error parsing query: " + err.Error())
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeInvalid, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	cp, err := GetConnectorPayload(r)
	if err != nil {
		errMsg := log.ErrMsg("gGetConnectorPayload error:  " + err.Error())
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	// FhirId := GetFhirId(r)
	// fhirSystem, err := GetFhirSystem(FhirId)
	// if err != nil {
	// 	err = fmt.Errorf("fhirConnetcor error:  %s", err.Error())
	// 	errMsg := err.Error()
	// 	fmt.Printf("getResource:680 -  GetFhirSystem err = %s\n", errMsg)
	// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	//fhirSystem := cp.FhirSystem
	// uriParts := strings.Split(r.RequestURI, "v1/")
	// fmt.Printf("\nuriParts: %v\n", uriParts)
	// uriParts1 := strings.Split(uriParts[1], "/")
	// resource = uriParts1[0]

	//fmt.Printf("getResource:775 - cp to use : %s\n", spew.Sdump(cp))
	// params := mux.Vars(r)
	// fmt.Printf("getResource:274 - params  %v\n", params)
	// resourceId := params["id"]
	log.Debug3(fmt.Sprintf("Retrieving %s Record for id: [%s]", resourceType, resourceId))
	if resourceId == "" {
		errMsg := log.ErrMsg(fmt.Sprintf("GetResource %s specific ID string is required", Resource))
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverity(fhir.IssueTypeInvalid), &errMsg))
		return
	}
	//TODO: Handle Get Resource by specific ID.  All Resources including Binary.
	resp := common.ResourceResponse{}
	results, err := GetResource(cp, Resource, resourceId, JWToken)
	if err == nil {
		resp.Status = 200
		resp.Message = "Ok"
	} else {
		errMsg := log.ErrMsg("GetResource error: " + err.Error())
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverity(fhir.IssueTypeInvalid), &errMsg))
		return
	}

	log.Debug3("FillResourceResponse for " + strings.ToLower(resourceType))
	switch strings.ToLower(resourceType) {
	case "patient":
		log.Debug3("Processing Patient")
		//fmt.Printf("GetResource:801  --  patient raw = %v\n", results)
		patient, err := fhir.UnmarshalPatient(results)
		if err != nil {
			errMsg := log.ErrMsg("UnmarshalPatient error:  " + err.Error())
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
			return
		}
		hdr := common.ResourceHeader{}
		narrative := common.Narrative{}
		narrative.Div = patient.Text.Div
		hdr.Narrative = &narrative
		ds := []common.KVData{}
		ds = append(ds, common.KVData{Name: "Name", Value: *patient.Name[0].Text})
		ds = append(ds, common.KVData{Name: "id", Value: *patient.Id})
		//TODO: Add MRN to display Fields
		//TODO: Add DOB to display Fields

		resp.Resource.ResourceHeader = &hdr
		resp.Resource.ResourceHeader.DisplayFields = ds
		//resp.Patient = &patient
		resp.ResourceType = resourceType
		resp.Resource.ResourceType = resourceType
		resp.Resource.Patient = &patient
		resp.Resource.ResourceId = *patient.Id
		resp.ResourceId = *patient.Id
		resp.Message = "Ok"
		resp.PageNumber = 1
		resp.TotalPages = 1
		resp.CountInPage = 1
		resp.QueryId = primitive.NewObjectID().Hex()
		resp.Status = 200
		log.Debug3("Patient case final " + spew.Sdump(resp))
	case "binary":
		log.Debug3("Processing Binary")
		//fmt.Printf("GetResource:831  --  patient raw = %v\n", results)
		binary, err := fhir.UnmarshalBinary(results)
		if err != nil {
			errMsg := log.ErrMsg("UnmarshalBinary error: " + err.Error())
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
			return
		}
		//mt.Printf("GetResource:840  --  Binary: %s\n", spew.Sdump(binary))
		resp.ResourceType = resourceType
		resp.Resource.ResourceType = resourceType
		resp.Resource.Binary = binary
		resp.Resource.ResourceId = *binary.Id
		resp.ResourceId = *binary.Id
		resp.Message = "Ok"
		resp.PageNumber = 1
		resp.TotalPages = 1
		resp.CountInPage = 1
		resp.QueryId = primitive.NewObjectID().Hex()
		resp.Status = 200
	case "documentreference":
		log.Debug3("Processing DocumentReference")
		//fmt.Printf("GetResource:831  --  patient raw = %v\n", results)
		data, err := fhir.UnmarshalDocumentReference(results)
		if err != nil {
			errMsg := log.ErrMsg(fmt.Sprintf("Unmarshal %s error: %s", resourceType, err.Error()))
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
			return
		}
		log.Debug3(fmt.Sprintf("Resource %s   contains: %s", resourceType, spew.Sdump(data)))
		resp.ResourceType = resourceType
		resp.Resource.ResourceType = resourceType

		resp.Resource.DocumentReference = &data
		resp.Resource.ResourceId = *data.Id
		resp.ResourceId = *data.Id
		resp.Message = "Ok"
		resp.PageNumber = 1
		resp.TotalPages = 1
		resp.CountInPage = 1
		resp.QueryId = primitive.NewObjectID().Hex()
		resp.Status = 200
	}
	// resp.ResourceType = Resource
	// resp.Resource.Resource = results
	//log.Debug3("returning resource: " + spew.Sdump(resp))
	WriteFhirResponse(w, resp.Status, &resp)
}

// func getCachePage(w http.ResponseWriter, r *http.Request) {
// 	var err error
// 	//fmt.Printf("getCachePage:300 - %s \n", spew.Sdump(r))

// 	params := mux.Vars(r)
// 	fmt.Printf("getCachePage:484  -- %v\n", params)
// 	queryId := params["queryId"] // The id assigned to the query that created the cache
// 	pageNumber := params["pageNum"]
// 	//pageId := params["page_id"]
// 	fmt.Printf("Retrieving bundle for id: [%s]  Page: [%s]\n", queryId, pageNumber)
// 	FhirId := GetFhirId(r)
// 	fmt.Printf("getCachePage:490 - FhirKey - [%s]\n", FhirId)
// 	fhirSystem, err := GetFhirSystem(FhirId)
// 	if err != nil {
// 		fmt.Printf("GetFhirSystem failed with : %s\n", err.Error())
// 		err = fmt.Errorf("fhirSystem error:  %s", err.Error())
// 		errMsg := err.Error()
// 		fmt.Printf("getCachePage::496 - %s\n", errMsg)
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	fmt.Printf("getCachePage:500 -  %s/n", spew.Sdump(fhirSystem))
// 	if queryId == "" || pageNumber == "" {
// 		err = fmt.Errorf("GetCachedPage queryId: %s, pageNumber: %s -  error:  %s", queryId, pageNumber, "query_id and pageNumber are required")
// 		errMsg := err.Error()
// 		fmt.Printf("getCachePage:312 - %s\n", errMsg)
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeIncomplete, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	} else {
// 		fmt.Printf("Call GetCache for  queryId: %s  pageNumber: %s\n", queryId, pageNumber)
// 		pageId, err := strconv.Atoi(pageNumber)
// 		if err != nil {
// 			err = fmt.Errorf("PageNumber: [%s] is invalid %s", pageNumber, err.Error())
// 			errMsg := err.Error()
// 			fmt.Printf("getCachePage:321 - %s\n", errMsg)
// 			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeInvalid, fhir.IssueSeverityFatal, &errMsg))

// 		}
// 		totalPages, bundle, header, err := GetCache(queryId, pageId)
// 		if err != nil {
// 			err = fmt.Errorf("GetCachePage QueryId: %s, page: %s -  error:  %s", queryId, pageNumber, err.Error())
// 			errMsg := err.Error()
// 			fmt.Printf("GetCachePage: 329 - %s\n", errMsg)
// 			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityFatal, &errMsg))
// 			return
// 		}
// 		//parts := strings.Split(r.RequestURI, "v1/")
// 		resp := common.ResourceResponse{}
// 		resp.Bundle = bundle
// 		resp.ResourceType = header.ResourceType
// 		resp.BundleId = *bundle.Id
// 		resp.Status = 200
// 		resp.QueryId = header.QueryId
// 		resp.PageNumber = header.PageId
// 		resp.CountInPage = len(bundle.Entry)
// 		resp.TotalPages = totalPages
// 		resp.Header = header

// 		//host := common.GetKVData(GetConfig().Data, "cacheHost")
// 		//header.CacheUrl = fmt.Sprintf("%s%sv1/Cache/%s/", host, parts[0], header.QueryId)
// 		header.CacheUrl = fmt.Sprintf("%s/Cache/%s/", fhirSystem.UcUrl, header.QueryId)
// 		//header.CacheUrl = fmt.Sprintf("%s/%sv1/Cache/%s/", r.Host, parts[0], header.QueryId)

// 		resp.Message = "Ok"
// 		fmt.Printf("getCachePage:351 - returning a cached %s bundle\n", header.ResourceType)
// 		WriteFhirResourceBundle(w, resp.Status, &resp)

// 	}
// }

// // func getCacheStatus(w http.ResponseWriter, r *http.Request) {
// // 	var err error
// // 	fmt.Printf("getCacheStatus:335 - %s \n", spew.Sdump(r))
// // 	params := mux.Vars(r)
// // 	fmt.Printf("params:337 - %v\n", params)
// // 	queryId := params["queryId"] // The id assigned to the query that created the cache
// // 	fmt.Printf("Count how many pages of cache are in an ID\n")

// // 	if queryId == "" {
// // 		err = fmt.Errorf("GetCacheStatus queryId: %s -  error:  %s", queryId, "query_id is required")
// // 		errMsg := err.Error()
// // 		fmt.Printf("Handler:344 - %s\n", errMsg)
// // 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeIncomplete, fhir.IssueSeverityFatal, &errMsg))
// // 		return
// // 	} else {
// // 		fmt.Printf("Count Pages for queryId: %s\n", queryId)
// // 		totalPages, err := TotalCacheForQuery(queryId)
// // 		if err != nil {
// // 			err = fmt.Errorf("GetCacheStatus QueryId: %s -  error:  %s", queryId, err.Error())
// // 			errMsg := err.Error()
// // 			fmt.Printf("GetCacheStatus: 353 - %s\n", errMsg)
// // 			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityFatal, &errMsg))
// // 			return
// // 		}
// // 		parts := strings.Split(r.RequestURI, "v1/")
// // 		resp := common.ResourceResponse{}

// // 		//resp.ResourceType = header.ResourceType
// // 		header := &common.CacheHeader{}
// // 		resp.Status = 200
// // 		resp.QueryId = queryId
// // 		//resp.PageNumber = header.PageId
// // 		//resp.CountInPage = len(bundle.Entry)
// // 		resp.TotalPages = totalPages
// // 		resp.Header = header
// // 		resp.Header.QueryId = queryId
// // 		host := common.GetKVData(GetConfig().Data, "cacheHost")
// // 		resp.Header.CacheUrl = fmt.Sprintf("%s/%sv1/Cache/%s/", host, parts[0], header.QueryId)

// // 		resp.Message = "Ok"
// // 		//fmt.Printf("$$$:373 - returning  cached %s bundle\n", header.ResourceType)
// // 		WriteFhirResourceBundle(w, resp.Status, &resp)
// // 	}
// // }

// // func checkStatus(w http.ResponseWriter, r *http.Request) {
// // 	var err error
// // 	fmt.Printf("checkStatus:301 - %s \n", spew.Sdump(r))
// // 	params := mux.Vars(r)
// // 	fmt.Printf("params:303 - %v\n", params)
// // 	queryId := params["queryId"] // The id assigned to the query that created the cache
// // 	fmt.Printf("Count how many pages of cache are in an ID\n")

// // 	if queryId == "" {
// // 		err = fmt.Errorf("GetCacheStatus queryId: %s -  error:  %s", queryId, "query_id is required")
// // 		errMsg := err.Error()
// // 		fmt.Printf("Handler:344 - %s\n", errMsg)
// // 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeIncomplete, fhir.IssueSeverityFatal, &errMsg))
// // 		return
// // 	} else {
// // 		fmt.Printf("Count Pages for queryId: %s\n", queryId)
// // 		totalPages, err := TotalCacheForQuery(queryId)
// // 		if err != nil {
// // 			err = fmt.Errorf("GetCacheStatus QueryId: %s -  error:  %s", queryId, err.Error())
// // 			errMsg := err.Error()
// // 			fmt.Printf("GetCacheStatus: 353 - %s\n", errMsg)
// // 			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityFatal, &errMsg))
// // 			return
// // 		}
// // 		parts := strings.Split(r.RequestURI, "v1/")
// // 		resp := common.ResourceResponse{}

// // 		//resp.ResourceType = header.ResourceType
// // 		header := &common.CacheHeader{}
// // 		resp.Status = 200
// // 		resp.QueryId = queryId
// // 		//resp.PageNumber = header.PageId
// // 		//resp.CountInPage = len(bundle.Entry)
// // 		resp.TotalPages = totalPages
// // 		resp.Header = header
// // 		resp.Header.QueryId = queryId
// // 		host := common.GetKVData(GetConfig().Data, "cacheHost")
// // 		resp.Header.CacheUrl = fmt.Sprintf("%s%sv1/Cache/%s/", host, parts[0], header.QueryId)

// // 		resp.Message = "Ok"
// // 		//fmt.Printf("$$$:373 - returning  cached %s bundle\n", header.ResourceType)
// // 		WriteFhirResourceBundle(w, resp.Status, &resp)
// // 	}
// // }

func DetermineResource(url string, prefix string) string {
	parts := strings.SplitAfter(url, prefix)
	resourceType := parts[len(parts)-1]
	log.Debug3("DetermineResource:1011  --  resourceType: " + resourceType)
	return resourceType
}
func DetermineGetResource(url string, prefix string) string {
	log.Debug3(fmt.Sprintf("  --  url: " + url))
	parts := strings.SplitAfter(url, prefix)
	resourceType := strings.TrimRight(parts[len(parts)-2], "/")
	// var subs []string
	log.Debug3("--  resourceType: " + resourceType)
	return resourceType
}
func FillResourceResponse(resp *common.ResourceResponse, resourceType string) error {
	resType := strings.ToLower(resourceType)
	log.Debug3("resType: " + resType)
	resp.Status = 200
	resp.Message = "Ok"
	resp.ResourceType = resourceType
	switch strings.ToLower(resourceType) {
	case "patient":
		//pats := []fhir.Patient{}
		//resData := common.ResourceData{}
		for _, item := range resp.Bundle.Entry {
			resData := common.ResourceData{}
			pat, err := fhir.UnmarshalPatient(item.Resource)
			if err != nil {
				return log.Errorf("UnMarshal(Patient) error = " + err.Error())
			}
			resData.Patient = &pat
			log.Debug3(" Added PatientId: " + *pat.Id)
			resp.Resources = append(resp.Resources, resData)
			//pats = append(resData.Patient, pat)
		}
		//resp.Resources = pats
		log.Debug3(fmt.Sprintf("-- Set %d Patients  Bundle had %d entries\n", len(resp.Resources), len(resp.Bundle.Entry)))
	case "documentreference":
		for _, item := range resp.Bundle.Entry {
			resData := common.ResourceData{}
			docRef, err := fhir.UnmarshalDocumentReference(item.Resource)
			if err != nil {
				return log.Errorf("Unmarshal DocumentReference error = " + err.Error())
			}
			resData.DocumentReference = &docRef
			log.Debug3("Added DocumentReferenceId: " + *docRef.Id)
			resp.Resources = append(resp.Resources, resData)
		}
		log.Debug3(fmt.Sprintf("-- Set %d DocumentReferences  Bundle had %d entries\n", len(resp.Resources), len(resp.Bundle.Entry)))
	case "diagnosticreport":
		for _, item := range resp.Bundle.Entry {
			resData := common.ResourceData{}
			data, err := fhir.UnmarshalDiagnosticReport(item.Resource)
			if err != nil {
				return log.Errorf(" -- error = " + err.Error())
			}
			resData.DiagnosticReport = &data
			fmt.Printf("FillResourceResponse:1064  --  Added DiagnosticReporteId: %s\n", *data.Id)
			resp.Resources = append(resp.Resources, resData)
		}
		log.Debug3(fmt.Sprintf("-- Set %d DiagnosticReport  Bundle had %d entries\n", len(resp.Resources), len(resp.Bundle.Entry)))
	}
	//TODO: Make Switch smarter.
	return nil
}
