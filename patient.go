package main

import (
	// 	//"context"
	// 	"encoding/json"
	//"bytes"
	"encoding/json"
	//"net/http"

	// 	//"errors"
	"fmt"
	//"os"

	"github.com/davecgh/go-spew/spew"
	fhir "github.com/dhf0820/fhir4"
	common "github.com/dhf0820/uc_core/common"

	//common "github.com/dhf0820/uc_core/common"
	//"io"
	"strings"
	"time"

	log "github.com/dhf0820/vslog"
	//"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func GetResourceBytes(cp *common.ConnectorPayload, resourceName, resourceId string, token string) ([]byte, int, error) {
// 	//startTime := time.Now()
// 	//log.Printf("GetResource:23 - cp: %s\n\n", spew.Sdump(cp))
// 	//url := fmt.Sprintf("%s/%s%s", fhirSystem.FhirUrl, resourceName, resourceId)
// 	qry := resourceId //fmt.Sprintf("%s", resourceId)
// 	log.Debug3("accept: " + cp.ConnectorConfig.AcceptValue)
// 	log.Debug3("Final Query: " + qry)
// 	log.Info("cp.System.Url: " + cp.ConnectorConfig.HostUrl)
// 	c := New(cp.ConnectorConfig.HostUrl, cp.ConnectorConfig.AcceptValue)
// 	log.Debug2(" Calling c.GetFhir with qry: %s  resource: %s", qry, resourceName)
// 	bodyBytes, resourceType, status, err := c.GetFhirBytes(qry, resourceName, token)
// 	log.Debug3(fmt.Sprintf("resourceType: Patient status: %d",  ResourceType, status))
// 	if bodyBytes != nil {
// 		log.Debug3("bodyBytes: " + string(bodyBytes))
// 		switch strings.ToLower(resourceName) {
// 		case "OperationOutcome":
// 			opOut, err := fhir.UnmarshalOperationOutcome(bodyBytes)
// 			if err != nil {
// 				log.Debug3("Response --  Error Decoding OperationOutcone: " + err.Error())
// 				return bodyBytes, status, log.Errorf("Response --  Error Decoding Patient: " + err.Error())
// 			}
// 			log.Debug3("Response --  OperationOutcome: " + spew.Sdump(opOut))
// 			return bodyBytes, status, nil
// 		case "patient":
// 			patient, err := fhir.UnmarshalPatient(bodyBytes)
// 			if err != nil {
// 				log.Debug3("Response --  Error Decoding Patient: " + err.Error())
// 				return bodyBytes, status, log.Errorf("Response --  Error Decoding Patient: " + err.Error())
// 			}
// 			log.Debug5("Response --  Patient: " + spew.Sdump(patient))
// 			return bodyBytes, status, nil

// 		case "documentreference":
// 			docRef, err := fhir.UnmarshalDocumentReference(bodyBytes)
// 			if err != nil {
// 				log.Debug3("Response --  Error Decoding DocumentReference: " + err.Error())
// 				return bodyBytes, status, log.Errorf("Response --  Error Decoding DocumentReference: " + err.Error())
// 			}
// 			log.Debug5("Response --  DocumentReference: " + spew.Sdump(docRef))
// 			return bodyBytes, status, nil
// 		case "diagnosticreport":
// 			diagRept, err := fhir.UnmarshalDiagnosticReport(bodyBytes)
// 			if err != nil {
// 				log.Debug3("Response --  Error Decoding DiagnosticReport: " + err.Error())
// 				return bodyBytes, status, log.Errorf("Response --  Error Decoding DiagnosticReport: " + err.Error())
// 			}
// 			log.Debug5("Response --  DiagnosticReport: " + spew.Sdump(diagRept))
// 			return bodyBytes, status, nil

// 		default:
// 			log.Debug3("ResponseType --  Not supported: " + resourceName)
// 			return bodyBytes, 400, log.Errorf("ResponseType --  Not supported: " + resourceName)

// 		}
// 		// diagRept, err := fhir.UnmarshalDiagnosticReport(byte)
// 		// if err != nil {
// 		// 	log.Debug3("Response --  Error Decoding DiagnosticReport: " + err.Error())
// 		// 	return byte, 400, log.Errorf("Response --  Error Decoding DiagnosticReport: " + err.Error())
// 		// }
// 		// log.Debug3("Response --  DiagnosticReport: " + spew.Sdump(diagRept))
// 		// return byte, resp.StatusCode, nil

// 		//TODO: Write test for this
// 	}
// 	//rawMessage, err := c.GetFhir(qry, resourceName, token)

// 	return nil, 400, err

// }

// //Request a specific resource by id
func GetPatient(cp *common.ConnectorPayload, resourceName, resourceId string, token string) (json.RawMessage, error) {
	log.Info("GetResource entered")
	qry := resourceId //fmt.Sprintf("%s", resourceId)
	log.Debug1("resoureName: " + resourceName)
	log.Debug3("accept: " + cp.ConnectorConfig.AcceptValue)
	log.Debug3("Final Query: " + qry)
	log.Info("cp.System.Url: " + cp.ConnectorConfig.HostUrl)
	log.Debug1("cp.ConnectorConfig.CacheUrl: " + cp.ConnectorConfig.CacheUrl)
	c := New(cp.ConnectorConfig.HostUrl, cp.ConnectorConfig.AcceptValue)
	log.Debug3(fmt.Sprintf("Calling c.GetFhir with qry: %s  resource: %s", qry, resourceName))
	bodyBytes, resourceType, status, err := c.GetFhirBytes(qry, resourceName, token)
	if err != nil {
		errMsg := log.ErrMsg("Error calling GetFhirBytes: " + err.Error())
		log.Error(errMsg)
		return nil, log.Errorf(errMsg)
	}
	if bodyBytes != nil {
		//log.Debug3("bodyBytes: " + string(bodyBytes))
		log.Debug3("resourceType: " + resourceType)
		log.Debug3("status: " + fmt.Sprint(status))
		if bodyBytes != nil {

			//log.Debug3("bodyBytes: " + string(bodyBytes))
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
				log.Debug3("Returning DocumentReference")
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

//var PageSize int

func FindPatient(connPayLoad *common.ConnectorPayload, userId, query, JWToken string) (int64, *fhir.Bundle, *common.CacheHeader, error) {
	fmt.Printf("\n\n\n")
	log.Debug2("Starting FindPatient")
	PageSize = 10
	page := 1
	resource := "Patient"
	connConfig := connPayLoad.System.ConnectorConfig
	//log.Debug3("connPayload: " + spew.Sdump(connPayLoad))

	//systemCfg := connPayLoad.System

	//log.Debug2("query: " + query)
	fullQuery := query //fmt.Sprintf("Patient?%s", query)
	//fmt.Printf("FindRecource:84  --  ConectorPayload: %s\n", spew.Sdump(connPayLoad))
	log.Debug3("UserId: " + userId)
	log.Debug2("fullQuery: " + fullQuery)
	// fmt.Printf("FindResource:87  --  FhirSystem: %s\n", spew.Sdump(fhirSystem))

	//log.Debug3("FindResource:89  -- Page: %d\n", page)
	//fmt.Printf("FindResource:90  --  ConnectorConfig: %s\n", spew.Sdump(connConfig))
	//fmt.Printf("FindResource:91  --  query: %s\n", query)

	//TODO: Process the query in the background filling the resourceCache and BundleCache. Assign a cacheId on the call
	//Once background is started wait in a loop checking the ResourceCache Status using the assigned cacheId until either
	// Have count documents or status is finished.
	// check every 10 seconds.  Should be a FhirSystem variable value to avoid code change
	c := New(connPayLoad.System.ConnectorConfig.HostUrl, "application/json")
	//c := New(connPayLoad.ConnectorConfig.HostUrl, "application/json")
	startTime := time.Now()
	//Get the first bundle(page)
	bundle, err := c.GetFhirBundle(fullQuery, JWToken)
	log.Info("GetFhirBundle returned")
	if err != nil {
		msg := log.ErrMsg("GetFhirBundle error: " + err.Error())
		fmt.Println(msg)
		fmt.Printf("error: %s\n", err.Error())
		return 0, nil, nil, err
	}
	if bundle == nil {
		//log.Debug3("bundle is nil")
		return 0, nil, nil, log.Errorf("bundle is nil")
	}
	//Have the first bundle(page).  Cache it and start the background process to get the rest of the bundles

	header := &common.CacheHeader{}
	header.SystemCfg = connPayLoad.System
	header.ResourceType = "Patient"
	header.UserId = userId
	header.PageId = page
	queryId := primitive.NewObjectID()
	header.QueryId = queryId
	header.CacheStatusUrl = fmt.Sprintf("queue/%s/CacheStatus", header.QueryId.Hex())
	header.CachePageUrl = fmt.Sprintf("queue/%s/CachePatientPage", header.QueryId.Hex())

	//log.Debug3("connConfig: " + spew.Sdump(connConfig))
	//log.Debug3("header: " + spew.Sdump(header))
	log.Info("CacheUrl: " + connConfig.CacheUrl)
	header.CacheUrl = connConfig.CacheUrl
	//header.CacheUrl = fmt.Sprintf("%s/system/%s", connConfig.CacheUrl, header.SystemCfg.ID.Hex())
	fmt.Printf("\n\n		######")
	log.Info("### CacheUrl: " + header.CacheUrl)

	//header.ResourceCacheUrl = fmt.Sprintf("%s/%s/%s/BundleTransaction", connConfig.CacheUrl, header.FhirSystem.ID.Hex())
	//header.GetBundleCacheUrl = fmt.Sprintf("%s/%s/BundleTransaction", header.CacheUrl, header.SystemCfg.ID.Hex())
	//header.GetResourceCacheUrl = fmt.Sprintf("%s/%s/CachePage", header.CacheUrl, header.SystemCfg.ID.Hex())

	cacheBundle := common.CacheBundle{}
	cacheBundle.PageId = header.PageId
	cacheBundle.Header = header
	cacheBundle.ID = primitive.NewObjectID()
	cacheBundle.Bundle = bundle
	log.Debug1(fmt.Sprintf("Query %s for Patients took %s", connConfig.Label, time.Since(startTime)))
	cacheBundle.Bundle = bundle
	startTime = time.Now()

	// if UseCache() {
	log.Info("CacheHeader: " + spew.Sdump(cacheBundle.Header))
	log.Info(fmt.Sprintf("Calling CacheViaCore for page %d ", page))
	cacheURL := header.CacheUrl
	log.Info("CacheURL: " + cacheURL)
	log.Info("cacheBundle: " + spew.Sdump(cacheBundle))
	log.Info("FindResource calling	CacheViaCore")
	err = CacheViaCore(bundle, queryId, cacheBundle.Header.ResourceType, JWToken, cacheURL, page)
	log.Debug2("CacheViaCore returned")
	if err != nil {
		log.Error(err.Error())

	}
	// log.Debug3(fmt.Sprintf("Length of page %d is %d", page, len(bundle.Entry)))
	// if page == 1 {
	// 	log.Debug3("")
	// 	PageSize = len(bundle.Entry)
	// 	log.Debug3("$$$$$$ length of Page 1: " + fmt.Sprint(PageSize))

	// }
	// if UseCache() {
	// 	log.Debug3("Calling CacheViaCore for page: " + fmt.Sprint(page))
	// 	log.Debug3("Cache current Page: " + fmt.Sprint(page))
	// 	err = CacheViaCore(bundle, queryId, CurrentToken, "ALL", page)
	// 	if err != nil {
	// 		log.Error("CacheViaCore err: " + err.Error())
	// 		return int64(page), 1, bundle, cacheBundle.Header, err
	// 	}

	// 	//log.Debug3("Calling CacheResourceBundleAndEntries with token: ") // + JWToken)
	// 	// pg, err := CacheResourceBundleAndEntries(&cacheBundle, JWToken, page)
	// 	// log.Debug3(fmt.Sprintf("pg = %d  page = %d", pg, page))
	// 	// log.Debug3(fmt.Sprintf("CacheResource returned %d %ss in page: %d for %s  took %s\n", len(cacheBundle.Bundle.Entry), resource, page, systemCfg.DisplayName, time.Since(startTime)))
	// 	// if err != nil {
	// 	// 	//return err and done
	// 	// 	return int64(pg + 1), bundle, cacheBundle.Header, err
	// 	// }
	// 	//log.Debug3("--  links: " + spew.Sdump(bundle.Link))
	// 	//Follow the bundle links to retrieve all bundles(pages) in the query response
	// 	nextURL := GetNextResourceUrl(bundle.Link)
	// 	total := int64(0)
	// 	if nextURL == "" {
	// 		log.Debug3("-- GetNext" + resource + " initialy No Next - One page only ")

	// 		//total, err = TotalCacheForQuery(cacheBundle.QueryId)
	// 		cacheBundle.Header.PageId = int(1)
	// 		//page++
	// 		log.Debug3("total: " + fmt.Sprint(total))
	// 		return int64(page), 1, bundle, cacheBundle.Header, err
	// 	}
	// 	// PageSize = len(bundle.Entry)
	// 	// log.Debug3("$$$$$$ length of current page: " + fmt.Sprint(PageSize))

	// 	page++
	// 	fmt.Printf("\n\n\n")
	// 	log.Debug3("Calling GetNextResource in the background for page: " + fmt.Sprint(page))
	// 	go c.GetNextResource(header, nextURL, resource, queryId, JWToken, int(page))
	// 	log.Debug3(fmt.Sprintf("Page 1 total time: %s", time.Since(startTime)))
	// 	// There is one full page and possibly more. Respond with two so the user will create two page buttons and update every
	// 	// 10 secnds.
	// 	//return int64(page), bundle, cacheBundle.Header, err
	// 	if bundle == nil {
	// 		return 0, 0, nil, nil, log.Errorf("Bundle is nil")
	// 	}
	// 	if len(bundle.Entry) == 0 {
	// 		log.Warn("Bundle length should never be 0")
	// 		return 0, 0, bundle, cacheBundle.Header, log.Errorf("No resources found")
	// 	} else {
	// 		//There are at least two pages since there is a next on the first page
	// 		return int64(len(bundle.Entry)), 2, bundle, cacheBundle.Header, err
	// 	}
	// } else {
	// 	log.Debug3("Not using cache for page: " + fmt.Sprint(page))

	log.Debug3("--  links: " + spew.Sdump(bundle.Link))
	//Follow the bundle links to retrieve all bundles(pages) in the query response
	log.Info("-- Call GetNextResourceUrl")
	//Call GetNextResourceURL to get the next page of resources
	nextURL := GetNextResourceUrl(bundle.Link)
	total := int64(0)
	if nextURL == "" {
		log.Debug1("-- GetNext" + resource + " initialy No Next - One page only ")

		total, err = TotalCacheForQuery(cacheBundle.QueryId)
		cacheBundle.Header.PageId = int(page)
		//page++
		log.Debug2("total: " + fmt.Sprint(total))
		log.Debug2("Returning page: " + fmt.Sprint(page))
		return int64(page), bundle, cacheBundle.Header, err
	}
	page++
	log.Debug2("--  Calling GetNextResource as go routine	")
	go c.GetNextResource(header, nextURL, resource, queryId, JWToken, int(page))
	log.Debug2(fmt.Sprintf("Page 1 total time: %s", time.Since(startTime)))
	// There is one full page and possibly more. Respond with two so the user will create two page buttons and update every
	// 2 secnds.
	//return int64(page), bundle, cacheBundle.Header, err
	if bundle == nil {
		return 0, nil, nil, log.Errorf("Bundle is nil")
	}
	if len(bundle.Entry) == 0 {
		return 0, bundle, cacheBundle.Header, log.Errorf("No resources found")
	}
	//else {
	//	//TODO: need to cache the first page
	//	return int64(len(bundle.Entry)), bundle, cacheBundle.Header, err
	//}

	if bundle == nil {
		log.Warn("bundle is nil")
	}
	if cacheBundle.Header == nil {
		log.Warn("cacheBundle.Header is nil")
	}
	enteries := bundle.Entry
	numEnteries := len(enteries)

	log.Debug2("Number of entries:  " + fmt.Sprint(numEnteries))
	return int64(numEnteries), bundle, cacheBundle.Header, err
}

// func GetNextResourceUrl(link []fhir.BundleLink) string {
// 	log.Debug3("link: " + spew.Sdump(link))
// 	for _, lnk := range link {
// 		if lnk.Relation == "next" {
// 			log.Debug2("--  There is a next page to get: " + lnk.Url)
// 			return lnk.Url
// 		}
// 	}
// 	return ""
// }

// //GetNextResource: fetches the resource at provided url, processes it and checks if more to call.
// func (c *Connection) GetNextResource(header *common.CacheHeader, url, resource string, queryId primitive.ObjectID, token string, page int) {
// 	fmt.Printf("\n\n\n\n####################  GetNextResource page: %d   ###############\n", page)
// 	log.Debug2("header.SystemCfg.CacheUrl: " + header.SystemCfg.ConnectorConfig.CacheUrl)
// 	header.CacheUrl = header.SystemCfg.ConnectorConfig.CacheUrl
// 	log.Debug2(fmt.Sprintf("--  resource: %s  -  header.CacheUrl: %s", resource, header.CacheUrl))

// 	//Call Remote FHIR server for the resource bundle
// 	//queryId := header.QueryId
// 	log.Info("queryId: " + queryId.Hex())
// 	startTime := time.Now()
// 	bundle, err := c.GetFhirBundle(url, JWToken)
// 	if err != nil {
// 		log.Error("c.GetFhirBundle error: " + err.Error())
// 		return
// 	}
// 	log.Info(fmt.Sprintf("--  Query Next Set from %s of %s time: %s\n", header.SystemCfg.DisplayName, header.ResourceType, time.Since(startTime)))
// 	// // fmt.Printf("GetNextResource:175  --  UnmarshalBundle\n")
// 	// // bundle, err := fhir4.UnmarshalBundle(bytes)
// 	// // if err != nil {
// 	// // 	msg := fmt.Sprintf("GetNextResource:178 unmarshal : %s", err.Error())
// 	// // 	//fmt.Printf(msg)
// 	// // 	fmt.Println(msg)
// 	// // 	return
// 	// // }

// 	// //unMarshalResource(resource, bundle)
// 	// header.PageId += 1
// 	// tn := time.Now()
// 	// header.CreatedAt = &tn
// 	// cacheBundle := common.CacheBundle{}
// 	// cacheBundle.ID = primitive.NewObjectID()
// 	// cacheBundle.Header = header
// 	// cacheBundle.Bundle = bundle
// 	// log.Debug3("-- Calling CacheResourceBundleAndEntries for page: " + fmt.Sprint(page))
// 	// pg, err := CacheResourceBundleAndEntries(&cacheBundle, token, int64(page))
// 	// if err != nil {
// 	// 	log.Errorf("GetNextResource returned err: " + err.Error())
// 	// 	return
// 	// }
// 	// log.Debug3(fmt.Sprintf("pg: %d  Page: %d", pg, page))

// 	log.Debug2("Cache current Page: " + fmt.Sprint(page))
// 	//log.Info("header: " + spew.Sdump(header))
// 	cacheURL := header.CacheUrl
// 	log.Debug2("CacheURL: " + cacheURL)
// 	err = CacheViaCore(bundle, queryId, token, cacheURL, page)
// 	if err != nil {
// 		log.Error("CacheViaCore err: " + err.Error())
// 		return
// 	}
// 	log.Debug2("-- Calling GetNextResourceUrl")
// 	nextURL := GetNextResourceUrl(bundle.Link)
// 	if nextURL == "" {
// 		//pageSize := 10 //TODO: Change pageSize from constant
// 		onPage := len(bundle.Entry)
// 		log.Warn(fmt.Sprintf("GetNextResource Last page had %d Resources processed ", onPage))
// 		log.Debug2("Send post to tell core the query is done and to complete it.")
// 		err = FinishCache(header.SystemCfg, queryId, token, page, onPage)

// 		if err != nil {
// 			log.Error("GetNextResource err: " + err.Error())
// 			return
// 		}
// 		return
// 	} else {
// 		page = page + 1
// 		log.Debug2("--go c.GetNextResource is being called in the background")
// 		go c.GetNextResource(header, nextURL, resource, queryId, token, page)
// 		log.Debug2("-- GetNextResource Returned and background started")
// 	}
// 	log.Debug2("GetNextResource is returning")
// }

// func GetHeaderInfoFromBundle(resource string, hdr *common.CacheHeader, bundle *fhir4.Bundle) (string, string, error) {
// 	fmt.Printf("\n\n\n################## GetHederInfoFromBundle:220  --  for Resource: [%s]\n\n", resource)
// 	//resHeader := common.ResourceHeader{}
// 	switch resource {
// 	case "Patient":
// 		fsIdentifiers := hdr.SystemCfg.Identifiers
// 		res, err := fhir4.UnmarshalPatient(bundle.Entry[0].Resource)
// 		if err != nil {
// 			log.Errorf("%s unmarshal : %s", resource, err.Error())
// 			return "", "", err
// 		}
// 		hdr.PatientId = *res.Id
// 		hdr.ResourceId = *res.Id
// 		GetPatientIdentifier(res.Identifier, fsIdentifiers, "mrn")
// 		fmt.Printf("\nGetHeaderInfoFromBundle:221 -- Patient: %s\n", spew.Sdump((res)))
// 		fmt.Printf("\nGetHeaderInfoFromBundle:222")
// 		res1, err := fhir4.UnmarshalPatient(bundle.Entry[1].Resource)
// 		if err != nil {
// 			log.Errorf("%s unmarshal : %s", resource, err.Error())
// 			return "", "", err
// 		}
// 		hdr.PatientId = *res.Id
// 		hdr.ResourceId = *res.Id
// 		fmt.Printf("\nGetHeaderInfoFromBundle:230 -- Patient: %s\n", spew.Sdump((res1)))
// 		return *res.Id, *res.Id, nil
// 	case "DocumentReference":
// 		//log.Printf("GetHederInfoFromBundle312: --  Raw data: %s\n", string(bundle.Entry[0].Resource))
// 		res, err := fhir4.UnmarshalDocumentReference(bundle.Entry[0].Resource)
// 		if err != nil {
// 			log.Errorf("%s unmarshal : %s", resource, err.Error())
// 			return "", "", err
// 		}
// 		//fmt.Printf("\n###DocumentReference: %s\n", spew.Sdump(res))
// 		parts := strings.Split(*res.Subject.Reference, "/")
// 		hdr.PatientId = parts[1]
// 		hdr.ResourceId = *res.Id
// 		fmt.Printf("GetHeaderInfoFromBundle:243 -- DocumentReference: %s\n", spew.Sdump((res)))
// 		return parts[1], *res.Id, nil
// 	case "DiagnosticReport":
// 		res, err := fhir4.UnmarshalDiagnosticReport(bundle.Entry[0].Resource)
// 		if err != nil {
// 			log.Errorf("%s unmarshal : %s", resource, err.Error())
// 			return "", "", err
// 		}
// 		//fmt.Printf("\n###DiagnosticReport: %s\n", spew.Sdump(res))
// 		parts := strings.Split(*res.Subject.Reference, "/")
// 		hdr.PatientId = parts[1]
// 		hdr.ResourceId = *res.Id
// 		fmt.Printf("GetHeaderInfoFromBundle:255 -- DiagnosticReport: %s\n", spew.Sdump((res)))
// 		return parts[1], "", nil
// 	case "Observation":
// 		res, err := fhir4.UnmarshalObservation(bundle.Entry[0].Resource)
// 		if err != nil {
// 			log.Errorf("%s unmarshal : %s", resource, err.Error())
// 			return "", "", err
// 		}
// 		//fmt.Printf("\n###Observation: %s\n", spew.Sdump(res))
// 		parts := strings.Split(*res.Subject.Reference, "/")
// 		hdr.PatientId = parts[1]
// 		hdr.ResourceId = *res.Id
// 		fmt.Printf("GetHeaderInfoFromBundle:267 -- Observation:%s\n", spew.Sdump((res)))
// 		return parts[1], "", nil
// 	}
// 	return "", "", nil
// }

// func GetCacheStatus(ucUrl, queryId string) int {
// 	// coreURL := ucUrl + "/BundleTransaction"
// 	// client := &http.Client{}
// 	// //fmt.Printf("Send Status to: [%s]\n", statusURL)
// 	// fmt.Printf("CacheResourceBundleAndEntries:99  --  Using CoreUrl: %s\n", coreURL)
// 	// req, _ := http.NewRequest("GET", coreURL, bytes.NewBuffer(cacheBundle))
// 	// //req, _ := http.NewRequest("POST", coreURL, bytes.NewBuffer(cacheBundle))
// 	// //r, _ := http.NewRequest("POST", coreURL, nil)
// 	// fmt.Printf("\nCacheResourceBundleEntries:339  --  Req: %s\n\n\n", spew.Sdump(req))
// 	// req.Header.Set("Accept", "application/json")
// 	// req.Header.Set("Content-Type", "application/json")
// 	// req.Header.Set("Authorization", token)
// 	// fmt.Printf("\nCacheResourceBundleEntries:107  --  Calling core: %s\n", coreURL)
// 	// ///fmt.Printf("Using Token: %s\n", token)
// 	// _, err = client.Do(req)
// 	return 0
// }

// func GetConnectorPayload(r *http.Request) (*common.ConnectorPayload, error) {
// 	log.Info("GetConnectorPayload entered")
// 	body, err := io.ReadAll(r.Body) // Should be ConnectorPayload
// 	if err != nil {
// 		return nil, log.Errorf("ReadAll FhirSystem error " + err.Error())
// 	}
// 	//mt.Printf("GetConnectorPayload:717  -- Got Body Now Unmarshal ConnectorPayload\n")
// 	b := string(body)
// 	log.Debug3("GetConnectorPayload Body: " + b)
// 	conPayload := &common.ConnectorPayload{}
// 	err = json.Unmarshal(body, &conPayload)
// 	if err != nil {
// 		log.Error(" --  unmarshal err = " + err.Error())
// 		// errMsg := err.Error()
// 		// WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
// 		return nil, err
// 	}
// 	log.Info("ConPayload unmarshaled, now check it")
// 	if conPayload == nil {
// 		return nil, log.Errorf("conPayload is nil ")
// 	}
// 	//fmt.Printf("GetConnectorPayload:860  --  ConnectorPayLoad = %s\n", spew.Sdump(conPayload))
// 	return conPayload, err
// }

// func CreateFhirQuery(r *http.Request) (string, error) {
// 	query := ""
// 	values := r.URL.Query()
// 	log.Debug3(fmt.Sprintf("CreateFhirQuery  values : %v", values))
// 	if len(values) < 1 {
// 		err := log.Errorf("Url.Querys are missing")
// 		return "", err
// 	}
// 	//fmt.Printf("\nCreateFhirQuery:713  --  Keys : %v\n\n", keys)
// 	for k, v := range values {
// 		log.Info("Key:  " + k + " => " + v[0])
// 		s := strings.TrimLeft(v[0], "[]")
// 		if query == "" {
// 			//for _, kv := range v {
// 			query = fmt.Sprintf("%s=%s", k, s)
// 			//}
// 		} else {
// 			query = fmt.Sprintf("%s&%s=%s", query, k, s)
// 		}
// 		log.Info("CreateFhirQuery = " + query)
// 	}
// 	return query, nil
// }

// func UseCache() bool {
// 	useCache := os.Getenv("USE_CACHE")
// 	log.Debug3("useCache: " + useCache)
// 	if useCache == "TRUE" {
// 		log.Debug3("USE_CACHE: true")
// 		return true
// 	}
// 	log.Debug3("Do Not Use Cache")
// 	return false
// }
