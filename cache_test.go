package main

import (
	//"context"
	"encoding/json"
	"net/http"

	fhir "github.com/dhf0820/fhir4"

	//"github.com/samply/golang-fhir-models/fhir-models/fhir"
	"fmt"
	//"go.mongodb.org/mongo-driver/bson"
	"os"
	"testing"

	jw_token "github.com/dhf0820/golangJWT"

	//"time"
	"io"
	"strings"

	"github.com/davecgh/go-spew/spew"
	//"github.com/dhf0820/token"
	//"github.com/dhf0820/uc_core/util"

	common "github.com/dhf0820/uc_common"
	log "github.com/dhf0820/vslog"
	. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//const pid = "Tbt3KuCY0B5PSrJvCu2j-PlK.aiHsu2xUjUM8bWpetXoB"

// //const ordercode = "8310-5"

// //const pid = "4342009"
const baseurl = "https://fhir-open.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d/"

func TestPatientCache(t *testing.T) {
	log.Debug3("Test run a FHIR query")
	//c := New(baseurl, "application/json+fhir")
	Convey("Run a query", t, func() {

		//caFhirId := "62f14531ba5395278cd530c4"
		//patient, err := c.GetPatient("375", caFhirId)
		jwt, payload, err := jw_token.CreateTestJWToken("10s")
		So(err, ShouldBeNil)
		So(jwt, ShouldNotBeNil)
		So(payload, ShouldNotBeNil)
		jwt = "Bearer " + jwt
		req, err := http.NewRequest("GET", "system/640ba5e3bd4105586a6dda74/Patient?patient=12724067", nil)
		So(err, ShouldBeNil)
		So(req, ShouldNotBeNil)
		req.Header.Set("Authorization", jwt)
		cp := CreateCP(false)
		cpb, err := json.Marshal(cp)
		So(err, ShouldBeNil)
		cps := string(cpb)
		rc := io.NopCloser(strings.NewReader(cps))
		req.Body = rc
		bundle, err := PatientSearch(cp, "family=sm", jwt)
		So(err, ShouldBeNil)
		So(bundle, ShouldNotBeNil)
		// data, err := c.Query("Patient/12724066")
		// So(err, ShouldBeNil)
		// So(data, ShouldNotBeNil)
		pat, err := fhir.UnmarshalPatient(bundle.Entry[0].Resource)
		So(err, ShouldBeNil)
		So(pat, ShouldNotBeNil)
		fmt.Printf("PatientSearch returned: %s\n", spew.Sdump(pat))
	})
}

// func TestDocumentReferenceCacheBundle(t *testing.T) {
// 	fmt.Printf("Test run a FHIR query\n")
// 	c := New(baseurl)
// 	Convey("Run a query", t, func() {
// 		//cerFhirId := "62f1c5dab3070d0b40e7aac1"
// 		//caFhirId := "62f14531ba5395278cd530c4"
// 		newToken, payload, err := createJWT()
// 		So(err, ShouldBeNil)
// 		So(newToken, ShouldNotBeNil)
// 		So(payload, ShouldNotBeNil)
// 		//req.Header.Set("AUTHORIZATION", newToken)
// 		bundle, err := c.DocumentReferenceSearch("patient=12724066", newToken) //family=smart")
// 		So(err, ShouldBeNil)
// 		So(bundle, ShouldNotBeNil)
// 		// data, err := c.Query("Patient/12724066")
// 		// So(err, ShouldBeNil)
// 		// So(data, ShouldNotBeNil)
// 		pat, err := fhir.UnmarshalPatient(bundle.Entry[0].Resource)
// 		So(err, ShouldBeNil)
// 		So(pat, ShouldNotBeNil)
// 		//fmt.Printf("PatientSearch returned: %s\n", spew.Sdump(pat))
// 	})
// }

func TestDocumentReferenceCache(t *testing.T) {
	log.Debug3("--  Test FHIR adding DocumentReference to query\n")
	//c := New(baseurl)
	Convey("Run a query", t, func() {
		os.Setenv("CONFIG_ADDRESS", "http://localhost:30300/api/rest/v1")
		conf, err := GetServiceConfig("uc_ca3", "local_test", "test") //GetConfig("delivery", "test")
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)
		//c := New(baseurl)
		err = os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!") //util.RandomString(32))
		So(err, ShouldBeNil)
		maker, err := jw_token.NewJWTMaker(os.Getenv("ACCESS_SECRET"))
		So(err, ShouldBeNil)
		So(maker, ShouldNotBeNil)
		os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
		jwt, payload, err := jw_token.CreateTestJWToken("10s")
		So(err, ShouldBeNil)
		So(jwt, ShouldNotBeNil)
		So(payload, ShouldNotBeNil)
		userId := "62d0af5dec383ade03a96b7e"

		//newToken := "Bearer " + jwt
		patientId := "12724066"
		//caFhirId := "62f14531ba5395278cd530c4"
		//cerFhirId := "62f1c5dab3070d0b40e7aac1"
		cerSysCfgId := "640ba5e3bd4105586a6dda74"
		//fhirSystem, err := GetFhirSystem(cerFhirId)
		cerSysCfgID, err := primitive.ObjectIDFromHex(cerSysCfgId)
		So(err, ShouldBeNil)
		So(cerSysCfgID, ShouldNotBeNil)
		cerSysCfg, err := GetSystemConfigById(cerSysCfgID)
		So(err, ShouldBeNil)
		So(cerSysCfg, ShouldNotBeNil)
		cc := common.ConnectorConfig{}
		cp := common.ConnectorPayload{}
		cc.ID, _ = primitive.ObjectIDFromHex("62f1c5dab3070d0b40e7aac1")
		cc.Name = "uc_ca3"
		cc.Version = "local_test"
		cc.CacheUrl = "http://uc_cache:9200"
		// "cacheurl" : "http://uc_cache:9200",
		// "cache_url" : "http://uc_cache:9200"
		data := []*common.KVData{}
		cacheServer := common.KVData{}
		cacheServer.Name = "cacheServer"
		cacheServer.Value = "http://universalcharts.com:30201"
		data = append(data, &cacheServer)
		hostServer := common.KVData{}
		hostServer.Name = "cacheHost"
		hostServer.Value = "http://ucCache:9200"
		data = append(data, &hostServer)
		cc.Data = data
		//cp.FhirSystem = fhirSystem
		cp.ConnectorConfig = &cc
		//FindResource(fhirSystem *common.FhirSystem, url, resource, userId, query)
		cnt, bundle, hdr, err := FindResource(&cp, "DocumentReference", userId, "DocumentReference?patient="+patientId, jwt)

		// //data, err := c.Query("DocumentReference?patient=12724066")
		So(err, ShouldBeNil)
		So(cnt, ShouldNotEqual, 0)
		So(hdr, ShouldNotBeNil)
		So(bundle, ShouldNotBeNil)
		So(len(bundle.Entry), ShouldEqual, cnt)
		fmt.Printf("Count = %d\n", cnt)
		fmt.Printf("Header : %s\n", spew.Sdump(hdr))
		// data, err = c.Query("DocumentReference?patient=12724066")
		// So(err, ShouldBeNil)
		// So(data, ShouldNotBeNil)
		// drBundle, err := fhir.UnmarshalBundle(bundle)
		// So(err, ShouldBeNil)
		// So(bundle, ShouldNotBeNil)
		doc, err := fhir.UnmarshalDocumentReference(bundle.Entry[0].Resource)
		So(err, ShouldBeNil)
		So(doc, ShouldNotBeNil)
		docId := *doc.Id
		log.Debug3("Number of DocRef: " + fmt.Sprint(len(bundle.Entry)))
		log.Debug3("Id of first Document: " + docId)
		// err = CacheResource(context.Background(), "queryId", userId, patientId, fhirSystem, &doc, "DocumentReference", docId)
		// So(err, ShouldBeNil)
	})
}

func TestDocumentReferenceBundleCache(t *testing.T) {
	fmt.Printf("\n\n\n\ncacheTest:68  --  Test FHIR ading DocumentReference to query\n")
	//c := New(baseurl)
	Convey("Run a query", t, func() {
		jwt, payload, err := jw_token.CreateTestJWToken("10s")
		So(err, ShouldBeNil)
		So(jwt, ShouldNotBeNil)
		So(payload, ShouldNotBeNil)

		// newToken, payload, err := createJWT()
		// So(err, ShouldBeNil)
		// So(newToken, ShouldNotBeNil)
		// So(payload, ShouldNotBeNil)
		//req.Header.Set("AUTHORIZATION", newToken)
		os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:30300/api/rest/v1")
		conf, err := GetServiceConfig("uc_ca3_conn", "local_test", "test") //GetConfig("delivery", "test")
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)
		//c := New(baseurl)
		//caFhirId := "62f14531ba5395278cd530c4"
		//cerFhirId := "62f1c5dab3070d0b40e7aac1"
		// fhirSystem, err := GetFhirSystem(cerFhirId)
		// So(err, ShouldBeNil)
		// So(fhirSystem, ShouldNotBeNil)
		// bundle, err := c.Query("DocumentReference?patient=63ed93c8bd78ae6b013a502b", jwt)
		// So(err, ShouldBeNil)
		// So(bundle, ShouldNotBeNil)
		// bundle, err := fhir.UnmarshalBundle(data)
		// So(err, ShouldBeNil)
		// So(bundle, ShouldNotBeNil)

		// doc, err := fhir4.UnmarshalDocumentReference(bundle.Entry[0].Resource)
		// So(err, ShouldBeNil)
		// So(doc, ShouldNotBeNil)
		//docId := *doc.Id
		// patientId := "12724066"
		// userId := "62d0af5dec383ade03a96b7e"
		// resourceType := "DocumentReference"
		//log.Debug3("--  Number of DocRefs: "+ fmt.Sprint(len(bundle.Entry)))
		//fmt.Printf("Id of first Document: %s\n", docId)
		//fmt.Printf("Document[0] returned: %s\n", spew.Sdump(doc))
		// CacheResourceBundleElements(context.Background(), userId,
		// 	patientId, fhirSystem, &bundle, resourceType)
	})
}
func TestGetCache(t *testing.T) {
	fmt.Printf("Test Get Cache for page")
	//c := New(baseurl)
	Convey("Run a query", t, func() {
		os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:20100/api/rest/v1")
		conf, err := GetServiceConfig("uc_ca3", "linode", "test") //GetConfig("delivery", "test")
		//conf, err := Initialize()
		So(err, ShouldBeNil)
		So(conf, ShouldNotBeNil)
		Convey("Run a query", func() {
			// pageId := 5
			// //queryId := "62ddb9f691f15a1e2d5206f7"
			// queryId := "62ddcd3261024840e7244591"
			// startTime := time.Now()
			// total, bundle, header, err := GetCache(queryId, pageId)
			// fmt.Printf("GetCache elapsed Time: %s\n", time.Since(startTime))
			// //
			// //fmt.Printf("Error: %s\n", err.Error())
			// So(err, ShouldBeNil)
			// So(header, ShouldNotBeNil)
			// So(bundle, ShouldNotBeNil)
			// So(total, ShouldNotEqual, 0)
			// fmt.Printf("GetCache Returned header returned: %s\n", spew.Sdump(header))
		})
	})
}

// func TestGetDocumentReferenceCachePage(t *testing.T) {
// 	fmt.Printf("Test Get Cache for page")
// 	//c := New(baseurl)
// 	Convey("Run a query", t, func() {
// 		os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:20100/api/rest/v1")
// 		conf, err := GetServiceConfig("uc_ca3", "linode", "test") //GetConfig("delivery", "test")
// 		So(err, ShouldBeNil)
// 		So(conf, ShouldNotBeNil)
// 		Convey("Query for page 1  of num per page of 3", func() {
// 			pageId := int64(1)
// 			perPage := int64(1)                  //count
// 			userId := "62d0af5dec383ade03a96b7e" //e"
// 			docs, err := GetDocumentReferenceCachePage(userId, perPage, pageId)
// 			So(err, ShouldBeNil)
// 			So(docs, ShouldNotBeNil)
// 			fmt.Printf("GetCache Returned resource: %s\n", spew.Sdump(docs))
// 		})
// 	})
// }

// func TestGetObservationCachePage(t *testing.T) {
// 	fmt.Printf("Test Get Observation Cache for page")
// 	//c := New(baseurl)
// 	Convey("Run a query", t, func() {
// 		os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:20100/api/rest/v1")
// 		conf, err := GetServiceConfig("uc_ca3", "linode", "test")
// 		So(err, ShouldBeNil)
// 		So(conf, ShouldNotBeNil)
// 		Convey("Query for page 1  of num per page of 2", func() {
// 			pageId := int64(1)
// 			perPage := int64(2) //count
// 			userId := "62d0af5dec383ade03a96b7e"
// 			obs, err := GetObservationCachePage(userId, perPage, pageId)
// 			So(err, ShouldBeNil)
// 			So(obs, ShouldNotBeNil)
// 			fmt.Printf("GetObservationCache Returned resource: %s\n", spew.Sdump(obs))
// 		})
// 	})
// }
