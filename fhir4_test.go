package main

import (
	//log "github.com/sirupsen/logrus"
	//. "github.com/smartystreets/goconvey/convey"
	//"encoding/json"
	"fmt"
	//"os"
	"testing"
	//"time"

	"github.com/davecgh/go-spew/spew"
	fhir "github.com/dhf0820/fhir4"

	//"github.com/dhf0820/token"
	jw_token "github.com/dhf0820/golangJWT"
	//github.com/davecgh/go-spew/spew"
	//log "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

//const pid = "Tbt3KuCY0B5PSrJvCu2j-PlK.aiHsu2xUjUM8bWpetXoB"

//const ordercode = "8310-5"
//const baseurl = "https://open-ic.epic.com/FHIR/api/FHIR/DSTU2/"

// const pid = "63ed93c8bd78ae6b013a502b"
//const baseurl = "http://universalcharts.com:4000/api/rest/v1"

func TestQuery(t *testing.T) {
	fmt.Printf("Test run a FHIR query")
	c := New(baseurl, "application/json+fhir")
	Convey("Run a query", t, func() {
		newToken, payload, err := jw_token.CreateTestToken("10s")
		So(err, ShouldBeNil)
		So(newToken, ShouldNotBeNil)
		So(payload, ShouldNotBeNil)
		bundle, err := c.Query("/Patient/63ed93c8bd78ae6b013a502b", newToken)
		So(err, ShouldBeNil)
		So(bundle, ShouldNotBeNil)
		// patient, err := fhir.UnmarshalPatient(bundle.Entry[0])
		// So(err, ShouldBeNil)
		// So(patient, ShouldNotBeNil)
		// fmt.Printf("TestQuery:43  --  Patient: %s\n", spew.Sdump(patient))
		//fmt.Printf("Patient returned: %s\n", spew.Sdump(data))
	})
}

func TestDocumentReferenceQuery(t *testing.T) {
	fmt.Printf("\n\n\n\nFhir4Test:38  --  Test run a FHIR query")
	c := New(baseurl, "application/json+fhir")
	Convey("Run a query", t, func() {
		newToken, payload, err := jw_token.CreateTestToken("10s")
		So(err, ShouldBeNil)
		So(newToken, ShouldNotBeNil)
		So(payload, ShouldNotBeNil)
		data, err := c.Query("DocumentReference?patient=12724066", newToken)
		So(err, ShouldBeNil)
		So(data, ShouldNotBeNil)
		//fmt.Printf("Patient returned: %s\n", spew.Sdump(data))
	})
}

// func TestGetFhirPdf(t *testing.T) {
// 	Convey("Subject: GetFhirPdf", t, func() {
// 		fmt.Printf("TestGetFhirPDF\n")
// 		newToken, payload, err := createJWT()
// 		So(err, ShouldBeNil)
// 		So(newToken, ShouldNotBeNil)
// 		So(payload, ShouldNotBeNil)
// 		url := fmt.Sprintf("%s%s%s", baseurl, "/Binary/XR-", "197198634")
// 		//url := "https://fhir-open.cerner.com/dstu2/ec2458f2-1e24-41c8-b71b-0e701af7583d/Patient?-pageContext=2d61b0b7-805d-4fd5-bb1d-a111f942f7a5&-pageDirection=NEXT"
// 		c := New(baseurl)

//			data, err := c.GetFhir(url, newToken)
//			So(err, ShouldBeNil)
//			So(data, ShouldNotBeNil)
//		})
//	}
//
// TestGetFhir returns a single Resource.
func TestGetFHIR(t *testing.T) {
	Convey("Subject: GetFHIR", t, func() {
		fmt.Printf("TestGetFHIR\n")
		newToken, payload, err := jw_token.CreateTestToken("10s")
		So(err, ShouldBeNil)
		So(newToken, ShouldNotBeNil)
		So(payload, ShouldNotBeNil)
		fmt.Printf("TestGetFHIR:86  --  baseURL: %s\n", baseurl)
		//url := "http://universalcharts.com:4000/api/rest/v1/Patient/180275"
		query := "/Patient/180275"
		c := New("http://192.168.1.152:30300/api/rest/v1", "application/json+fhir")
		//query := "/Patient?family=SMART&given=ANNE"
		//url := "/Patient?family=smart&given=ANNE"
		rawJson, err := c.GetFhir(query, "Patient", newToken)
		So(err, ShouldBeNil)
		So(rawJson, ShouldNotBeNil)
		// bundle, err := fhir.UnmarshalBundle(data)
		// So(err, ShouldBeNil)
		//fmt.Printf("TestGetFhir:96  --  Bundle: %s\n", spew.Sdump(bundle))
		//resource := bundle.Entry[0].Resource
		//fmt.Printf("TestGetFhir:99  --  raw patient: %s\n", resource)
		//patient := fhir.Patient{}
		//err = json.NewDecoder(resource).Decode(patient)
		//patient, err := fhir.UnmarshalPatient(bundle.Entry[0].Resource)
		patient, err := fhir.UnmarshalPatient(rawJson)
		So(err, ShouldBeNil)
		So(patient, ShouldNotBeNil)
		fmt.Printf("TestGetFhir:106  --  Patient: %s\n", spew.Sdump(patient))
		fmt.Printf("TestGetFhir:107  --  Patient.Name: %s\n", spew.Sdump(patient.Name))
		//fhir.HumanName.Given

	})
}

// TestGetFhir returns a single Resource.
func TestGetSrcFHIRBundle(t *testing.T) {
	Convey("Subject: GetSrcFHIRBundle", t, func() {
		srcBaseUrl := "http://192.168.1.152:4000/api/rest/v1"
		fmt.Printf("TestGetSrcFHIR:116\n")
		newToken, payload, err := jw_token.CreateTestToken("10s")
		So(err, ShouldBeNil)
		So(newToken, ShouldNotBeNil)
		So(payload, ShouldNotBeNil)
		fmt.Printf("TestGetSrcFHIRBundle:122  --  baseURL: %s\n", srcBaseUrl)
		//url := "http://universalcharts.com:4000/api/rest/v1/Patient?family=SMART&given=ANNE"
		c := New(srcBaseUrl, "application/json+fhir")
		//query := "/Patient?family=SMART&given=ANNE"
		//url := "/Patient/63ed93c8bd78ae6b013a502b"
		url := "Patient?family=SMART&given=ANNE"
		bundle, err := c.GetFhirBundle(url, newToken)
		So(err, ShouldBeNil)
		So(bundle, ShouldNotBeNil)
		// bundle, err := fhir.UnmarshalBundle(data)
		// So(err, ShouldBeNil)
		//fmt.Printf("TestGetFhir:96  --  Bundle: %s\n", spew.Sdump(bundle))
		//resource := bundle.Entry[0].Resource
		//fmt.Printf("TestGetFhir:99  --  raw patient: %s\n", resource)
		//patient := fhir.Patient{}
		//err = json.NewDecoder(resource).Decode(patient)
		//patient, err := fhir.UnmarshalPatient(bundle.Entry[0].Resource)
		patient, err := fhir.UnmarshalPatient(bundle.Entry[0].Resource)
		So(err, ShouldBeNil)
		So(patient, ShouldNotBeNil)
		fmt.Printf("TestGetFhir:141  --  Patient: %s\n", spew.Sdump(patient))
		fmt.Printf("TestGetFhir:142  --  Patient.Name: %s\n", spew.Sdump(patient.Name))
		//fhir.HumanName.Given

	})
}

// func TestGetDocumentImage(t *testing.T) {

// 	Convey("Get the imag of a document", t, func() {
// 		//m.DeleteDocuments(session.CacheName)
// 		Convey("Authorized with a valid session", func() {
// 			url := "https://fhir-open.cerner.com/dstu2/ec2458f2-1e24-41c8-b71b-0e701af7583d/Binary/XR-197293272"
// 			c := New(baseurl)
// 			//bytes, err := c.GetFhir(url, "application/pdf")
// 			bytes, err := c.GetDiagnosticPDF(url)
// 			So(err, ShouldBeNil)
// 			So(bytes, ShouldNotBeNil)
// 			fmt.Printf("Writting the Datafile\n")
// 			if err := os.WriteFile("./debbie.data", bytes, 0666); err != nil {
// 				log.Fatal(err)
// 			}

// 		})
// 	})
// }

// func TestPatientDocumentReference(t *testing.T) {
// 	c := New(baseurl)
// 	Convey("Get GetpatientDocumentReference", t, func() {
// 		//https://fhir-open.sandboxcerner.com/dstu2/0b8a0111-e8e6-4c26-a91c-5069cbc6b1ca/DiagnosticReport?patient=1316020&_count=10
// 		data, err := c.GetPatientDiagnosticReports("12724066")
// 		So(err, ShouldBeNil)
// 		So(data, ShouldNotBeNil)
// 		fmt.Printf("Document: %s\n", spew.Sdump(data))
// 	})

// 	//data, err := c.GetDocumentReference(pid)

// }

// func TestPatientDiagnosticReports(t *testing.T) {
// 	c := New(baseurl)
// 	Convey("Get GetpatientDiagnosticReoprts", t, func() {
// 		//https://fhir-open.sandboxcerner.com/dstu2/0b8a0111-e8e6-4c26-a91c-5069cbc6b1ca/DiagnosticReport?patient=1316020&_count=10
// 		data, err := c.GetPatientDiagnosticReports("12724066")
// 		So(err, ShouldBeNil)
// 		So(data, ShouldNotBeNil)
// 		fmt.Printf("Document: %s\n", spew.Sdump(data))
// 	})

// 	//data, err := c.GetDocumentReference(pid)

// }
// func TestPatientDiagnosticReport(t *testing.T) {
// 	c := New(baseurl)
// 	// Provide the PatientID , DocumentID and DateRange of the document
// 	// R2 cerner/Epic Does not support direct query by DocumentID. Cerner does not have an R4 DiagRept

// 	Convey("Get GetpatientDiagnosticReoprts", t, func() {
// 		//https://fhir-open.sandboxcerner.com/dstu2/0b8a0111-e8e6-4c26-a91c-5069cbc6b1ca/DiagnosticReport?patient=1316020&_count=10
// 		data, err := c.GetPatientDiagnosticReports("4342009")
// 		So(err, ShouldBeNil)
// 		So(data, ShouldNotBeNil)
// 		//data, err := c.GetDocumentReference(pid)

// 		//fmt.Printf("Document: %s\n", spew.Sdump(data))
// 		// if data.Total == 0 {
// 		// 	t.Error("Expected > 0 got 0")
// 		// }
// 	})
// }

// func TestDocument(t *testing.T) {
// 	c := New(baseurl)
// 	data, err := c.GetDocumentReference("12724066")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Printf("Document: %s\n", spew.Sdump((data)))
// 	// if data.Total == 0 {
// 	// 	t.Error("Expected > 0 got 0")
// 	// }
// }

// // func TestCondition(t *testing.T) {
// // 	c := New(baseurl)
// // 	data, err := c.GetCondition(pid)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	if data.Total == 0 {
// // 		t.Error("Expected > 0 got 0")
// // 	}
// // }

// // func TestProcedure(t *testing.T) {
// // 	c := New(baseurl)
// // 	data, err := c.GetProcedure(pid)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	if data.Total == 0 {
// // 		t.Error("Expected > 0 got 0")
// // 	}
// // }

// // func TestMedication(t *testing.T) {
// // 	c := New(baseurl)
// // 	data, err := c.GetMedication(pid)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	if data.Total == 0 {
// // 		t.Error("Expected > 0 got 0")
// // 	}
// // }

// // func TestObservation(t *testing.T) {
// // 	c := New(baseurl)
// // 	data, err := c.GetObservation(pid, ordercode)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	if data.Total == 0 {
// // 		t.Error("Expected > 0 got 0")
// // 	}
// // }

// // func TestImmunization(t *testing.T) {
// // 	c := New(baseurl)
// // 	data, err := c.GetImmunization(pid)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	if data.Total == 0 {
// // 		t.Error("Expected > 0 got 0")
// // 	}
// // }

// // func TestAllergy(t *testing.T) {
// // 	c := New(baseurl)
// // 	data, err := c.GetAllergyIntolerence(pid)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	if data.Total == 0 {
// // 		t.Error("Expected > 0 got 0")
// // 	}
// // }

// // func TestFamilyHx(t *testing.T) {
// // 	c := New(baseurl)
// // 	data, err := c.GetFamilyMemberHistory(pid)
// // 	if err != nil {
// // 		t.Fatal(err)
// // 	}
// // 	if data.Total == 0 {
// // 		t.Error("Expected > 0 got 0")
// // 	}
// // }

// func createJWT() (string, *token.Payload, error) {
// 	os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
// 	maker, err := token.NewJWTMaker(os.Getenv("ACCESS_SECRET"))
// 	So(err, ShouldBeNil)
// 	So(maker, ShouldNotBeNil)
// 	username := "DHarman0127"
// 	duration := 5 * time.Minute
// 	userId := "user123456"
// 	role := "Provider"
// 	ip := "192.168.1.1.99"
// 	fullName := "Debbie Harman MD"
// 	newToken, payload, err := maker.CreateToken(ip, username, duration, userId, fullName, role)
// 	return newToken, payload, err

// }

func TestGetRemoteFhirPatient(t *testing.T) {
	Convey("Subject: GetRemoteFhirPatient", t, func() {
		fmt.Printf("TestGetRemoteFhirPatient:321\n")
		newToken, payload, err := jw_token.CreateTestToken("10s")
		So(err, ShouldBeNil)
		So(newToken, ShouldNotBeNil)
		So(payload, ShouldNotBeNil)
		fmt.Printf("TestGetRemoteFhirPatient:326  --  baseURL: %s\n", baseurl)
		query := "/Patient/63ed93c8bd78ae6b013a502b"
		//fhirUrl := "http://192.168.1.152:30300/api/rest/v1"
		ca3Url := "http://universalcharts.com:4000/api/rest/v1"
		//query := "/Patient?family=SMART&given=ANNE"
		//url := "/Patient?family=smart&given=ANNE"
		patient, err := GetRemoteFhirPatient(query, ca3Url, newToken)
		So(err, ShouldBeNil)
		So(patient, ShouldNotBeNil)
		//fmt.Printf("TestGetRemoteFhirPatient:335  --  Patient: %s\n", spew.Sdump(patient))
		fmt.Printf("TestGetRemoteFhirPatient:336  --  Patient.Name: %s\n", spew.Sdump(patient.Name))
		//fhir.HumanName.Given

	})
}
