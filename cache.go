package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	//"os"

	//"strconv"
	//"github.com/davecgh/go-spew/spew"

	//"strings"
	//"time"

	//"time"

	"github.com/davecgh/go-spew/spew"
	common "github.com/dhf0820/uc_core/common"
	log "github.com/dhf0820/vslog"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//"github.com/sirupsen/logrus"
	//"github.com/samply/golang-fhir-models/fhir-models/fhir"
	fhir "github.com/dhf0820/fhir4"
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

type ResourceResponse struct {
	Status       int                 `json:"status"`
	Message      string              `json:"message"`
	ResourceType string              `json:"resourceType"`
	PageNumber   int                 `json:"pageNumber"`
	TotalPages   int64               `json:"totalPages"`
	CountInPage  int                 `json:"countInPage"`
	BundleId     string              `json:"bundleId"`
	QueryId      string              `json:"queryId"`
	Header       *common.CacheHeader `json:"header"`
	Bundle       *fhir.Bundle        `json:"bundle"`
	Resource     interface{}         `json:"resource"`
}

func CacheResourceBundleElements(ctx context.Context, userId,
	patientId string, sysConfig *common.SystemConfig, bundle *fhir.Bundle,
	resourceType string) error {
	log.Debug3("CacheResourceBundleElements: --  Is not implemented")
	return log.Errorf("CacheResourceBundleElements:  --  Not implemented")

	// var lastError error
	// fmt.Printf("\n\n\nThere are %d documents to cache\n\n\n", len(bundle.Entry))
	// for _, entry := range bundle.Entry {
	// 	//fmt.Printf("Caching document  = %s\n", spew.Sdump(entry))
	// 	doc, err := fhir.UnmarshalDocumentReference(entry.Resource)
	// 	if err != nil {
	// 		log.Error(fmt.Sprintf("cacheResource error on resource: %s Patient: %s resourceiId: %s  userId: %s err= %s",
	// 			resourceType, patientId, *entry.Id, userId, err.Error()))
	// 		lastError = err
	// 	} else {
	// 		docId := doc.Id
	// 		err = CacheResource(ctx, "", userId, patientId, sysConfig, doc, resourceType, *docId)
	// 		if err != nil {
	// 			log.Error(fmt.Sprintf("cacheResource error on resource: %s Patient: %s resourceiId: %s  userId: %s err= %s",
	// 				resourceType, patientId, *entry.Id, userId, err.Error()))
	// 			lastError = err
	// 		}
	// 	}
	// }
	// return lastError
}

// // CacheResourceBundleAndEntries: accepts a cacheBundle and JWToken, submiting it to the caching system returning the QueryId and err
// // sends bundle to cache which caches the Bundle  in BundleCache, then caches each entry in ResourceCacheCaches both the bundle and the individual entries cached in
func CacheResourceBundleAndEntries(cbdl *common.CacheBundle, token string, page int64) (int, error) {

	fmt.Printf("\n\n$$$$$$$$$$ Entering CaheResourceBundleAndEntries Use CacheViaCore instead $$$$$$\n")
	return 0, log.Errorf("CaheResourceBundleAndEntries is retured. Use CacheViaCore instead")
}

// 	log.Debug2("caching resource type: " + cbdl.Header.ResourceType)
// 	// if !UseCache() {
// 	// 	return -1, errors.New("NO CACHE")
// 	// }
// 	log.Info("Starting CacheResourceBundleAndEnteries. Page: " + fmt.Sprint(page))
// 	fmt.Println()
// 	//log.Debug3("cbdl:  " + spew.Sdump(cbdl))
// 	header := *cbdl.Header
// 	//log.Debug3(fmt.Sprintf("--  Starting for ResourceType: %s  Page: %d\n", header.ResourceType, page))
// 	//fmt.Printf("CacheResourceBundleAndEntries:77  -- Header = %s\n", spew.Sdump(header))
// 	log.Info(fmt.Sprintf("-- CashBase: %s\n", header.CacheUrl))

// 	//log.Debug3("cbdl:  " + spew.Sdump(cbdl))
// 	//fhirSystem := header.FhirSystem
// 	// fhirSystem, err := GetFhirSystem(header.FhirSystem.Hex())
// 	// if err != nil {
// 	// 	log.Errorf("getFhrSystem in CacheResourceBundleAndEntries failed: %v", err.Error())
// 	// 	return
// 	//fmt.Printf("CacheResourceBundleAndEntries:85  --  header = %s\n", spew.Sdump(header))
// 	//CacheServer := "http://192.168.1.117:30201"
// 	//GetDataByName()
// 	header.PageId = int(page)
// 	fmt.Printf("\n\n")
// 	log.Debug2("Header: " + spew.Sdump(header))
// 	//log.Debug2("cbdl.Header: " + spew.Sdump(cbdl.Header))
// 	//header.CacheUrl = fmt.Sprintf("%s/ResourceCache/%s", CacheServer, header.QueryId)
// 	//fmt.Printf("CacheResourceBundleAndEntries:89  --  CacheUrl = %s\n", header.CacheUrl)
// 	//log.Debug3("CacheResourceBundleAndEntries --  Number of Entries: " + fmt.Sprint(len(cbdl.Bundle.Entry)))
// 	log.Debug2("Marshaling cbdl")
// 	cacheBundle, err := json.Marshal(cbdl)
// 	if err != nil {
// 		err = log.Errorf("-- Error marshaling CacheBundle into json: " + err.Error())
// 		fmt.Printf("%s\n", err.Error())
// 		return 0, err
// 	}
// 	//log.Debug3("CacheResourceBundleAndEntries  --  CacheBundle: " + string(cacheBundle))
// 	//fmt.Printf("Send Status to: [%s]\n", statusURL)
// 	//fmt.Printf("CacheResourceBundleAndEntries:99  --  Using CoreUrl: %s\n", coreURL)
// 	//coreURL := cbdl.Header.FhirSystem.UcUrl + "/BundleTransaction"
// 	time.Sleep(3 * time.Second)
// 	//cacheURL := "http://UniversalCharts.com:30300/system/640ba5e3bd4105586a6dda74" + "/BundleTransaction"
// 	//cacheURL := cbdl.Header.CacheUrl + "/BundleTransaction"
// 	log.Info("CacheResourceBundleAndEntries  --  POST cacheURL: " + header.CacheUrl)
// 	cacheURL := "http://UniversalCharts.com:30300/system/640ba5e3bd4105586a6dda74" + "/BundleTransaction"
// 	fmt.Println()
// 	fmt.Printf("### cacheURL: %s  ###\n\n", cacheURL)
// 	fmt.Println()
// 	log.Info("cbdl.header:  " + spew.Sdump(cbdl.Header))
// 	log.Info("Preparing to post to core  --  POST cacheURL: " + cacheURL)
// 	req, _ := http.NewRequest("POST", cacheURL, bytes.NewBuffer(cacheBundle))
// 	req.Header.Set("Accept", "application/json")
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", token)
// 	req.Header.Set("Resource", cbdl.Header.ResourceType)
// 	token = "Bearer " + token
// 	//log.Debug3("CacheResourceBundleAndEntries  --  Token: " + token)
// 	client := &http.Client{}

// 	fmt.Printf("\n\n$$$$$$ sending %s to core$$$$ \n\n", cbdl.Header.ResourceType)
// 	log.Debug1(fmt.Sprintf("Sending to Core  --  Method: %s  URL: %s", req.Method, cacheURL))
// 	resp, err := client.Do(req)
// 	//log.Debug3("CacheResourceBundleAndEntries  --  resp: " + spew.Sdump(resp))
// 	//defer resp.Body.Close()
// 	if err != nil {
// 		err = log.Errorf("CacheResourceBundleAndEntries  -- Error uc_cache Request: " + err.Error())
// 		fmt.Println(err.Error())
// 		return 0, nil
// 	}
// 	if resp.StatusCode < 200 || resp.StatusCode > 299 {
// 		err = log.Errorf(fmt.Sprintf("CacheResourceBundleAndEntries -- Invalid uc_cache Status: " + fmt.Sprint(resp.StatusCode) + " -- " + resp.Status))
// 		fmt.Println(err.Error())
// 		return 0, nil
// 	}
// 	log.Debug3("Bundle Sent to uc_core Successful")
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		err = fmt.Errorf("CacheResourceBundleAndEntries: ReadAllBody : error: " + err.Error())
// 		fmt.Println(err.Error())
// 		return 0, nil
// 	}
// 	bundleResp := common.BundleCacheResponse{}
// 	err = json.Unmarshal(body, &bundleResp)
// 	if err != nil {
// 		err = log.Errorf("CacheResourceBundleAndEntries: Unmarshal BundleResponse failed: " + err.Error())
// 		fmt.Println(err.Error())
// 		return 0, nil
// 	}
// 	//page = page + 1
// 	return int(page), nil
// 	//c := common.Facility{}
// 	//c.CoreUrl = "http://	"

// }

// <<<<<<< HEAD
func FinishCache(systemConfig *common.SystemConfig, queryID primitive.ObjectID, token string, onPage int, pageSize int) error {
	//systemId := systemConfig.ID.Hex()
	cacheURL := systemConfig.ConnectorConfig.CacheUrl + "FinishCache"
	//cacheURL := "http://UniversalCharts.com:30300/system/" + systemId + "/BundleTransaction"
	// =======
	// func FinishCache(systemConfig *common.SystemConfig, queryID primitive.ObjectID, token string, pageNum int, onPage int, pageSize int) error {
	// 	systemId := systemConfig.ID.Hex()
	// 	log.Debug3("@!!!SystemConfig: " + spew.Sdump(systemConfig))
	// 	log.Debug3("pageSize: " + fmt.Sprint(pageSize))
	// 	log.Debug3("onPage: " + fmt.Sprint(onPage))
	// 	log.Debug3("PageNum: " + fmt.Sprint(pageNum))
	// 	cacheURL := "http://UniversalCharts.com:30300/system/" + systemId + "/FinishCache"
	// >>>>>>> 2062832a583c454d91d8636cf44ef1efded64a16
	fmt.Println()
	fmt.Println()
	fmt.Println()
	log.Info("FinishCache  --  POST cacheURL: " + cacheURL)
	finishedCache := common.FinishCachePayload{}
	// cacheSavePayload.Bundle = bundle
	// cacheSavePayload.Option = option
	//finishedCache.PageNum = pageNum
	finishedCache.OnPage = onPage
	finishedCache.PageSize = pageSize
	finishedCache.QueryId = queryID.Hex()
	log.Debug3("finishedCache to send to core: " + spew.Sdump(finishedCache))
	payload, err := json.Marshal(finishedCache)
	//bndl, err := bundle.MarshalJSON()
	if err != nil {
		err = log.Errorf("Marshal Bundle error: " + err.Error())
		return err
	}
	req, _ := http.NewRequest("POST", cacheURL, bytes.NewBuffer(payload))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	//req.Header.Set("Resource", *bundle.ResourceType)
	req.Header.Set("Final", "final")
	client := &http.Client{}
	//log.Debug3(fmt.Sprintf("CacheResourceBundleAndEntries  --  Method: %s  URL: %s", req.Method, cacheURL))
	fmt.Printf("\n\n@@@@@@@@@@@@@@@@ Send To Core @@@@@@\n")
	resp, err := client.Do(req)
	//log.Debug3("CacheResourceBundleAndEntries  --  resp: " + spew.Sdump(resp))
	//defer resp.Body.Close()
	if err != nil {
		err = log.Errorf("FinishCache  -- Error uc_cache Request: " + err.Error())

		fmt.Println(err.Error())
		return nil
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = log.Errorf(fmt.Sprintf("CacheResourceBundleAndEntries -- Invalid uc_cache Status: " + fmt.Sprint(resp.StatusCode) + " -- " + resp.Status))
		fmt.Println(err.Error())
		return nil
	}
	log.Debug3("FinishCache Sent to uc_core Successful")
	return nil
}

func CacheViaCore(bundle *fhir.Bundle, queryID primitive.ObjectID, resourceType, token string, cacheURL string, page int) error {
	//cacheURL := "http://UniversalCharts.com:30300/system/640ba5e3bd4105586a6dda74" + "/BundleTransaction"
	fmt.Printf("\n\n@@@@@@@@@@@@@@@@@ CacheViaCore @@@@@@@@@@@@@@@@@\n")
	log.Debug2("ResourceType: " + resourceType)
	if bundle == nil {
		return log.Errorf("Bundle is nil")
	}
	queryId := queryID.Hex()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	cacheURL = cacheURL + "BundleTransaction"
	//log.Info("CacheViaCore  --  POST cacheURL: " + cacheURL)
	cacheSavePayload := common.CacheSavePayload{}
	// <<<<<<< HEAD
	// =======
	// 	cacheSavePayload.Bundle = bundle
	// 	cacheSavePayload.Option = option
	// 	cacheSavePayload.PageNum = page
	// 	cacheSavePayload.ResourceType = *bundle.ResourceType
	// >>>>>>> 2062832a583c454d91d8636cf44ef1efded64a16

	cacheSavePayload.PageNum = page
	cacheSavePayload.ResourceType = resourceType
	cacheSavePayload.QueryId = queryId
	log.Debug2("cacheSavePayload without Bundle: " + spew.Sdump(cacheSavePayload))
	cacheSavePayload.Bundle = bundle
	payload, err := json.Marshal(cacheSavePayload)
	//bndl, err := bundle.MarshalJSON()
	if err != nil {
		err = log.Errorf("Marshal Bundle error: " + err.Error())
		return err
	}

	log.Debug1("CacheViaCore  --  POST cacheURL: " + cacheURL)
	req, _ := http.NewRequest("POST", cacheURL, bytes.NewBuffer(payload))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	req.Header.Set("Resource", resourceType)
	//req.Header.Set("Final", final)
	//token = "Bearer " + token
	//log.Debug3("CacheResourceBundleAndEntries  --  Token: " + token)
	client := &http.Client{}
	//log.Debug3(fmt.Sprintf("CacheResourceBundleAndEntries  --  Method: %s  URL: %s", req.Method, cacheURL))
	fmt.Printf("\n\n\n!!!!!!!!!! Sending %s page: %d to core!!!!!!!!!!!!\n\n", resourceType, cacheSavePayload.PageNum)
	log.Debug1("!!!!! Sending to core !!!!!")
	resp, err := client.Do(req)
	//log.Debug3("CacheResourceBundleAndEntries  --  resp: " + spew.Sdump(resp))
	//defer resp.Body.Close()
	if err != nil {
		err = log.Errorf("CacheResourceBundleAndEntries  -- Error uc_cache Request: " + err.Error())
		fmt.Println(err.Error())
		return nil
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = log.Errorf(fmt.Sprintf("CacheResourceBundleAndEntries -- Invalid uc_cache Status: " + fmt.Sprint(resp.StatusCode) + " -- " + resp.Status))
		fmt.Println(err.Error())
		return nil
	}
	log.Debug1("Bundle Sent to uc_core Successful")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("CacheResourceBundleAndEntries: ReadAllBody : error: " + err.Error())
		fmt.Println(err.Error())
		return nil
	}
	bundleResp := common.BundleCacheResponse{}
	err = json.Unmarshal(body, &bundleResp)
	if err != nil {
		err = log.Errorf("CacheResourceBundleAndEntries: Unmarshal BundleResponse failed: " + err.Error())
		fmt.Println(err.Error())
		return nil
	}

	//page = page + 1
	return nil
}

// func CacheResource(ctx context.Context, queryId, userId,
// 	patientId string, fhirSystem *common.FhirSystem, resource json.RawMessage,
// 	resourceType string, resourceId string) error {
// 	var collection *mongo.Collection
// 	var err error
// 	//collection, err = GetCollection(strings.ToLower(resourceType) + "_cache")
// 	fmt.Printf("CacheResource:167 --  starting\n")
// 	collection, err = GetCollection("resource_cache")
// 	if err != nil {
// 		return err
// 	}
// 	//userID, _ := primitive.ObjectIDFromHex(userId)
// 	hdr := common.CacheHeader{}
// 	cache := common.ResourceCache{}
// 	timeNow := time.Now()
// 	hdr.CreatedAt = &timeNow
// 	hdr.UserId = userId
// 	hdr.PatientId = patientId
// 	hdr.ResourceType = resourceType
// 	hdr.ResourceId = resourceId
// 	hdr.FhirSystem = fhirSystem
// 	hdr.FhirId = fhirSystem.ID.Hex()
// 	hdr.CacheUrl = fmt.Sprintf("%s/ResourceCache/%s", fhirSystem.UcUrl, queryId)
// 	fmt.Printf("CacheResource:249  -- CacheUrl = %s\n", hdr.CacheUrl)
// 	fmt.Printf("CacheResource:250  -- Header = %s\n", spew.Sdump(hdr))
// 	cache.ID = primitive.NewObjectID()
// 	fmt.Printf("ID: %s\n", cache.ID.Hex())
// 	cache.ResourceType = resourceType
// 	cache.Resource = resource
// 	cache.Header = &hdr

// 	_, err = collection.InsertOne(ctx, cache)
// 	if err != nil {
// 		err = fmt.Errorf("Insert ResourceCache InsertOne failed: %v", err.Error())
// 		return err
// 	}
// 	return nil
// }

//todo: go routine to  cache each resource in the bundle.
//pass the header already filled in and blank the page number
// Question, can you have  Go ROUTINE AND MAIN WORKING ON DIFFERENT COLLECTIONS
// func Insert(ctx context.Context, cbdl *common.CacheBundle, token string) error {
// 	fmt.Printf("\n$$$Insert:i78 -  %s - queryId: %s page: %d\n", cbdl.Header.ResourceType, cbdl.Header.QueryId, cbdl.Header.PageId)
// 	// if cbdl.Header.ResourceType != "Patient" { // Only cache the non Patient Resources
// 	// 	CacheResourceBundleAndEntries(cbdl)
// 	// }
// 	fmt.Printf("\n\n\n\n$$$Insert:360  calling CacheResourceBundleAndEntries\n")
// 	fmt.Printf("Insert:361  --  cbdl.Header : %s\n", spew.Sdump(cbdl.Header))
// 	CacheResourceBundleAndEntries(cbdl, token)
// 	return nil

// //entry := cbdl.Bundle.Entry[0]

// //fmt.Printf("Insert:155 - Entry[0] = %s\n", spew.Sdump(entry.Resource))

// collection, err := GetCollection("cache_bundle")
// if err != nil {
// 	return err
// }
// //fmt.Printf("Insert:157 -- header = %s\n", spew.Sdump(cbdl.Header))

// timeNow := time.Now()
// cbdl.Header.CreatedAt = &timeNow
// //data.UpdatedAt = data.CreatedAt
// cbdl.ID = primitive.NewObjectID()
// fmt.Printf("ID: %s\n", cbdl.ID.Hex())
// cbdl.QueryId = cbdl.Header.QueryId
// cbdl.PageId = cbdl.Header.PageId
// fmt.Printf("cache.Insert:315 -- cbdl.QueryId: %s, page: %d Number on Page: %d\n", cbdl.QueryId, cbdl.PageId, len(cbdl.Bundle.Entry))
// //fmt.Printf("Inserting: %s\n", spew.Sdump(cbdl))

// _, err = collection.InsertOne(ctx, cbdl)
// if err != nil {
// 	err = fmt.Errorf("Insert CacheBundle InsertOne failed: %v", err.Error())
// 	return err
// }
// return nil
//}

type FhirResource interface {
}

func GetResourceCachePage(resource, userId string, perPage, pageNum int64) ([]Interface, error) {
	return nil, errors.New("GetResourceCachePage: Not implemented")
	// collection, err := GetCollection("resource_cache")
	// if err != nil {
	// 	return nil, err
	// }
	// if perPage == 0 {
	// 	perPage = 10
	// }
	// var results []Interface
	// offset := int64((pageNum - 1) * perPage)
	// //query := bson.D{{"Header.UserId", userId}, {"ResourceType", resource}}
	// query := bson.D{{"header.resourceType", resource}}
	// fmt.Printf("GetResourceCachePage:314  --  query = %v\n", query)
	// cur, err := collection.Find(context.Background(), query, options.Find().SetSkip(offset).SetLimit(perPage))
	// i := 0
	// if cur.RemainingBatchLength() > int(0) {
	// 	for cur.Next(context.Background()) {
	// 		i++
	// 		var data common.ResourceCache //fhir.DocumentReference
	// 		fmt.Printf("Decoding\n")
	// 		err := cur.Decode(&data)
	// 		fmt.Printf("Decode finished\n")
	// 		if err != nil {
	// 			fmt.Printf("Find CacheResource decode failed: %v", err)
	// 			return nil, fmt.Errorf("find CacheResource decode failed: %v", err)
	// 		} //
	// 		//fmt.Printf("Appending document %s\n", spew.Sdump(d))
	// 		results = append(results, data.Resource)
	// 	}
	// } else {
	// 	//fmt.Printf("Skipping cursor, returning nil, nil\n")
	// 	//no existing document
	// 	return nil, fmt.Errorf("no %s matching %v found", resource, query)
	// }
	// return results, nil
}

// 	return results, err
// switch resource {
// case "DocumentReference":
// 	cur, err := collection.Find(context.Background(), query, options.Find().SetSkip(offset).SetLimit(perPage))
// 	i := 0
// 	results := []fhir.DocumentReference{}
// 	if cur.RemainingBatchLength() > int(0) {
// 		for cur.Next(context.Background()) {
// 			i++
// 			var data fhir.DocumentReference
// 			fmt.Printf("Decoding\n")
// 			err := cur.Decode(&data)
// 			fmt.Printf("Decode finished\n")
// 			if err != nil {
// 				fmt.Printf("Find CacheResource decode failed: %v", err)
// 				return nil, fmt.Errorf("find CacheResource decode failed: %v", err)
// 			}
// 			//fmt.Printf("Appending document %s\n", spew.Sdump(d))
// 			results = append(results, data)
// 		}
// 	} else {
// 		//fmt.Printf("Skipping cursor, returning nil, nil\n")
// 		//no existing document
// 		return nil, fmt.Errorf("no DocumentReferences matching %v found", query)
// 	}
// 	return results, err
// case "Observation":
// 	cur, err := collection.Find(context.Background(), query, options.Find().SetSkip(offset).SetLimit(perPage))
// 	i := 0

// 	if cur.RemainingBatchLength() > int(0) {
// 		for cur.Next(context.Background()) {
// 			i++
// 			var data interface{}
// 			fmt.Printf("Decoding\n")
// 			err := cur.Decode(&data)
// 			fmt.Printf("Decode finished\n")
// 			if err != nil {
// 				fmt.Printf("Find CacheResource decode failed: %v", err)
// 				return nil, fmt.Errorf("find CacheResource decode failed: %v", err)
// 			}
// 			//fmt.Printf("Appending document %s\n", spew.Sdump(d))
// 			results = append(results, data)
// 		}
// 	} else {
// 		//fmt.Printf("Skipping cursor, returning nil, nil\n")
// 		//no existing document
// 		return nil, fmt.Errorf("no Observations matching %v found", query)

// 	}
// 	return results, err
// case "Condition":
// 	cur, err := collection.Find(context.Background(), query, options.Find().SetSkip(offset).SetLimit(perPage))
// 	i := 0

// 	if cur.RemainingBatchLength() > int(0) {
// 		for cur.Next(context.Background()) {
// 			i++
// 			var data interface{}
// 			fmt.Printf("Decoding\n")
// 			err := cur.Decode(&data)
// 			fmt.Printf("Decode finished\n")
// 			if err != nil {
// 				fmt.Printf("Find CacheResource decode failed: %v", err)
// 				return nil, fmt.Errorf("find CacheResource decode failed: %v", err)
// 			}
// 			//fmt.Printf("Appending document %s\n", spew.Sdump(d))
// 			results = append(results, data)
// 		}
// 	} else {
// 		//fmt.Printf("Skipping cursor, returning nil, nil\n")
// 		//no existing document
// 		return nil, fmt.Errorf("no Conditions matching %v found", query)

// 	}
// 	return results, err
// }

//TODO: Call Core to get cacheBundlePage
//func GetCache(queryId string, pageId int) (int64, *fhir.Bundle, *common.CacheHeader, error) {
// collection, err := GetCollection("cache_bundle")
// if err != nil {
// 	return 0, nil, nil, err
// }
// fmt.Printf("\n$$$GetCache: 70 - Using mongo database: %s\n", DB.DatabaseName)
//oid, err := primitive.ObjectIDFromHex(id)
// if err != nil {
// 	return nil, nil, fmt.Errorf("invalid FhirId: [%s]", id)
// }
//pageNum, err := strconv.Atoi(pageId)
// if err != nil {
// 	return nil, nil, fmt.Errorf("pageId invalid: %s", err.Error())
// }
//oid, _ := primitive.ObjectIDFromHex("62ddb9f891f15a1e2d5206fd")
//query := bson.D{}
//query := bson.D{{"header.queryId", queryId}, {"header.pageId", pageId}}
//total, err := TotalCacheForQuery(queryId)
//query := bson.D{{"queryId", queryId}, {"pageId", pageId}}
// var query []bson.M
// if queryId != "" {
// 	query = append(query, bson.M{"_id": oid})
// }
// if pageId != 0 {
// 	query = append(query, bson.M{"pageId": pageId})
// }

// filter := bson.D{{"header.queryId", queryId}, {"header.pageId", pageNum}}
// filterM := bson.M{"header.queryId": queryId, "header.pageId": pageNum}
//fhirConfig := &common.FhirConfig{}
//fmt.Printf("   Now Calling GetCache:106 FindOne CacheBundle: bson.D %v\n", query)
//cacheBundle := []*CacheBundle{}
// cacheBundle := &common.CacheBundle{}
// err = collection.FindOne(context.Background(), query).Decode(cacheBundle)
// if err != nil {
// 	fmt.Printf("GetCache:111 FindOne %v NotFound\n", query)
// 	return 0, nil, nil, fmt.Errorf("GetCacheBundle:112  FindOne %v NotFound\n", query)
// }
// fmt.Printf("CacheBundle.Header: %s\n", spew.Sdump(cacheBundle.Header))
// if err != nil {
// 	fmt.Printf("   Now Calling GetCache:80  FindOne caheBundle bson.M %v\n", filterM)
// 	err = collection.FindOne(context.Background(), filterM).Decode(cacheBundle)
// }
// if err != nil {
// 	fmt.Printf("GetCache:115 FindOne %v NotFound\n", query)
// 	return nil, nil, fmt.Errorf("GetCacheBundle:116  FindOne %v NotFound\n", query)
// }

//return cacheBundle[0].Bundle, cacheBundle[0].Header, nil
//return total, cacheBundle.Bundle, cacheBundle.Header, err
//}

func TotalCacheForQuery(queryId primitive.ObjectID) (int64, error) {
	//TODO: Call Core to get cache status
	total := int64(99999)
	// startTime := time.Now()
	// collection, err := GetCollection("cache_bundle")
	// if err != nil {
	// 	return 0, err
	// }
	// fmt.Printf("GetCollection elapsed time: %s\n", time.Since(startTime))
	// startTime = time.Now()
	// filter := bson.D{{"queryId", queryId}}
	// total, err := collection.CountDocuments(context.TODO(), filter)
	// fmt.Printf("Count Cache elspded time: %s\n", time.Since(startTime))
	return total, nil
}

// func GetDocumentReferenceCachePage(userId string, perPage, pageNum int64) ([]fhir.DocumentReference, error) {
// 	collection, err := GetCollection("resource_cache")
// 	if err != nil {
// 		return nil, err
// 	}
// 	if perPage == 0 {
// 		perPage = 10
// 	}
// 	//var results []common.CacheResource
// 	//var results []common.ResourceCache
// 	offset := int64((pageNum - 1) * perPage)
// 	//query := bson.D{{"Header.UserId", userId}, {"ResourceType", resource}}
// 	query := bson.D{{"header.userId", userId}, {"resourceType", "DocumentReference"}}
// 	fmt.Printf("GetDocumentReferencePage:497  --  query = %v\n", query)
// 	cur, err := collection.Find(context.Background(), query, options.Find().SetSkip(offset).SetLimit(perPage))
// 	if err != nil {
// 		return nil, fmt.Errorf("DocumentReference query[%v] failed err=%s\n", query, err.Error())
// 	}
// 	i := 0
// 	fmt.Printf("RemainingBatch = %d\n", cur.RemainingBatchLength())
// 	results := []fhir.DocumentReference{}
// 	if cur.RemainingBatchLength() > int(0) {
// 		for cur.Next(context.Background()) {
// 			i++
// 			var resData common.ResourceCache
// 			err := cur.Decode(&resData)
// 			fmt.Printf("Decode finished : %s\n", spew.Sdump(resData.Resource))
// 			if err != nil {
// 				fmt.Printf("Find GetDocumentReferenceCachePage:507 decode failed: %v", err)
// 				return nil, fmt.Errorf("find GetDocumentReferenceCachePage:508  decode failed: %v", err)
// 			} //
// 			dr := fhir.DocumentReference{}
// 			res, err := bson.Marshal(resData.Resource)
// 			if err != nil {
// 				fmt.Printf("bson.Marshal error: %s\n", err.Error())
// 			}
// 			err = bson.Unmarshal(res, &dr)
// 			if err != nil {
// 				fmt.Printf("bson.Unmarshal error: %s\n", err.Error())
// 			}
// 			results = append(results, dr)
// 		}
// 	} else {
// 		//fmt.Printf("Skipping cursor, returning nil, nil\n")
// 		//no existing document
// 		return nil, fmt.Errorf("no DocumentReferencesmatching %v found", query)
// 	}
// 	fmt.Printf("documentRefernces: %s\n", spew.Sdump(results))
// 	return results, nil
// }

// func GetObservationCachePage(userId string, perPage, pageNum int64) ([]fhir.Observation, error) {
// 	collection, err := GetCollection("resource_cache")
// 	if err != nil {
// 		return nil, err
// 	}
// 	if perPage == 0 {
// 		perPage = 10
// 	}
// 	offset := int64((pageNum - 1) * perPage)
// 	query := bson.D{{"header.userId", userId}, {"resourceType", "Observation"}}
// 	fmt.Printf("GetObservationPage:545  --  query = %v\n", query)
// 	cur, err := collection.Find(context.Background(), query, options.Find().SetSkip(offset).SetLimit(perPage))
// 	if err != nil {
// 		return nil, fmt.Errorf("Observation query[%v] failed err=%s\n", query, err.Error())
// 	}
// 	i := 0
// 	results := []fhir.Observation{}
// 	if cur.RemainingBatchLength() > int(0) {
// 		for cur.Next(context.Background()) {
// 			i++
// 			var resData common.ResourceCache
// 			err := cur.Decode(&resData)
// 			if err != nil {
// 				err = fmt.Errorf("find GetObservationCachePage:560  decode failed: %v", err)
// 				return nil, err
// 			} //
// 			obs := fhir.Observation{}
// 			res, err := bson.Marshal(resData.Resource)
// 			if err != nil {
// 				fmt.Printf("bson.Marshal error: %s\n", err.Error())
// 			}
// 			err = bson.Unmarshal(res, &obs)
// 			if err != nil {
// 				fmt.Printf("bson.Unmarshal error: %s\n", err.Error())
// 			}
// 			results = append(results, obs)
// 		}
// 	} else {
// 		return nil, fmt.Errorf("no Observations matching %v found", query)
// 	}
// 	fmt.Printf("Observations: %s\n", spew.Sdump(results))
// 	return results, nil
// }

// func GetPatientIdentifier(patIdentifiers []fhir.Identifier, fsIdentifiers []*common.KVData, name string) (string, error) {
// 	var id *string
// 	for _, idType := range fsIdentifiers {
// 		if idType.Name == name {
// 			id = &idType.Value
// 			break
// 		}
// 	}
// 	if id == nil {
// 		return "", fmt.Errorf("Identifier type: %s is not registered", name)
// 	}
// 	for _, patId := range patIdentifiers {
// 		if patId.System == id {
// 			return *patId.Value, nil
// 		}
// 	}
// 	return "", fmt.Errorf("Identifier type: %s is invalid", name)
// }

// func CreateCacheHeaders(resourceType string, data json.RawMessage) (*common.CacheHeader,
// 	*common.ResourceHeader, error) {

// 	//for _, entry := range cbdl.Bundle.Entry {
// 	hdr := common.CacheHeader{}
// 	resHeader := common.ResourceHeader{}
// 	cache := common.ResourceCache{}
// 	cache.ID = primitive.NewObjectID()
// 	// if header.ResourceType == "cache:110  -- Patient" {
// 	// 	pat := fhir4.Patient{}
// 	// 	err := json.Unmarshal(entry.Resource, &pat)
// 	// 	if err != nil {
// 	// 		fmt.Printf("cache:110  -- Unmarshal err: %s\n", err.Error())
// 	// 		continue
// 	// 	} else {
// 	// 		fmt.Printf("cache:113  -- patient: %s\n", spew.Sdump(pat))
// 	// 		cache.Resource = entry.Resource
// 	// 	}
// 	// 	continue

// 	// }
// 	// cache.Resource = data  //entry.Resource
// 	// cache.ResourceType = resourceType
// 	//fmt.Printf("entry : %s\n", spew.Sdump(entry))
// 	//fmt.Printf("Entry = %d -- entry.id = %s\n", i, spew.Sdump(entry.Id))
// 	//byte := data  //entry.Resource
// 	//var res Interface
// 	switch resourceType {
// 	// case "Patient":
// 	// 	//fmt.Printf("GetReource:38  -  Patient= %s\n", spew.Sdump(pat))
// 	// 	pat, err := fhir4.UnmarshalPatient(byte)
// 	// 	if err != nil {
// 	// 		fmt.Printf("UnmarshalPatient err = %v\n", err)
// 	// 	} else {
// 	// 		hdr.PatientId = *pat.Id
// 	// 		hdr.ResourceId = *pat.Id

// 	// 		cache.Header = &hdr
// 	// 		cache.Resource = pat
// 	// 		_, err = collection.InsertOne(context.Background(), cache)
// 	// 		if err != nil {
// 	// 			log.Errorf("Insert ResourceCache InsertOne failed: %v", err.Error())
// 	// 			return
// 	// 		}
// 	// 	}
// 	case "DocumentReference":
// 		fmt.Printf("CreateCacheHeaders:743  --  Handle DocumentReference\n")
// 		doc, err := fhir4.UnmarshalDocumentReference(data)
// 		if err != nil {
// 			err = fmt.Errorf("CreateCacheHeaders:746  --  UnmarshalDocumentReference:nerr = %v", err)
// 			fmt.Printf("%s\n", err.Error())
// 			return nil, nil, err
// 		} else {
// 			hdr.ResourceType = "DocumentReference"
// 			parts := strings.Split(*doc.Subject.Reference, "/")
// 			hdr.PatientId = parts[1]
// 			hdr.ResourceId = *doc.Id
// 			fmt.Printf("CreateCacheHeaders:754  --  RecourceId = %s  PatientId = %s\n", hdr.ResourceId, hdr.PatientId)
// 			return &hdr, &resHeader, nil
// 		}
// 	case "Observation":
// 		fmt.Printf("CreateCacheHeaders:758  --  Handle Observation\n")
// 		obs, err := fhir4.UnmarshalObservation(data)
// 		if err != nil {
// 			err = fmt.Errorf("CreateCacheHeaders:761  --  UnmarshalObservation err = %v", err)
// 			fmt.Printf("%s\n", err.Error())
// 			return nil, nil, err
// 		}
// 		hdr.ResourceType = "Observation"
// 		parts := strings.Split(*obs.Subject.Reference, "/")
// 		hdr.PatientId = parts[1]
// 		hdr.ResourceId = *obs.Id
// 		fmt.Printf("CreateCacheHeaders:769  --  RecourceId = %s  PatientId = %s\n", hdr.ResourceId, hdr.PatientId)
// 		cache.Header = &hdr
// 		fmt.Printf("CreateCacheHeaders: cache.Header = %s\n", spew.Sdump(cache.Header))
// 		return &hdr, &resHeader, nil

// 	case "Condition":
// 		cond, err := fhir4.UnmarshalCondition(data)
// 		fmt.Printf("CreateCacheHeaders:776  --  Handle Condition\n")
// 		if err != nil {
// 			err = fmt.Errorf("CreateCacheHeaders:778  UnmarshalCondition err = %v", err)
// 			fmt.Printf("%s\n", err.Error())
// 			return nil, nil, err
// 		}
// 		hdr.ResourceType = "Condition"
// 		parts := strings.Split(*cond.Subject.Reference, "/")
// 		hdr.PatientId = parts[1]
// 		hdr.ResourceId = *cond.Id
// 		fmt.Printf("CreateCacheHeaders:786  --  RecourceId = %s  -  PatientId = %s\n", hdr.ResourceId, hdr.PatientId)
// 		return &hdr, &resHeader, nil

// 	case "Patient":
// 		pat, err := fhir4.UnmarshalPatient(data)
// 		fmt.Printf("CreateCacheHeaders:791  --  Handle Patient = %s\n", spew.Sdump(pat))
// 		if err != nil {
// 			err = fmt.Errorf("CreateCacheHeaders:793  UnmarshalPatient err = %v", err)
// 			fmt.Printf("%s\n", err.Error())
// 			return nil, nil, err
// 		}
// 		kvData := common.KVData{}
// 		kvData.Name = "name"
// 		kvData.Value = *pat.Name[0].Text
// 		resHeader.DisplayFields = append(resHeader.DisplayFields, kvData)
// 		fsIdentifiers := hdr.SystemCfg.Identifiers
// 		GetPatientIdentifier(pat.Identifier, fsIdentifiers, "mrn")
// 		hdr.ResourceType = "Patient"
// 		hdr.PatientId = *pat.Id
// 		fmt.Printf("CreateCacheHeaders:805  --  PatientId = %s\n", hdr.PatientId)
// 		hdr.ResourceId = *pat.Id
// 		return &hdr, &resHeader, nil

// 	default:
// 		err := fmt.Errorf("CreateCacheHeaders:810 unknown type of Resource: [%s]", resourceType)
// 		fmt.Printf("%s\n", err.Error())
// 		return nil, nil, err
// 	}
// }

// TODO: Call Cache Service to cache a single Resource
// func CacheResource(sysCfg *common.SystemConfig, resourceCache *common.ResourceCache, token string) (string, error) {
// 	coreURL := fhirSystem.UcUrl + "/Cache"
// 	byte, err := json.Marshal(resourceCache)
// 	if err != nil {
// 		err = fmt.Errorf("CacheResource:812 Marshal %s error: %s", resourceCache.ResourceType, err.Error())
// 		fmt.Printf("%s\n", err)
// 		return "", err
// 	}
// 	req, _ := http.NewRequest("POST", coreURL, bytes.NewBuffer(byte))
// 	req.Header.Set("Accept", "application/json")
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", token)
// 	client := &http.Client{}
// 	fmt.Printf("\nCacheResource:821  --  Calling core: %s\n", coreURL)
// 	resp, err := client.Do(req)
// 	defer resp.Body.Close()
// 	if err != nil {
// 		fmt.Println("CacheResource:825  --  Error Core Cache Request: ", err.Error())
// 		return "", err
// 	} else {
// 		fmt.Printf("CacheResource:827  --  Resource Sent to Core Successful\n")
// 	}
// 	if err != nil {
// 		log.Errorf("GetFhir:105  --  !!!fhir query returned err: %s\n", err)
// 		return "", err
// 	}
// 	//fmt.Printf("GetFhir:108  --  resp = %s\n", spew.Sdump(resp))
// 	if resp.StatusCode < 200 || resp.StatusCode > 299 {
// 		log.Errorf("GetFhir:110  --  returned error of %d - %s\n", resp.StatusCode, resp.Status)
// 		err = fmt.Errorf("%d|fhir:105 %s", resp.StatusCode, resp.Status)
// 		//log.Errorf("%s", err.Error())
// 		return "", err
// 	}

// 	return "", nil
// }

// func GetCachedResource(cp *common.ConnectorPayload, token string) (*common.ResourceCache, error) {
// 	if cp.System == nil {
// 		return nil, fmt.Errorf("CachedResource:681  --  FhirSystem is Nil")
// 	}
// 	coreURL := fmt.Sprintf("http://uc_cache:30201/%s/CacheEntry?query_id=%s", cp.System.ID.Hex(), cp.SavePayload.QueryId)

// 	req, _ := http.NewRequest("GET", coreURL, nil)
// 	req.Header.Set("Accept", "application/json")
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", token)
// 	client := &http.Client{}
// 	fmt.Printf("\nCachedResource:690  --  Calling core: %s\n", coreURL)
// 	resp, err := client.Do(req)
// 	defer resp.Body.Close()
// 	if err != nil {
// 		fmt.Println("CachedResource:694  --  Error Core Cache Request: ", err.Error())
// 		return nil, err
// 	} else {
// 		fmt.Printf("CachedResource:697  --  Resource Sent to Core Successful\n")
// 	}
// 	if err != nil {
// 		log.Errorf("CachedResource:700  --  !!!fhir query returned err: %s\n", err)
// 		return nil, err
// 	}
// 	//fmt.Printf("GetFhir:108  --  resp = %s\n", spew.Sdump(resp))
// 	if resp.StatusCode < 200 || resp.StatusCode > 299 {
// 		err = fmt.Errorf("CachedResource:705  --  returned error of %d - %s\n", resp.StatusCode, resp.Status)
// 		return nil, err
// 	}
// 	body, err = ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		err = fmt.Errorf("CachedResource:708  --  ReadAllBody : error: %s\n", err.Error())
// 		fmt.Printf("%s\n", err.Error())
// 		return nil, err
// 	}
// 	res := &common.ResourceCache{}
// 	err = json.Unmarshal(body, res)
// 	if err != nil {
// 		err = fmt.Errorf("CacheResourceBundleAndEntries:131 -- Unmarshal BundleResponse failed: %s", err.Error())
// 		fmt.Printf("%s\n", err.Error())
// 		return res, err
// 	}

// 	return res, nil
// }
