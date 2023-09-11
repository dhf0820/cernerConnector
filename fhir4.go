package main

import (
	//"encoding/json"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	fhir "github.com/dhf0820/fhir4"
	//common "github.com/dhf0820/uc_common"
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
}

// TODO: Consider including the FhirSystem in the Connection
// New creates a new connection
func New(baseurl string) *Connection {
	return &Connection{
		BaseURL: baseurl,
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
func (c *Connection) GetFhir(qry string, resourceType, token string) (json.RawMessage, error) {
	logrus.Printf("GetFhir:102  --  ResourceType = %s Query = %s  BaseUrl = %s", resourceType, qry, c.BaseURL)
	fullUrl := fmt.Sprintf("%s/%s%s", c.BaseURL, resourceType, qry)
	logrus.Printf("GetFhir:104  --  FullURL Requested: %s", fullUrl)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		logrus.Errorf("GetFhir:107  --  !!!NewRequest failed: %s\n", err.Error())
		return nil, err
	}
	req.Header.Set("ACCEPT", "application/json+fhir")
	req.Header.Set("AUTHORIZATION", token)
	fmt.Printf("getFhir:112  --  req: %s\n", spew.Sdump(req))
	resp, err := c.client.Do(req)
	if err != nil {
		logrus.Errorf("GetFhir:115  --  !!!fhir query returned err: %s", err.Error())
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		logrus.Errorf("GetFhir:119  --  returned error of %d - %n", resp.StatusCode, resp.Status)
		err = fmt.Errorf("%d|GetFhir: %s", resp.StatusCode, resp.Status)
		//log.Errorf("%s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	byte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("getFhir:127  --  Error readying body: %s", err.Error())
	}
	return byte, nil
	// fmt.Printf("GetFhir:124  --  Raw bundle: %s\n", string(byte))
	// bundle, err := fhir.UnmarshalBundle(byte)
	// if err != nil {
	// 	fmt.Printf("GetFhir:127  --  Error Decoding bundle: %s\n", err.Error())
	// 	return nil, err
	// }
	// // err = json.NewDecoder(resp.Body).Decode(&data)
	// // if err != nil {
	// // 	fmt.Printf("NewDecoder error: %s\n", err.Error())
	// // }
	// // fmt.Printf("NewDecoder: %s\n\n", spew.Sdump(data))
	// // bundle := &fhir.Bundle{}
	// // err = json.NewDecoder(resp.Body).Decode(bundle)
	// // if err != nil {
	// // 	fmt.Printf("GetFhir:131  --  Error Decoding bundle: %s\n", err.Error())
	// // 	return nil, err
	// // }

	// fmt.Printf("GetFhir:135  --  Bundle: %s\n", spew.Sdump(bundle))
	// patient, err := fhir.UnmarshalPatient(bundle.Entry[0].Resource)
	// fmt.Printf("GetFhir:137  --  patient:  %s\n", spew.Sdump(patient))
	// return &bundle, nil
	// // body, err := ioutil.ReadAll(resp.Body)
	// // if err != nil {
	// // 	fmt.Printf("ReadBody Error:119 %s\n", err.Error())
	// // 	return nil, err
	// // }
	// // fmt.Printf("\nGetFhir:129  -- %s\n", string(body))
	// // fmt.Printf("GetFhir:130 returning no error and length of data: %d\n", len(body))
	// // return body, nil
}
func (c *Connection) GetFhirBundle(url string, token string) (*fhir.Bundle, error) {
	fmt.Printf("GetFhirBundle:162  --  BaseUrl - %s  add url: %s\n", c.BaseURL, url)
	//besure first character of partial url is /
	// if url[0:1] != "/" {
	// 	url = "/" + url
	// }
	fmt.Printf("GetFhirBundle:167  --  url = %s\n", url)
	fullUrl := c.BaseURL + url
	log.Info("GetFhirBundle:169 FullURL Requested: " + fullUrl)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		logrus.Errorf("GetFhirBundle:172  --  !!!NewRequest failed: %s\n", err.Error())
		return nil, err
	}
	req.Header.Set("Accept", "application/json+fhir")
	//req.Header.Set("AUTHORIZATION", token)
	//fmt.Printf("GetFhirBundle:175  --  req: %s\n", spew.Sdump(req))
	resp, err := c.client.Do(req)
	if err != nil {
		logrus.Errorf("GetFhirBundle:180  --  !!!fhir query returned err: %s\n", err)
		return nil, err
	}
	//fmt.Printf("GetFhir:181  --  resp = %s\n", spew.Sdump(resp))
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		logrus.Errorf("GetFhirBundle:185  --  returned error of %d - %s\n", resp.StatusCode, resp.Status)
		err = fmt.Errorf("%d|GetFhirBundle:186 %s", resp.StatusCode, resp.Status)
		//log.Errorf("%s", err.Error())
		return nil, err
	}
	//tbundle := &fhir.Bundle{}
	defer resp.Body.Close()
	byte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("GetFhirBundle:194  --  Error Decoding bundle: %s", err.Error())
	}
	bundle, err := fhir.UnmarshalBundle(byte)
	if err != nil {
		fmt.Printf("GetFhirBundle:198  --  Error Decoding bundle: %s\n", err.Error())
		return nil, err
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
func (c *Connection) QueryBundle(q string, token string) (*fhir.Bundle, error) {
	fmt.Printf("\n\n\n\nQueryBundle:233  --  BaseUrl: %s  -  Query param: %s\n\n\n\n", c.BaseURL, q)
	if q == "" {
		return nil, fmt.Errorf("c.QueryBundle:231  --  query parameter missing")
	}
	url := fmt.Sprintf("%s/%s", c.BaseURL, q)
	//fmt.Printf("fhir4_query:60  --  c.BaseUrl = %s\n", c.BaseURL)
	fmt.Printf("c.QueryBundle:239  -- url = %s\n", url)

	return c.GetFhirBundle(url, token)
}

func GetRemoteFhirPatient(qry string, fhirUrl string, token string) (*fhir.Patient, error) {
	fmt.Printf("GetRemotePatient:245  --  fhirUrl = %s\n", fhirUrl)
	c := New(fhirUrl)
	rawPatient, err := c.GetFhir(qry, "Patient", token)
	if err != nil {
		return nil, err
	}
	patient, err := fhir.UnmarshalPatient(rawPatient)
	return &patient, err
}

func (c *Connection) PostFhir(qry, resourceType, token string, patient *fhir.Patient) (json.RawMessage, error) {
	fmt.Printf("PostFhir:256  --  ResourceType = %s,  URLQuery = %s  BaseUrl = %s\n", resourceType, qry, c.BaseURL)
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
	//fmt.Printf("getFhir:102  --  req: %s\n", spew.Sdump(req))
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
	byte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("PostFhir:283  --  Error readying body: %s", err.Error())
	}
	//fmt.Printf("PostFhir:284  --  Raw body: %s\n", string(byte))

	pat, err := fhir.UnmarshalPatient(byte)
	if err != nil {
		fmt.Printf("PostFhir:289  --  Error Decoding Patient: %s\n", err.Error())
		return nil, err
	}
	fmt.Printf("PostFhir:292  --  Patient =  %s\n", spew.Sdump(pat))
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
	//fmt.Printf("getFhir:102  --  req: %s\n", spew.Sdump(req))
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
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("postFhir:327  --  Error readying body: %s\n", err.Error())
			return nil, fmt.Errorf("postFhir:327  --  Error readying body: %s", err.Error())
		}
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
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
	// pat, err := fhir.UnmarshalPatient(byte)
	// if err != nil {
	// 	fmt.Printf("postFhir:288  --  Error Decoding Patient: %s\n", err.Error())
	// 	return nil, err
	// }
	// fmt.Printf("postFhir:291  --  Patient =  %s\n", spew.Sdump(pat))
	// return byte, nil
}

// https://fhir-open.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d/Binary/XR-197574792
// https://fhir-open.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d/Binary/XR-197574792
