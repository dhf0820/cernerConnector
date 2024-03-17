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
	//"strconv"
	"strings"
	"time"

	jw_token "github.com/dhf0820/golangJWT"
)

func findObservation(w http.ResponseWriter, r *http.Request) {
	JWToken = r.Header.Get("Authorization")
	Payload, status, err := jw_token.ValidateToken(JWToken, "")
	if err != nil {
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	userId := Payload.UserId
	resourceType := "Observation"
	log.Debug1(" --  Resource = " + resourceType)
	log.Debug2("--  Reading Body")
	body, err := io.ReadAll(r.Body) // Should be ConnectorPayload
	if err != nil {
		errMsg := log.ErrMsg("--  ReadAll Connector error: " + err.Error())
		fmt.Println(errMsg)
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	connectorPayload := common.ConnectorPayload{}
	err = json.Unmarshal(body, &connectorPayload)
	if err != nil {
		errMsg := log.ErrMsg("--  unmarshal ConnectorPayload err = " + err.Error())
		fmt.Println(errMsg)
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}

	connectorConfig := connectorPayload.ConnectorConfig
	log.Debug5("-- ConnectorPayload = " + spew.Sdump(connectorPayload))
	uri := r.URL.RequestURI()
	log.Debug3("--  uri: " + uri)
	log.Debug3("--  URL.Path() = " + r.URL.Path)
	log.Debug3("--  query = " + r.URL.RawQuery)
	if err := r.ParseForm(); err != nil {
		errMsg := log.ErrMsg("error parsing query: " + err.Error())
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
		return
	}

	queryStr := ""
	log.Debug3("-- Resource: " + resourceType)
	queryStr = fmt.Sprintf("%s?%s", resourceType, r.URL.RawQuery) //BuildDiagnosticQuery(r)
	QueryString = queryStr
	log.Debug3("--   QueryString = " + QueryString)
	log.Debug3("--  queryStr = " + queryStr)
	var bundle *fhir.Bundle
	var header *common.CacheHeader
	// qryStr := r.URL.RawQuery

	// log.Debug3(fmt.Sprintf(" - resource = %s  uri = %s", resourceType, qryStr))
	//url := connectorPayload.ConnectorConfig.HostUrl + queryStr

	log.Debug3(" - calling " + queryStr)
	var totalPages int64
	startTime := time.Now()
	totalPages, bundle, header, err = FindObservation(&connectorPayload, userId, queryStr, JWToken)
	log.Debug3(" - FindObservation returned")
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

		log.Debug3("OpOutcome: " + spew.Sdump(oo))
		WriteFhirOperationOutcome(w, finalStatus, oo)
		//CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityInformation, &errMsg))
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
	log.Debug3(fmt.Sprintf("--  Number of entries in bundle: %d", len(resp.Bundle.Entry)))
	log.Debug3(fmt.Sprintf("--  QueryId: " + header.QueryId))
	FillResourceResponse(&resp, resourceType)
	//fmt.Printf("findResource:614  --  Returning Bundle: %s\n", spew.Sdump(bundle))
	//WriteFhirResourceBundle(w, resp.Status, &resp)
	WriteFhirResponse(w, resp.Status, &resp)
}

// type BasicResource struct {
// 	Id           string         `json:"id"`
// 	Text         fhir.Narrative `json:"text"`
// 	ResourceType string         `json:"resourceType"`
// }

func getObservation(w http.ResponseWriter, r *http.Request) {
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
	results := json.RawMessage{}
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
	results, err = GetResource(cp, Resource, resourceId, JWToken)
	if err == nil {
		resp.Status = 200
		resp.Message = "Ok"
	} else {
		errMsg := log.ErrMsg("GetResource error: " + err.Error())
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverity(fhir.IssueTypeInvalid), &errMsg))
		return
	}
	basicResource := BasicResource{}
	err = json.Unmarshal(results, &basicResource)
	if err != nil {
		errMsg := log.ErrMsg("UnmarshalBasicResource error:  " + err.Error())
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityError, &errMsg))
		return
	}
	log.Debug3("Basic Resource: " + spew.Sdump(basicResource))
	resourceType = basicResource.ResourceType
	//TODO: unmarshal into a basic fhir resource (id, text)
	log.Debug3("FillResourceResponse for " + strings.ToLower(resourceType))
	switch strings.ToLower(resourceType) {
	case "patient":
		log.Debug3("Processing Patient")
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
		log.Debug3("Processing default: " + resourceType)
		basicResource := BasicResource{}
		err = json.Unmarshal(results, &basicResource)
		if err != nil {
			errMsg := log.ErrMsg("UnmarshalBasicResource error:  " + err.Error())
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityError, &errMsg))
			return
		}
		log.Debug3("Basic Resource: " + spew.Sdump(basicResource))
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

func GetObservation(cp *common.ConnectorPayload, resourceName, resourceId string, token string) (json.RawMessage, error) {
	log.Debug3("resourceName: " + resourceName)
	qry := resourceId //fmt.Sprintf("%s", resourceId)
	log.Debug3("accept: " + cp.ConnectorConfig.AcceptValue)
	log.Debug3("Final Query: " + qry)
	log.Info("cp.System.Url: " + cp.ConnectorConfig.HostUrl)
	c := New(cp.ConnectorConfig.HostUrl, cp.ConnectorConfig.AcceptValue)
	log.Debug3(fmt.Sprintf("Calling c.GetFhir with qry: %s  resource: %s", qry, resourceName))
	bodyBytes, resourceType, status, err := c.GetFhirBytes(qry, resourceName, token)
	if err != nil {
		errMsg := log.ErrMsg("Error calling GetFhirBytes: " + err.Error())
		log.Error(errMsg)
		return nil, log.Errorf(errMsg)
	}

	if bodyBytes != nil {
		log.Debug3("bodyBytes: " + string(bodyBytes))
		log.Debug3("resourceType: " + resourceType)
		log.Debug3("status: " + fmt.Sprint(status))
		if bodyBytes != nil {
			log.Debug3("bodyBytes: " + string(bodyBytes))
			switch strings.ToLower(resourceName) {
			case "OperationOutcome":
				opOut, err := fhir.UnmarshalOperationOutcome(bodyBytes)
				if err != nil {
					log.Debug3("Response --  Error Decoding OperationOutcone: " + err.Error())
					return bodyBytes, log.Errorf("Response --  Error Decoding Patient: " + err.Error())
				}
				log.Debug5("Response --  OperationOutcome: " + spew.Sdump(opOut))
				return bodyBytes, nil
			case "patient":
				patient, err := fhir.UnmarshalPatient(bodyBytes)
				if err != nil {
					log.Debug3("Response --  Error Decoding Patient: " + err.Error())
					return bodyBytes, log.Errorf("Response --  Error Decoding Patient: " + err.Error())
				}
				log.Debug5("Response --  Patient: " + spew.Sdump(patient))
				return bodyBytes, nil

			case "documentreference":
				docRef, err := fhir.UnmarshalDocumentReference(bodyBytes)
				if err != nil {
					log.Debug3("Response --  Error Decoding DocumentReference: " + err.Error())
					return bodyBytes, log.Errorf("Response --  Error Decoding DocumentReference: " + err.Error())
				}
				log.Debug5("Response --  DocumentReference: " + spew.Sdump(docRef))
				return bodyBytes, nil
			case "diagnosticreport":
				diagRept, err := fhir.UnmarshalDiagnosticReport(bodyBytes)
				if err != nil {
					log.Debug3("Response --  Error Decoding DiagnosticReport: " + err.Error())
					return bodyBytes, log.Errorf("Response --  Error Decoding DiagnosticReport: " + err.Error())
				}
				log.Debug5("Response --  DiagnosticReport: " + spew.Sdump(diagRept))
				return bodyBytes, nil

			default:
				log.Debug3("ResponseType:  " + resourceName)
				return bodyBytes, nil //log.Errorf("ResponseType --  Not supported: " + resourceName)

			}
		}
		return bodyBytes, nil
	}
	return nil, log.Errorf("No body read from GetFhirBytes")
	// if bodyBytes != nil {
	// rawMessage, err := c.GetFhir(qry, resourceName, token)
	// if err != nil {
	// 	return nil, err
	// }
	// err = nil
	// //fmt.Printf("GetResource:33  --  bundle: %s\n", spew.Sdump(bundle))
	// return rawMessage, nil

	// var res interface{}
	// //var binary *fhir4.Binary
	// //var vsResource *fhir4.VsResource
	// switch resourceName {
	// case "Binary":
	// 	res, err = fhir4.UnmarshalBinary(bytes)
	// case "Condition":
	// 	res, err = fhir4.UnmarshalCondition(bytes)
	// case "DocumentReference":
	// 	res, err = fhir4.UnmarshalDocumentReference(bytes)
	// 	if err != nil {
	// 		fmt.Printf("UnmarshalDocumentReference direct to DocumentReference err = %v\n", err)
	// 	}
	// case "Observation":
	// 	res, err = fhir4.UnmarshalObservation(bytes)
	// case "Patient":
	// 	fmt.Printf("Returning ByteArray")
	// 	return bytes, nil
	// 	pat, err := fhir4.UnmarshalPatient(bytes)
	// 	if err != nil {
	// 		fmt.Printf("UnmarshalPatient direct to patient err = %v\n", err)
	// 	}
	// 	bytes, err := json.Marshal(&pat)
	// 	if err != nil {
	// 		fmt.Printf("MarshalPatient to RawMessage err = %v\n", err)
	// 	}
	// 	rm := string(bytes)
	// 	fmt.Println(rm)

	// default:
	// 	err := fmt.Errorf("unknown type of Resource: [%s]", resourceName)
	// 	return nil, err
	// }
	// //fmt.Printf("GetResource is returning err: %v\n", err)
	// fmt.Printf("  QueryTime: %s\n", time.Since(startTime))
	// return bytes, err
}

func FindObservation(connPayLoad *common.ConnectorPayload, userId, query, JWToken string) (int64, *fhir.Bundle, *common.CacheHeader, error) {
	resource := "Observation"
	page := 1

	//sysCfg := connPayLoad.System
	connConfig := connPayLoad.ConnectorConfig
	log.Debug3("--  query: " + query)
	fullQuery := fmt.Sprintf("%s", query)
	//fmt.Printf("FindRecource:84  --  ConectorPayload: %s\n", spew.Sdump(connPayLoad)))
	log.Debug3("--  UserId: " + userId)
	log.Debug3("--  fullQuery: " + fullQuery)
	//log.Debug3("-- Page: %d\n", page)
	//fmt.Printf("FindResource:90  --  ConnectorConfig: %s\n", spew.Sdump(connConfig))
	//fmt.Printf("FindResource:91  --  query: %s\n", query)

	//TODO: Process the query in the background filling the resourceCache and BundleCache. Assign a cacheId on the call
	//Once background is started wait in a loop checking the ResourceCache Status using the assigned cacheId until either
	// Have count documents or status is finished.
	// check every 10 seconds.  Should be a FhirSystem variable value to avoid code change
	c := New(connPayLoad.ConnectorConfig.HostUrl, "application/json")
	startTime := time.Now()
	log.Debug3("query: " + fullQuery)
	bundle, err := c.GetFhirBundle(fullQuery, JWToken)
	if err != nil {
		// msg := log.ErrMsg("GetNextResource error: " + err.Error())
		// fmt.Println(msg)
		fmt.Println(err.Error())
		return 0, nil, nil, err
	}
	// bundle, err := c.Query(query, JWToken) // Perform the actul query of the fhir server
	// if err != nil {
	// 	return 0, nil, nil, err
	// }
	log.Debug5("bundle: " + spew.Sdump(bundle))
	header := &common.CacheHeader{}
	header.SystemCfg = connPayLoad.System
	header.ResourceType = resource
	header.UserId = userId
	header.PageId = page
	queryId := primitive.NewObjectID().Hex()
	header.QueryId = queryId
	log.Debug5("connConfig: " + spew.Sdump(connConfig))
	//header.CacheBase = fmt.Sprintf("%s/%s", connConfig.CacheUrl, header.SystemCfg.ID.Hex())
	//header.ResourceCacheBase = fmt.Sprintf("%s/%s/%s/BundleTransaction", connConfig.CacheUrl, header.FhirSystem.ID.Hex())
	//header.GetBundleCacheBase = fmt.Sprintf("%s/%s/BundleTransaction", header.CacheBase, header.SystemCfg.ID.Hex())
	//header.GetResourceCacheBase = fmt.Sprintf("%s/%s/CachePage", header.CacheBase, header.SystemCfg.ID.Hex())

	cacheBundle := common.CacheBundle{}
	cacheBundle.PageId = header.PageId
	cacheBundle.Header = header
	cacheBundle.ID = primitive.NewObjectID()
	//fmt.Printf("\n\n\n\n$$$ FindResource:110 calling CacheResourceBundleAndEntries (without bundle) - %s \n", spew.Sdump(cacheBundle))
	//fmt.Printf("FindResource:126  --  bundle = %s\n", spew.Sdump(bundle))
	//Cache the first bundle(page)
	log.Debug3(fmt.Sprintf("--  Query %s for %ss took %s\n\n\n", connPayLoad.ConnectorConfig.Label, resource, time.Since(startTime)))
	log.Debug3("--  UnmarshalBundle")
	// bundle := fhir4.Bundle{}
	// bundle, err = fhir4.UnmarshalBundle(byte)
	// if err != nil {
	// 	return 0, nil, nil, err
	// }
	log.Debug3("bundle: " + spew.Sdump(bundle))
	cacheBundle.Bundle = bundle
	startTime = time.Now()
	if !UseCache() {
		log.Info("Using Caching")
		pg, err := CacheResourceBundleAndEntries(&cacheBundle, JWToken, int64(page))
		//log.Debug3(fmt.Sprintf("CacheResourceBundleAndEntries returned %d %ss in page: %d for %s  took %s", len(cacheBundle.Bundle.Entry), resource, page, sysCfg.DisplayName, time.Since(startTime)))
		if err != nil {
			//return err and done
			return int64(pg + 1), bundle, cacheBundle.Header, err
		}
		log.Debug3("links: " + spew.Sdump(bundle.Link))
		//Follow the bundle links to retrieve all bundles(pages) in the query response
		nextURL := GetNextObservationUrl(bundle.Link)
		total := int64(0)
		if nextURL == "" {
			log.Debug3(fmt.Sprintf("GetNext%sUrl initialy No Next - One page only ", resource))
			total, err = TotalCacheForQuery(cacheBundle.QueryId)
			cacheBundle.Header.PageId = pg
			log.Debug3("total: " + fmt.Sprint(total))
			//page++
			return int64(pg), bundle, cacheBundle.Header, err
		}
		page++
		go c.GetNextObservation(header, nextURL, resource, JWToken, page)
	} else {
		log.Info("Not Using Caching")
	}
	log.Info("Used queery: " + fullQuery)
	log.Debug3("--  Page 1 total time: " + fmt.Sprint(time.Since(startTime)))
	// There is one full page and possibley more. Respond with two aso they user will create two page buttons and update every
	// 10 secnds.
	//return int64(page), bundle, cacheBundle.Header, err
	if bundle == nil {
		return 0, bundle, cacheBundle.Header, log.Errorf("No resources found")
	}
	if len(bundle.Entry) == 0 {
		return 0, bundle, cacheBundle.Header, log.Errorf("No resources found")
	} else {
		return int64(len(bundle.Entry)), bundle, cacheBundle.Header, err

	}
	return 0, bundle, cacheBundle.Header, err
}

func GetNextObservationUrl(link []fhir.BundleLink) string {
	for _, lnk := range link {
		if lnk.Relation == "next" {
			log.Info("--  There is a next page to get")
			return lnk.Url
		}
	}
	return ""
}

// //GetNextResource: fetches the resource at provided url, processes it and checks if more to call.
func (c *Connection) GetNextObservation(header *common.CacheHeader, url, resource, token string, page int) {
	log.Debug3("-GetNextObservation- page:  " + fmt.Sprint(page))
	//fmt.Printf("GetNextResource:155  --  resource: %s\n", resource) //spew.Sdump(header))
	//Call Remote FHIR server for the resource bundle
	startTime := time.Now()
	bundle, err := c.GetFhirBundle(url, JWToken)
	if err != nil {
		msg := fmt.Sprintf("--  error: " + err.Error())
		fmt.Println(msg)
		return
	}
	log.Debug3(fmt.Sprintf("--  Query Next Set from %s of %s time: %s\n", header.SystemCfg.DisplayName, header.ResourceType, time.Since(startTime)))
	// fmt.Printf("GetNextResource:175  --  UnmarshalBundle\n")
	// bundle, err := fhir4.UnmarshalBundle(bytes)
	// if err != nil {
	// 	msg := fmt.Sprintf("GetNextResource:178 unmarshal : %s", err.Error())
	// 	//fmt.Printf(msg)
	// 	fmt.Println(msg)
	// 	return
	// }

	//unMarshalResource(resource, bundle)
	header.PageId += 1
	tn := time.Now()
	header.CreatedAt = &tn
	cacheBundle := common.CacheBundle{}
	cacheBundle.ID = primitive.NewObjectID()
	cacheBundle.Header = header
	cacheBundle.Bundle = bundle
	log.Debug3("-- Calling CacheResourceBundleAndEntries")
	pg, err := CacheResourceBundleAndEntries(&cacheBundle, token, int64(page))
	if err != nil {
		log.Error("GetNextObservation: returned err: " + err.Error())
		return
		//return int64(pg + 1), &bundle, cacheBundle.Header, err
	}

	log.Debug3("--  GetNextResourceUrl")
	nextURL := GetNextConditionUrl(bundle.Link)
	if nextURL == "" {
		msg := fmt.Sprintf("GetNextObservationUrl Last page had %d Resources processed ", len(bundle.Entry))
		// fmt.Println(msg)
		log.Debug3("--  Should return:  " + msg)
		return
	} else {
		log.Debug3("-- GetNextObservation is being called in the background")
		go c.GetNextObservation(header, nextURL, resource, token, pg+1)
		log.Debug3("GetNextObservation Returned")
	}
	log.Debug3("GetNextObservation is returning")
}
