package main

/*import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/dhf0820/vslog"

	"github.com/davecgh/go-spew/spew"
	fhir "github.com/dhf0820/fhir4"

	//common "github.com/dhf0820/uc_core/common"
	common "github.com/dhf0820/uc_core/common"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//"os"
	//"strconv"
	"strings"
	"time"

	jw_token "github.com/dhf0820/golangJWT"
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
// // 	//CacheUrlURL := fmt.Sprintf("%s/%s/v1/", r.Host, parts[0])
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
	log.Debug2(" --  findResource	")
	JWToken = r.Header.Get("Authorization")
	Payload, status, err := jw_token.ValidateToken(JWToken, "")
	if err != nil {
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}

	CurrentToken = JWToken
	uri := r.URL.RequestURI()
	log.Debug2(fmt.Sprintf(" --  uri = %s", uri))
	params := mux.Vars(r)
	log.Info(fmt.Sprintf("mux.Vars: %s", spew.Sdump(params)))
	// SystemId := params["SystemId"] // The id of the fhirSystem to use
	// if SystemId == "" {
	// 	SystemId = r.Header.Get("SystemId")
	// }
	// //log.Debug2("SystemId: " + SystemId)
	// CurrentSystemId = SystemId
	// //CurrentSystemId = r.Header.Get("SYSTEMID")
	// if CurrentSystemId == "" {
	// 	errMsg := log.ErrMsg("--  Header SystemId is required ")
	// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// log.Debug2("CurrentSystemId: " + CurrentSystemId)
	userId := Payload.UserId
	userID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		errMsg := log.ErrMsg("--  ReadAll FhirSystem error " + err.Error())
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	CurrentUserID = userID
	resourceType := ""
	connectorPayload := common.ConnectorPayload{}

	//fmt.Printf("findResource:470  --  body = %s\n", string(body))
	err = json.Unmarshal(body, &connectorPayload)
	if err != nil {
		errMsg := log.ErrMsg("--  unmarshal ConnectorPayload err = " + err.Error())
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	log.Debug3("-- ConnectorPayload = " + spew.Sdump(connectorPayload))

	connectorConfig := connectorPayload.System.ConnectorConfig

	params = mux.Vars(r)
	log.Debug2(fmt.Sprintf("Params: %v\n", params))
	if params["resource"] != "" {
		log.Debug2(" --  Using Resource in params")
		resourceType = params["resource"]
	} else {
		// url := r.URL
		// <<<<<<< HEAD
		log.Debug2(fmt.Sprintf("--  url = %s", r.URL.Path))
		uri := r.URL.RequestURI()
		log.Debug2(fmt.Sprintf("--  url = %s", uri))

		URIparts := strings.Split(uri, "&")
		log.Debug2(fmt.Sprintf("  uri = %s", URIparts[len(URIparts)-1]))
		//countURI := URIparts[len(URIparts)-1]
		// uri= URIparts[0 : len(URIparts)-1]+"&_count="
		// log.Debug2(fmt.Sprintf("  newUri = %s", newUri))
		count := URIparts[len(URIparts)-1]
		log.Debug2(fmt.Sprintf("count = %s", count))
		//countParts[0]
		countParts := strings.Split(count, "=")
		log.Debug2(fmt.Sprintf("countParts = %v", countParts))
		if countParts[0] == "_count" {
			errMsg := log.ErrMsg("_count should not be used in this context")
			fmt.Printf("findResource:453  --  %s\n", errMsg)
			WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
			return
		}
		// if countParts[0] == "_count" {
		// 	count := os.Getenv("COUNT")
		// 	if count == "" {
		// 		count = "50"
		// 		//newCountPart := fmt.Sprintf("_count=%s", count)
		// 	}
		// }
		// uri= URIparts[0 : len(URIparts)-1]+"&_count="
		// log.Debug2(fmt.Sprintf("  newUri = %s", newUri))
		//if HasSufix(uri/)) == "/api/rest/v1/" {
		log.Debug2(fmt.Sprintf("--  uri = %s", r.URL.RequestURI()))
		resourceType = DetermineResource(r.URL.Path, "/api/rest/v1/")
		// =======
		// 		log.Debug3(fmt.Sprintf("--  url = %s", r.URL.Path))
		// 		log.Debug3(fmt.Sprintf("--  uri = %s", r.URL.RequestURI()))
		// 		resourceType = DetermineResource(r.URL.Path, "")
		// >>>>>>> 2062832a583c454d91d8636cf44ef1efded64a16
		if resourceType == "" {
			errMsg := log.ErrMsg("Resource not found in URL")
			WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
			return
		}
	}
	log.Debug3("ResourceType : " + resourceType)
	log.Debug3(" --  Resource = " + resourceType)

	//log.Debug1(" -- being called for resource: [%s]\n", Resource)
	log.Debug3("--  Reading Body")
	//fmt.Printf("findResource:453  --  r = %s\n", spew.Sdump(r))
	body, err := io.ReadAll(r.Body) // Should be ConnectorPayload
	if err != nil {
		// <<<<<<< HEAD
		errMsg := log.ErrMsg("--  ReadAll FhirSystem error " + err.Error())
		// =======
		// 		errMsg := log.ErrMsg("--  ReadAll FhirSystem error: " + err.Error())
		// >>>>>>> 2062832a583c454d91d8636cf44ef1efded64a16
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
	log.Debug3("-- ConnectorPayload = " + spew.Sdump(connectorPayload))

	connectorConfig := connectorPayload.System.ConnectorConfig
	// <<<<<<< HEAD

	// =======
	// 	//log.Debug3("-- ConnectorPayload = " + spew.Sdump(connectorPayload))
	// >>>>>>> 2062832a583c454d91d8636cf44ef1efded64a16
	//uri := r.URL.RequestURI()
	log.Debug3(fmt.Sprintf("--  uri: %s\n", uri))
	log.Debug3(fmt.Sprintf("--  URL.Path() = %s\n", r.URL.Path))
	log.Debug3(fmt.Sprintf("--  query = %s\n", r.URL.RawQuery))
	QueryString = r.URL.RawQuery
	log.Debug3("QueryString: " + QueryString)
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
	//CacheUrlURL := fmt.Sprintf("%s/%s/v1/", r.Host, parts[0])
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
			idents := connectorPayload.System.ConnectorConfig.Identifiers
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
	//log.Debug3(" - connectorPayload = " + spew.Sdump(connectorPayload))
	qryStr := r.URL.RawQuery

	log.Debug3(fmt.Sprintf(" - resource = %s  uri = %s", resourceType, qryStr))
	log.Debug3("")
	url := connectorPayload.System.ConnectorConfig.HostUrl + resourceType + "?" + qryStr
	//url := connectorPayload.System.ConnectorConfig.HostUrl + resourceType + "?" + qryStr
	//url := connectorPayload.System.Url + Resource + "?" + qryStr
	log.Debug3(" - calling " + url)
	var totalPages int64
	var inPage int64
	log.Debug3(fmt.Sprintf(" --  Search %s with %s", resourceType, qryStr))
	startTime := time.Now()
	inPage, bundle, header, err = FindResource(&connectorPayload, resourceType, userId, qryStr, JWToken)
	log.Debug3(" - FindResource returned Inpage: " + fmt.Sprint(inPage))
	finalStatus := status
	if err != nil {
		errMsg := log.ErrMsg(fmt.Sprintf("error:  %s", err.Error()))
		fmt.Println(errMsg)
		// errParts := strings.Split(err.Error(), "|")
		// log.Debug3(" - errParts = " + spew.Sdump(errParts))
		// if len(errParts) > 1 { //Not a valid error message
		// 	errMsg = errParts[1]
		// 	finalStatus, err = strconv.Atoi(errParts[0])
		// 	if err != nil {
		// 		finalStatus = 413
		// 	}
		// 	log.Debug3("finalStatus: " + fmt.Sprint(finalStatus))
		// }
		oo := CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityInformation, &errMsg)
		//log.Debug3("OpOutcome: " + spew.Sdump(oo))
		WriteFhirOperationOutcome(w, finalStatus, oo)
		//CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityInformation, &errMsg))
		return
	}

	//finalStatus = status
	log.Debug3(fmt.Sprintf(" - Get %s bundle successful in %s", resourceType, time.Since(startTime)))
	log.Debug3(fmt.Sprintf(" - Total Pages: %d", totalPages))
	log.Debug3(" - inPage :" + fmt.Sprint(inPage))
	log.Debug3(fmt.Sprintf(" - Number in page: %d", len(bundle.Entry)))
	log.Debug3(fmt.Sprintf(" - PageNumber: %d", header.PageId))
	log.Debug3(fmt.Sprintf(" - QueryId: %s", header.QueryId))
	resp := common.ResourceResponse{}
	//fmt.Printf("findResource:628 - Header: %s\n", spew.Sdump(header))
	host := connectorPayload.System.ConnectorConfig.HostUrl
	//host := common.GetKVData(GetConfig().Data, "cacheHost")
	log.Debug3(" --  host: " + host)
	cacheBundleUrl := fmt.Sprintf("%s/%s/BundleTransaction", connectorConfig.CacheUrl, header.QueryId)
	//header.CacheUrl = fmt.Sprintf("%s%sv1/Cache/%s/", host, parts[0], header.QueryId)
	log.Debug3("  --  CacheUrl = " + cacheBundleUrl)
	header.CacheUrl = cacheBundleUrl
	//resp.Resource = header.ResourceType
	resp.BundleId = *bundle.Id
	resp.ResourceType = resourceType
	resp.Status = 200
	resp.QueryId = header.QueryId.Hex()
	resp.PageNumber = header.PageId
	resp.CountInPage = len(bundle.Entry)
	resp.TotalPages = totalPages
	resp.Header = header
	resp.Message = "Ok"
	logTime := time.Now()
	//log.Debug3(fmt.Sprintf("--  resp without bundle: " + spew.Sdump(resp)))
	log.Debug3(fmt.Sprintf("--  Time to log = %s", time.Since(logTime)))
	resp.Bundle = bundle
	log.Debug3(fmt.Sprintf("--  Number of entries in bundle: %d", len(bundle.Entry)))
	log.Debug3(fmt.Sprintf("--  QueryId: " + header.QueryId.Hex()))
	FillResourceResponse(&resp, resourceType)
	//fmt.Printf("findResource:614  --  Returning Bundle: %s\n", spew.Sdump(bundle))
	//WriteFhirResourceBundle(w, resp.Status, &resp)
	WriteFhirResponse(w, resp.Status, &resp)
}

type BasicResource struct {
	Id           string         `json:"id"`
	Text         fhir.Narrative `json:"text"`
	ResourceType string         `json:"resourceType"`
}

func getResource(w http.ResponseWriter, r *http.Request) {
	log.Info("getResource  Entered")
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
	log.Info(fmt.Sprintf("params = %s", params))
	log.Info(fmt.Sprintf("param resource : %s - %s  ResId: %s", Resource, resourceType, resourceId))
	//results := json.RawMessage{}
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
	log.Info("getResource calling GetConnectorPayload")
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
	log.Info("ResourceType: " + resourceType)
	log.Debug3(fmt.Sprintf("Retrieving %s Record for id: [%s]", resourceType, resourceId))
	if resourceId == "" {
		errMsg := log.ErrMsg(fmt.Sprintf("GetResource %s specific ID string is required", Resource))
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverity(fhir.IssueTypeInvalid), &errMsg))
		return
	}
	//TODO: Handle Get Resource by specific ID.  All Resources including Binary.
	resp := common.ResourceResponse{}
	log.Info("getResource calling GetResource")
	results, err := GetResource(cp, Resource, resourceId, JWToken)
	if err == nil {
		resp.Status = 200
		resp.Message = "Ok"
	} else {
		errMsg := log.ErrMsg("GetResource error: " + err.Error())
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverity(fhir.IssueTypeInvalid), &errMsg))
		return
	}
	log.Info("getResource determining the type of resource")
	basicResource := BasicResource{}
	err = json.Unmarshal(results, &basicResource)
	if err != nil {
		errMsg := log.ErrMsg("UnmarshalBasicResource error:  " + err.Error())
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityError, &errMsg))
		return
	}
	log.Info("Basic Resource: " + spew.Sdump(basicResource))
	resourceType = basicResource.ResourceType
	//TODO: unmarshal into a basic fhir resource (id, text)
	log.Debug3("FillResourceResponse for " + strings.ToLower(resourceType))
	switch strings.ToLower(resourceType) {
	case "patient":
		log.Info("Processing Patient")
		//fmt.Printf("GetResource:801  --  patient raw = %v\n", results)
		patient, err := fhir.UnmarshalPatient(results)
		if err != nil {
			errMsg := log.ErrMsg("UnmarshalPatient error:  " + err.Error())
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityError, &errMsg))
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
		//resp.Resource.ResourceType = resourceType
		//rawPat, err := json.Marshal(patient)
		//resp.Resource.Resource = rawPat
		resp.RawResource = results
		resp.ResourceId = *patient.Id
		resp.Message = "Ok"
		resp.PageNumber = 1
		resp.TotalPages = 1
		resp.CountInPage = 1
		resp.QueryId = primitive.NewObjectID().Hex()
		resp.Status = 200
		log.Debug5("Patient case final " + spew.Sdump(resp))
	case "binary":
		log.Info("Processing Binary")
		//fmt.Printf("GetResource:831  --  patient raw = %v\n", results)
		binary, err := fhir.UnmarshalBinary(results)
		if err != nil {
			errMsg := log.ErrMsg("UnmarshalBinary error: " + err.Error())
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
			return
		}
		//mt.Printf("GetResource:840  --  Binary: %s\n", spew.Sdump(binary))
		resp.ResourceType = resourceType
		//resp.Resource.ResourceType = resourceType
		//resp.Resource.Resource = results
		resp.RawResource = results
		resp.ResourceId = *binary.Id
		resp.Message = "Ok"
		resp.PageNumber = 1
		resp.TotalPages = 1
		resp.CountInPage = 1
		resp.QueryId = primitive.NewObjectID().Hex()
		resp.Status = 200
	case "documentreference":
		log.Info("Processing DocumentReferene")
		log.Debug5("Processing DocumentReference results: " + spew.Sdump(results))
		data, err := fhir.UnmarshalDocumentReference(results)
		if err != nil {
			errMsg := log.ErrMsg(fmt.Sprintf("Unmarshal %s error: %s", resourceType, err.Error()))
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
			return
		}
		log.Debug3(fmt.Sprintf("Resource %s   contains: %s", resourceType, spew.Sdump(data)))
		resp.ResourceType = resourceType
		//resp.Resource.ResourceType = resourceType

		//resp.Resource.Resource = results
		resp.RawResource = results
		//resp.Resource.ResourceId = *data.Id
		resp.ResourceId = *data.Id
		resp.Message = "Ok"
		resp.PageNumber = 1
		resp.TotalPages = 1
		resp.CountInPage = 1
		resp.QueryId = primitive.NewObjectID().Hex()
		resp.Status = 200
	default:
		log.Info("Processing default: " + resourceType)
		basicResource := BasicResource{}
		err = json.Unmarshal(results, &basicResource)
		if err != nil {
			errMsg := log.ErrMsg("UnmarshalBasicResource error:  " + err.Error())
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityError, &errMsg))
			return
		}
		log.Info("Basic Resource: " + spew.Sdump(basicResource))
		resp.ResourceType = resourceType
		//resp.Resource.ResourceType = resourceType

		//resp.Resource.Resource = results
		resp.RawResource = results
		//resp.Resource.ResourceId = basicResource.Id
		resp.ResourceId = basicResource.Id
		resp.Message = "Ok"
		resp.PageNumber = 1
		resp.TotalPages = 1
		resp.CountInPage = 1
		resp.QueryId = primitive.NewObjectID().Hex()
		resp.Status = 200
	}
	// resp.ResourceType = Resource
	// resp.Resource.Resource = results
	log.Debug5("returning resource: " + spew.Sdump(resp))
	WriteFhirResponse(w, resp.Status, &resp)
}
*/
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

// func DetermineResource(url string, prefix string) string {
// 	parts := strings.Split(url, "/")
// 	//parts := strings.SplitAfter(url, prefix)
// 	resourceType := parts[len(parts)-1]
// 	log.Debug3("--  resourceType: " + resourceType)
// 	return resourceType
// }
// func DetermineGetResource(url string, prefix string) string {
// 	log.Debug3(fmt.Sprintf("  --  url: " + url))
// 	parts := strings.Split(url, "/")
// 	//parts := strings.SplitAfter(url, prefix)
// 	resourceType := strings.TrimRight(parts[len(parts)-1], "/")
// 	// var subs []string
// 	log.Debug3("--  resourceType: " + resourceType)
// 	return resourceType
// }
// func FillResourceResponse(resp *common.ResourceResponse, resourceType string) error {
// 	resType := strings.ToLower(resourceType)
// 	log.Debug3("resType: " + resType)
// 	resp.Status = 200
// 	resp.Message = "Ok"
// 	resp.ResourceType = resourceType
// 	switch strings.ToLower(resourceType) {
// 	case "patient":
// 		//pats := []fhir.Patient{}
// 		//resData := common.ResourceData{}
// 		for i, item := range resp.Bundle.Entry {
// 			log.Debug3("Bundle.Entry: " + fmt.Sprint(i))
// 			resData := common.ResourceData{}

// 			pat, err := fhir.UnmarshalPatient(item.Resource)
// 			if err != nil {
// 				err = log.Errorf("UnMarshal(Patient) error = " + err.Error())
// 				fmt.Println(err.Error())
// 				return err
// 			}
// 			log.Debug3("Patient: " + spew.Sdump(pat))
// 			fmt.Printf("\n\n\n")
// 			log.Debug3("CurrentUser: " + spew.Sdump(CurrentUser))
// 			log.Info("HIPPALog access user: " + CurrentUserID.Hex() + " of patient: " + pat.ID.Hex())
// 			logMsg := fmt.Sprintf("HIPPA log User: %s - %s accessed  Patient: %s", CurrentUser.ID.Hex(), CurrentUser.UserName, pat.ID.Hex())
// 			log.Info(logMsg)
// 			//log.Info(fmt.Sprintf("HIPPA log User: %s - %s accessed  Patient: %s", CurrentUser.ID.Hex(), CurrentUser.UserName, pat.ID.Hex()))
// 			hl := &common.HippaLog{}
// 			hl.UserId = CurrentUser.ID.Hex()
// 			hl.PatientId = pat.ID.Hex()
// 			hl.ResourceType = "Patient"
// 			hl.ResourceId = hl.PatientId
// 			hl.SystemId = CurrentSystemId
// 			hl.LogType = "Core-Listed"
// 			hl.LogTime = time.Now().UTC()
// 			hl.LogMessage = logMsg
// 			err = LogHippa(hl)
// 			if err != nil {
// 				log.Warn("HippaLogging failed: " + err.Error())
// 			}
// 			resData.Patient = &pat
// 			log.Debug3(" Added PatientId: " + *pat.Id)
// 			resp.Resources = append(resp.Resources, resData)
// 			//pats = append(resData.Patient, pat)
// 		}
// 		//resp.Resources = pats
// 		log.Debug3(fmt.Sprintf("-- Set %d Patients  Bundle had %d entries\n", len(resp.Resources), len(resp.Bundle.Entry)))
// 	case "documentreference":
// 		for _, item := range resp.Bundle.Entry {
// 			resData := common.ResourceData{}
// 			docRef, err := fhir.UnmarshalDocumentReference(item.Resource)
// 			if err != nil {
// 				return log.Errorf("Unmarshal DocumentReference error = " + err.Error())
// 			}
// 			resData.DocumentReference = &docRef
// 			log.Debug3("Added DocumentReferenceId: " + *docRef.Id)
// 			resp.Resources = append(resp.Resources, resData)
// 		}
// 		log.Debug3(fmt.Sprintf("-- Set %d DocumentReferences  Bundle had %d entries\n", len(resp.Resources), len(resp.Bundle.Entry)))
// 	case "diagnosticreport":
// 		for _, item := range resp.Bundle.Entry {
// 			resData := common.ResourceData{}
// 			data, err := fhir.UnmarshalDiagnosticReport(item.Resource)
// 			if err != nil {
// 				return log.Errorf(" -- error = " + err.Error())
// 			}
// 			resData.DiagnosticReport = &data
// 			fmt.Printf("FillResourceResponse:1064  --  Added DiagnosticReporteId: %s\n", *data.Id)
// 			resp.Resources = append(resp.Resources, resData)
// 		}
// 		log.Debug3(fmt.Sprintf("-- Set %d DiagnosticReport  Bundle had %d entries\n", len(resp.Resources), len(resp.Bundle.Entry)))
// 	case "observation":
// 		for _, item := range resp.Bundle.Entry {
// 			resData := common.ResourceData{}
// 			data, err := fhir.UnmarshalObservation(item.Resource)
// 			if err != nil {
// 				return log.Errorf(" -- error = " + err.Error())
// 			}
// 			resData.Observation = &data
// 			log.Debug3(fmt.Sprintf(" --  Added DiagnosticReporteId: %s\n", *data.Id))
// 			resp.Resources = append(resp.Resources, resData)
// 		}
// 		log.Debug3(fmt.Sprintf("-- Set %d Observation  Bundle had %d entries\n", len(resp.Resources), len(resp.Bundle.Entry)))

// 	}
// 	//TODO: Make Switch smarter.
// 	return nil
// }

// package main

// import (
// 	"encoding/json"
// 	//"errors"
// 	"fmt"
// 	"io"

// 	"github.com/davecgh/go-spew/spew"
// 	cm "github.com/dhf0820/baseConnector/common"

// 	fhir "github.com/dhf0820/fhir4"
// 	jw_token "github.com/dhf0820/golangJWT"
// 	common "github.com/dhf0820/uc_core/common"

// 	//"github.com/gorilla/mux"

// 	//"github.com/gorilla/schema"
// 	"net/http"

// 	log "github.com/dhf0820/vslog"

// 	//"os"
// 	"reflect"
// 	//"strconv"
// 	//"strings"
// 	"time"
// )

// //####################################### Response Writers Functions #######################################

// //################################### FHIR Responses ####################################

// // ####################################### Route Handlers #######################################
// // getPatient - By patientId returning one single patient matching the ID
// // otherwise return OperationOutcome for NotFound
// var JWToken string

// func getPatient(w http.ResponseWriter, r *http.Request) {
// 	//Resource := "Patient"
// 	//buildFieldsByTagMap("schema", *psp)
// 	JWToken = r.Header.Get("Authorization")
// 	if JWToken == "" {
// 		err := fmt.Errorf("getPatient:41  --  Authorization is blank")
// 		errMsg := err.Error()
// 		WriteFhirOperationOutcome(w, 401, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	// Payload, status, err := token.ValidateToken(JWToken, "")
// 	// if err != nil {
// 	// 	errMsg := err.Error()
// 	// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 	// 	return
// 	// }
// 	// userId := Payload.UserId
// 	// fmt.Printf("getPatient:47  --  userId: %s\n", userId)
// 	// cp, err := GetConnectorPayload(r)
// 	// fmt.Printf("getPatient:48  --  ConnectorPayLoad = %s\n", spew.Sdump(cp))
// 	// //defer r.Body.Close()
// 	// params := mux.Vars(r)
// 	// fmt.Printf("getPatient:52  --  Params: %v\n", params)
// 	// id := params["id"]
// 	// //uri := r.URL.RequestURI()
// 	// fmt.Printf("getPatient:55  --  raw: %s\n", r.URL.RawQuery)
// 	// fmt.Printf("getPatient:56  --  query values: %v\n", r.URL.Query())
// 	// values := r.URL.Query()
// 	// for k, v := range values {
// 	// 	fmt.Println(k, " => ", v)
// 	// }
// 	// fmt.Printf("getPatient:61  --  id = %s\n", id)
// 	// patient, err := GetPatient(id)
// 	// if err != nil {
// 	// 	err = fmt.Errorf("getPatient:64  --  GetPatient error: %s\n", err.Error())
// 	// 	errMsg := err.Error()
// 	// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeIncomplete, fhir.IssueSeverityFatal, &errMsg))
// 	// 	return
// 	// }
// 	// fmt.Printf("\n\n\ngetPatient:69  --  Returning Patient: %s\n", spew.Sdump(patient))
// 	// resp := common.ResourceResponse{}

// 	// resp.Resource.Patient = patient
// 	// resp.Patient = *patient
// 	// resp.Status = 200
// 	// resp.Message = "Ok"
// 	// resp.ResourceType = Resource
// 	// WriteFhirResource(w, 200, &resp)

// 	// //defer resp.Body.Close()
// 	// //cfg = mod.ServiceConfig{}
// 	// fmt.Printf("Reading Body\n")
// 	// body, err = ioutil.ReadAll(r.Body)
// 	// if err != nil {
// 	// 	log.Printf("getPatient:53  --  ReadAllBody : error: %s\n", err.Error())
// 	// 	//err = errors.New("invalid FHIR URL")
// 	// 	errMsg := err.Error()
// 	// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 	// 	return
// 	// }
// 	// connectPayload := common.ConnectorPayload{}
// 	// fmt.Printf("raw string: %s\n", string(body))
// 	// fmt.Printf("GetPatient:61  --  Unmarshal ConnectorPayload\n")
// 	// err = json.Unmarshal(body, &connectPayload)
// 	// if err != nil {
// 	// 	log.Printf("getPatient:64  --  Unmarshal connectPayload error: %s\n", err.Error())
// 	// 	//err = errors.New("invalid FHIR URL")
// 	// 	errMsg := err.Error()
// 	// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 	// 	return
// 	// }
// 	// params := mux.Vars(r)
// 	// id := params["id"]
// 	// if id != "" {
// 	// 	fmt.Printf("GetPatient:73  -- Specific Query: %s\n", id)
// 	// 	//patient, err := GetPatient(id)
// 	// }

// 	// //fmt.Printf("getPatient:70 -- connectPayload: %s\n", spew.Sdump(connectPayload))
// 	// // fhirId := GetFhirId(r)
// 	// // fhirSystem, err := GetFhirSystem(fhirId)
// 	// // if err != nil {
// 	// // 	log.Printf("searchPatient:50  --  FhirId : [%s] error: %s\n", fhirId, err.Error())
// 	// // 	err = errors.New("invalid FHIR URL")
// 	// // 	errMsg := err.Error()
// 	// // 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 	// // 	return
// 	// // }

// 	// //  Separate connector for each emr Vendor.  CernerConnector, EpicConnector, CAConnector,...
// 	// // // handles query andsaveeither via fhir or direct API (AllScripts, Athena)

// 	// // Payload, status, err := token.ValidateToken(r.Header.Get("Authorization"), "")
// 	// // if err != nil {
// 	// // 	errMsg := err.Error()
// 	// // 	fmt.Printf("getPatient:55  --  Err: %s\n", errMsg)
// 	// // 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 	// // 	return
// 	// // }
// 	// // userId := Payload.UserId
// 	// // log.Printf("getPatient:59  --  UserId: %s\n", userId)
// 	// // fhirId := GetFhirId(r)
// 	// // fhirSystem, err := GetFhirSystem(fhirId)
// 	// // if err != nil {
// 	// // 	log.Printf("getPatient:63  --  FhirId : [%s] error: %s\n", fhirId, err.Error())
// 	// // 	err = errors.New("invalid FHIR URL")
// 	// // 	errMsg := err.Error()
// 	// // 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 	// // 	return
// 	// // }
// 	// fmt.Printf("getPatient:109  --  Request: [%s]\n", r.RequestURI)
// 	// urlA, err := r.URL.Parse(r.RequestURI)
// 	// if err != nil {
// 	// 	err = fmt.Errorf("error parsing patient URI: %s", err.Error())
// 	// 	errMsg := err.Error()
// 	// 	fmt.Printf("getPatient:114 - r.URL.Parse error = %s\n", errMsg)
// 	// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 	// 	return
// 	// }
// 	// fmt.Printf("getPatient:118 - r.URL.Parse = %v\n", urlA)
// 	// urlB := *urlA
// 	// uriValues := urlB.Query()
// 	// fmt.Printf("getPatient:121 - uriValues= %v\n", uriValues)

// 	// uri := r.RequestURI
// 	// log.Printf("uri = %s\n", uri)
// 	// parts := strings.Split(uri, Resource)
// 	// uri = parts[1]
// 	// log.Printf("getPatient:127 - URI = %s\n", uri)
// 	// //patient := fhir.Patient{}
// 	// resource, err := GetResource(connectPayload.FhirSystem, Resource, uri)
// 	// resp := common.ResourceResponse{}
// 	// if err != nil {
// 	// 	resp.Status = 400
// 	// 	resp.Message = err.Error()
// 	// } else {
// 	// 	resp.Status = 200
// 	// 	resp.Message = "Ok"
// 	// }
// 	// //var patient fhir.Patient
// 	// //patient := resource.(fhir.Patient)
// 	// resp.Resource.Resource = resource
// 	// // var res []interface{}
// 	// // res = append(res, &resource)
// 	// // resp.Resources = res
// 	// resp.ResourceType = Resource
// 	// //resp.ResourceId = *patient.Id
// 	// //log.Printf("\nGetPatient:139  --  resp: %s\n", spew.Sdump(resp))

// 	// WriteFhirResource(w, resp.Status, &resp)
// }

// // postPatient: Stores the fhir patient payload in the url {Fhir-System} specified fhirSystem.
// /* func savePatient(w http.ResponseWriter, r *http.Request) {
// 	//Resource := "Patient"
// 	//fmt.Printf("postPatient:182 - Post: %s \n", spew.Sdump(r))
// 	JWToken := r.Header.Get("Authorization")
// 	payload, status, err := jw_token.ValidateToken(JWToken, "")
// 	if err != nil {
// 		errMsg := err.Error()
// 		fmt.Printf("savePatient:190  - ValidateToken err = %s\n", errMsg)
// 		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	log.Info("payload: " + spew.Sdump(payload))
// 	body, err := io.ReadAll(r.Body) // Should be ConnectorPayload
// 	if err != nil {
// 		fmt.Printf("savePatient:197  --  ReadAll FhirSystem error %s\n", err.Error())
// 		errMsg := err.Error()
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	fmt.Printf("savePatient:202  --  ConnectorPayload: %s\n", spew.Sdump())
// 	//b := string(body)
// 	//fmt.Printf("SavePatient:204  Body: %s\n", b)
// 	conPayload := &common.ConnectorPayload{}
// 	err = json.Unmarshal(body, &conPayload)
// 	if err != nil {
// 		fmt.Printf("\nsavePatient:208  --  unmarshal err = %s\n", err.Error())
// 		errMsg := err.Error()
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	log.Debug3("--  Check ConPayload")
// 	if conPayload == nil {
// 		fmt.Printf("\n\nconPayload is nil\n")
// 	}
// 	fmt.Printf("savePatient:217  --  Check SavePayload\n")
// 	if conPayload.SavePayload == nil {
// 		fmt.Printf("\n\nconPayload.SavePayload is nil\n")
// 	}
// 	log.Debug3("--  conPayload: " + spew.Sdump(conPayload))
// 	//
// 	//fmt.Printf("SavePatient:223  --  Check ConPayload srcPatient: %s\n", spew.Sdump(conPayload))
// 	// if conPayload.SavePayload.SrcPatient == nil {
// 	// 	res, err := GetCachedResource(conPayload, JWToken)
// 	// 	if res != nil {
// 	// 		errMsg := err.Error()
// 	// 		fmt.Printf("savePatient:228 - r.URL.Parse error = %s\n", errMsg)
// 	// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 	// 		return
// 	// 	}
// 	// 	conPayload.SavePayload.SrcPatient = res.Patient
// 	// }

// 	fmt.Printf("savePatient:235  --  calling SavePatient with Payload: %s\n", spew.Sdump(conPayload))
// 	httpResp, err := SavePatient("", conPayload, JWToken)
// 	if err != nil {
// 		fmt.Printf("savePatient:238  --  SavePatient error %s\n", err.Error())
// 	}
// 	if httpResp.StatusCode != 201 {
// 		fmt.Printf("savePatient:241  --  httpResp.Status = %s   StatusCode: %d\n", httpResp.Status, httpResp.StatusCode)
// 		fmt.Printf("savePatient:242  --  error: %s\n", httpResp.Status)
// 		defer httpResp.Body.Close()
// 		bodyBytes, err := io.ReadAll(httpResp.Body)
// 		if err != nil {
// 			errMsg := err.Error()
// 			fmt.Printf("savePatient:247  --  ReadAll SavePatient ReadBody error %s\n", err.Error())
// 			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 			return
// 		}
// 		opOutcome := fhir.OperationOutcome{}
// 		err = json.Unmarshal(bodyBytes, &opOutcome)
// 		if err != nil {
// 			errMsg := log.ErrMsg("ErrorMessage ReadBody error " + err.Error())
// 			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 			return
// 		}
// 		WriteFhirOperationOutcome(w, httpResp.StatusCode, &opOutcome)
// 		// issueType := fhir.IssueTypeException
// 		// if httpResp.StatusCode == 409 {
// 		// 	issueType = fhir.IssueTypeDuplicate
// 		// }

// 		// WriteFhirOperationOutcome(w, httpResp.StatusCode, CreateOperationOutcome(issueType, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	defer httpResp.Body.Close()
// 	byte, err := io.ReadAll(httpResp.Body)
// 	if err != nil {
// 		errMsg := err.Error()
// 		fmt.Printf("savePatient:272  --  ReadAll SavePatient ReadBody error %s\n", err.Error())
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	fmt.Printf("savePatient:276  --  Check if patient\n")

// 	// saveResp := common.SaveResponse{}
// 	// err = json.Unmarshal(byte, &saveResp)
// 	saveResp := common.SaveResponse{}
// 	err = json.Unmarshal(byte, &saveResp)
// 	if err != nil {
// 		fmt.Printf("savePatient: 283  --  Response is not a Patient err: %s\n", err.Error())
// 		HandleOperationOutcome(w, byte)
// 		return
// 	}
// 	fmt.Printf("savePatient:287 returned saveResp: %s\n", spew.Sdump(saveResp))
// 	WriteSaveResponse(w, 201, &saveResp)

// 	// if patient.ResourceType == nil {
// 	// 	fmt.Printf("savePatient: 269  --  Response is not a Patient\n")
// 	// 	HandleOperationOutcome(w, byte)
// 	// 	return
// 	// }
// 	//resp := &common.ResourceResponse{}

// 	//resp.Patient = patient
// 	// resp := &common.SaveResponse{}
// 	// resp.Id = *patient.Id
// 	// resp.Text = patient.Text.Div
// 	// resp.Mrn = GetMrn(&patient, "urn:oid:1.3.6.1.4.1.54392.5.1593.1", "mrn")
// 	//fmt.Printf("\n\n\nsavePatient:298  --  Returning %s\n", spew.Sdump(resp))
// 	//location := Conf.Server.Host + ":" + Conf.Server.Port + "/system/640ba66cbd4105586a6dda75/Patient/" + patient.Id
// 	// baseUrl := Conf.BaseURL
// 	// location := baseUrl + "/system/640ba66cbd4105586a6dda75/Patient/" // + *patient.Id
// 	// fmt.Printf("savePatient:274  --  Location: %s\n", location)
// 	//WriteFhirPatient(w, 200, &patient)
// 	// opOutcom := fhir.OperationOutcome{}
// 	// err = json.Unmarshal(byte, &opOutcom)
// 	// if err == nil {
// 	// 	fmt.Printf("SavePatient:253  --  opOutcom: %s\n", spew.Sdump(opOutcom))
// 	// 	WriteHttpResponse(w, httpResp.StatusCode, httpResp)
// 	// 	//WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 	// 	return
// 	// }
// 	// patient := fhir.Patient{}
// 	// err = json.Unmarshal(byte, &patient)
// 	// if err != nil {
// 	// 	fmt.Printf("savePatient: 258  --  Response is not a Patient\n")
// 	// 	opOutcom := fhir.OperationOutcome{}
// 	// 	err = json.Unmarshal(byte, &opOutcom)
// 	// 	if err == nil {
// 	// 		fmt.Printf("SavePatient:262  --  opOutcom: %s\n", spew.Sdump(opOutcom))
// 	// 		WriteHttpResponse(w, httpResp.StatusCode, httpResp)
// 	// 		//WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 	// 		return
// 	// 	}
// 	// 	fmt.Printf("savePatient:267  --   Unmarshal Patient Body error %s\n", err.Error())
// 	// 	return
// 	// } else {
// 	// 	fmt.Printf("savePatient:270  --  Patient Unmarshal error: %s\n", spew.Sdump(patient))
// 	// }
// 	// if err != nil {
// 	// 	errMsg := err.Error()
// 	// 	fmt.Printf("savePatient:249  --  ReadAll SavePatient Unmarshal Body error %s\n", err.Error())
// 	// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 	// 	return
// 	// }

// 	// if err != nil {
// 	// 	errMsg := err.Error()
// 	// 	fmt.Printf("SavePatient:239 - SavePatient error = %s\n", errMsg)
// 	// 	values := strings.Split(errMsg, "|")
// 	// 	//errNum, _ := strconv.Atoi(values[0])
// 	// 	//fmt.Printf("SavePatient:240 - values = %s\n", spew.Sdump(values))
// 	// 	fmt.Printf("SavePatient:241 - values[0] = %s\n", values[0])
// 	// 	if values[0] == "409" {
// 	// 		WriteFhirOperationOutcome(w, 409, CreateOperationOutcome(409, fhir.IssueSeverityFatal, &errMsg))
// 	// 	} else {
// 	// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 	// 	}
// 	// 	return
// 	// }
// 	// fmt.Printf("savePatient:293 returned PatientId: %s\n", *patient.Id)
// 	// resp := &common.SaveResponse{}
// 	// resp.Id = *patient.Id
// 	// resp.Text = patient.Text.Div
// 	// resp.Mrn = GetMrn(&patient, "urn:oid:1.3.6.1.4.1.54392.5.1593.1", "mrn")
// 	// fmt.Printf("\n\n\nsavePatient:298  --  Returning %s\n", spew.Sdump(resp))
// 	// WriteSaveResponse(w, 200, resp)

// 	// fhirId := GetFhirId(r)                   // Get the Fhir-System ID portion of the URL
// 	// fhirSystem, err := GetFhirSystem(fhirId) // Get the actual FhirSystem Configuration
// 	// if err != nil {
// 	// 	log.Printf("postPatient:162  --  FhirId : [%s] error: %s\n", fhirId, err.Error())
// 	// 	err = errors.New("url contains Invalid FHIR identifier: " + fhirId)
// 	// 	errMsg := err.Error()
// 	// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 	// 	return
// 	// }
// 	// urlA, err := r.URL.Parse(r.RequestURI)
// 	// if err != nil {
// 	// 	err = fmt.Errorf("error parsing patient URI: [%s]  error:%s", r.RequestURI, err.Error())
// 	// 	errMsg := err.Error()
// 	// 	fmt.Printf("postPatient:172 - r.URL.Parse error = %s\n", errMsg)
// 	// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 	// 	return
// 	// }
// 	// fmt.Printf("postPatient:176 - r.URL.Parse = %v\n", urlA)
// 	// urlB := *urlA
// 	// uriValues := urlB.Query()
// 	// fmt.Printf("postPatient:179 - uriValues= %v\n", uriValues)

// 	// uri := r.RequestURI
// 	// log.Printf("uri = %s\n", uri)
// 	// parts := strings.Split(uri, Resource)
// 	// uri = parts[1]
// 	// log.Printf("postPatient:185 - URI = %s\n", uri)
// 	// //patient := fhir.Patient{}
// 	//WriteFhirResource(w, 200, resp)
// 	// errMsg := "SavePatient to CA not implemented"
// 	// WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 	// return
// 	// resource, err := GetResource(fhirSystem, Resource, uri)
// 	// resp := common.ResourceResponse{}
// 	// if err != nil {
// 	// 	resp.Status = 400
// 	// 	resp.Message = err.Error()
// 	// } else {
// 	// 	resp.Status = 200
// 	// 	resp.Message = "Ok"
// 	// }
// 	// var patient fhir.Patient
// 	// //patient = resource.(fhir.Patient)
// 	// resp.Resource.Resource = resource
// 	// // var res []interface{}
// 	// // res = append(res, &resource)
// 	// // resp.Resources = res
// 	// resp.ResourceType = Resource
// 	// resp.ResourceId = *patient.Id
// 	// log.Printf("\nGetPatient:204  --  resp: %s\n", spew.Sdump(resp))
// 	// WriteFhirResourceBundle(w, resp.Status, &resp)
// } */

// // searchPatient uses the systemId url parameter to determin the FhirSystem to use
// func searchPatient(w http.ResponseWriter, r *http.Request) {
// 	srchParams, err := cm.PatientSearchparams(r) //FhirPatientSearch(r)
// 	if err != nil {
// 		fmt.Printf("searchPatient:397  -- Error %s \n", err.Error())
// 		errMsg := err.Error()
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	fmt.Printf("searchPatient:402  --  srchParams: %s\n", spew.Sdump(srchParams))

// 	// var pspTags map[string]string
// 	// tagFields := make(map[string]string)
// 	// var Limit int
// 	// var Skip int
// 	//Resource := "Patient"
// 	body, err := io.ReadAll(r.Body) // Should be ConnectorPayload
// 	if err != nil {
// 		errMsg := log.ErrMsg("ReadAll FhirSystem error: " + err.Error())
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	connectorPayload := common.ConnectorPayload{}
// 	//fhirSystem := common.FhirSystem{}
// 	err = json.Unmarshal(body, &connectorPayload)
// 	if err != nil {
// 		fmt.Printf("\nsearchPatient:420  --  unmarshal err = %s\n", err.Error())
// 		errMsg := log.ErrMsg("Unmarshal ConnectorPayload error: " + err.Error())
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	//fhirSystem := connectorPayload.FhirSystem
// 	//connConfig := connectorPayload.System.ConnectorConfig
// 	//buildFieldsByTagMap("schema", *psp)
// 	JWToken = r.Header.Get("Authorization")
// 	log.Debug3("searchPatient - JWToken: " + JWToken)
// 	Payload, status, err := jw_token.ValidateToken(r.Header.Get("Authorization"), "")
// 	if err != nil {
// 		errMsg := err.Error()
// 		fmt.Printf("searchPatient:433  --  ValidateToken err: %s\n", errMsg)
// 		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	//fhirId := fhirSystem.ID.String()
// 	userId := Payload.UserId
// 	log.Info("UserId: " + userId)

// 	log.Info("raw: " + r.URL.RawQuery)
// 	//fmt.Printf("query values: %v\n", r.URL.Query())
// 	values := r.URL.Query()
// 	for k, v := range values {
// 		fmt.Println(k, " => ", v)
// 	}

// 	/*
// 		// fhirId := GetFhirId(r)
// 		// fhirSystem, err := GetFhirSystem(fhirId)
// 		// if err != nil {
// 		// 	log.Printf("searchPatient:232  --  FhirId : [%s] error: %s\n", fhirId, err.Error())
// 		// 	err = errors.New("invalid FHIR URL")
// 		// 	errMsg := err.Error()
// 		// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 		// 	return
// 		// }
// 		uri := r.RequestURI
// 		log.Printf("searchPatient:314  --  r.RequestURI = %s\n", uri)
// 		parts := strings.Split(uri, Resource)
// 		uri = parts[1]
// 		log.Printf("\nsearchPatient:260 - URI = %s\n", uri)

// 		urlA, err := r.URL.Parse(r.RequestURI)
// 		if err != nil {
// 			err = fmt.Errorf("error parsing patient URI: %s", err.Error())
// 			errMsg := err.Error()
// 			fmt.Printf("searchPatient:266 - r.URL.Parse error = %s\n", errMsg)
// 			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 			return
// 		}
// 		fmt.Printf("searchPatient:270 - r.URL.Parse = %v\n", urlA)
// 		urlB := *urlA
// 		uriValues := urlB.Query()
// 		fmt.Printf("searchPatient:273 - uriValues= %v\n", uriValues)
// 		//ident := uriValues.Get("identifier")
// 		// if ident != "" { // There is identifier Search, use it
// 		// 	fmt.Printf("searchPatient:102 - using Identifier: %s to search\n", ident)
// 		// } else {
// 		// 	fmt.Printf("searchPatient:104 - using other search params: %v\n", uriValues)
// 		// }

// 		// //}
// 		// //fhirVersion := GetFHIRVersion(r)
// 		// //CacheUrlURL := fmt.Sprintf("%s/%s/v1/", r.Host, parts[0])
// 		// if err := r.ParseForm(); err != nil {
// 		// 	err = fmt.Errorf("error parsing query: %s", err.Error())
// 		// 	errMsg := err.Error()
// 		// 	fmt.Printf("searchPatient:113 - %s\n", errMsg)
// 		// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 		// 	return
// 		// }
// 		// FhirId := GetFhirId(r)
// 		// fmt.Printf("searchPatient:79 - FhirKey - [%s]\n", FhirId)
// 		// fhirSystem, err := GetFhirSystem(FhirId)
// 		// if err != nil {
// 		// 	fmt.Printf("GetFhirSystem failed with : %s\n", err.Error())
// 		// 	err = fmt.Errorf("fhirSystem error:  %s", err.Error())
// 		// 	errMsg := err.Error()
// 		// 	fmt.Printf("searchPatient:86 - %s\n", errMsg)
// 		// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityFatal, &errMsg))
// 		// 	return
// 		// }
// 		//fmt.Printf("searchPatient:90 -  %s/n", spew.Sdump(fhirSystem))

// 		// if Resource == "Patient" {
// 		log.Printf("\n\nsearchPatient:305  --  Resource Is Patient\n\n")
// 		//urlA, err := r.URL.Parse(r.RequestURI)
// 		if err != nil {
// 			err = fmt.Errorf("error parsing patient URI: %s", err.Error())
// 			errMsg := err.Error()
// 			fmt.Printf("searchPatient:310 - r.URL.Parse error = %s\n", errMsg)
// 			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 			return
// 		}
// 		fmt.Printf("searchPatient:314 - r.URL.Parse = %v\n", urlA)
// 		//urlB := *urlA
// 		//uriValues := urlB.Query()
// 		fmt.Printf("searchPatient:317 - uriValues= %v\n", uriValues)
// 		idSearch := uriValues.Get("identifier")
// 		idValue := ""
// 		if idSearch != "" { // There is identifier Search, use it
// 			fmt.Printf("searchPatient:321- using Identifier: %s to search\n", idSearch)
// 			ids := strings.Split(idSearch, "|")
// 			if len(ids) != 2 {
// 				err = fmt.Errorf("invalid identifier: %s", idSearch)
// 				errMsg := err.Error()
// 				fmt.Printf("searchPatient:326 - r.URL.Parse error = %s\n", errMsg)
// 				WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 				return
// 			}
// 			idName := ids[0]
// 			idSearchValue := ids[1]
// 			idents := fhirSystem.Identifiers
// 			for _, id := range idents {
// 				fmt.Printf("searchPatient:334  --  Looking at %s = %s\n", id.Name, idName)
// 				if id.Name == idName {
// 					idValue = id.Value
// 					break
// 				}
// 			}
// 			//
// 			if idValue == "" { //Not configured identifier
// 				err = fmt.Errorf("identifier type: %s is not configured", idName)
// 				errMsg := err.Error()
// 				fmt.Printf("searchPatient:344 - Identifiers = %s\n", errMsg)
// 				WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 				return
// 			}
// 			uri = fmt.Sprintf("?identifier=%s", idValue+idSearchValue)
// 			fmt.Printf("searchPatient:349 - New Identifier search Value: %s\n", uri)
// 		} else {
// 			fmt.Printf("searchPatient:351 - using other search params: %v\n", uriValues)
// 		}
// 		var bundle *fhir.Bundle
// 		var header *common.CacheHeader
// 		fmt.Printf("\nsearchPatient:355 - resource = %s  uri = %s\n", Resource, uri)
// 		url := fmt.Sprintf("%s/%s%s", fhirSystem.FhirUrl, Resource, uri) //" + "/" + uri
// 		fmt.Printf("searchPatient:357 - calling %s \n", url)
// 		var totalPages int64
// 		fmt.Printf("searchPatient:359 Search %s\n", url)
// 		uri = "/" + Resource + uri
// 		totalPages, bundle, header, err = FindResource(&connectorPayload, Resource, userId, uri, JWToken)
// 		if err != nil {
// 			err = fmt.Errorf("searchPatient:363 --  fhirSearch url: %s error:  %s", url, err.Error())
// 			errMsg := err.Error()
// 			fmt.Printf("searchPatient:365 - %s\n", errMsg)
// 			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityInformation, &errMsg))
// 			return
// 		}
// 		if bundle == nil {
// 			log.Printf("searchPatient:370  --  bundle is nil")
// 		} else {
// 			log.Printf("searchPatient:372  --  bundle is not nil \n")
// 		}
// 		fmt.Printf("searchPatient:374 - Get %s bundle successful\n", Resource)
// 		fmt.Printf("searchPatient:375 - Number in page: %d\n", len(bundle.Entry))
// 		fmt.Printf("searchPatient:376 - PageNumber: %d\n", header.PageId)
// 		resp := common.ResourceResponse{}
// 		header.CacheUrl = fmt.Sprintf("%s/%s/BundleTransaction", connConfig.CacheUrl, header.FhirSystem.ID.Hex())
// 		log.Printf("\n\nsearchPatient:379  --  CacheUrl = %s\n", header.CacheUrl)
// 		header.FhirId = fhirId
// 		header.UserId = userId
// 		resp.Bundle = bundle
// 		resp.Resource.Resource = bundle.Entry[0].Resource
// 		resp.BundleId = *bundle.Id
// 		resp.ResourceType = Resource
// 		resp.Status = 200
// 		resp.QueryId = header.QueryId
// 		resp.PageNumber = header.PageId
// 		if bundle.Entry == nil {
// 			err = fmt.Errorf("searchPatient:390 --  fhirSearch url: %s error:  %s", url, "Bundle.Entry is nil")
// 			errMsg := err.Error()
// 			fmt.Printf("searchPatient:392 - %s\n", errMsg)
// 			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityInformation, &errMsg))
// 			return
// 		}
// 		resp.CountInPage = len(bundle.Entry)
// 		resp.TotalPages = totalPages
// 		resp.Header = header
// 		resp.Message = "Ok"
// 		//fmt.Printf("searchPatient:400 - returning a resource bundle: %s\n", spew.Sdump(resp))
// 		WriteFhirResourceBundle(w, resp.Status, &resp)
// 		//WriteFhirBundle(w, resp.Status, bundle)
// 	*/
// }

// //func searchPatient(w http.ResponseWriter, r *http.Request) {
// // 	// var pspTags map[string]string
// // 	// tagFields := make(map[string]string)
// // 	// var Limit int
// // 	// var Skip int
// // 	fmt.Printf("Request: %s \n", spew.Sdump(r))
// // 	//buildFieldsByTagMap("schema", *psp)
// // 	//facility = "demo"
// // 	resource := GetFHIRResource(r)
// // 	fmt.Printf("search%s called with %s\n", resource, r.URL.RawQuery)
// // 	if err := r.ParseForm(); err != nil {
// // 		err = fmt.Errorf("error parsing query: %s", err.Error())
// // 		errMsg := err.Error()
// // 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// // 		return
// // 	}
// // 	params := mux.Vars(r)
// // 	fmt.Printf("params: %v\n", params)
// // 	resourceId := params["id"]
// // 	fmt.Printf("Retrieving Patient Record for id: %s\n", resourceId)
// // 	// psp := new(PatientSearchParams)
// // 	// psp.Limit = Limit
// // 	// psp.Skip = Skip
// // 	// psp.CurrentFacility = GetDeploymentFacility(r)
// // 	// psp.BaseUrl = GetCurrentURL(r)
// // 	//FhirVersion := GetFHIRVersion(r)
// // 	FhirId := GetFhirId(r)
// // 	_, err := GetFhirConnector(FhirId)
// // 	if err != nil {
// // 		err = fmt.Errorf("fhirConnecor error:  %s", err.Error())
// // 		errMsg := err.Error()
// // 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// // 	}

// // 	//bundle, err := FindResource(fhirConnector, resource, r.URL.RawQuery)

// // }

// // func WriteFhirOperationOutcome(w http.ResponseWriter, status int, resp *fhir.OperationOutcome) error {
// // 	w.Header().Set("Content-Type", "application/json")

// // 	switch status {
// // 	case 200:
// // 		w.WriteHeader(http.StatusOK)
// // 	case 400:
// // 		w.WriteHeader(http.StatusBadRequest)
// // 	case 401:
// // 		w.WriteHeader(http.StatusUnauthorized)
// // 	case 403:
// // 		w.WriteHeader(http.StatusForbidden)
// // 	}
// // 	err := json.NewEncoder(w).Encode(resp)
// // 	if err != nil {
// // 		fmt.Println("Error marshaling JSON:", err)
// // 		return err
// // 	}
// // 	return nil
// // }

// func WriteFhirPatient(w http.ResponseWriter, status int, resp *fhir.Patient) error {
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
// 	}
// 	err := json.NewEncoder(w).Encode(resp)
// 	if err != nil {
// 		fmt.Println("WriteFhirPatient:682  --  Error marshaling JSON:", err)
// 		return err
// 	}
// 	return nil
// }

// type PatientFilter struct {
// 	Id         string    `json:"id" schema:"_id"`
// 	MRN        string    `json:"mrn" schema:"mrn"`
// 	SSN        string    `json:"ssn" schema:"ssn"`
// 	Identifier string    `json:"identifier" schema:"identifier"`
// 	Gender     string    `json:"gender" schema:"gender"`
// 	BirthDate  string    `json:"birthdate" schema:"birthdate"`
// 	Name       string    `json:"name" schema:"name"`
// 	Family     string    `json:"family" schema:"family"`
// 	Given      string    `json:"given" schema:"given"`
// 	Phone      string    `json:"phone" schema:"phone"`
// 	Email      string    `json:"email" schema:"email"`
// 	PostalCode string    `json:"address-postalcode" schema:"address-postalcode"`
// 	Active     string    `json:"active" schema:"active"`
// 	DOB        time.Time `json:"dob" schema:"dob"`
// 	BaseUrl    string    `json:"base_url"`
// 	RequestURI string    `json:"request_uri"`
// 	Limit      uint32    `json:"limit"`
// 	Skip       uint32    `json:"skip" schema:"skip"`
// 	Count      uint32    `json:"count" schema:"_count"`
// 	OffSet     uint32    `json:"offset" schema:"_offset"`
// 	// Order      SearchParam `json:"order" schema:"_order"`
// 	// Sort       SearchParam `json:"sort" schema:"_sort"`
// 	// Page       SearchParam `json:"page" schema:"_page"`
// }

// type SearchParam struct {
// 	Schema   string
// 	Modifier string
// 	Value    string
// }

// func findPatient(w http.ResponseWriter, r *http.Request) {
// 	JWToken := r.Header.Get("Authorization")
// 	Payload, status, err := jw_token.ValidateToken(JWToken, "")
// 	if err != nil {
// 		errMsg := err.Error()
// 		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	userId := Payload.UserId
// 	log.Info(" UserId: " + userId)
// 	r.ParseForm()
// 	decoder.RegisterConverter(time.Now(), func(value string) reflect.Value {
// 		result := reflect.Value{}
// 		if t, err := time.Parse("2006-01-02", value); err == nil {
// 			result = reflect.ValueOf(t)
// 		}
// 		return result
// 	})
// 	cp, err := GetConnectorPayload(r)
// 	if err != nil {
// 		errMsg := log.ErrMsg("GetConnectorPayload error: " + err.Error())
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeForbidden, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	//fmt.Printf("findPatient:645  --  ConnectorPayload: %s\n", spew.Sdump(cp))
// 	//fmt.Printf("findPatient:646  --  r.Form = %v\n", r.Form)
// 	//fmt.Printf("\nfindPatient:647  --  qry = %v   len of URL.Query = %d\n\n", r.URL.Query(), len(r.URL.Query()))
// 	patFilter := new(PatientFilter)
// 	u := r.URL.Query()
// 	fmt.Printf("findPatient:779  --  u: %v\n", u)
// 	// qry, err := CreateFhirQuery(r)
// 	// if err != nil {
// 	// 	if err != nil {
// 	// 		status := 400
// 	// 		errMsg := fmt.Sprintf("findPatient:655  --  CreateFhirQuery err: %s\n", err.Error())
// 	// 		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 	// 		return
// 	// 	}
// 	// }
// 	err = decoder.Decode(patFilter, r.Form)
// 	if err != nil {
// 		status := 400
// 		errMsg := log.ErrMsg(" decode err: " + err.Error())
// 		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	fmt.Printf("findPatient:796  --  newSearchParams: %s\n", spew.Sdump(patFilter))
// 	fmt.Printf("findPatient:797  --  raw: %v\n", r.URL.RawQuery)
// 	fmt.Printf("findPatient:798  --  u: %v\n", u)
// 	//fmt.Printf("findPatient:673  --  SearchParams: %s\n", spew.Sdump(newSearchParams))
// 	fmt.Printf("findPatient:800  --  r.Form: %v\n\n\n\n", r.Form)
// 	if patFilter.Id != "" {
// 		pat, err := patFilter.FindById()
// 		if err != nil {
// 			status := 400
// 			errMsg := fmt.Sprintf("findPatient:805  --  decode err: %s\n", err.Error())
// 			WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 			return
// 		}
// 		resp := common.ResourceResponse{}
// 		resp.Patient = pat
// 		WriteFhirResource(w, 200, &resp)
// 		return
// 	}
// 	fmt.Printf("findPatient:814  --  Calling patFilter.Find\n")
// 	qry := r.URL.RawQuery
// 	bundle, cacheHeader, err := PatientSearch(cp, qry, JWToken)
// 	if err != nil {
// 		fmt.Printf("findPatient:818  -- PatientSearch returned err: %s\n", err.Error())
// 		status := 400
// 		errMsg := fmt.Sprintf("findPatient:820  --  decode err: %s\n", err.Error())
// 		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	log.Debug3("Cacheheader: " + spew.Sdump(cacheHeader))
// 	resp := common.ResourceResponse{}
// 	resp.Bundle = bundle
// 	FillResourceResponse(&resp, "patient") // Fills the response resource fields
// 	//fmt.Printf("\n\nfindPatient:830  --  patients: %s\n", spew.Sdump(resp.Patients))
// 	fmt.Printf("\n\nfindPatient:828  --  resp: %s\n", spew.Sdump(resp))
// 	WriteFhirResponse(w, 200, &resp)
// }

// // func GetConnectorPayload(r *http.Request) (*common.ConnectorPayload, error) {
// // 	body, err := io.ReadAll(r.Body) // Should be ConnectorPayload
// // 	if err != nil {
// // 		return nil, log.Errorf("ReadAll FhirSystem error " + err.Error())
// // 	}
// // 	//mt.Printf("GetConnectorPayload:717  -- Got Body Now Unmarshal ConnectorPayload\n")
// // 	b := string(body)
// // 	fmt.Printf("GetConnectorPayload:845  Body: %s\n", b)
// // 	conPayload := &common.ConnectorPayload{}
// // 	err = json.Unmarshal(body, &conPayload)
// // 	if err != nil {
// // 		fmt.Printf("\nGetConnectorPayload:849  --  unmarshal err = %s\n", err.Error())
// // 		// errMsg := err.Error()
// // 		// WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// // 		return nil, err
// // 	}
// // 	log.Info("Check ConPayload")
// // 	if conPayload == nil {
// // 		return nil, log.Errorf("conPayload is nil ")
// // 	}
// // 	//fmt.Printf("GetConnectorPayload:860  --  ConnectorPayLoad = %s\n", spew.Sdump(conPayload))
// // 	return conPayload, err
// // }

// // func CreateFhirQuery(r *http.Request) (string, error) {
// // 	query := ""
// // 	values := r.URL.Query()
// // 	log.Info(fmt.Sprintf("CreateFhirQuery  values : %v", values))
// // 	if len(values) < 1 {
// // 		err := log.Errorf("Url.Querys are missing")
// // 		return "", err
// // 	}
// // 	//fmt.Printf("\nCreateFhirQuery:713  --  Keys : %v\n\n", keys)
// // 	for k, v := range values {
// // 		log.Info("Key:  " + k + " => " + v[0])
// // 		s := strings.TrimLeft(v[0], "[]")
// // 		if query == "" {
// // 			//for _, kv := range v {
// // 			query = fmt.Sprintf("%s=%s", k, s)
// // 			//}
// // 		} else {
// // 			query = fmt.Sprintf("%s&%s=%s", query, k, s)
// // 		}
// // 		log.Info("CreateFhirQuery = " + query)
// // 	}
// // 	return query, nil
// // }

// func HandleOperationOutcome(w http.ResponseWriter, body []byte) {
// 	fmt.Printf("HandleOperationOutcome:892  --  body = %s\n", string(body))
// 	opOutcome := &fhir.OperationOutcome{}
// 	err := json.Unmarshal(body, &opOutcome)
// 	if err != nil {
// 		errMsg := log.ErrMsg(" Error: " + err.Error())
// 		WriteFhirOperationOutcome(w, 401, CreateOperationOutcome(fhir.IssueTypeForbidden, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	log.Info("opOutcome = " + spew.Sdump(opOutcome))
// 	issue := opOutcome.Issue[0]
// 	code := opOutcome.Issue[0].Code
// 	fmt.Printf("HandleOperationOutcome:904  --  code = %s,  Issue %s\n", code.Display(), *issue.Details.Text)
// 	//if code.Display() == "Duplicate" {
// 	//fmt.Printf("HandleOperationOutcome:906  --  code.Display = %s,  Issue %s\n", code.Display(), *issue.Details.Text)
// 	//WriteFhirOperationOutcome(w, 409, CreateOpOutcome(fhir.IssueTypeDuplicate, fhir.IssueSeverity(fhir.IssueTypeConflict), issue.Details.Text))

// 	WriteFhirOperationOutcome(w, 409, CreateOpOutcome(opOutcome.Issue))
// }

// func DetermineOutComeErr(body []byte) error {
// 	log.Info("HandleOperationOutcome  body n" + string(body))
// 	opOutcome := &fhir.OperationOutcome{}
// 	err := json.Unmarshal(body, &opOutcome)
// 	if err != nil {
// 		return log.Errorf("Unmarshal err = " + err.Error())
// 	}
// 	if opOutcome.Id == nil {
// 		return log.Errorf("opOutcome.Id is nil")
// 	}
// 	//fmt.Printf("HandleOperationOutcome:924  --  opOutcome = %s\n", spew.Sdump(opOutcome))
// 	issue := opOutcome.Issue[0]
// 	code := opOutcome.Issue[0].Code
// 	log.Info(fmt.Sprintf("HandleOperationOutcome:   code = %s,  Issue %s\n", code.Display(), *issue.Details.Text))
// 	return nil
// }
