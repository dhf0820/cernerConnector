package main

import (
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

var JWToken string

func findPatient(w http.ResponseWriter, r *http.Request) {
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
	body, err := io.ReadAll(r.Body) // Should be ConnectorPayload
	err = json.Unmarshal(body, &connectorPayload)
	if err != nil {
		errMsg := log.ErrMsg("--  unmarshal ConnectorPayload err = " + err.Error())
		fmt.Println(errMsg)
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	log.Debug3("-- ConnectorPayload = " + spew.Sdump(connectorPayload))

	connectorConfig := connectorPayload.System.ConnectorConfig
	CurrentSystem = *connectorPayload.System
	CurrentSystemID = connectorPayload.System.ID
	params = mux.Vars(r)
	log.Debug2(fmt.Sprintf("Params: %v\n", params))
	if params["resource"] != "" {
		log.Debug2(" --  Using Resource in params")
		resourceType = params["resource"]
	} else {
		log.Debug2(fmt.Sprintf("--  url = %s", r.URL.Path))
		uri := r.URL.RequestURI()
		log.Debug2(fmt.Sprintf("--  url = %s", uri))

		URIparts := strings.Split(uri, "&")
		log.Debug2(fmt.Sprintf("  uri = %s", URIparts[len(URIparts)-1]))
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
		log.Debug2(fmt.Sprintf("--  uri = %s", r.URL.RequestURI()))
		resourceType = DetermineResource(r.URL.Path, "/api/rest/v1/")
		if resourceType == "" {
			errMsg := log.ErrMsg("Resource not found in URL")
			WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
			return
		}
	}
	//log.Debug3("ResourceType : " + resourceType)
	log.Debug2(" --  ResourceType = " + resourceType)

	//log.Debug1(" -- being called for resource: [%s]\n", Resource)
	// log.Debug3("--  Reading Body")
	// body, err := io.ReadAll(r.Body) // Should be ConnectorPayload
	// if err != nil {
	// 	errMsg := log.ErrMsg("--  ReadAll FhirSystem error " + err.Error())
	// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// connectorPayload := common.ConnectorPayload{}

	//fmt.Printf("findResource:470  --  body = %s\n", string(body))
	// err = json.Unmarshal(body, &connectorPayload)
	// if err != nil {
	// 	errMsg := log.ErrMsg("--  unmarshal ConnectorPayload err = " + err.Error())
	// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	log.Debug3("-- ConnectorPayload = " + spew.Sdump(connectorPayload))

	// connectorConfig := connectorPayload.System.ConnectorConfig
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
