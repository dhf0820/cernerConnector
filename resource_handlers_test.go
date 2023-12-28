package main

// import (
// 	//"bytes"
// 	"context"
// 	"encoding/json"
// 	"io"

// 	//"github.com/gorilla/mux"
// 	"github.com/joho/godotenv"
// 	//"gitlab.com/dhf0820/ids_model/common"

// 	log "github.com/sirupsen/logrus"
// 	//. "github.com/smartystreets/goconvey/convey"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"strings"
// 	"testing"

// 	//"time"

// 	jw_token "github.com/dhf0820/golangJWT"
// 	common "github.com/dhf0820/uc_core/common"

// 	//"github.com/dhf0820/uc_core/service"

// 	"github.com/davecgh/go-spew/spew"
// 	fhir "github.com/dhf0820/fhir4"
// 	"go.mongodb.org/mongo-driver/bson"

// 	//fhirR4go "github.com/dhf0820/fhirR4go"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	//log "github.com/sirupsen/logrus"
// 	. "github.com/smartystreets/goconvey/convey"
// )

// func TestSimpleFindResource(t *testing.T) {
// 	Convey("TestSimpleFindResource", t, func() {
// 		req, err := http.NewRequest("GET", "/6329112852f3616990e2f763/Patient?family=smart", nil)
// 		//req, err := http.NewRequest("GET", "/api/rest/v1/GetPatient/12345", nil)
// 		So(err, ShouldBeNil)
// 		//w := httptest.NewRecorder()
// 		os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
// 		jwt, payload, err := jw_token.CreateTestJWToken("10s")
// 		So(err, ShouldBeNil)
// 		So(jwt, ShouldNotBeNil)
// 		So(payload, ShouldNotBeNil)
// 		req.Header.Set("Authorization", jwt)
// 		fmt.Printf("\nCalling Router\n")
// 		cp := CreateCP(false)
// 		cpb, err := json.Marshal(cp)
// 		So(err, ShouldBeNil)
// 		cps := string(cpb)
// 		rc := io.NopCloser(strings.NewReader(cps))
// 		req.Body = rc
// 		fmt.Printf("TestResourceHandlerGet:55  --  url: %s\n", req.URL)
// 		rr := executeRequest(req)
// 		checkResponseCode(t, http.StatusOK, rr.Code)
// 		resp := rr.Result()
// 		defer resp.Body.Close()
// 		byte, err := io.ReadAll(resp.Body)
// 		So(err, ShouldBeNil)
// 		So(byte, ShouldNotBeNil)
// 		resResp := &common.ResourceResponse{}
// 		err = json.Unmarshal(byte, resResp)
// 		//patient, err := fhir.UnmarshalPatient(byte)
// 		So(err, ShouldBeNil)
// 		So(resResp, ShouldNotBeNil)
// 		So(resResp.Bundle, ShouldNotBeNil)
// 		So(resResp.Patients, ShouldNotBeNil)
// 		So(resResp.Resources, ShouldBeNil)
// 		//fmt.Printf("\nTestSimpleFindResource:68  --  Patients: %s\n", spew.Sdump(resResp.Patients))
// 		//fmt.Printf("\nTestSimpleFindResource:69  --  Bundle: %s\n", spew.Sdump(resResp.Bundle))
// 	})
// }

// func TestResourceHandlerGet(t *testing.T) {
// 	Convey("TestResourceHandlerGet", t, func() {
// 		fmt.Printf("Creating Get Request\n")
// 		jwt, payload, err := jw_token.CreateTestJWToken("10s")
// 		So(err, ShouldBeNil)
// 		So(jwt, ShouldNotBeNil)
// 		So(payload, ShouldNotBeNil)
// 		req, err := http.NewRequest("GET", "/api/rest/v1/Patient/63ed93c8bd78ae6b013a502b", nil)
// 		So(err, ShouldBeNil)
// 		So(req, ShouldNotBeNil)
// 		req.Header.Set("Authorization", jwt)
// 		req.Header.Set("facility", "demo")
// 		req.Header.Set("Fhir_Version", "r4")
// 		cp := CreateCernerCP(false)
// 		cpb, err := json.Marshal(cp)
// 		So(err, ShouldBeNil)
// 		cps := string(cpb)
// 		rc := io.NopCloser(strings.NewReader(cps))
// 		req.Body = rc
// 		fmt.Printf("TestResourceHandlerGet:71  --  url: %s\n", req.URL)
// 		rr := executeRequest(req)
// 		checkResponseCode(t, http.StatusOK, rr.Code)
// 		resp := rr.Result()
// 		defer resp.Body.Close()
// 		byte, err := io.ReadAll(resp.Body)
// 		So(err, ShouldBeNil)
// 		So(byte, ShouldNotBeNil)
// 		//bundle := fhir.Bundle{}
// 		patient, err := fhir.UnmarshalPatient(byte)
// 		So(err, ShouldBeNil)
// 		So(patient, ShouldNotBeNil)
// 		fmt.Printf("TestResourceHandlerGet:83  --  Patient: %s\n", spew.Sdump(patient))

// 		// //req, err := http.NewRequest("GET", "/api/rest/v1/healthcheck", nil) //"api/Patient?family=smart", nil)
// 		// // req, err := http.NewRequest("GET", "/api/rest/v1/Patient?family=smart", nil)
// 		// // So(err, ShouldBeNil)
// 		// //req, err := http.NewRequest("GET", "/api/rest/v1/Patient/12345678", nil)
// 		// // vars := map[string]string{
// 		// // 	"resource": "Patient",
// 		// // }
// 		// // req = mux.SetURLVars(req, vars)
// 		// // fmt.Printf("\n\nSetting FhirSystem body\n")
// 		// // fs := `{
// 		// // 		"id": "6329112852f3616990e2f763",
// 		// // 		"facilityId": "62d0af5dec383ade03a96b7f",
// 		// // 		"facilityName": "Harman Clinic",
// 		// // 		"displayName": "Harman",
// 		// // 		"description": "Medical Records from 2023-",
// 		// // 		"fhirVersion": "r4",
// 		// // 		"authUrl": "",
// 		// // 		"ucUrl": "http://192.168.1.152:30300/6329112852f3616990e2f763",
// 		// // 		"fhirUrl": "http://api/rest/v1",
// 		// // 		"insert": "true",
// 		// // 		"identifiers": [
// 		// // 			{
// 		// // 				"Name": "Mrn",
// 		// // 				"Value": "urn:oid:1.3.6.1.4.1.54392.5.1593.1|"
// 		// // 			}
// 		// // 		],
// 		// // 		"facilityCode": "",
// 		// // 		"serviceName" : "uc_ca3",
// 		// // 		"returnBundle" : "true",
// 		// // 		"connector" : "uc_ca3:local_test"
// 		// // 	}`
// 		// // //}`

// 		// // fsByte := []byte(fs)
// 		// // fhirSystem := &uc_core/common.FhirSystem{}
// 		// // err := json.Unmarshal(fsByte, fhirSystem)
// 		// // So(err, ShouldBeNil)
// 		// // fmt.Printf("\n\nTest:89  --  fhirSystem = %s\n", spew.Sdump(fhirSystem))

// 		// // // fsBody := json.RawMessage(fs)
// 		// // // //(err, ShouldBeNil)
// 		// // // fmt.Printf("fsBody RawMessage String: %s\n\n", fsBody)
// 		// // // fst := common.FhirSystem{}
// 		// // // json.Unmarshal(fsBody, &fst)
// 		// // //So(err, ShouldBeNil)
// 		// // //fmt.Printf("\n\ntest:106 -- FhirSystem = %s\n", spew.Sdump(fhirSystem))
// 		// // cp := uc_core/common.ConnectorPayload{}
// 		// // cp.FhirSystem = fhirSystem
// 		// // //cp := common.ConnectorPayload{}
// 		// // cc := uc_core/common.ConnectorConfig{}
// 		// // cc.ID, _ = primitive.ObjectIDFromHex("62f1c5dab3070d0b40e7aac1")
// 		// // cc.Name = "uc_ca3"
// 		// // cc.Version = "local_test"
// 		// // cc.CacheUrl = "http://localhost:30201"
// 		// // // "cacheurl" : "http://uc_cache:9200",
// 		// // // "cache_url" : "http://uc_cache:9200"
// 		// // data := []*uc_core/common.KVData{}
// 		// // cacheServer := uc_core/common.KVData{}
// 		// // cacheServer.Name = "cacheServer"
// 		// // cacheServer.Value = "http://192.168.1.152:30201"
// 		// // data = append(data, &cacheServer)
// 		// // hostServer := uc_core/common.KVData{}
// 		// // hostServer.Name = "cacheHost"
// 		// // hostServer.Value = "http://ucCache:9200"
// 		// // data = append(data, &hostServer)
// 		// // cc.Data = data
// 		// // cp.FhirSystem = fhirSystem
// 		// // cp.ConnectorConfig = &cc
// 		// // cps, err := json.Marshal(cp)
// 		// // So(err, ShouldBeNil)
// 		// // rc := io.NopCloser(strings.NewReader(string(cps)))
// 		// req, err := http.NewRequest("GET", "/api/rest/v1/Patient/63ed93c8bd78ae6b013a502b", nil)
// 		// So(err, ShouldBeNil)
// 		// cp := CreateCP()
// 		// //cpb, err := json.Marshal(cp)
// 		// So(err, ShouldBeNil)
// 		// //cps := string(cpb)
// 		// //rc := io.NopCloser(strings.NewReader(cps))
// 		// rc := CreateTestFileCloser()
// 		// //req.Body, err = json.Marshal([]byte(fs))
// 		// req.Body = rc
// 		// //rc := CreateTestFileCloser()
// 		// // values := req.URL.Query()
// 		// // values.Add("resource", "Patient")
// 		// // req.URL.RawQuery = values.Encode()
// 		// // vars := map[string]string{
// 		// // 	"resource": "Patient",
// 		// // }
// 		// //req = mux.SetURLVars(req, vars)
// 		// // fmt.Printf("test request: ")
// 		// // spew.Dump(req)
// 		// //So(err, ShouldBeNil)
// 		// w := httptest.NewRecorder()
// 		// jwt, err := CreateJWToken()
// 		// So(err, ShouldBeNil)
// 		// So(jwt, ShouldNotBeNil)
// 		// req.Header.Set("Authorization", jwt)

// 		// fmt.Printf("\nTestGetResource:148  --  Calling Router with %s\n", req.URL)
// 		// //StartTime := time.Now()
// 		// NewRouter().ServeHTTP(w, req)
// 		// resp := w.Result()
// 		// fmt.Printf("\n\nTestGetResource:152  --  Results.Status: %s\n body: %v\n", w.Result().Status, w.Result().Body)
// 		// defer resp.Body.Close()
// 		// byte, err := ioutil.ReadAll(resp.Body)
// 		// //bundle := fhir.Bundle{}
// 		// patient, err := fhir.UnmarshalPatient(byte)
// 		// So(err, ShouldBeNil)
// 		// So(patient, ShouldNotBeNil)
// 		// //So(bundle.ResourceType, ShouldEqual, resource)
// 		// // fmt.Printf("########################## Wait for Background Elapsed Time = %s\n\n\n\n\n", time.Since(StartTime))
// 		// // time.Sleep(20 * time.Second)
// 	})
// }
// func TestFindCa3ResourceHandler(t *testing.T) {
// 	Convey("Subject: Find all Resources matching Filter", t, func() {
// 		godotenv.Load("./.env.uc_ca3_test")
// 		//os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:30100/api/rest/v1")

// 		//os.Setenv("COMPANY", "demo")

// 		// conf, err := service.InitCore("uc_ca3", "local_test", "test")
// 		// So(err, ShouldBeNil)
// 		// So(conf, ShouldNotBeNil)
// 		// So(err, ShouldBeNil)
// 		// So(conf, ShouldNotBeNil)
// 		// _, err := service.Initialize()
// 		// //, err := service.InitCore("uc_core","test", "test")
// 		// if err != nil {
// 		// 	t.Fatalf("InitCore failed: %s", err.Error())
// 		// }

// 		Convey("Given a valid patient Family", func() {
// 			//resource := "Patient"
// 			fmt.Printf("\n\nGiven a valid Family Name\n")
// 			jwt, payload, err := jw_token.CreateTestJWToken("10s")
// 			So(err, ShouldBeNil)
// 			So(jwt, ShouldNotBeNil)
// 			So(payload, ShouldNotBeNil)

// 			//w := httptest.NewRecorder()
// 			fmt.Printf("Creating Find Request\n")
// 			//req, _ := http.NewRequest("GET", "http://192.168.1.152:20104/6329112852f3616990e2f763/Patient?family=smart", nil)
// 			req, _ := http.NewRequest("GET", "/6329112852f3616990e2f763/Patient?family=smart", nil)
// 			req.Header.Set("Authorization", jwt)
// 			req.Header.Set("facility", "demo")
// 			req.Header.Set("Fhir_Version", "r4")
// 			cp := CreateCP(false)
// 			cpb, err := json.Marshal(cp)
// 			So(err, ShouldBeNil)
// 			cps := string(cpb)
// 			rc := io.NopCloser(strings.NewReader(cps))
// 			req.Body = rc
// 			fmt.Printf("TestFindCaResource:239  --  url: %s\n", req.URL)
// 			rr := executeRequest(req)
// 			checkResponseCode(t, http.StatusOK, rr.Code)
// 			resp := rr.Result()

// 			//resp, err := client.Do(req)
// 			So(resp, ShouldNotBeNil)
// 			So(resp.StatusCode, ShouldEqual, 200)
// 			//fmt.Printf("TestFindCa3Resource:247  --  body:  %v\n", resp.BodyGetFhirBundle)
// 			defer resp.Body.Close()
// 			//byte, err := ioutil.ReadAll(resp.Body)
// 			//bundle := fhir.Bundle{}
// 			resResp := &common.ResourceResponse{}
// 			err = json.NewDecoder(resp.Body).Decode(resResp)
// 			So(err, ShouldBeNil)
// 			So(resResp, ShouldNotBeNil)
// 			So(len(resResp.Patients), ShouldEqual, 2)
// 			// bundle, err := fhir.UnmarshalBundle(byte)
// 			// So(err, ShouldBeNil)
// 			patient, err := fhir.UnmarshalPatient(resResp.Bundle.Entry[0].Resource)
// 			fmt.Printf("TestFindCa3Resource:259  --  Bundle Patient: %s\n", spew.Sdump(patient))
// 			fmt.Printf("TestFindCa3Resource:260  --  Patients[0]: %s\n", spew.Sdump(resResp.Patients[0]))

// 			// // if err != nil {
// 			// // 	return 0, nil, nil, err
// 			// // }
// 			// //patientId, err := GetPatientFromBundle(resource, &bundle)

// 			// if err != nil {
// 			// 	return 0, nil, nil, err
// 			// }
// 			// // if err != nil {
// 			// // 	log.Println("CacheResourceBundleAndEntries:108  --  Error FindResource Request: ", err.Error())
// 			// // } else {
// 			// // 	log.Println("CacheResourceBundleAndEntries:106  --  FindResource Successful")
// 			// // }
// 			// findResource(w, req)
// 			// resp := w.Result()
// 			// So(resp.StatusCode, ShouldEqual, http.StatusOK)
// 			// //var bundle []fhir.Bundle

// 			// err = json.NewDecoder(w.Body).Decode(&bundle)
// 			// So(err, ShouldBeNil)
// 			// fmt.Printf("bundle: %s\n", spew.Sdump(bundle))

// 		})
// 	})
// }

// // func TestPostFhirPatient(t *testing.T) {
// // 	Convey("Subject: GetDocument For Patient returns documents for a patient", t, func() {
// // 		godotenv.Load("./.env.uc_ca3_test")
// // 		//os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:30100/api/rest/v1")

// // 		//os.Setenv("COMPANY", "demo")

// // 		conf, err := InitCore("ca3_conn", "local_test", "test")
// // 		So(err, ShouldBeNil)
// // 		So(conf, ShouldNotBeNil)
// // 		So(err, ShouldBeNil)
// // 		So(conf, ShouldNotBeNil)
// // 		// _, err := service.Initialize()
// // 		// //, err := service.InitCore("uc_core","test", "test")
// // 		// if err != nil {
// // 		// 	t.Fatalf("InitCore failed: %s", err.Error())
// // 		// }

// // 		Convey("Given a valid patient Family", func() {
// // 			fmt.Printf("\n\nTestPostFhirPatient:290  --  Given a valid DocumentId\n")
// // 			os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
// // 			dur := time.Duration(300) * time.Second
// // 			jwt, err := token.CreateToken("192.168.1.2", "DHarman", dur, "userId1234", "Debbie Harman", "Physician")
// // 			So(err, ShouldBeNil)
// // 			So(jwt, ShouldNotBeNil)
// // 			w := httptest.NewRecorder()
// // 			fmt.Printf("TestPostFhirPatient:297  --  Creating Search Request\n")
// // 			req, _ := http.NewRequest("GET", "62f1c5dab3070d0b40e7aac1/Patient?family=smart&given=sandy", nil)
// // 			//req, _ := http.NewRequest("GET", "62f1c5dab3070d0b40e7aac1/api/rest/v1/Documents?patient=1&_count=2", nil)
// // 			//fmt.Printf("n\nTest Request = %s\n\n\n", spew.Sdump(req))
// // 			//req.Header.Add("tracing-id", "123")
// // 			// vars := map[string]string{
// // 			// 	"doc_id" : "40441",
// // 			// }
// // 			//req = mux.SetURLVars(req, vars)
// // 			req.Header.Set("Authorization", jwt)
// // 			req.Header.Set("facility", "demo")
// // 			req.Header.Set("Fhir_Version", "r4")
// // 			fmt.Printf("")

// // 			findResource(w, req)
// // 			resp := w.Result()
// // 			So(resp.StatusCode, ShouldEqual, http.StatusOK)
// // 			var bundle []fhir.Bundle

// // 			err = json.NewDecoder(w.Body).Decode(&bundle)
// // 			So(err, ShouldBeNil)
// // 			fmt.Printf("bundle: %s\n", spew.Sdump(bundle))

// // 		})
// // 	})
// // }

// func TestFhirDocumentForPatient(t *testing.T) {
// 	Convey("Subject: GetDocument For Patient returns documents for a patient", t, func() {
// 		godotenv.Load("./.env.uc_ca3_test")
// 		//os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:30100/api/rest/v1")

// 		//os.Setenv("COMPANY", "demo")

// 		// conf, err := service.InitCore("uc_ca3", "local_test", "test")
// 		// So(err, ShouldBeNil)
// 		// So(conf, ShouldNotBeNil)
// 		// So(err, ShouldBeNil)
// 		// So(conf, ShouldNotBeNil)
// 		// _, err := service.Initialize()
// 		// //, err := service.InitCore("uc_core","test", "test")
// 		// if err != nil {
// 		// 	t.Fatalf("InitCore failed: %s", err.Error())
// 		// }

// 		Convey("Given a valid patient Family", func() {
// 			fmt.Printf("\n\nTestFhirDocumentForPatient:343  --  Given a valid DocumentId\n")
// 			os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
// 			jwt, payload, err := jw_token.CreateTestJWToken("10s")
// 			So(err, ShouldBeNil)
// 			So(jwt, ShouldNotBeNil)
// 			So(payload, ShouldNotBeNil)
// 			w := httptest.NewRecorder()
// 			fmt.Printf("TestFhirDocumentForPatient:350  --  Creating Search Request\n")
// 			cps := CreateTestFileCloser()
// 			req, _ := http.NewRequest("GET", "62f1c5dab3070d0b40e7aac1/Find/DocumentReference?patient=12748336", cps)
// 			//req, _ := http.NewRequest("GET", "62f1c5dab3070d0b40e7aac1/api/rest/v1/Documents?patient=1&_count=2", nil)
// 			//fmt.Printf("n\nTest Request = %s\n\n\n", spew.Sdump(req))
// 			//req.Header.Add("tracing-id", "123")
// 			// vars := map[string]string{
// 			// 	"doc_id" : "40441",
// 			// }
// 			//req = mux.SetURLVars(req, vars)

// 			req.Header.Set("Authorization", jwt)
// 			req.Header.Set("facility", "demo")
// 			req.Header.Set("Fhir_Version", "r4")
// 			fmt.Printf("")

// 			findResource(w, req)
// 			resp := w.Result()
// 			So(resp.StatusCode, ShouldEqual, http.StatusOK)
// 			//var bundle []fhir4.Bundle
// 			var resResp common.ResourceResponse

// 			err = json.NewDecoder(w.Body).Decode(&resResp)
// 			So(err, ShouldBeNil)
// 			fmt.Printf("ResResp: %s\n", spew.Sdump(resResp))

// 		})
// 	})
// }

// func TestFhirEncountersForPatient(t *testing.T) {
// 	Convey("Subject: GetDocument For Patient returns documents for a patient", t, func() {
// 		godotenv.Load("./.env.uc_ca3_test")
// 		// conf, err := service.InitCore("uc_ca3", "local_test", "test")
// 		// So(err, ShouldBeNil)
// 		// So(conf, ShouldNotBeNil)
// 		// So(err, ShouldBeNil)
// 		// So(conf, ShouldNotBeNil)

// 		Convey("Given a valid patient Family", func() {
// 			fmt.Printf("\n\nGiven a valid PatientId\n")
// 			os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
// 			jwt, payload, err := jw_token.CreateTestJWToken("10s")
// 			So(payload, ShouldNotBeNil)
// 			So(err, ShouldBeNil)
// 			So(jwt, ShouldNotBeNil)
// 			w := httptest.NewRecorder()
// 			fmt.Printf("Creating Search Request\n")
// 			cps := CreateTestFileCloser()
// 			req, _ := http.NewRequest("GET", "62f1c5dab3070d0b40e7aac1/Find/Encounter?patient=12748336", cps)
// 			req.Header.Set("Authorization", jwt)
// 			req.Header.Set("facility", "demo")
// 			req.Header.Set("Fhir_Version", "r4")
// 			fmt.Printf("")

// 			findResource(w, req)
// 			resp := w.Result()
// 			So(resp.StatusCode, ShouldEqual, http.StatusOK)
// 			var resResp common.ResourceResponse
// 			err = json.NewDecoder(w.Body).Decode(&resResp)
// 			So(err, ShouldBeNil)
// 			fmt.Printf("TestFhirEncountersForPatient:441  --  ResResp: %s\n", spew.Sdump(resResp))
// 		})
// 	})
// }

// func TestFhirProceduresForPatient(t *testing.T) {
// 	Convey("Subject: GetDocument For Patient returns documents for a patient", t, func() {
// 		godotenv.Load("./.env.baseConnecor")
// 		// conf, err := service.InitCore("uc_ca3", "local_test", "test")
// 		// So(err, ShouldBeNil)
// 		// So(conf, ShouldNotBeNil)
// 		// So(err, ShouldBeNil)
// 		// So(conf, ShouldNotBeNil)

// 		Convey("Given a valid patient Family", func() {
// 			fmt.Printf("\n\nGiven a valid PatientId\n")
// 			os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
// 			jwt, payload, err := jw_token.CreateTestJWToken("10s")
// 			So(err, ShouldBeNil)
// 			So(jwt, ShouldNotBeNil)
// 			So(payload, ShouldNotBeNil)
// 			w := httptest.NewRecorder()
// 			fmt.Printf("TestFhirProceduresForPatient:433  --  Creating Search Request\n")
// 			cps := CreateTestFileCloser()
// 			req, _ := http.NewRequest("GET", "62f1c5dab3070d0b40e7aac1/Find/Procedure?patient=12748336", cps)
// 			req.Header.Set("Authorization", jwt)
// 			req.Header.Set("facility", "demo")
// 			req.Header.Set("Fhir_Version", "r4")
// 			fmt.Printf("")

// 			findResource(w, req)
// 			resp := w.Result()
// 			So(resp.StatusCode, ShouldEqual, http.StatusOK)
// 			var resResp common.ResourceResponse
// 			err = json.NewDecoder(w.Body).Decode(&resResp)
// 			So(err, ShouldBeNil)
// 			fmt.Printf("TestFhirProceduresForPatient:447  --  ResResp: %s\n", spew.Sdump(resResp))
// 		})
// 	})
// }

// //func CreateTestFhirSystem() *common.FhirSystem {
// // fs1 := `{
// // 	"_id" : ObjectId("6329112852f3616990e2f763"),
// // 	"facilityId" : ObjectId("62d0af5dec383ade03a96b7f"),
// // 	"facilityName" : "Harman Clinic",
// // 	"displayName" : "Harman ChartArchive",
// // 	"description" : "Current records for Dr. Harman",
// // 	"fhirVersion" : "r4",
// // 	"identifiers" : [
// // 		{
// // 			"name" : "mrn",
// // 			"value" : "mrn|"
// // 		},
// // 		{
// // 			"name" : "ssn",
// // 			"value" : "ssn|"
// // 		},
// // 		{
// // 			"name" : "id",
// // 			"value" : "id|"
// // 		}
// // 	],
// // 	"authUrl" : "",
// // 	"fhirUrl" : "http://universalcharts.com:4000/api/rest/v1",
// // 	"insert" : "true",
// // 	"ucUrl" : "http://test.universalcharts.com/6329112852f3616990e2f763",
// // 	"facilityCode" : "harman",
// // 	"baseUrl" : "https://UniversalCharts/api/rest/v1",
// // 	"cacheUrl" : "http://UniversalCharts.com:30201",
// // 	"connector" : "uc_ca3:cloud",
// // 	"returnBundle" : "true",
// // 	"serviceName" : "uc_ca3"
// // }`

// // 	fs := &common.FhirSystem{}
// // 	fs.ID, _ = primitive.ObjectIDFromHex("6329112852f3616990e2f763")
// // 	fs.FacilityId, _ = primitive.ObjectIDFromHex("62d0af5dec383ade03a96b7f")
// // 	fs.FacilityName = "Harman Clinic"
// // 	fs.DisplayName = "Harman ChartArchive"
// // 	fs.Description = "Current records for Dr. Harman"
// // 	fs.FhirVersion = "r4"
// // 	fs.Identifiers = []*uc_core/common.KVData{}
// // 	mrn := uc_core/common.KVData{}
// // 	mrn.Name = "mrn"
// // 	mrn.Value = "mrn|"
// // 	fs.Identifiers = append(fs.Identifiers, &mrn)
// // 	fs.FhirUrl = "http://192.168.1.152:4000/api/rest/v1"
// // 	fs.Insert = "true"
// // 	fs.UcUrl = "http://192.168.1.152/6329112852f3616990e2f763"
// // 	fs.FacilityCode = "harman"
// // 	fs.Facility.BaseUrl = "http://192.168.1.152/api/rest/v1"
// // 	fs.Connector = "uc_ca3:cloud"
// // 	fhirSystem := fs
// // 	fmt.Printf("CreateTestFhirSystem:505  --  fhirSystem: %s\n", spew.Sdump(fhirSystem))
// // 	return fhirSystem
// // }

// func CreateTestFileCloser() io.ReadCloser {
// 	cp := CreateCP(false)
// 	cpb, err := json.Marshal(cp)
// 	if err != nil {
// 		fmt.Printf("CreateTestFileCloser:513  --  Marshal cp failed: %s\n", err.Error())
// 	}
// 	//cps := string(cpb)
// 	log.Printf("CreateTestFileCloser:516  -- Starting \n\n\n\n")
// 	log.Printf("CreateTestCloser:554  --  cpb: %s\n", string(cpb))
// 	fmt.Println()
// 	rc := io.NopCloser(strings.NewReader(string(cpb)))
// 	return rc
// }

// func TestDeterminePatientResource(t *testing.T) {
// 	Convey("Subject: Determine Patient Resource from url", t, func() {
// 		godotenv.Load("./.env.uc_ca3_test")
// 		Convey("Determine PatientResource of GET", func() {
// 			resource := DetermineResource("https://chartarchivefhir.com:4000/api/rest/v1/Patient/63ebfb080d5ee398ec3b66c4", "https://chartarchivefhir.com:4000/api/rest/v1/")
// 			fmt.Printf("TestDetermineResource:566  --  Found Resource: %s\n", resource)
// 			So(resource, ShouldEqual, "Patient")
// 		})
// 		Convey("Subject: Determine Resource of search from url", func() {
// 			godotenv.Load("./.env.uc_ca3_test")
// 			resource := DetermineResource("https://chartarchivefhir.com:4000/api/rest/v1/Patient?family=smart", "https://chartarchivefhir.com:4000/api/rest/v1/")
// 			fmt.Printf("TestDetermineResource:572  --  Found Resource: %s\n", resource)
// 			So(resource, ShouldEqual, "Patient")
// 		})
// 		Convey("Subject: Bad resource search", func() {
// 			godotenv.Load("./.env.uc_ca3_test")
// 			resource := DetermineResource("https://chartarchivefhir.com:4000/api/rest/v1/Patient-family=smart", "https://chartarchivefhir.com:4000/api/rest/v1/")
// 			fmt.Printf("TestDetermineResource:578  --  Found Resource: %s\n", resource)
// 			So(resource, ShouldEqual, "")
// 		})
// 	})
// }

// // func CreateCP() *uc_core/common.ConnectorPayload {
// // 	sc := uc_core/common.SystemConfig{}
// // 	sc.Name = "uc_ca3"
// // 	ids := []*uc_core/common.KVData{}

// // 	mrnIdent := uc_core/common.KVData{}
// // 	mrnIdent.Name = "mrn"
// // 	mrnIdent.Value = "mrn|"
// // 	ids = append(ids, &mrnIdent)
// // 	ssnIdent := uc_core/common.KVData{}
// // 	ssnIdent.Name = "ssn"
// // 	ssnIdent.Value = "ssn|"
// // 	ids = append(ids, &ssnIdent)
// // 	idIdent := uc_core/common.KVData{}
// // 	idIdent.Name = "id"
// // 	idIdent.Value = "id|"
// // 	ids = append(ids, &idIdent)
// // 	sc.Identifiers = ids

// // 	cp := uc_core/common.ConnectorPayload{}
// // 	cc := uc_core/common.ConnectorConfig{}
// // 	cc.ID, _ = primitive.ObjectIDFromHex("62f1c5dab3070d0b40e7aac1")
// // 	cc.Name = "uc_ca3"
// // 	cc.Version = "local"
// // 	cc.Label = "CA3FhirConnector"
// // 	cc.Credentials = ""
// // 	cc.HostUrl = "http://192.168.1.148:4100/system/640ba66cbd4105586a6dda75"
// // 	cc.URL = "192.168.1.152:20103"
// // 	data := []*uc_core/common.KVData{}
// // 	cacheServer := uc_core/common.KVData{}
// // 	cacheServer.Name = "cacheServer"
// // 	cacheServer.Value = "http://universalcharts.com:30201"
// // 	data = append(data, &cacheServer)
// // 	hostServer := uc_core/common.KVData{}
// // 	hostServer.Name = "cacheHost"
// // 	hostServer.Value = "http://ucCache:9200"
// // 	data = append(data, &hostServer)
// // 	cc.Data = data
// // 	cc.CacheUrl = "http://universalcharts.com:30201"

// // 	//TODO: AddFhirAuthToken
// // 	cp.ConnectorConfig = &cc
// // 	cp.System = &sc
// // 	cp.SavePayload = &uc_core/common.SavePayload{}
// // 	cp.SavePayload.SrcResource = SamplePatient()
// // 	cp.SavePayload.ResourceType = "Patient"
// // 	cp.SavePayload.SrcPatient = SampleFhirPatient()

// // 	//cp.BaseAddress = "http://192.168.1.148:4100"
// // 	// fmt.Printf("CreateCP:617  --  returning cp: %s\n\n", spew.Sdump(cp))
// // 	// fmt.Println()
// // 	// finalJson, err := json.Marshal(cp)
// // 	// if err != nil {
// // 	// 	fmt.Printf("CreateCP:621  --  json.Marshal error: %s\n", err.Error())
// // 	// 	return nil
// // 	// }
// // 	//fmt.Printf("\n\n\nCreateCP:623  --  finalJson: %s\n", finalJson)
// // 	fmt.Printf("\n\n\n")
// // 	return &cp
// // }

// // func CreateJWToken() (string, error) {
// // 	err := os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
// // 	if err != nil {
// // 		return "", err
// // 	}
// // 	os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
// // 	dur := time.Duration(300) * time.Second
// // 	jwt, err := token.CreateToken("192.168.1.2", "DHarman", dur, "dharman0127", "Debbie Harman", "Physician")
// // 	//maker, err := token.NewJWTMaker(os.Getenv("ACCESS_SECRET"))
// // 	if err != nil {
// // 		return "", err
// // 	}
// // 	return jwt, nil
// // }

// func GetSystemConfigById(id primitive.ObjectID) (*common.SystemConfig, error) {
// 	fmt.Printf("GetSystemConfig:12 --   Id: %s\n", id)
// 	collection, err := GetCollection("systemConfig")
// 	if err != nil {
// 		return nil, err
// 	}
// 	filter := bson.D{bson.E{"_id", id}}
// 	sysConfig := &common.SystemConfig{}
// 	fmt.Printf("GetSystemConfig:24  --  Calling FindOne SystemConfig: %v\n", filter)
// 	err = collection.FindOne(context.Background(), filter).Decode(&sysConfig)
// 	return sysConfig, err
// }

// func CreateCP(includeSave bool) *common.ConnectorPayload {
// 	fmt.Printf("CreateCP:248  --  includeSave: %v\n", includeSave)
// 	id, _ := primitive.ObjectIDFromHex("640ba5e3bd4105586a6dda74")
// 	sc, err := GetSystemConfigById(id)
// 	if err != nil {
// 		log.Errorf("CreateCP:663  --  GetSystemConfigById error: %s\n", err.Error())
// 		return nil
// 	}
// 	// sc := uc_core/common.SystemConfig{}
// 	// sc.Name = "uc_ca3"
// 	// ids := []*uc_core/common.KVData{}

// 	// mrnIdent := uc_core/common.KVData{}
// 	// mrnIdent.Name = "mrn"
// 	// mrnIdent.Value = "mrn|"
// 	// ids = append(ids, &mrnIdent)
// 	// ssnIdent := uc_core/common.KVData{}
// 	// ssnIdent.Name = "ssn"
// 	// ssnIdent.Value = "ssn|"
// 	// ids = append(ids, &ssnIdent)
// 	// idIdent := uc_core/common.KVData{}
// 	// idIdent.Name = "id"
// 	// idIdent.Value = "id|"
// 	// ids = append(ids, &idIdent)
// 	// sc.Identifiers = ids

// 	cp := common.ConnectorPayload{}
// 	cc := common.ConnectorConfig{}
// 	cc.ID, _ = primitive.ObjectIDFromHex("62f1c5dab3070d0b40e7aac1")
// 	cc.Name = "uc_ca3"
// 	cc.Version = "local"
// 	cc.Label = "CA3FhirConnector"
// 	cc.Credentials = ""
// 	cc.HostUrl = "http://192.168.1.148:4100/system/640ba66cbd4105586a6dda75"
// 	cc.URL = "192.168.1.152:20103"
// 	data := []*common.KVData{}
// 	cacheServer := common.KVData{}
// 	cacheServer.Name = "cacheServer"
// 	cacheServer.Value = "http://universalcharts.com:30201"
// 	data = append(data, &cacheServer)
// 	hostServer := common.KVData{}
// 	hostServer.Name = "cacheHost"
// 	hostServer.Value = "http://ucCache:9200"
// 	data = append(data, &hostServer)
// 	cc.Data = data
// 	cc.CacheUrl = "http://universalcharts.com:30201"

// 	//TODO: AddFhirAuthToken
// 	cp.ConnectorConfig = &cc
// 	cp.System = sc
// 	if includeSave {
// 		cp.SavePayload = &common.SavePayload{}
// 		cp.SavePayload.SrcResource = SamplePatient()
// 		cp.SavePayload.ResourceType = "Patient"
// 		cp.SavePayload.SrcPatient = SampleFhirPatient()
// 	}

// 	//cp.BaseAddress = "http://192.168.1.148:4100"
// 	// fmt.Printf("CreateCP:617  --  returning cp: %s\n\n", spew.Sdump(cp))
// 	// fmt.Println()
// 	// finalJson, err := json.Marshal(cp)
// 	// if err != nil {
// 	// 	fmt.Printf("CreateCP:621  --  json.Marshal error: %s\n", err.Error())
// 	// 	return nil
// 	// }
// 	//fmt.Printf("\n\n\nCreateCP:623  --  finalJson: %s\n", finalJson)
// 	fmt.Printf("\n\n\n")
// 	return &cp
// }

// func CreateCernerCP(includeSave bool) *common.ConnectorPayload {
// 	fmt.Printf("CreateCernerCP:744  --  includeSave: %v\n", includeSave)
// 	id, _ := primitive.ObjectIDFromHex("640ba5e3bd4105586a6dda74")
// 	sc, err := GetSystemConfigById(id)
// 	if err != nil {
// 		fmt.Printf("CreateCernerCP:729  --  GetSystemConfigById error: %s\n", err.Error())
// 		return nil
// 	}
// 	//ids := []*uc_core/common.KVData{}

// 	// mrnIdent := uc_core/common.KVData{}
// 	// mrnIdent.Name = "mrn"
// 	// mrnIdent.Value = "mrn|"
// 	// ids = append(ids, &mrnIdent)
// 	// ssnIdent := uc_core/common.KVData{}
// 	// ssnIdent.Name = "ssn"
// 	// ssnIdent.Value = "ssn|"
// 	// ids = append(ids, &ssnIdent)
// 	// idIdent := uc_core/common.KVData{}
// 	// idIdent.Name = "id"
// 	// idIdent.Value = "id|"
// 	// ids = append(ids, &idIdent)
// 	// sc.Identifiers = ids

// 	cp := common.ConnectorPayload{}
// 	cc := common.ConnectorConfig{}
// 	cc.ID, _ = primitive.ObjectIDFromHex("6488a9580403ff647fca2f7e")
// 	cc.Name = "uc_cerner"
// 	cc.Version = "local"
// 	cc.Label = "CernerConnector"
// 	cc.Credentials = ""
// 	cc.HostUrl = "https://fhir-open.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d"
// 	//cc.HostUrl = "http://192.168.1.148:4100/system/640ba66cbd4105586a6dda75"
// 	cc.URL = "192.168.1.152:20103"
// 	data := []*common.KVData{}
// 	cacheServer := common.KVData{}
// 	cacheServer.Name = "cacheServer"
// 	cacheServer.Value = "http://universalcharts.com:30201"
// 	data = append(data, &cacheServer)
// 	hostServer := common.KVData{}
// 	hostServer.Name = "cacheHost"
// 	hostServer.Value = "http://ucCache:9200"
// 	data = append(data, &hostServer)
// 	cc.Data = data
// 	cc.CacheUrl = "http://universalcharts.com:30201"

// 	//TODO: AddFhirAuthToken
// 	cp.ConnectorConfig = &cc
// 	cp.System = sc
// 	if includeSave {
// 		cp.SavePayload = &common.SavePayload{}
// 		cp.SavePayload.SrcResource = SamplePatient()
// 		cp.SavePayload.ResourceType = "Patient"
// 		cp.SavePayload.SrcPatient = SampleFhirPatient()
// 	}

// 	//cp.BaseAddress = "http://192.168.1.148:4100"
// 	// fmt.Printf("CreateCP:617  --  returning cp: %s\n\n", spew.Sdump(cp))
// 	// fmt.Println()
// 	// finalJson, err := json.Marshal(cp)
// 	// if err != nil {
// 	// 	fmt.Printf("CreateCP:383  --  json.Marshal error: %s\n", err.Error())
// 	// 	return nil
// 	// }
// 	//fmt.Printf("\n\n\nCreateCP:386  --  finalJson: %s\n", finalJson)
// 	fmt.Printf("\n\n\n")
// 	return &cp
// }

// // func TestCernerResourceSearch(t *testing.T) {
// // 	fmt.Printf("Test run a FHIR query")
// // 	//c := New(baseurl)
// // 	Convey("RunCernerDocumentResourceQuery", t, func() {
// // 		os.Setenv("CONFIG_ADDRESS", "http://192.168.1.148:30100/api/rest/v1")
// // 		// This should not be used. It is now the ConnectorPayload
// // 		// _, err := GetServiceConfig("uc_ca3", "local_test", "test") //GetConfig("delivery", "test")
// // 		// So(err, ShouldBeNil)
// // 		cp := CreateCernerCP(false)
// // 		So(cp, ShouldNotBeNil)

// // 		userId := "62d0af5dec383ade03a96b7e"
// // 		//userID, err := primitive.ObjectIDFromHex("62d0af5dec383ade03a96b7e")
// // 		//So(err, ShouldBeNil)

// // 		err := os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
// // 		So(err, ShouldBeNil)
// // 		maker, err := token.NewJWTMaker(os.Getenv("ACCESS_SECRET"))
// // 		So(err, ShouldBeNil)
// // 		So(maker, ShouldNotBeNil)
// // 		username := "dHarman"
// // 		duration := time.Minute
// // 		//userId := "user123456"
// // 		role := "Provider"
// // 		ip := "192.168.1.1.99"
// // 		fullName := "Debbie Harman MD"
// // 		//issuedAt := time.Now()
// // 		//expiredAt := issuedAt.Add(duration)

// // 		newToken, payload, err := maker.CreateToken(ip, username, duration, userId, fullName, role)
// // 		So(err, ShouldBeNil)
// // 		So(newToken, ShouldNotBeNil)
// // 		So(payload, ShouldNotBeNil)
// //TODO:  Add user Id toAll handlers which will get it fom the header

// // Use id of known patient
// // _, bundle, _, err := SearchPatient(fhirSystem, "Patient", userId, "Patient?family=dawg&given=joel")
// // //_, bundle, _, err := SearchPatient(fhirSystem, "Patient", userId, "Patient?family=smart&given=na")
// // //bundle, err := c.PatientSearch(caFhirId, "family=smart&given=na&_count=12")
// // So(err, ShouldBeNil)
// // So(bundle, ShouldNotBeNil)
// // //fmt.Printf("bundle: %s\n", spew.Sdump(bundle.Entry[0]))
// // // pat := fhir.Patient{}
// // // json.Unmarshal(bundle.Entry[0], &pat)
// // // fmt.Printf("patient[0] = %s\n", spew.Sdump(pat))

// // //fmt.Printf("PatientSearch returned: %s\n", spew.Sdump(bundle))
// // // data, err := c.Query("Patient/12724066")
// // // So(err, ShouldBeNil)
// // // So(data, ShouldNotBeNil)

// // pat, err := fhir.UnmarshalPatient(bundle.Entry[0].Resource)
// // So(err, ShouldBeNil)
// // So(pat, ShouldNotBeNil)

// //fmt.Printf("PatientSearch[0] returned: %s\n", spew.Sdump(pat))
// //patientId := "642e0767243f6740a3847c1c"
// //log.Printf("TestResourceSearch:854 -- Patient.ID = %s\n\n\n\n", patientId)
// //fmt.Printf("Patient: - %s\n", spew.Sdump(pat))
// //time.Sleep(10 * time.Second)
// // //Get Documents for this patient

// // qry := "/DocumentReference?patient=" + patientId
// // log.Printf("TestResourceSearch:860  --  Process DocumentReference")
// // cnt, bundle, hdr, err := FindResource(cp, "DocumentReference", userId, qry, newToken)
// // So(err, ShouldBeNil)
// // So(bundle, ShouldNotBeNil)
// // So(cnt, ShouldNotEqual, 0)
// // So(hdr, ShouldNotBeNil)
// // //fmt.Printf("TestResourceSearch:74  --  Number of entries = %d\n", len(bundle.Entry))
// // //fmt.Printf("TestResourceSearch:75  --  Entry = %s\n", spew.Sdump(bundle.Entry))
// // //rawDoc := bundle.Entry[0].Resource
// // //doc, _ := fhir.UnmarshalDocumentReference(rawDoc)
// // //fmt.Printf("TestResourceSearch:78  --  Doc: %s\n", spew.Sdump(doc))
// // //fmt.Printf("Doc[0] = %s\n", spew.Sdump(doc))

// // // func CachePatient(ctx context.Context, queryId string, userID primitive.ObjectID,
// // // 	patientId string, fhirSystem *common.FhirSystem, resource *Interface,
// // // 	resourceType string) error {

// // //CacheResource(context.Background(), hdr.QueryId, userId, patientId, fhirSystem, &doc, "Documentreference", doc)

// // qry = "/Observation?patient=" + patientId
// // cnt, bundle, _, err = FindResource(cp, "Observation", userId, qry, newToken)
// // So(err, ShouldBeNil)
// // So(bundle, ShouldNotBeNil)
// // So(cnt, ShouldNotEqual, 0)
// // // rawObs := bundle.Entry[0].Resource
// // // obs, _ := fhir.UnmarshalObservation(rawObs)
// // //fmt.Printf("Observation[0] = %s\n", spew.Sdump(obs))
// // qry = fmt.Sprintf("/Condition?patient=%s", "12743119")
// // //qry = fmt.Sprintf("/Condition?patient=%s", patientId)
// // cnt, bundle, _, err = FindResource(cp, "Condition", userId, qry, newToken)
// // So(err, ShouldBeNil)
// // So(bundle, ShouldNotBeNil)
// // So(cnt, ShouldNotEqual, 0)

// //})
// //}

// func TestCernerPatientResourceSearch(t *testing.T) {
// 	fmt.Printf("TestCernerPatientResourceSearch:915")
// 	//c := New(baseurl)
// 	os.Setenv("COMPANY", "test")
// 	mongo := OpenDBUrl("mongodb+srv://dhfadmin:Sacj0nhati@cluster1.24b12.mongodb.net/test?retryWrites=true&w=majority")
// 	if mongo == nil {
// 		log.Errorf("TestCernerSearchPatient:294  --  Unable to connect to Mongo")
// 		return
// 	}
// 	Convey("RunCernerPatientResourceSearch", t, func() {
// 		os.Setenv("CONFIG_ADDRESS", "http://192.168.1.148:30100/api/rest/v1")
// 		// This should not be used. It is now the ConnectorPayload
// 		// _, err := GetServiceConfig("uc_ca3", "local_test", "test") //GetConfig("delivery", "test")
// 		// So(err, ShouldBeNil)
// 		req, err := http.NewRequest("GET", "/api/rest/v1/Patient?family=smart&given=w&_count=1", nil)
// 		So(err, ShouldBeNil)
// 		cp := CreateCernerCP(false)
// 		So(cp, ShouldNotBeNil)
// 		fmt.Printf("\nTestCernerPatientResourceSearch:930  --  cp: %s\n", spew.Sdump(cp))
// 		cpb, err := json.Marshal(cp)
// 		So(err, ShouldBeNil)
// 		cps := string(cpb)
// 		fmt.Printf("\nTestCernerPatientResourceSearch:934  --  cpb: %s\n", cpb)
// 		fmt.Printf("\nTestCernerPatientResourceSearch:935  --  cps: %s\n", cps)
// 		rc := io.NopCloser(strings.NewReader(cps))
// 		req.Body = rc

// 		err = os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
// 		So(err, ShouldBeNil)
// 		newToken, payload, err := jw_token.CreateTestJWToken("10s")
// 		So(err, ShouldBeNil)
// 		So(newToken, ShouldNotBeNil)
// 		So(payload, ShouldNotBeNil)
// 		req.Header.Set("Authorization", newToken)
// 		fmt.Printf("TestCernerPatientResourceSearch:953\n")
// 		rr := executeRequest(req)
// 		checkResponseCode(t, http.StatusOK, rr.Code)
// 		resp := rr.Result()
// 		defer resp.Body.Close()
// 		byte, err := io.ReadAll(resp.Body)
// 		So(err, ShouldBeNil)
// 		So(byte, ShouldNotBeNil)
// 		resResp := &common.ResourceResponse{}
// 		err = json.Unmarshal(byte, resResp)
// 		//patient, err := fhir.UnmarshalPatient(byte)
// 		So(err, ShouldBeNil)
// 		So(resResp, ShouldNotBeNil)
// 		So(resResp.Bundle, ShouldNotBeNil)
// 		So(resResp.Patients, ShouldNotBeNil)
// 		//So(resResp.Resources, ShouldBeNil)
// 		fmt.Printf("TestCernerPatientResourceSearch:965  --  resResp: %s\n", spew.Sdump(resResp))
// 		//_, bundle, _, err := SearchPatient(fhirSystem, "Patient", userId, "Patient?family=dawg&given=joel")

// 	})
// }

// func TestCernerResourceHandlerGet(t *testing.T) {
// 	Convey("TestCernerResourceHandlerGet", t, func() {
// 		fmt.Printf("Creating Get Request\n")
// 		os.Setenv("COMPANY", "test")
// 		mongo := OpenDBUrl("mongodb+srv://dhfadmin:Sacj0nhati@cluster1.24b12.mongodb.net/test?retryWrites=true&w=majority")
// 		if mongo == nil {
// 			log.Errorf("TestCernerSearchPatient:988  --  Unable to connect to Mongo")
// 			return
// 		}
// 		jwt, payload, err := jw_token.CreateTestJWToken("10s")
// 		So(err, ShouldBeNil)
// 		So(jwt, ShouldNotBeNil)
// 		So(payload, ShouldNotBeNil)
// 		req, err := http.NewRequest("GET", "/api/rest/v1/Patient/12792842", nil)
// 		So(err, ShouldBeNil)
// 		So(req, ShouldNotBeNil)
// 		req.Header.Set("Authorization", jwt)
// 		req.Header.Set("facility", "demo")
// 		req.Header.Set("Fhir_Version", "r4")
// 		cp := CreateCernerCP(false)
// 		cpb, err := json.Marshal(cp)
// 		So(err, ShouldBeNil)
// 		cps := string(cpb)
// 		fmt.Printf("\nTestResourceHandlerGet:1004  --  cps: %s\n", cps)
// 		rc := io.NopCloser(strings.NewReader(cps))
// 		req.Body = rc
// 		fmt.Printf("TestResourceHandlerGet:1007  --  url: %s\n", req.URL)
// 		rr := executeRequest(req)
// 		checkResponseCode(t, http.StatusOK, rr.Code)
// 		resp := rr.Result()
// 		defer resp.Body.Close()
// 		byte, err := io.ReadAll(resp.Body)
// 		So(err, ShouldBeNil)
// 		So(byte, ShouldNotBeNil)
// 		//bundle := fhir.Bundle{}
// 		patient, err := fhir.UnmarshalPatient(byte)
// 		So(err, ShouldBeNil)
// 		So(patient, ShouldNotBeNil)
// 		fmt.Printf("TestResourceHandlerGet:1019  --  Patient: %s\n", spew.Sdump(patient))
// 	})
// }
