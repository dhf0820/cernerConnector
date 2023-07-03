package main

// import (
// 	//"bytes"
// 	"encoding/json"
// 	"io"
// 	"io/ioutil"

// 	//"github.com/gorilla/mux"
// 	//"github.com/joho/godotenv"
// 	//"gitlab.com/dhf0820/ids_model/common"

// 	//log "github.com/sirupsen/logrus"
// 	//. "github.com/smartystreets/goconvey/convey"
// 	"fmt"
// 	"net/http"
// 	//"net/http/httptest"
// 	"os"
// 	"strings"
// 	"testing"
// 	"time"

// 	"github.com/dhf0820/token"
// 	//"github.com/dhf0820/uc_common"
// 	//"github.com/dhf0820/uc_core/service"

// 	//"github.com/davecgh/go-spew/spew"
// 	//fhir "github.com/dhf0820/fhir4"

// 	//fhirR4go "github.com/dhf0820/fhirR4go"
// 	//"go.mongodb.org/mongo-driver/bson/primitive"
// 	//log "github.com/sirupsen/logrus"
// 	. "github.com/smartystreets/goconvey/convey"
// )

// func TestHealthCheckHandler(t *testing.T) {
// 	Convey("TestHealthCheckHandler", t, func() {
// 		req, err := http.NewRequest("GET", "/api/rest/v1/healthcheck", nil)
// 		So(err, ShouldBeNil)

// 		os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
// 		dur := time.Duration(300) * time.Second
// 		newToken, payload, err := jw_token.CreateTestJWToken("10s")
// 		So(err, ShouldBeNil)
// 		So(jwt, ShouldNotBeNil)
// 		req.Header.Set("Authorization", jwt)
// 		fmt.Printf("\nCalling Router\n")
// 		cp := CreateCP()
// 		cpb, err := json.Marshal(cp)
// 		So(err, ShouldBeNil)
// 		cps := string(cpb)
// 		fmt.Printf("\nTestHealthCheckHandlerGet:52  --  cpb: %s\n", cpb)
// 		rc := io.NopCloser(strings.NewReader(cps))
// 		req.Body = rc
// 		fmt.Printf("TestHealthCheckHandlerGet:55  --  url: %s\n", req.URL)
// 		rr := executeRequest(req)
// 		checkResponseCode(t, http.StatusOK, rr.Code)
// 		resp := rr.Result()
// 		defer resp.Body.Close()
// 		byte, err := ioutil.ReadAll(resp.Body)
// 		So(err, ShouldBeNil)
// 		So(resp.StatusCode, ShouldEqual, 200)
// 		So(byte, ShouldNotBeNil)
// 		fmt.Printf("\nTestHealthCheckHandlerGet:63  --  byte: %s\n", byte)

// 		//fmt.Printf("\nTestSimpleFindResource:68  --  Patients: %s\n", spew.Sdump(resResp.Patients))
// 		//fmt.Printf("\nTestSimpleFindResource:69  --  Bundle: %s\n", spew.Sdump(resResp.Bundle))
// 	})
// }
