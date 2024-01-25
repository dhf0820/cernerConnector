package main

import (
	//"bytes"
	"encoding/json"
	"io"

	"github.com/joho/godotenv"

	//"github.com/gorilla/mux"
	//"github.com/joho/godotenv"
	//"gitlab.com/dhf0820/ids_model/common"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//log "github.com/sirupsen/logrus"
	//. "github.com/smartystreets/goconvey/convey"
	"fmt"
	"net/http"

	//"net/http/httptest"
	"os"
	"strings"
	"testing"

	//"time"

	//"github.com/dhf0820/token"
	common "github.com/dhf0820/uc_common"
	//"github.com/dhf0820/uc_core/service"
	jw_token "github.com/dhf0820/golangJWT"
	//"github.com/davecgh/go-spew/spew"
	//fhir "github.com/dhf0820/fhir4"

	//fhirR4go "github.com/dhf0820/fhirR4go"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	log "github.com/dhf0820/vslog"
	//"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/davecgh/go-spew/spew"
)

func TestSearchPatient(t *testing.T) {

	log.SetDebuglevel("DEBUG3")
	err := godotenv.Load("./.env.cerner_conn_test")
	if err != nil {
		log.Error("./.env.cerner_conn_test was not found")
		t.FailNow()
	}
	Conf, err = GetServiceConfig("cerner_conn", "ssd", "test")
	if err != nil {
		log.Error("GetServiceConfig failed: " + err.Error())
		t.FailNow()
	}
	log.Debug3("TestSearchPatient Conf: " + spew.Sdump(Conf))
	Convey("TestSearchPatient", t, func() {
		req, err := http.NewRequest("GET", "/api/rest/v1/Patient?family=harman", nil)
		So(err, ShouldBeNil)

		os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")

		jwt, payload, err := jw_token.CreateTestToken("10s")
		So(err, ShouldBeNil)
		So(jwt, ShouldNotBeNil)
		So(payload, ShouldNotBeNil)
		req.Header.Set("Authorization", jwt)
		log.Debug3("Calling Router")
		cp := CreateCP(false)
		fmt.Printf("\nTestSearchPatient:51  --  cp: %s\n", spew.Sdump(cp))
		cpb, err := json.Marshal(cp)
		So(err, ShouldBeNil)
		cps := string(cpb)
		fmt.Printf("\nTestSearchPatient:54  --  cpb: %s\n", cpb)
		rc := io.NopCloser(strings.NewReader(cps))
		req.Body = rc
		fmt.Printf("TestSearchPatient:55  --  url: %s\n", req.URL)
		rr := executeRequest(req)
		checkResponseCode(t, http.StatusOK, rr.Code)
		resp := rr.Result()
		defer resp.Body.Close()
		byte, err := io.ReadAll(resp.Body)
		So(err, ShouldBeNil)
		So(resp.StatusCode, ShouldEqual, 200)
		So(byte, ShouldNotBeNil)
		fmt.Printf("\nTestSearchPatient:66  --  byte: %s\n", byte)

		//fmt.Printf("\nTestSimpleFindResource:68  --  Patients: %s\n", spew.Sdump(resResp.Patients))
		//fmt.Printf("\nTestSimpleFindResource:69  --  Bundle: %s\n", spew.Sdump(resResp.Bundle))
	})
}

// func TestSPatient(t *testing.T) {
// 	Convey("Subject: SearchForPatient", t, func() {
// 		godotenv.Load("./.env.uc_ca3_test")
// 		os.Setenv("CONFIG_ADDRESS", "http://192.168.1.152:30300/api/rest/v1")
// 		//os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:30100/api/rest/v1")
// 		os.Setenv("COMPANY", "test")

// 		// conf, err := InitCore("uc_core", "go_test", "test")
// 		// log.Printf("\n\n\nInitCore returned\n\n\n")
// 		// if err != nil {
// 		// 	log.Printf("Err = %s\n\n", err.Error())
// 		// }
// 		// if conf == nil {
// 		// 	log.Println("Conf is nil")
// 		// }
// 		// So(err, ShouldBeNil)
// 		// So(conf, ShouldNotBeNil)

// 		_, err := Initialize("ca3_conn", "test")
// 		//, err := service.InitCore("uc_core","test", "test")
// 		if err != nil {
// 			t.Fatalf("TestSearchPatient:56  --  InitCore failed: %s", err.Error())
// 		}

// 		/*   Handle query for patient, create a cache header info for the queryquerying for results cache each element of the bundle
// 		in the back ground check on the size of the results aver couple of seconds.  When the number _count requested, Return them
// 		in a new standart cache results
// 		*/
// 		fmt.Printf("\n\n\nTestSearchPatient63  --  testing TestSearchFhirForPatient\n\n\n")
// 		Convey("Given a valid family/given name", func() {
// 			fmt.Printf("\n\nTestSearchPatient:65  --  Given a valid family/given name\n\n\n")
// 			w := httptest.NewRecorder()
// 			fmt.Printf("Creating Search Request\n")

// 			//req, _ := http.NewRequest("GET", "/634f0ec03240a53a52a83a9d/Patient?family=smart&given=fred&dob=1958-01-27&_count=2&_offset=10", nil)
// 			//req, _ := http.NewRequest("GET", "/634f0ec03240a53a52a83a9d/Patient?_id=12747609", nil)
// 			//req, _ := http.NewRequest("GET", "/634f0ec03240a53a52a83a9d/Patient?_id=12345678", nil)
// 			cp := CreateTestFileCloser()
// 			req, _ := http.NewRequest("GET", "/facility/6329112852f3616990e2f763/Patient?family=SMART&given=C&_count=2", cp)
// 			godotenv.Load("./.env.core_test")
// 			newToken, payload, err := createJWT()
// 			So(err, ShouldBeNil)
// 			So(newToken, ShouldNotBeNil)
// 			So(payload, ShouldNotBeNil)
// 			req.Header.Set("AUTHORIZATION", newToken)
// 			//fmt.Printf("n\nTest Request = %s\n\n\n", spew.Sdump(req))
// 			//req.Header.Add("tracing-id", "123")
// 			// vars := map[string]string{
// 			// 	"doc_id" : "40441",
// 			// }
// 			//req = mux.SetURLVars(req, vars)
// 			req.Header.Set("facility", "demo")
// 			//req.Header.Set("FhirVersion", "r4")
// 			//req.Header.Set("FhirSystemId", "6329112852f3616990e2f763")

// 			fmt.Printf("TestSearchPatient:90  --  Calling findPatient URL = %s\n", req.URL)
// 			findPatient(w, req)
// 			resp := w.Result()
// 			So(resp.StatusCode, ShouldEqual, http.StatusOK)
// 			//var bundle commmon.ResourceResponse{}  // fhir.Bundle
// 			var resResp common.ResourceResponse
// 			fmt.Printf("TestSearchPatient:95 --  Resp = %s\n", spew.Sdump(resp))
// 			err = json.NewDecoder(resp.Body).Decode(&resResp)
// 			So(err, ShouldBeNil)
// 			//fmt.Printf("TestSearchPatient:99  --  ResourceResponse: %s\n", spew.Sdump(resResp))

// 		})
// 	})
// }

func TestPatientGet(t *testing.T) {
	log.SetDebuglevel("DEBUG3")
	Convey("Subject: GetPatient", t, func() {
		godotenv.Load("./.env.uc_ca3_conn_test")
		// os.Setenv("CONFIG_ADDRESS", "http://192.168.1.117:30300/api/rest/v1")
		// //os.Setenv("CONFIG_ADDRESS", "http://universalcharts.com:30100/api/rest/v1")
		// os.Setenv("COMPANY", "test")

		// conf, err := service.InitCore("uc_core", "go_test", "test")
		// log.Printf("\n\n\nInitCore returned\n\n\n")
		// if err != nil {
		// 	log.Printf("Err = %s\n\n", err.Error())
		// }
		// if conf == nil {
		// 	log.Println("Conf is nil")
		// }
		// So(err, ShouldBeNil)
		// So(conf, ShouldNotBeNil)

		// _, err := service.Initialize()
		// //, err := service.InitCore("uc_core","test", "test")
		// if err != nil {
		// 	t.Fatalf("InitCore failed: %s", err.Error())
		// }

		///   Handle query for patient, create a cache header info for the queryquerying for results cache each element of the bundle
		//in the back ground check on the size of the results aver couple of seconds.  When the number _count requested, Return them
		// in a new standart cache results

		fmt.Printf("\n\n\ntesting TestGetPatient\n\n\n")
		Convey("Given a valid patient ID", func() {
			fmt.Printf("\n\nTestPatientGet:173  --  Given a valid PatientId\n\n\n")

			fmt.Printf("TestPatientGet:175  --  Creating Get Request\n")
			//req := httptest.NewRequest("GET", "/62f14531ba5395278cd530c4/Patient/12724066", nil)
			//req, _ := http.NewRequest("GET", "/api/rest/v1/Patient?family=smart&given=fred&_count=2", nil)
			godotenv.Load("./.env_uc_ca3_conn_test")
			jwt, payload, err := jw_token.CreateTestToken("10s")
			So(err, ShouldBeNil)
			So(jwt, ShouldNotBeNil)
			So(payload, ShouldNotBeNil)

			//fmt.Printf("n\nTest Request = %s\n\n\n", spew.Sdump(req))
			//req.Header.Add("tracing-id", "123")
			// vars := map[string]string{
			// 	"doc_id" : "40441",
			// }
			//req = mux.SetURLVars(req, vars)
			//fmt.Printf("testGetPatient:180  --  req: %s\n", spew.Sdump(req))
			//connectorPayload := common.ConnectorPayload{}
			//fs, err := GetFhirSystem("62f14531ba5395278cd530c4")
			//So(err, ShouldBeNil)
			//So(fs, ShouldNotBeNil)
			cp := CreateCP(false)
			cpb, err := json.Marshal(cp)
			So(err, ShouldBeNil)
			cps := string(cpb)
			fmt.Printf("\nTestPatientGet:198  --  cpb: %s\n", cpb)
			rc := io.NopCloser(strings.NewReader(cps))

			//cp := CreateTestFileCloser()
			// cp := CreateCP()
			// fs := common.FhirSystem{}
			// //fs.FhirUrl = "https://fhir-open.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d"
			// fs.FhirUrl = "/api/rest/v1"
			// fs.UcUrl = "http://192.168.1.152:4000/6329112852f3616990e2f763"
			// fs.FacilityName = "HarmanClinic"
			// fs.DisplayName = "ChartArchive"
			// fs.FacilityCode = "harman"
			// fs.FhirVersion = "r4"
			// fs.Identifiers = []*common.KVData{}
			// ident := common.KVData{}
			// ident.Name = "mrn"
			// ident.Value = "urn:oid:2.16.840.1.113883.6.1000|"
			// fs.Identifiers = append(fs.Identifiers, &ident)
			// fs.ID, err = primitive.ObjectIDFromHex("6329112852f3616990e2f763")
			So(err, ShouldBeNil)
			So(rc, ShouldNotBeNil)
			// fs.FacilityId, err = primitive.ObjectIDFromHex("6329112852f3616990e2f763")
			// fs.Insert = "false"
			So(err, ShouldBeNil)

			//connectorPayload.FhirSystem = &fs
			//requestBody, err := json.Marshal(cp)
			// cpb, err := json.Marshal(cp)
			// So(err, ShouldBeNil)
			// cps := string(cpb)
			// log.Printf("\nTestPatientGet:221  --  cpb: %s\n", cpb)
			// rc := io.NopCloser(strings.NewReader(cps))
			//req, err := http.NewRequest("GET", "/facility/6329112852f3616990e2f763/Patient/63c5d9289ab9792cc6d2e22e", bytes.NewBuffer(requestBody))
			//req, err := http.NewRequest("GET", "/system/640ba66cbd4105586a6dda75/Patient/642e0767243f6740a3847c1c", nil)

			req, err := http.NewRequest("GET", "/api/rest/v1/Patient/642e3f225b444c831e8ff529", nil)
			fmt.Printf("TestPatientGet:234  --  url: %s\n", req.URL)
			req.Body = rc
			So(err, ShouldBeNil)
			req.Header.Set("Authorization", jwt)
			fmt.Printf("TestPatientGet:238   --  url = %s\n", req.URL)
			rr := executeRequest(req)
			checkResponseCode(t, http.StatusOK, rr.Code)

			//fmt.Printf("Setting Headers\n")
			// req.Header.Set("facility", "demo")
			// req.Header.Set("FhirVersion", "r4")
			// req.Header.Set("FhirSystemId", "62f14531ba5395278cd530c4")
			// req.Header.Set("AUTHORIZATION", newToken)

			//h := map[string][]string{
			// "facility":    {"demo"},
			// "FhirVersion": {"r4"},
			// //"FhirSystemId":  {"6329112852f3616990e2f763"},
			// "Authorization": {jwt},
			//}
			//req.Header = h
			// fmt.Printf("TestPatientGet:195  --  Calling execute\n")
			// rr := executeRequest(req)
			// checkResponseCode(t, http.StatusOK, rr.Code)
			// resp := rr.Result()
			// So(resp.StatusCode, ShouldEqual, 200)
			//w := httptest.NewRecorder()
			// // if req == nil {
			// // 	fmt.Printf("req is nil\n")
			// // }
			// // if w == nil {
			// // 	fmt.Printf("w is nil\n")
			// // }

			// // //fmt.Printf("testGetPatient:228  --  req: %s\n", spew.Sdump(req))
			// fmt.Printf("testGetPatient:260  --  Calling Router\n")
			// NewRouter().ServeHTTP(w, req)
			// // // //getPatient(w, req)
			// fmt.Printf("testGetPatient:263  --  router Returned\n")
			// resp := w.Result()
			fmt.Printf("TestPatientGet:274  --  resp: %d\n", rr.Result().StatusCode)
			So(rr.Result().StatusCode, ShouldEqual, http.StatusOK)
			//var bundle []fhir.Bundle
			var patient common.ResourceResponse
			//fmt.Printf("testGetPatient:229  --  Resp = %s\n", spew.Sdump(resp))
			//err = json.NewDecoder(resp.Body).Decode(&bundle)
			err = json.NewDecoder(rr.Result().Body).Decode(&patient)
			So(err, ShouldBeNil)
			fmt.Printf("TestPatientGet:283  --  Patient: %s\n", spew.Sdump(patient))

		})
	})
}

func TestCernerSearchPatient(t *testing.T) {
	log.SetDebuglevel("DEBUG3")
	Convey("TestCernerSearchPatient", t, func() {
		os.Setenv("COMPANY", "test")
		mongo := OpenDBUrl("mongodb+srv://dhfadmin:Sacj0nhati@cluster1.24b12.mongodb.net/test?retryWrites=true&w=majority")
		if mongo == nil {
			log.Errorf("TestCernerSearchPatient:294  --  Unable to connect to Mongo")
			return
		}
		req, err := http.NewRequest("GET", "/api/rest/v1/Patient?family=smart&given=w&_count=1", nil)
		So(err, ShouldBeNil)

		os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
		jwt, payload, err := jw_token.CreateTestToken("10s")
		So(err, ShouldBeNil)
		So(jwt, ShouldNotBeNil)
		So(payload, ShouldNotBeNil)
		req.Header.Set("Authorization", jwt)
		fmt.Printf("\nTestCernerSearchPatient:306  --  Calling Router\n")
		cp := CreateCernerCP(false)
		fmt.Printf("\nTestCernerSearchPatient:308  --  cp: %s\n", spew.Sdump(cp))
		cpb, err := json.Marshal(cp)
		So(err, ShouldBeNil)
		cps := string(cpb)
		fmt.Printf("\nTestCernerSearchPatient:312  --  cpb: %s\n", cpb)
		rc := io.NopCloser(strings.NewReader(cps))
		req.Body = rc
		fmt.Printf("TestCernerSearchPatient:315  --  url: %s\n", req.URL)
		rr := executeRequest(req)
		checkResponseCode(t, http.StatusOK, rr.Code)
		resp := rr.Result()
		defer resp.Body.Close()
		byte, err := io.ReadAll(resp.Body)
		So(err, ShouldBeNil)
		So(resp.StatusCode, ShouldEqual, 200)
		So(byte, ShouldNotBeNil)
		fmt.Printf("\nTestCernerSearchPatient:324  --  cps: %s\n", cps)
		resResp := common.ResourceResponse{}
		//bundle := fhir.Bundle{}
		fmt.Printf("TestSearchPatient:327 --  Resp = %s\n", spew.Sdump(resp))
		err = json.NewDecoder(resp.Body).Decode(&resResp)
		So(err, ShouldBeNil)
		So(resp.StatusCode, ShouldEqual, 200)
		So(resResp, ShouldNotBeNil)
		fmt.Printf("\nTestCernerSearchPatient:332  --  resResp: %s\n", spew.Sdump(resResp))

		//fmt.Printf("\nTestSimpleFindResource:68  --  Patients: %s\n", spew.Sdump(resResp.Patients))
		//fmt.Printf("\nTestSimpleFindResource:69  --  Bundle: %s\n", spew.Sdump(resResp.Bundle))
	})
}
