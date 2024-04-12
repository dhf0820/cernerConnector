package main

import (
	//"encoding/json"
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/davecgh/go-spew/spew"
	fhir "github.com/dhf0820/fhir4"
	"net"
	"net/http"
	//"os"

	//common "github.com/dhf0820/uc_core/common"
	"strings"
	//"github.com/davecgh/go-spew/spew"
	"time"

	log "github.com/dhf0820/vslog"
	"github.com/sirupsen/logrus"
)

const (
	timeout = 15
)

var (
	//err error
	body []byte
)

// RetData is the mapped json of the request
type RetData map[string]interface{}

// Connection is a FHIR connection
type Connection struct {
	BaseURL string
	client  *http.Client
	Accept  string
}

// TODO: Consider including the FhirSystem in the Connection
// New creates a new connection
func New(baseurl, accept string) *Connection {
	return &Connection{
		BaseURL: baseurl,
		Accept:  accept,
		client: &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   time.Duration(timeout*120) * time.Second,
					KeepAlive: time.Duration(timeout*120) * time.Second,
				}).Dial,
				TLSHandshakeTimeout:   time.Duration(timeout) * time.Second,
				ResponseHeaderTimeout: time.Duration(timeout) * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		},
	}
}

// Query sends a query to the base url
func (c *Connection) Query(q, token string) (*fhir.Bundle, error) {
	fmt.Printf("\n\n\n\nQuery:62  --  BaseUrl: %s  -  Query param: %s\n\n\n\n", c.BaseURL, q)
	if q == "" {
		return nil, fmt.Errorf("c.Query:64  --  query parameter missing")
	}
	url := fmt.Sprintf("%s/%s", c.BaseURL, q)
	//fmt.Printf("fhir4_query:60  --  c.BaseUrl = %s\n", c.BaseURL)
	fmt.Printf("c.Query:68  -- url = %s\n", url)

	return c.GetFhirBundle(q, token)
	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	return nil, err
	// }
	// req.Header.Add("Accept", "application/json+fhir")
	// //fmt.Println("Calling the GET request")
	// resp, err := c.client.Do(req)
	// if err != nil {
	// 	log.Errorf(" !!!fhir query returned err: %s\n", err)
	// 	return nil, err
	// }
	// //fmt.Printf("resp: %s\n", spew.Sdump(resp))
	// //defer resp.Body.Close()
	// if resp.StatusCode < 200 || resp.StatusCode > 299 {
	// 	err = fmt.Errorf("%d|%s", resp.StatusCode, string(body))
	// 	return nil, err
	// }
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("Query Error: %v\n", err)
	// 	return nil, err
	// }

	// //fmt.Printf("fhir length of Body: %d\n", len(body))
	// return body, nil
}

// func (c *Connection) GetById(id string)([]byte, error ) {
// }
func (c *Connection) GetFhirBytes(qry string, resourceType, token string) ([]byte, string, int, error) {
	log.Debug3(fmt.Sprintf("ResourceType = %s Query = %s  BaseUrl = %s", resourceType, qry, c.BaseURL))
	fullUrl := fmt.Sprintf("%s/%s/%s", c.BaseURL, resourceType, qry)
	log.Debug3("--  FullURL Requested: " + fullUrl)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		errMsg := log.ErrMsg("NewRequest failed: " + err.Error())
		log.Debug3(errMsg)
		return nil, resourceType, 400, log.Errorf(errMsg)
	}
	req.Header.Set("Accept", "application/json+fhir")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, resourceType, 500, log.Errorf("!!!fhir query returned err: " + err.Error())
	}
	byte := []byte{}

	if resp.Body != nil {
		byte, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, resourceType, 400, log.Errorf("Response  --  Error readying body: " + err.Error())
		}
		//log.Debug3("Response --  Raw body: " + string(byte))
		switch resourceType {
		case "OperationOutcome":
			opOut, err := fhir.UnmarshalOperationOutcome(byte)
			if err != nil {
				log.Debug3("Response --  Error Decoding OperationOutcone: " + err.Error())
				return byte, resourceType, resp.StatusCode, log.Errorf("Response --  Error Decoding Patient: " + err.Error())
			}
			log.Debug3("Response --  OperationOutcome: " + spew.Sdump(opOut))

		case "Patient":
			patient, err := fhir.UnmarshalPatient(byte)
			if err != nil {
				log.Debug3("Response --  Error Decoding Patient: " + err.Error())
				return byte, resourceType, resp.StatusCode, log.Errorf("Response --  Error Decoding Patient: " + err.Error())
			}
			log.Debug5("Response --  Patient: " + spew.Sdump(patient))
		case "DocumentReference":
			docRef, err := fhir.UnmarshalDocumentReference(byte)
			if err != nil {
				log.Debug3("Response --  Error Decoding DocumentReference: " + err.Error())
				return byte, resourceType, resp.StatusCode, log.Errorf("Response --  Error Decoding DocumentReference: " + err.Error())
			}
			log.Debug5("Response --  DocumentReference: " + spew.Sdump(docRef))
		case "DiagnosticReport":
			diagRept, err := fhir.UnmarshalDiagnosticReport(byte)
			if err != nil {
				log.Debug3("Response --  Error Decoding DiagnosticReport: " + err.Error())
				return byte, resourceType, resp.StatusCode, log.Errorf("Response --  Error Decoding DiagnosticReport: " + err.Error())
			}
			log.Debug5("Response --  DiagnosticReport: " + spew.Sdump(diagRept))

		default:
			log.Debug3("ResponseType --  " + resourceType)
			//return byte, resourceType, resp.StatusCode, nil
			//return byte, resourceType, http.StatusNotImplemented, nil
		}
		// diagRept, err := fhir.UnmarshalDiagnosticReport(byte)
		// if err != nil {
		// 	log.Debug3("Response --  Error Decoding DiagnosticReport: " + err.Error())
		// 	return byte, 400, log.Errorf("Response --  Error Decoding DiagnosticReport: " + err.Error())
		// }
		//log.Debug3("Response --  DiagnosticReport: " + spew.Sdump(diagRept))
		return byte, resourceType, resp.StatusCode, nil

	} else {
		return nil, resourceType, resp.StatusCode, log.Errorf("Response body is nil ")
	}
}
func (c *Connection) GetFhir(qry string, resourceType, token string) (json.RawMessage, error) {
	logrus.Printf("GetFhir:102  --  ResourceType = %s Query = %s  BaseUrl = %s", resourceType, qry, c.BaseURL)
	fullUrl := fmt.Sprintf("%s/%s/%s", c.BaseURL, resourceType, qry)
	logrus.Printf("GetFhir:104  --  FullURL Requested: %s", fullUrl)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		logrus.Errorf("GetFhir:107  --  !!!NewRequest failed: %s\n", err.Error())
		return nil, err
	}
	req.Header.Set("Accept", "application/json+fhir")
	resp, err := c.client.Do(req)
	if err != nil {
		logrus.Errorf("GetFhir:115  --  !!!fhir query returned err: %s", err.Error())
		return nil, err
	}
	byte := []byte{}
	if resp.Body != nil {
		byte, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, log.Errorf("Response  --  Error readying body: " + err.Error())
		}
		return byte, nil
	}
	return nil, log.Errorf("Response body is nil ")
}

func (c *Connection) GetFhirReq(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", "application/json+fhir")
	resp, err := c.client.Do(req)
	if err != nil {
		log.Error("--  !!!fhir query returned err: " + err.Error())
		return nil, err
	}

	return resp, err
}

func (c *Connection) GetFhirBundle(url string, token string) (*fhir.Bundle, error) {
	log.Info("GetFhirBundle URL Requested: " + url + "&_count=25")
	fullUrl := ""
	if strings.Contains(url, "https") {
		fullUrl = url
	} else {
		fullUrl = c.BaseURL + "/" + url
	}
	// limit := os.Getenv("PAGE_SIZE")
	// if limit == "" {
	// 	limit = "10"
	// }
	// limit = "10"
	// //besure first character of partial url is /
	// // if url[0:1] != "/" {
	// // 	url = "/" + url
	// // }
	// fullUrl = fullUrl + "&_count=" + limit
	log.Info("GetFhirBundle FullURL Requested: " + fullUrl)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		log.Error("--  !!!NewRequest failed: " + err.Error())
		return nil, err
	}
	log.Info("Calling c.GetFhirReq")
	resp, err := c.GetFhirReq(req)
	// req.Header.Set("Accept", "application/json+fhir")
	// resp, err := c.client.Do(req)
	if err != nil {
		log.Error("--  !!!fhir query returned err: " + err.Error())
		return nil, err
	}
	log.Info("resp.StatusCode: " + fmt.Sprint(resp.StatusCode))
	if resp.StatusCode == 401 {
		defer resp.Body.Close()
		byte, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, log.Errorf("--  Error Reading resp.Body: " + err.Error())
		}
		log.Debug3("body: " + string(byte))
		return nil, log.Errorf("--401 unauthorized")
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Error("GetFhirBundle returned statusCode of " + fmt.Sprint(resp.StatusCode))
		return nil, err
	}
	//tbundle := &fhir.Bundle{}
	defer resp.Body.Close()
	byte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, log.Errorf("--  Error Reading resp.Body: " + err.Error())
	}
	//log.Debug3("body: " + string(byte))
	bundle, err := fhir.UnmarshalBundle(byte)
	if err != nil {
		return nil, log.Errorf("--  Error Decoding bundle: " + err.Error())
	}
	//fmt.Printf("GetFhirBundle:201  --  Bundle =  %s\n", spew.Sdump(bundle))
	// err = json.NewDecoder(resp.Body).Decode(&data)
	// if err != nil {
	// 	fmt.Printf("NewDecoder error: %s\n", err.Error())
	// }
	// fmt.Printf("NewDecoder: %s\n\n", spew.Sdump(data))
	// bundle := &fhir.Bundle{}
	// err = json.NewDecoder(resp.Body).Decode(bundle)
	// if err != nil {
	// 	fmt.Printf("GetFhir:131  --  Error Decoding bundle: %s\n", err.Error())
	// 	return nil, err
	// }

	// fmt.Printf("GetFhirBundle:209  --  Bundle: %s\n", spew.Sdump(bundle))
	// patient, err := fhir.UnmarshalPatient(bundle.Entry[0].Resource)
	// fmt.Printf("GetFhirBundle:211  --  patient:  %s\n", spew.Sdump(patient))
	log.Info("GetFhirBundle Returning Bundle")
	return &bundle, nil
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("ReadBody Error:215 %s\n", err.Error())
	// 	return nil, err
	// }
	// fmt.Printf("\nGetFhir:218  -- %s\n", string(body))
	// fmt.Printf("GetFhir:219 returning no error and length of data: %d\n", len(body))
	// return body, nil
}

func (c *Connection) GetFhirResults(url string, token string) (*fhir.Bundle, error) {
	fullUrl := ""
	if strings.Contains(url, "https") {
		fullUrl = url
	} else {
		fullUrl = c.BaseURL + "/" + url
	}

	//besure first character of partial url is /
	// if url[0:1] != "/" {
	// 	url = "/" + url
	// }

	log.Info("GetFhirBundle FullURL Requested: " + fullUrl)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		log.Error("--  !!!NewRequest failed: " + err.Error())
		return nil, err
	}
	log.Info("Calling c.GetFhirReq")
	resp, err := c.GetFhirReq(req)
	// req.Header.Set("Accept", "application/json+fhir")
	// resp, err := c.client.Do(req)
	if err != nil {
		log.Error("--  !!!fhir query returned err: " + err.Error())
		return nil, err
	}
	log.Debug3("resp.StatusCode: " + fmt.Sprint(resp.StatusCode))
	if resp.StatusCode == 401 {
		defer resp.Body.Close()
		byte, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, log.Errorf("--  Error Reading resp.Body: " + err.Error())
		}
		log.Debug3("body: " + string(byte))
		return nil, log.Errorf("--401 unauthorized")
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Error("GetFhirBundle returned statusCode of " + fmt.Sprint(resp.StatusCode))
		return nil, err
	}
	//tbundle := &fhir.Bundle{}
	defer resp.Body.Close()
	byte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, log.Errorf("--  Error Reading resp.Body: " + err.Error())
	}
	//log.Debug3("body: " + string(byte))
	bundle, err := fhir.UnmarshalBundle(byte)
	if err != nil {
		return nil, log.Errorf("--  Error Decoding bundle: " + err.Error())
	}
	//fmt.Printf("GetFhirBundle:201  --  Bundle =  %s\n", spew.Sdump(bundle))
	// err = json.NewDecoder(resp.Body).Decode(&data)
	// if err != nil {
	// 	fmt.Printf("NewDecoder error: %s\n", err.Error())
	// }
	// fmt.Printf("NewDecoder: %s\n\n", spew.Sdump(data))
	// bundle := &fhir.Bundle{}
	// err = json.NewDecoder(resp.Body).Decode(bundle)
	// if err != nil {
	// 	fmt.Printf("GetFhir:131  --  Error Decoding bundle: %s\n", err.Error())
	// 	return nil, err
	// }

	// fmt.Printf("GetFhirBundle:209  --  Bundle: %s\n", spew.Sdump(bundle))
	// patient, err := fhir.UnmarshalPatient(bundle.Entry[0].Resource)
	// fmt.Printf("GetFhirBundle:211  --  patient:  %s\n", spew.Sdump(patient))
	return &bundle, nil
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("ReadBody Error:215 %s\n", err.Error())
	// 	return nil, err
	// }
	// fmt.Printf("\nGetFhir:218  -- %s\n", string(body))
	// fmt.Printf("GetFhir:219 returning no error and length of data: %d\n", len(body))
	// return body, nil
}

//	func (c *Connection)PatientNextPage(url string) {
//		bytes, err := c.GetFhir(url)
//	}
//
// Query sends a query to the base url
func (c *Connection) QueryBundle(q, token string) (*fhir.Bundle, error) {
	log.Info(fmt.Sprintf("--  BaseUrl: %s  -  Query param: %s\n\n\n\n", c.BaseURL, q))
	if q == "" {
		return nil, fmt.Errorf("c.QueryBundle:231  --  query parameter missing")
	}
	url := fmt.Sprintf("%s/%s", c.BaseURL, q)
	//fmt.Printf("fhir4_query:60  --  c.BaseUrl = %s\n", c.BaseURL)
	log.Info("calling GetFhirBundle url = " + url)
	bundle, err := c.GetFhirBundle(url, token)
	log.Info("GetFhirBundle returned")
	return bundle, err
}

func GetRemoteFhirPatient(qry string, fhirUrl string, token string) (*fhir.Patient, error) {
	fmt.Printf("GetRemotePatient:245  --  fhirUrl = %s\n", fhirUrl)
	c := New(fhirUrl, "application/json+fhir")
	rawPatient, err := c.GetFhir(qry, "Patient", token)
	if err != nil {
		return nil, err
	}
	patient, err := fhir.UnmarshalPatient(rawPatient)
	return &patient, err
}

func (c *Connection) PostFhir(qry, resourceType, token string, patient *fhir.Patient) (json.RawMessage, error) {
	log.Info(fmt.Sprintf("--  ResourceType = %s,  URLQuery = %s  BaseUrl = %s\n", resourceType, qry, c.BaseURL))
	fullUrl := c.BaseURL + qry
	logrus.Infof("PostFhir:258 FullURL Requested: %s\n", fullUrl)
	patb, err := json.Marshal(patient)
	rc := io.NopCloser(strings.NewReader(string(patb)))
	req, err := http.NewRequest("POST", fullUrl, rc)
	if err != nil {
		logrus.Errorf("PostFhir:263  --  !!!NewRequest failed: %s\n", err.Error())
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Set("AUTHORIZATION", token)

	resp, err := c.client.Do(req)
	if err != nil {
		logrus.Errorf("PostFhir:271  --  !!!fhir query returned err: %s\n", err)
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		logrus.Errorf("postFhir:275  --  returned error of %d - %s\n", resp.StatusCode, resp.Status)
		err = fmt.Errorf("%d|PostFhir:275 %s", resp.StatusCode, resp.Status)
		//log.Errorf("%s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	byte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("PostFhir:283  --  Error readying body: %s", err.Error())
	}
	//fmt.Printf("PostFhir:284  --  Raw body: %s\n", string(byte))

	_, err = fhir.UnmarshalPatient(byte) // Validate that it is a patient
	if err != nil {
		fmt.Printf("PostFhir:289  --  Error Decoding Patient: %s\n", err.Error())
		return nil, err
	}
	return byte, nil
}

func (c *Connection) postFhir(qry, resourceType, token string, patient *fhir.Patient) (*http.Response, error) {
	fmt.Printf("postFhir:298  --  ResourceType = %s,  URLQuery = %s  BaseUrl = %s\n", resourceType, qry, c.BaseURL)
	fullUrl := c.BaseURL + qry
	log.Info("FullURL Requested: " + fullUrl)
	patb, err := json.Marshal(patient)
	if err != nil {
		fmt.Printf("postFhir:302  --  Error Marshalling Patient: %s\n", err.Error())
		return nil, err
	}
	rc := io.NopCloser(strings.NewReader(string(patb)))
	req, err := http.NewRequest("POST", fullUrl, rc)
	if err != nil {
		logrus.Errorf("postFhir:308  --  !!!NewRequest failed: %s\n", err.Error())
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Set("AUTHORIZATION", token)
	resp, err := c.client.Do(req)

	if err != nil {
		logrus.Errorf("postFhir:317  --  !!!fhir query returned err: %s\n", err)
		return resp, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		logrus.Errorf("postFhir:321  --  returned error of %d - %s\n", resp.StatusCode, resp.Status)
		//err = fmt.Errorf("%d|postFhir:275 %s", resp.StatusCode, resp.Status)
		//log.Errorf("%s", err.Error())
		defer resp.Body.Close()
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("postFhir:327  --  Error readying body: %s\n", err.Error())
			return nil, fmt.Errorf("postFhir:327  --  Error readying body: %s", err.Error())
		}
		resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		fmt.Printf("postFhir:331  --  ReSet resp.Body to initial value\n")
		fmt.Printf("postFhir:332  --  Raw body: %s\n", string(bodyBytes))
		//resp.Body.Close()
		return resp, err
	}
	// defer resp.Body.Close()
	// byte, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("postFhir:328  --  Error readying body: %s\n", err.Error())
	// 	return nil, fmt.Errorf("postFhir:328  --  Error readying body: %s", err.Error())
	// }
	// fmt.Printf("postFhir:331  --  Raw body: %s\n", string(byte))
	return resp, err

}

// https://fhir-open.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d/Binary/XR-197574792
// https://fhir-open.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d/Binary/XR-197574792
