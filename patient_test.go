package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	//"github.com/davecgh/go-spew/spew"
	fhir "github.com/dhf0820/fhir4"
	//log "github.com/sirupsen/logrus"
	//. "github.com/smartystreets/goconvey/convey"
	"fmt"
	"os"
	"strings"
	"testing"

	"time"

	"github.com/dhf0820/token"
	common "github.com/dhf0820/uc_common"

	"github.com/davecgh/go-spew/spew"
	"github.com/dhf0820/uc_core/util"

	//log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestPostPatient(t *testing.T) {
	fmt.Printf("Test Add a new patient to Server")
	//c := New(baseurl)

	Convey("PostNewPatient", t, func() {
		fmt.Printf("\n\n\nPostNewPatient:  --  Start\n")

		os.Setenv("CONFIG_ADDRESS", "http://192.168.1.148:30300/api/rest/v1")
		_, err := GetServiceConfig("ca3_conn", "local_test", "test") //GetConfig("delivery", "test")
		//fmt.Printf("TestPostPatient:39  --  cfg = %s\n", spew.Sdump(cfg))
		So(err, ShouldBeNil)
		// mongodb, err := OpenMongoDB()
		// So(err, ShouldBeNil)
		// So(mongodb, ShouldNotBeNil)
		//conf, err := Initialize()
		So(err, ShouldBeNil)
		//caFhirId := "62d0ad3c9d0119afff9978b3"
		//cerFhirId := "62f1c5dab3070d0b40e7aac1"
		err = os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
		So(err, ShouldBeNil)
		maker, err := token.NewJWTMaker(os.Getenv("ACCESS_SECRET"))
		So(err, ShouldBeNil)
		So(maker, ShouldNotBeNil)
		username := util.RandomOwner()
		duration := time.Minute
		//userId := "user123456"
		userId := "62d0af5dec383ade03a96b7e"
		role := "Provider"
		ip := "192.168.1.1.99"
		fullName := "Debbie Harman MD"
		newToken, payload, err := maker.CreateToken(ip, username, duration, userId, fullName, role)
		So(err, ShouldBeNil)
		So(newToken, ShouldNotBeNil)
		So(payload, ShouldNotBeNil)
		cp := CreateCP(true)
		cp.SavePayload.SrcPatient = SampleFhirPatient()
		cp.SavePayload.SrcResource = SamplePatient()
		cpb, err := json.Marshal(cp)
		So(err, ShouldBeNil)
		cps := string(cpb)
		fmt.Printf("\nTestPostPatient:70  --  cps: %s\n", cps)
		req, err := http.NewRequest("POST", "/system/640ba66cbd4105586a6dda75/Patient", nil)
		So(err, ShouldBeNil)
		rc := io.NopCloser(strings.NewReader(cps))
		req.Body = rc
		req.Header.Set("AUTHORIZATION", newToken)
		fmt.Printf("TestPostNewPatient:76  --  url: %s\n", req.URL)
		rr := executeRequest(req)
		checkResponseCode(t, http.StatusCreated, rr.Code)
		resp := rr.Result()
		fmt.Printf("TestPostNewPatient:80  --  resp: %s\n", spew.Sdump(resp))
		So(resp, ShouldNotBeNil)
		So(resp.StatusCode, ShouldEqual, http.StatusCreated)

		// //defer resp.Body.Close()
		// opOutcome := fhir.OperationOutcome{}
		// err = json.NewDecoder(resp.Body).Decode(&opOutcome)
		// fmt.Printf("TestPostNewPatient:84  --  opOutcome: %s\n", spew.Sdump(opOutcome))
		// //So(err, ShouldBeNil)
		// So(opOutcome.Issue, ShouldNotBeNil)
		// if err != nil {
		// 	fmt.Printf("GetFhir:131  --  Error Decoding bundle: %s\n", err.Error())
		// 	return nil, err
		// }
		// byte, err := ioutil.ReadAll(resp.Body)
		// So(err, ShouldBeNil)
		// opOutcome := fhir.OperationOutcome{}
		// err = json.Unmarshal(byte, &opOutcome)
		// if err != nil {
		// 	fmt.Printf("PostNewPatient:87  --  Error unmarshalling response to OperationOutcome: %s\n", err.Error())
		// 	So(err, ShouldBeNil)
		// }

		//patient := fhir.Patient{}
		saveResp := common.SaveResponse{}

		err = json.NewDecoder(resp.Body).Decode(&saveResp)
		fmt.Printf("TestPostNewPatient:107  --  patient: %s\n", spew.Sdump(saveResp))
		So(err, ShouldBeNil)
		// err = json.Unmarshal(byte, &patient)
		// if err != nil {
		// 	fmt.Printf("PostNewPatient: 87  --  Response is not a Patient\n")
		// 	err = DetermineOutComeErr(byte)
		// 	So(err, ShouldBeNil)
		// }
		// if patient.ResourceType == nil {
		// 	fmt.Printf("PostNewPatient: 92  --  Response is not a Patient\n")
		// 	err = DetermineOutComeErr(byte)
		// 	So(err, ShouldBeNil)
		// }
		// fmt.Printf("PostNewPatient:96  --  byte: %s\n", string(byte))
		// opOutcom := fhir.OperationOutcome{}
		// err = json.Unmarshal(byte, &opOutcom)
		// So(err, ShouldBeNil)
		// fmt.Printf("PostNewPatient:100  --  opOutcome: %s\n", spew.Sdump(opOutcom))
		// So(opOutcom.Issue, ShouldBeNil)
		// fmt.Printf("PostNewPatient:101  --  patient: %s\n", spew.Sdump(patient))
		// So(resp.StatusCode, ShouldEqual, 200)
		// So(byte, ShouldNotBeNil)

		//location := Conf.Server.Host + ":" + services.Conf.Server.Port + "/system/640ba66cbd4105586a6dda75/Patient/" + patient.Id
		// fmt.Printf("\nTestPostNewPatient:106  --  byte: %s\n", byte)
		// fmt.Printf("\nTestPostNewPatient:107  --  Patient: %s\n", spew.Sdump(patient))
		// location := baseUrl + "/system/640ba66cbd4105586a6dda75/Patient/" // + *patient.Id
		// fmt.Printf("TestPostNewPatient:109  --  location: %s\n", location)
	})
}

func TestPostDuplicatePatient(t *testing.T) {
	fmt.Printf("Test Add a duplicate patient to Server")
	//c := New(baseurl)

	Convey("PostDuplicatePatient", t, func() {
		fmt.Printf("\n\n\nTestPostDuplicatePatient:127  --  Start\n")
		os.Setenv("CONFIG_ADDRESS", "http://192.168.1.148:30100/api/rest/v1")
		_, err := GetServiceConfig("ca3_conn", "local_test", "test") //GetConfig("delivery", "test")
		//fmt.Printf("TestPostDuplicatePatientL143  --  cfg = %s\n", spew.Sdump(cfg))
		So(err, ShouldBeNil)
		// mongodb, err := OpenMongoDB()
		// So(err, ShouldBeNil)
		// So(mongodb, ShouldNotBeNil)
		//conf, err := Initialize()
		So(err, ShouldBeNil)
		//caFhirId := "62d0ad3c9d0119afff9978b3"
		//cerFhirId := "62f1c5dab3070d0b40e7aac1"
		err = os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
		So(err, ShouldBeNil)
		maker, err := token.NewJWTMaker(os.Getenv("ACCESS_SECRET"))
		So(err, ShouldBeNil)
		So(maker, ShouldNotBeNil)
		username := util.RandomOwner()
		duration := time.Minute
		//userId := "user123456"
		userId := "62d0af5dec383ade03a96b7e"
		role := "Provider"
		ip := "192.168.1.1.99"
		fullName := "Debbie Harman MD"
		newToken, payload, err := maker.CreateToken(ip, username, duration, userId, fullName, role)
		So(err, ShouldBeNil)
		So(newToken, ShouldNotBeNil)
		So(payload, ShouldNotBeNil)
		cp := CreateCP(true)
		cp.SavePayload.SrcPatient = SampleFhirPatient()
		cp.SavePayload.SrcResource = SamplePatient()
		cpb, err := json.Marshal(cp)
		So(err, ShouldBeNil)
		cps := string(cpb)
		//fmt.Printf("\nTestPostPatient:174  --  cpb: %s\n", cpb)
		req, err := http.NewRequest("POST", "/system/640ba66cbd4105586a6dda75/Patient", nil)
		So(err, ShouldBeNil)
		rc := io.NopCloser(strings.NewReader(cps))
		req.Body = rc
		req.Header.Set("AUTHORIZATION", newToken)
		fmt.Printf("TestPostDuplicatePatient:180  --  url: %s\n", req.URL)
		rr := executeRequest(req)
		checkResponseCode(t, http.StatusOK, rr.Code)
		resp := rr.Result()
		fmt.Printf("TestPostDuplicatePatient:184  --  resp: %s\n", spew.Sdump(rr))
		defer resp.Body.Close()
		byte, err := ioutil.ReadAll(resp.Body)
		So(err, ShouldBeNil)
		patient := fhir.Patient{}
		err = json.Unmarshal(byte, &patient)
		if err != nil {
			fmt.Printf("TestPostDuplicatePatient:191  --  Response is not a Patient\n")
			err = DetermineOutComeErr(byte)
			So(err, ShouldBeNil)
		}
		if patient.ResourceType == nil {
			fmt.Printf("TestPostDuplicatePatient:196  --  Response is not a Patient\n")
			err = DetermineOutComeErr(byte)
			So(err, ShouldBeNil)
		}
		fmt.Printf("TestPostDuplicatePatient:200  --  byte: %s\n", string(byte))
		opOutcome := fhir.OperationOutcome{}
		err = json.Unmarshal(byte, &opOutcome)
		So(err, ShouldBeNil)
		fmt.Printf("TestPostDuplicatePatient:204  --  opOutcome: %s\n", spew.Sdump(opOutcome))

		So(resp.StatusCode, ShouldEqual, 409)
		So(byte, ShouldNotBeNil)
		So(opOutcome.Issue[0].Code, ShouldEqual, fhir.IssueTypeDuplicate)
		fmt.Printf("\nTestPostDuplictePatient:209  --  byte: %s\n", byte)
	})

	// So(err, ShouldBeNil)
	// So(newToken, ShouldNotBeNil)
	// So(payload, ShouldNotBeNil)
	// //fhirSystem, err := GetFhirSystem(cerFhirId)
	// So(err, ShouldBeNil)
	// cp = CreateCP()

	// // cc := common.ConnectorConfig{}
	// // cp = &common.ConnectorPayload{}
	// // cc.ID, _ = primitive.ObjectIDFromHex("62f1c5dab3070d0b40e7aac1")
	// // cc.Name = "ca3_conn"
	// // cc.Version = "local_test"
	// // cc.CacheUrl = "http://uc_cache:9200"
	// // // "cacheurl" : "http://uc_cache:9200",
	// // // "cache_url" : "http://uc_cache:9200"
	// // data := []*common.KVData{}
	// // cacheServer := common.KVData{}
	// // cacheServer.Name = "cacheServer"
	// // cacheServer.Value = "http://192.168.1.152:30201"
	// // data = append(data, &cacheServer)
	// // hostServer := common.KVData{}
	// // hostServer.Name = "cacheHost"
	// // hostServer.Value = "http://ucCache:9200"
	// // data = append(data, &hostServer)
	// // cc.Data = data
	// // //cp.FhirSystem = fhirSystem
	// // cp.ConnectorConfig = &cc
	// //cp.SavePayload.SrcPatient
	// // dlhFhirId := "6329112852f3616990e2f763"
	// // dlhFhirSystem, err := GetFhirSystem(dlhFhirId)
	// // So(err, ShouldBeNil)
	// // So(dlhFhirSystem, ShouldNotBeNil)
	// // fmt.Printf("dlhFhirSystemURL %s\n", dlhFhirSystem.FhirUrl)
	// //Get a Source Patient to save

	// cnt, bundle, header, err := FindResource(cp, "Patient", userId, "family=smart&_count=12", newToken)
	// //bundle, err := c.PatientSearch(fhirSystem, "family=smart&given=sally&_count=12", "Patient", newToken)
	// So(err, ShouldBeNil)
	// So(bundle, ShouldNotBeNil)
	// So(header, ShouldNotBeNil)
	// So(cnt, ShouldNotEqual, 0)
	// fmt.Printf("TestPostPatient:125 returned %d resources\n", cnt)
	// //fmt.Printf("PatientSearch returned: %s\n", spew.Sdump(bundle))
	// // data, err := c.Query("Patient/12724066")
	// // So(err, ShouldBeNil)
	// // So(data, ShouldNotBeNil)

	// pat, err := fhir.UnmarshalPatient(bundle.Entry[1].Resource)
	// //fmt.Printf("PATIENT: %s\n", spew.Sdump(pat))
	// So(err, ShouldBeNil)
	// So(pat, ShouldNotBeNil)
	// fmt.Printf("TestPostPatient:135  --  Number of entries: %d\n", len(bundle.Entry))

	// //fmt.Printf("PatientSearch[0] returned: %s\n", spew.Sdump(pat))
	// fmt.Printf("TestPostPatient:138  --  Patient.ID := %s\n", *pat.Id)
	// pat.Id = nil
	// pat.Meta = nil
	// newPat, err := c.PostPatient(cp, "210205", &pat)
	// So(err, ShouldBeNil)
	// So(newPat, ShouldNotBeNil)
	// fmt.Printf("TestPostPatient:144  --  NewPatient: %s\n", spew.Sdump(newPat))
	// //time.Sleep(15 * time.Second)
	// })
}

// func TestPatientSearch(t *testing.T) {
// 	fmt.Printf("Test run a FHIR query")
// 	//c := New(baseurl)
// 	Convey("RunPatientQuery", t, func() {
// 		// os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:20100/api/rest/v1")
// 		// _, err := GetServiceConfig("uc_ca3", "local_test", "test") //GetConfig("delivery", "test")
// 		// So(err, ShouldBeNil)
// 		// mongodb, err := OpenMongoDB()
// 		// So(err, ShouldBeNil)
// 		// So(mongodb, ShouldNotBeNil)
// 		// //conf, err := Initialize()
// 		// So(err, ShouldBeNil)
// 		// //caFhirId := "62d0ad3c9d0119afff9978b3"
// 		// cerFhirId := "62f1c5dab3070d0b40e7aac1"
// 		// fhirSystem, err := GetFhirSystem(cerFhirId)
// 		// So(err, ShouldBeNil)
// 		// cc := common.ConnectorConfig{}
// 		// cp := common.ConnectorPayload{}
// 		// cc.ID, _ = primitive.ObjectIDFromHex("62f1c5dab3070d0b40e7aac1")
// 		// cc.Name = "uc_ca3"
// 		// cc.Version = "local_test"
// 		// cc.CacheUrl = "http://uc_cache:9200"
// 		// // "cacheurl" : "http://uc_cache:9200",
// 		// // "cache_url" : "http://uc_cache:9200"
// 		// data := []*common.KVData{}
// 		// cacheServer := common.KVData{}
// 		// cacheServer.Name = "cacheServer"
// 		// cacheServer.Value = "http://192.168.1.152:30201"
// 		// data = append(data, &cacheServer)
// 		// hostServer := common.KVData{}
// 		// hostServer.Name = "cacheHost"
// 		// hostServer.Value = "http://ucCache:9200"
// 		// data = append(data, &hostServer)
// 		// cc.Data = data
// 		// cp.FhirSystem = fhirSystem
// 		// cp.ConnectorConfig = &cc
// 		// //err = os.Setenv("ACCESS_SECRET", util.RandomString(32))
// 		// err = os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
// 		// So(err, ShouldBeNil)
// 		// maker, err := token.NewJWTMaker(os.Getenv("ACCESS_SECRET"))
// 		// So(err, ShouldBeNil)
// 		// So(maker, ShouldNotBeNil)
// 		// username := util.RandomOwner()
// 		// duration := 10 * time.Minute
// 		// //userId := "user123456"
// 		// userId := "62d0af5dec383ade03a96b7e"
// 		// role := "Provider"
// 		// ip := "192.168.1.1.99"
// 		// fullName := "Debbie Harman MD"
// 		// newToken, payload, err := maker.CreateToken(ip, username, duration, userId, fullName, role)
// 		// So(err, ShouldBeNil)
// 		// So(newToken, ShouldNotBeNil)
// 		// So(payload, ShouldNotBeNil)
// 		cp := CreateCP()
// 		JWToken, err := CreateJWToken()
// 		So(err, ShouldBeNil)
// 		cnt, bundle, header, err := FindResource(cp, "Patient", "dharman0127", "family=smart&_count=12", JWToken)
// 		//bundle, err := c.PatientSearch(fhirSystem, "family=smart&_count=12", "patient", newToken)
// 		So(header, ShouldNotBeNil)
// 		So(cnt, ShouldNotEqual, 0)
// 		So(err, ShouldBeNil)
// 		So(bundle, ShouldNotBeNil)
// 		fmt.Printf("TestPatientSearch:177 returned %d resources\n", cnt)
// 		//fmt.Printf("PatientSearch returned: %s\n", spew.Sdump(bundle))
// 		// data, err := c.Query("Patient/12724066")
// 		// So(err, ShouldBeNil)
// 		// So(data, ShouldNotBeNil)
// 		pat, err := fhir4.UnmarshalPatient(bundle.Entry[0].Resource)
// 		So(err, ShouldBeNil)
// 		So(pat, ShouldNotBeNil)

// 		//fmt.Printf("PatientSearch[0] returned: %s\n", spew.Sdump(pat))
// 		fmt.Printf("Patient.ID := %s\n", *pat.Id)
// 		time.Sleep(15 * time.Second)
// 	})
// }

// // func TestCaPatientSearch(t *testing.T) {
// // 	fmt.Printf("Test query through to CA for patient")
// // 	//c := New(baseurl)
// // 	Convey("RunCaPatientQuery", t, func() {
// // 		os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:20100/api/rest/v1")
// // 		_, err := GetServiceConfig("uc_ca3", "local_test", "test") //GetConfig("delivery", "test")
// // 		So(err, ShouldBeNil)
// // 		mongodb, err := OpenMongoDB()
// // 		So(err, ShouldBeNil)
// // 		So(mongodb, ShouldNotBeNil)
// // 		//conf, err := Initialize()
// // 		So(err, ShouldBeNil)
// // 		caFhirId := "62f14531ba5395278cd530c4"
// // 		//cerFhirId := "62f1c5dab3070d0b40e7aac1"
// // 		fhirSystem, err := GetFhirSystem(caFhirId)
// // 		So(err, ShouldBeNil)
// // 		cc := common.ConnectorConfig{}
// // 		cp := common.ConnectorPayload{}
// // 		cc.ID, _ = primitive.ObjectIDFromHex("62f1c5dab3070d0b40e7aac1")
// // 		cc.Name = "uc_ca3"
// // 		cc.Version = "local_test"
// // 		cc.CacheUrl = "http://uc_cache:9200"
// // 		// "cacheurl" : "http://uc_cache:9200",
// // 		// "cache_url" : "http://uc_cache:9200"
// // 		data := []*common.KVData{}
// // 		cacheServer := common.KVData{}
// // 		cacheServer.Name = "cacheServer"
// // 		cacheServer.Value = "http://192.168.1.152:30201"
// // 		data = append(data, &cacheServer)
// // 		hostServer := common.KVData{}
// // 		hostServer.Name = "cacheHost"
// // 		hostServer.Value = "http://ucCache:9200"
// // 		data = append(data, &hostServer)
// // 		cc.Data = data
// // 		cp.FhirSystem = fhirSystem
// // 		cp.ConnectorConfig = &cc
// // 		maker, err := token.NewJWTMaker(os.Getenv("ACCESS_SECRET"))
// // 		So(err, ShouldBeNil)
// // 		So(maker, ShouldNotBeNil)
// // 		username := util.RandomOwner()
// // 		duration := time.Minute
// // 		//userId := "user123456"
// // 		userId := "62d0af5dec383ade03a96b7e"
// // 		role := "Provider"
// // 		ip := "192.168.1.1.99"
// // 		fullName := "Debbie Harman MD"
// // 		newToken, payload, err := maker.CreateToken(ip, username, duration, userId, fullName, role)
// // 		So(err, ShouldBeNil)
// // 		So(newToken, ShouldNotBeNil)
// // 		So(payload, ShouldNotBeNil)
// // 		// bundle, err := c.PatientSearch(fhirSystem, "family=HARMAN", "Patient", newToken)
// // 		// So(err, ShouldBeNil)
// // 		// So(bundle, ShouldNotBeNil)
// // 		cnt, bundle, header, err := FindResource(&cp, "Patient", userId, "family=smart&_count=12", newToken)
// // 		//bundle, err := c.PatientSearch(fhirSystem, "family=smart&_count=12", "patient", newToken)
// // 		So(header, ShouldNotBeNil)
// // 		So(cnt, ShouldNotEqual, 0)
// // 		So(err, ShouldBeNil)
// // 		So(bundle, ShouldNotBeNil)
// // 		fmt.Printf("TestCaPatientSearch:184 returned %d resources\n", cnt)
// // 		//fmt.Printf("PatientSearch returned: %s\n", spew.Sdump(bundle))
// // 		// data, err := c.Query("Patient/12724066")
// // 		// So(err, ShouldBeNil)
// // 		// So(data, ShouldNotBeNil)
// // 		pat, err := fhir4.UnmarshalPatient(bundle.Entry[0].Resource)
// // 		So(err, ShouldBeNil)
// // 		So(pat, ShouldNotBeNil)

// // 		fmt.Printf("PatientSearch[0] returned: %s\n", spew.Sdump(pat))
// // 		fmt.Printf("TestCaPatientSearch:256  --  Name: %s\n", spew.Sdump(pat.Name))
// // 		//fmt.Printf("Patient.ID = %s Name = %s\n", *pat.Id, *pat.Name[0].Family)
// // 		//ime.Sleep(15 * time.Second)
// // 	})
// // }

// // func TestGetPatient(t *testing.T) {
// // 	fmt.Printf("Test run a FHIR query")
// // 	//c := New(baseurl)
// // 	Convey("RunPatientGet", t, func() {
// // 		os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:30300/api/rest/v1")
// // 		_, err := GetServiceConfig("uc_ca3", "local_test", "test") //GetConfig("delivery", "test")
// // 		So(err, ShouldBeNil)
// // 		mongodb, err := OpenMongoDB()
// // 		So(err, ShouldBeNil)
// // 		So(mongodb, ShouldNotBeNil)
// // 		//conf, err := Initialize()
// // 		So(err, ShouldBeNil)
// // 		pat, err := GetPatient("12743944")
// // 		So(err, ShouldBeNil)
// // 		So(pat, ShouldNotBeNil)
// // 		//fmt.Printf("Found Patient: %s\n", spew.Sdump(pat))
// // 	})
// // }

// // func TestCreateIdentifier(t *testing.T) {
// // 	fmt.Printf("Test run a FHIR query")
// // 	//c := New(baseurl)
// // 	Convey("CreateIdentifier", t, func() {
// // 		// os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:30300/api/rest/v1")
// // 		// _, err := GetServiceConfig("uc_ca3", "local_test", "test") //GetConfig("delivery", "test")
// // 		// So(err, ShouldBeNil)
// // 		// mongodb, err := OpenMongoDB()
// // 		// So(err, ShouldBeNil)
// // 		// So(mongodb, ShouldNotBeNil)
// // 		// //conf, err := Initialize()
// // 		// So(err, ShouldBeNil)
// // 		ident := CreateIdentifier("63c703eac5cc538807e9b775")
// // 		So(ident, ShouldNotBeNil)
// // 		fmt.Printf("Ident: %s\n", spew.Sdump(ident))
// // 	})
// // }

// func SamplePatient() []byte {
// 	return []byte(`{
// 		"resourceType": "Patient",
// 		"id": "12742611",

// 		"text": {
// 			"status": "generated",
// 			"div": "<div><p><b>Patient</b></p><p><b>Name</b>: SMART, Debra</p><p><b>DOB</b>: Jan 27, 1950</p><p><b>Administrative Gender</b>: Female</p><p><b>Status</b>: Active</p></div>"
// 		},
// 		"identifier": [
// 			{
// 				"use": "usual",
// 				"type": {
// 					"coding": [
// 						{
// 							"system": "http://hl7.org/fhir/v2/0203",
// 							"code": "MR",
// 							"display": "Medical record number",
// 							"userSelected": false
// 						}
// 					],
// 					"text": "MRN"
// 				},
// 				"system": "urn:oid:2.16.840.1.113883.6.1000",
// 				"value": "106979",
// 				"_value": {
// 					"extension": [
// 						{
// 							"url": "http://hl7.org/fhir/StructureDefinition/rendered-value",
// 							"valueString": "00000106979"
// 						}
// 					]
// 				},
// 				"period": {
// 					"start": "2023-01-02T19:34:51.000Z"
// 				}
// 			}
// 		],
// 		"active": true,
// 		"name": [
// 			{
// 				"use": "official",
// 				"text": "SMART, Debra",
// 				"family": [
// 					"SMART"
// 				],
// 				"given": [
// 					"Debra"
// 				],
// 				"period": {
// 					"start": "2023-01-01T19:15:48.000Z"
// 				}
// 			}
// 		],
// 		"telecom": [
// 			{
// 				"system": "email",
// 				"value": "dsmart@yopmail.com",
// 				"use": "home",
// 				"period": {
// 					"start": "2023-01-01T19:15:47.000Z"
// 				}
// 			}
// 		],
// 		"gender": "female",
// 		"birthDate": "1958-01-27"
// 	}`)
// }

// func SampleFhirPatient() *fhir.Patient {
// 	//fp := fhir.Patient{}
// 	//fmt.Printf("\nfhirPatient: %s\n\n", spew.Sdump(fp))
// 	fhirPat, err := fhir.UnmarshalPatient(SamplePatient())
// 	if err != nil {
// 		fmt.Printf("UnmarshalPatient error : %s\n", err.Error())
// 		return nil
// 	}
// 	return &fhirPat
// }
