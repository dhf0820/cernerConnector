package main

import (
	"encoding/json"
	//"errors"
	"fmt"
	"io"

	"github.com/davecgh/go-spew/spew"
	cm "github.com/dhf0820/baseConnector/common"
	fhir "github.com/dhf0820/fhir4"
	jw_token "github.com/dhf0820/jwToken"
	common "github.com/dhf0820/uc_common"
	//"github.com/gorilla/mux"

	//"github.com/gorilla/schema"
	log "github.com/dhf0820/vslog"
	"net/http"

	//"os"
	"reflect"
	//"strconv"
	"strings"
	"time"
)

//####################################### Response Writers Functions #######################################

//################################### FHIR Responses ####################################

// ####################################### Route Handlers #######################################
// getPatient - By patientId returning one single patient matching the ID
// otherwise return OperationOutcome for NotFound
var JWToken string

func getPatient(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("getPatient:36 - req: %s\n", spew.Sdump(r))
	//Resource := "Patient"
	//buildFieldsByTagMap("schema", *psp)
	JWToken = r.Header.Get("Authorization")
	if JWToken == "" {
		err := fmt.Errorf("getPatient:41  --  Authorization is blank")
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, 401, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	// Payload, status, err := token.ValidateToken(JWToken, "")
	// if err != nil {
	// 	errMsg := err.Error()
	// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// userId := Payload.UserId
	// fmt.Printf("getPatient:47  --  userId: %s\n", userId)
	// cp, err := GetConnectorPayload(r)
	// fmt.Printf("getPatient:48  --  ConnectorPayLoad = %s\n", spew.Sdump(cp))
	// //defer r.Body.Close()
	// params := mux.Vars(r)
	// fmt.Printf("getPatient:52  --  Params: %v\n", params)
	// id := params["id"]
	// //uri := r.URL.RequestURI()
	// fmt.Printf("getPatient:55  --  raw: %s\n", r.URL.RawQuery)
	// fmt.Printf("getPatient:56  --  query values: %v\n", r.URL.Query())
	// values := r.URL.Query()
	// for k, v := range values {
	// 	fmt.Println(k, " => ", v)
	// }
	// fmt.Printf("getPatient:61  --  id = %s\n", id)
	// patient, err := GetPatient(id)
	// if err != nil {
	// 	err = fmt.Errorf("getPatient:64  --  GetPatient error: %s\n", err.Error())
	// 	errMsg := err.Error()
	// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeIncomplete, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// fmt.Printf("\n\n\ngetPatient:69  --  Returning Patient: %s\n", spew.Sdump(patient))
	// resp := common.ResourceResponse{}

	// resp.Resource.Patient = patient
	// resp.Patient = *patient
	// resp.Status = 200
	// resp.Message = "Ok"
	// resp.ResourceType = Resource
	// WriteFhirResource(w, 200, &resp)

	// //defer resp.Body.Close()
	// //cfg = mod.ServiceConfig{}
	// fmt.Printf("Reading Body\n")
	// body, err = ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	log.Printf("getPatient:53  --  ReadAllBody : error: %s\n", err.Error())
	// 	//err = errors.New("invalid FHIR URL")
	// 	errMsg := err.Error()
	// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// connectPayload := common.ConnectorPayload{}
	// fmt.Printf("raw string: %s\n", string(body))
	// fmt.Printf("GetPatient:61  --  Unmarshal ConnectorPayload\n")
	// err = json.Unmarshal(body, &connectPayload)
	// if err != nil {
	// 	log.Printf("getPatient:64  --  Unmarshal connectPayload error: %s\n", err.Error())
	// 	//err = errors.New("invalid FHIR URL")
	// 	errMsg := err.Error()
	// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// params := mux.Vars(r)
	// id := params["id"]
	// if id != "" {
	// 	fmt.Printf("GetPatient:73  -- Specific Query: %s\n", id)
	// 	//patient, err := GetPatient(id)
	// }

	// //fmt.Printf("getPatient:70 -- connectPayload: %s\n", spew.Sdump(connectPayload))
	// // fhirId := GetFhirId(r)
	// // fhirSystem, err := GetFhirSystem(fhirId)
	// // if err != nil {
	// // 	log.Printf("searchPatient:50  --  FhirId : [%s] error: %s\n", fhirId, err.Error())
	// // 	err = errors.New("invalid FHIR URL")
	// // 	errMsg := err.Error()
	// // 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// // 	return
	// // }

	// //  Separate connector for each emr Vendor.  CernerConnector, EpicConnector, CAConnector,...
	// // // handles query andsaveeither via fhir or direct API (AllScripts, Athena)

	// // Payload, status, err := token.ValidateToken(r.Header.Get("Authorization"), "")
	// // if err != nil {
	// // 	errMsg := err.Error()
	// // 	fmt.Printf("getPatient:55  --  Err: %s\n", errMsg)
	// // 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// // 	return
	// // }
	// // userId := Payload.UserId
	// // log.Printf("getPatient:59  --  UserId: %s\n", userId)
	// // fhirId := GetFhirId(r)
	// // fhirSystem, err := GetFhirSystem(fhirId)
	// // if err != nil {
	// // 	log.Printf("getPatient:63  --  FhirId : [%s] error: %s\n", fhirId, err.Error())
	// // 	err = errors.New("invalid FHIR URL")
	// // 	errMsg := err.Error()
	// // 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// // 	return
	// // }
	// fmt.Printf("getPatient:109  --  Request: [%s]\n", r.RequestURI)
	// urlA, err := r.URL.Parse(r.RequestURI)
	// if err != nil {
	// 	err = fmt.Errorf("error parsing patient URI: %s", err.Error())
	// 	errMsg := err.Error()
	// 	fmt.Printf("getPatient:114 - r.URL.Parse error = %s\n", errMsg)
	// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// fmt.Printf("getPatient:118 - r.URL.Parse = %v\n", urlA)
	// urlB := *urlA
	// uriValues := urlB.Query()
	// fmt.Printf("getPatient:121 - uriValues= %v\n", uriValues)

	// uri := r.RequestURI
	// log.Printf("uri = %s\n", uri)
	// parts := strings.Split(uri, Resource)
	// uri = parts[1]
	// log.Printf("getPatient:127 - URI = %s\n", uri)
	// //patient := fhir.Patient{}
	// resource, err := GetResource(connectPayload.FhirSystem, Resource, uri)
	// resp := common.ResourceResponse{}
	// if err != nil {
	// 	resp.Status = 400
	// 	resp.Message = err.Error()
	// } else {
	// 	resp.Status = 200
	// 	resp.Message = "Ok"
	// }
	// //var patient fhir.Patient
	// //patient := resource.(fhir.Patient)
	// resp.Resource.Resource = resource
	// // var res []interface{}
	// // res = append(res, &resource)
	// // resp.Resources = res
	// resp.ResourceType = Resource
	// //resp.ResourceId = *patient.Id
	// //log.Printf("\nGetPatient:139  --  resp: %s\n", spew.Sdump(resp))

	// WriteFhirResource(w, resp.Status, &resp)
}

// postPatient: Stores the fhir patient payload in the url {Fhir-System} specified fhirSystem.
func savePatient(w http.ResponseWriter, r *http.Request) {
	//Resource := "Patient"
	//fmt.Printf("postPatient:182 - Post: %s \n", spew.Sdump(r))
	JWToken := r.Header.Get("Authorization")
	payload, status, err := jw_token.ValidateToken(JWToken, "")
	if err != nil {
		errMsg := err.Error()
		fmt.Printf("savePatient:190  - ValidateToken err = %s\n", errMsg)
		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	log.Info("payload: " + spew.Sdump(payload))
	body, err := io.ReadAll(r.Body) // Should be ConnectorPayload
	if err != nil {
		fmt.Printf("savePatient:197  --  ReadAll FhirSystem error %s\n", err.Error())
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	fmt.Printf("savePatient:202  --  ConnectorPayload: %s\n", spew.Sdump())
	//b := string(body)
	//fmt.Printf("SavePatient:204  Body: %s\n", b)
	conPayload := &common.ConnectorPayload{}
	err = json.Unmarshal(body, &conPayload)
	if err != nil {
		fmt.Printf("\nsavePatient:208  --  unmarshal err = %s\n", err.Error())
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	fmt.Printf("savePatient:213  --  Check ConPayload\n")
	if conPayload == nil {
		fmt.Printf("\n\nconPayload is nil\n")
	}
	fmt.Printf("savePatient:217  --  Check SavePayload\n")
	if conPayload.SavePayload == nil {
		fmt.Printf("\n\nconPayload.SavePayload is nil\n")
	}
	fmt.Printf("\n\nsavePatient:221  --  conPayload: %s\n", spew.Sdump(conPayload))
	//
	//fmt.Printf("SavePatient:223  --  Check ConPayload srcPatient: %s\n", spew.Sdump(conPayload))
	// if conPayload.SavePayload.SrcPatient == nil {
	// 	res, err := GetCachedResource(conPayload, JWToken)
	// 	if res != nil {
	// 		errMsg := err.Error()
	// 		fmt.Printf("savePatient:228 - r.URL.Parse error = %s\n", errMsg)
	// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
	// 		return
	// 	}
	// 	conPayload.SavePayload.SrcPatient = res.Patient
	// }

	fmt.Printf("savePatient:235  --  calling SavePatient with Payload: %s\n", spew.Sdump(conPayload))
	httpResp, err := SavePatient("", conPayload, JWToken)
	if err != nil {
		fmt.Printf("savePatient:238  --  SavePatient error %s\n", err.Error())
	}
	if httpResp.StatusCode != 201 {
		fmt.Printf("savePatient:241  --  httpResp.Status = %s   StatusCode: %d\n", httpResp.Status, httpResp.StatusCode)
		fmt.Printf("savePatient:242  --  error: %s\n", httpResp.Status)
		defer httpResp.Body.Close()
		bodyBytes, err := io.ReadAll(httpResp.Body)
		if err != nil {
			errMsg := err.Error()
			fmt.Printf("savePatient:247  --  ReadAll SavePatient ReadBody error %s\n", err.Error())
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
			return
		}
		opOutcome := fhir.OperationOutcome{}
		err = json.Unmarshal(bodyBytes, &opOutcome)
		if err != nil {
			errMsg := log.ErrMsg("ErrorMessage ReadBody error " + err.Error())
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
			return
		}
		WriteFhirOperationOutcome(w, httpResp.StatusCode, &opOutcome)
		// issueType := fhir.IssueTypeException
		// if httpResp.StatusCode == 409 {
		// 	issueType = fhir.IssueTypeDuplicate
		// }

		// WriteFhirOperationOutcome(w, httpResp.StatusCode, CreateOperationOutcome(issueType, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	defer httpResp.Body.Close()
	byte, err := io.ReadAll(httpResp.Body)
	if err != nil {
		errMsg := err.Error()
		fmt.Printf("savePatient:272  --  ReadAll SavePatient ReadBody error %s\n", err.Error())
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	fmt.Printf("savePatient:276  --  Check if patient\n")

	// saveResp := common.SaveResponse{}
	// err = json.Unmarshal(byte, &saveResp)
	saveResp := common.SaveResponse{}
	err = json.Unmarshal(byte, &saveResp)
	if err != nil {
		fmt.Printf("savePatient: 283  --  Response is not a Patient err: %s\n", err.Error())
		HandleOperationOutcome(w, byte)
		return
	}
	fmt.Printf("savePatient:287 returned saveResp: %s\n", spew.Sdump(saveResp))
	WriteSaveResponse(w, 201, &saveResp)

	// if patient.ResourceType == nil {
	// 	fmt.Printf("savePatient: 269  --  Response is not a Patient\n")
	// 	HandleOperationOutcome(w, byte)
	// 	return
	// }
	//resp := &common.ResourceResponse{}

	//resp.Patient = patient
	// resp := &common.SaveResponse{}
	// resp.Id = *patient.Id
	// resp.Text = patient.Text.Div
	// resp.Mrn = GetMrn(&patient, "urn:oid:1.3.6.1.4.1.54392.5.1593.1", "mrn")
	//fmt.Printf("\n\n\nsavePatient:298  --  Returning %s\n", spew.Sdump(resp))
	//location := Conf.Server.Host + ":" + Conf.Server.Port + "/system/640ba66cbd4105586a6dda75/Patient/" + patient.Id
	// baseUrl := Conf.BaseURL
	// location := baseUrl + "/system/640ba66cbd4105586a6dda75/Patient/" // + *patient.Id
	// fmt.Printf("savePatient:274  --  Location: %s\n", location)
	//WriteFhirPatient(w, 200, &patient)
	// opOutcom := fhir.OperationOutcome{}
	// err = json.Unmarshal(byte, &opOutcom)
	// if err == nil {
	// 	fmt.Printf("SavePatient:253  --  opOutcom: %s\n", spew.Sdump(opOutcom))
	// 	WriteHttpResponse(w, httpResp.StatusCode, httpResp)
	// 	//WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// patient := fhir.Patient{}
	// err = json.Unmarshal(byte, &patient)
	// if err != nil {
	// 	fmt.Printf("savePatient: 258  --  Response is not a Patient\n")
	// 	opOutcom := fhir.OperationOutcome{}
	// 	err = json.Unmarshal(byte, &opOutcom)
	// 	if err == nil {
	// 		fmt.Printf("SavePatient:262  --  opOutcom: %s\n", spew.Sdump(opOutcom))
	// 		WriteHttpResponse(w, httpResp.StatusCode, httpResp)
	// 		//WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
	// 		return
	// 	}
	// 	fmt.Printf("savePatient:267  --   Unmarshal Patient Body error %s\n", err.Error())
	// 	return
	// } else {
	// 	fmt.Printf("savePatient:270  --  Patient Unmarshal error: %s\n", spew.Sdump(patient))
	// }
	// if err != nil {
	// 	errMsg := err.Error()
	// 	fmt.Printf("savePatient:249  --  ReadAll SavePatient Unmarshal Body error %s\n", err.Error())
	// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }

	// if err != nil {
	// 	errMsg := err.Error()
	// 	fmt.Printf("SavePatient:239 - SavePatient error = %s\n", errMsg)
	// 	values := strings.Split(errMsg, "|")
	// 	//errNum, _ := strconv.Atoi(values[0])
	// 	//fmt.Printf("SavePatient:240 - values = %s\n", spew.Sdump(values))
	// 	fmt.Printf("SavePatient:241 - values[0] = %s\n", values[0])
	// 	if values[0] == "409" {
	// 		WriteFhirOperationOutcome(w, 409, CreateOperationOutcome(409, fhir.IssueSeverityFatal, &errMsg))
	// 	} else {
	// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
	// 	}
	// 	return
	// }
	// fmt.Printf("savePatient:293 returned PatientId: %s\n", *patient.Id)
	// resp := &common.SaveResponse{}
	// resp.Id = *patient.Id
	// resp.Text = patient.Text.Div
	// resp.Mrn = GetMrn(&patient, "urn:oid:1.3.6.1.4.1.54392.5.1593.1", "mrn")
	// fmt.Printf("\n\n\nsavePatient:298  --  Returning %s\n", spew.Sdump(resp))
	// WriteSaveResponse(w, 200, resp)

	// fhirId := GetFhirId(r)                   // Get the Fhir-System ID portion of the URL
	// fhirSystem, err := GetFhirSystem(fhirId) // Get the actual FhirSystem Configuration
	// if err != nil {
	// 	log.Printf("postPatient:162  --  FhirId : [%s] error: %s\n", fhirId, err.Error())
	// 	err = errors.New("url contains Invalid FHIR identifier: " + fhirId)
	// 	errMsg := err.Error()
	// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// urlA, err := r.URL.Parse(r.RequestURI)
	// if err != nil {
	// 	err = fmt.Errorf("error parsing patient URI: [%s]  error:%s", r.RequestURI, err.Error())
	// 	errMsg := err.Error()
	// 	fmt.Printf("postPatient:172 - r.URL.Parse error = %s\n", errMsg)
	// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
	// 	return
	// }
	// fmt.Printf("postPatient:176 - r.URL.Parse = %v\n", urlA)
	// urlB := *urlA
	// uriValues := urlB.Query()
	// fmt.Printf("postPatient:179 - uriValues= %v\n", uriValues)

	// uri := r.RequestURI
	// log.Printf("uri = %s\n", uri)
	// parts := strings.Split(uri, Resource)
	// uri = parts[1]
	// log.Printf("postPatient:185 - URI = %s\n", uri)
	// //patient := fhir.Patient{}
	//WriteFhirResource(w, 200, resp)
	// errMsg := "SavePatient to CA not implemented"
	// WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// return
	// resource, err := GetResource(fhirSystem, Resource, uri)
	// resp := common.ResourceResponse{}
	// if err != nil {
	// 	resp.Status = 400
	// 	resp.Message = err.Error()
	// } else {
	// 	resp.Status = 200
	// 	resp.Message = "Ok"
	// }
	// var patient fhir.Patient
	// //patient = resource.(fhir.Patient)
	// resp.Resource.Resource = resource
	// // var res []interface{}
	// // res = append(res, &resource)
	// // resp.Resources = res
	// resp.ResourceType = Resource
	// resp.ResourceId = *patient.Id
	// log.Printf("\nGetPatient:204  --  resp: %s\n", spew.Sdump(resp))
	// WriteFhirResourceBundle(w, resp.Status, &resp)
}

// searchPatient uses the systemId url parameter to determin the FhirSystem to use
func searchPatient(w http.ResponseWriter, r *http.Request) {
	srchParams, err := cm.PatientSearchparams(r) //FhirPatientSearch(r)
	if err != nil {
		fmt.Printf("searchPatient:397  -- Error %s \n", err.Error())
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	fmt.Printf("searchPatient:402  --  srchParams: %s\n", spew.Sdump(srchParams))

	// var pspTags map[string]string
	// tagFields := make(map[string]string)
	// var Limit int
	// var Skip int
	//Resource := "Patient"
	body, err := io.ReadAll(r.Body) // Should be ConnectorPayload
	if err != nil {
		errMsg := log.ErrMsg("ReadAll FhirSystem error: " + err.Error())
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	connectorPayload := common.ConnectorPayload{}
	//fhirSystem := common.FhirSystem{}
	err = json.Unmarshal(body, &connectorPayload)
	if err != nil {
		fmt.Printf("\nsearchPatient:420  --  unmarshal err = %s\n", err.Error())
		errMsg := log.ErrMsg("Unmarshal ConnectorPayload error: " + err.Error())
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	//fhirSystem := connectorPayload.FhirSystem
	//connConfig := connectorPayload.ConnectorConfig
	//buildFieldsByTagMap("schema", *psp)
	JWToken = r.Header.Get("Authorization")
	//fmt.Printf("searchPatient:219 - JWToken: %s\n", JWToken)
	Payload, status, err := jw_token.ValidateToken(r.Header.Get("Authorization"), "")
	if err != nil {
		errMsg := err.Error()
		fmt.Printf("searchPatient:433  --  ValidateToken err: %s\n", errMsg)
		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	//fhirId := fhirSystem.ID.String()
	userId := Payload.UserId
	log.Info("UserId: " + userId)

	log.Info("raw: " + r.URL.RawQuery)
	//fmt.Printf("query values: %v\n", r.URL.Query())
	values := r.URL.Query()
	for k, v := range values {
		fmt.Println(k, " => ", v)
	}

	/*
		// fhirId := GetFhirId(r)
		// fhirSystem, err := GetFhirSystem(fhirId)
		// if err != nil {
		// 	log.Printf("searchPatient:232  --  FhirId : [%s] error: %s\n", fhirId, err.Error())
		// 	err = errors.New("invalid FHIR URL")
		// 	errMsg := err.Error()
		// 	WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		// 	return
		// }
		uri := r.RequestURI
		log.Printf("searchPatient:314  --  r.RequestURI = %s\n", uri)
		parts := strings.Split(uri, Resource)
		uri = parts[1]
		log.Printf("\nsearchPatient:260 - URI = %s\n", uri)

		urlA, err := r.URL.Parse(r.RequestURI)
		if err != nil {
			err = fmt.Errorf("error parsing patient URI: %s", err.Error())
			errMsg := err.Error()
			fmt.Printf("searchPatient:266 - r.URL.Parse error = %s\n", errMsg)
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
			return
		}
		fmt.Printf("searchPatient:270 - r.URL.Parse = %v\n", urlA)
		urlB := *urlA
		uriValues := urlB.Query()
		fmt.Printf("searchPatient:273 - uriValues= %v\n", uriValues)
		//ident := uriValues.Get("identifier")
		// if ident != "" { // There is identifier Search, use it
		// 	fmt.Printf("searchPatient:102 - using Identifier: %s to search\n", ident)
		// } else {
		// 	fmt.Printf("searchPatient:104 - using other search params: %v\n", uriValues)
		// }

		// //}
		// //fhirVersion := GetFHIRVersion(r)
		// //cacheBaseURL := fmt.Sprintf("%s/%s/v1/", r.Host, parts[0])
		// if err := r.ParseForm(); err != nil {
		// 	err = fmt.Errorf("error parsing query: %s", err.Error())
		// 	errMsg := err.Error()
		// 	fmt.Printf("searchPatient:113 - %s\n", errMsg)
		// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
		// 	return
		// }
		// FhirId := GetFhirId(r)
		// fmt.Printf("searchPatient:79 - FhirKey - [%s]\n", FhirId)
		// fhirSystem, err := GetFhirSystem(FhirId)
		// if err != nil {
		// 	fmt.Printf("GetFhirSystem failed with : %s\n", err.Error())
		// 	err = fmt.Errorf("fhirSystem error:  %s", err.Error())
		// 	errMsg := err.Error()
		// 	fmt.Printf("searchPatient:86 - %s\n", errMsg)
		// 	WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityFatal, &errMsg))
		// 	return
		// }
		//fmt.Printf("searchPatient:90 -  %s/n", spew.Sdump(fhirSystem))

		// if Resource == "Patient" {
		log.Printf("\n\nsearchPatient:305  --  Resource Is Patient\n\n")
		//urlA, err := r.URL.Parse(r.RequestURI)
		if err != nil {
			err = fmt.Errorf("error parsing patient URI: %s", err.Error())
			errMsg := err.Error()
			fmt.Printf("searchPatient:310 - r.URL.Parse error = %s\n", errMsg)
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
			return
		}
		fmt.Printf("searchPatient:314 - r.URL.Parse = %v\n", urlA)
		//urlB := *urlA
		//uriValues := urlB.Query()
		fmt.Printf("searchPatient:317 - uriValues= %v\n", uriValues)
		idSearch := uriValues.Get("identifier")
		idValue := ""
		if idSearch != "" { // There is identifier Search, use it
			fmt.Printf("searchPatient:321- using Identifier: %s to search\n", idSearch)
			ids := strings.Split(idSearch, "|")
			if len(ids) != 2 {
				err = fmt.Errorf("invalid identifier: %s", idSearch)
				errMsg := err.Error()
				fmt.Printf("searchPatient:326 - r.URL.Parse error = %s\n", errMsg)
				WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
				return
			}
			idName := ids[0]
			idSearchValue := ids[1]
			idents := fhirSystem.Identifiers
			for _, id := range idents {
				fmt.Printf("searchPatient:334  --  Looking at %s = %s\n", id.Name, idName)
				if id.Name == idName {
					idValue = id.Value
					break
				}
			}
			//
			if idValue == "" { //Not configured identifier
				err = fmt.Errorf("identifier type: %s is not configured", idName)
				errMsg := err.Error()
				fmt.Printf("searchPatient:344 - Identifiers = %s\n", errMsg)
				WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
				return
			}
			uri = fmt.Sprintf("?identifier=%s", idValue+idSearchValue)
			fmt.Printf("searchPatient:349 - New Identifier search Value: %s\n", uri)
		} else {
			fmt.Printf("searchPatient:351 - using other search params: %v\n", uriValues)
		}
		var bundle *fhir.Bundle
		var header *common.CacheHeader
		fmt.Printf("\nsearchPatient:355 - resource = %s  uri = %s\n", Resource, uri)
		url := fmt.Sprintf("%s/%s%s", fhirSystem.FhirUrl, Resource, uri) //" + "/" + uri
		fmt.Printf("searchPatient:357 - calling %s \n", url)
		var totalPages int64
		fmt.Printf("searchPatient:359 Search %s\n", url)
		uri = "/" + Resource + uri
		totalPages, bundle, header, err = FindResource(&connectorPayload, Resource, userId, uri, JWToken)
		if err != nil {
			err = fmt.Errorf("searchPatient:363 --  fhirSearch url: %s error:  %s", url, err.Error())
			errMsg := err.Error()
			fmt.Printf("searchPatient:365 - %s\n", errMsg)
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityInformation, &errMsg))
			return
		}
		if bundle == nil {
			log.Printf("searchPatient:370  --  bundle is nil")
		} else {
			log.Printf("searchPatient:372  --  bundle is not nil \n")
		}
		fmt.Printf("searchPatient:374 - Get %s bundle successful\n", Resource)
		fmt.Printf("searchPatient:375 - Number in page: %d\n", len(bundle.Entry))
		fmt.Printf("searchPatient:376 - PageNumber: %d\n", header.PageId)
		resp := common.ResourceResponse{}
		header.CacheBase = fmt.Sprintf("%s/%s/BundleTransaction", connConfig.CacheUrl, header.FhirSystem.ID.Hex())
		log.Printf("\n\nsearchPatient:379  --  CacheUrl = %s\n", header.CacheBase)
		header.FhirId = fhirId
		header.UserId = userId
		resp.Bundle = bundle
		resp.Resource.Resource = bundle.Entry[0].Resource
		resp.BundleId = *bundle.Id
		resp.ResourceType = Resource
		resp.Status = 200
		resp.QueryId = header.QueryId
		resp.PageNumber = header.PageId
		if bundle.Entry == nil {
			err = fmt.Errorf("searchPatient:390 --  fhirSearch url: %s error:  %s", url, "Bundle.Entry is nil")
			errMsg := err.Error()
			fmt.Printf("searchPatient:392 - %s\n", errMsg)
			WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeNotFound, fhir.IssueSeverityInformation, &errMsg))
			return
		}
		resp.CountInPage = len(bundle.Entry)
		resp.TotalPages = totalPages
		resp.Header = header
		resp.Message = "Ok"
		//fmt.Printf("searchPatient:400 - returning a resource bundle: %s\n", spew.Sdump(resp))
		WriteFhirResourceBundle(w, resp.Status, &resp)
		//WriteFhirBundle(w, resp.Status, bundle)
	*/
}

//func searchPatient(w http.ResponseWriter, r *http.Request) {
// 	// var pspTags map[string]string
// 	// tagFields := make(map[string]string)
// 	// var Limit int
// 	// var Skip int
// 	fmt.Printf("Request: %s \n", spew.Sdump(r))
// 	//buildFieldsByTagMap("schema", *psp)
// 	//facility = "demo"
// 	resource := GetFHIRResource(r)
// 	fmt.Printf("search%s called with %s\n", resource, r.URL.RawQuery)
// 	if err := r.ParseForm(); err != nil {
// 		err = fmt.Errorf("error parsing query: %s", err.Error())
// 		errMsg := err.Error()
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 		return
// 	}
// 	params := mux.Vars(r)
// 	fmt.Printf("params: %v\n", params)
// 	resourceId := params["id"]
// 	fmt.Printf("Retrieving Patient Record for id: %s\n", resourceId)
// 	// psp := new(PatientSearchParams)
// 	// psp.Limit = Limit
// 	// psp.Skip = Skip
// 	// psp.CurrentFacility = GetDeploymentFacility(r)
// 	// psp.BaseUrl = GetCurrentURL(r)
// 	//FhirVersion := GetFHIRVersion(r)
// 	FhirId := GetFhirId(r)
// 	_, err := GetFhirConnector(FhirId)
// 	if err != nil {
// 		err = fmt.Errorf("fhirConnecor error:  %s", err.Error())
// 		errMsg := err.Error()
// 		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(400, fhir.IssueSeverityFatal, &errMsg))
// 	}

// 	//bundle, err := FindResource(fhirConnector, resource, r.URL.RawQuery)

// }

// func WriteFhirOperationOutcome(w http.ResponseWriter, status int, resp *fhir.OperationOutcome) error {
// 	w.Header().Set("Content-Type", "application/json")

// 	switch status {
// 	case 200:
// 		w.WriteHeader(http.StatusOK)
// 	case 400:
// 		w.WriteHeader(http.StatusBadRequest)
// 	case 401:
// 		w.WriteHeader(http.StatusUnauthorized)
// 	case 403:
// 		w.WriteHeader(http.StatusForbidden)
// 	}
// 	err := json.NewEncoder(w).Encode(resp)
// 	if err != nil {
// 		fmt.Println("Error marshaling JSON:", err)
// 		return err
// 	}
// 	return nil
// }

func WriteFhirPatient(w http.ResponseWriter, status int, resp *fhir.Patient) error {
	w.Header().Set("Content-Type", "application/json")

	switch status {
	case 200:
		w.WriteHeader(http.StatusOK)
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 401:
		w.WriteHeader(http.StatusUnauthorized)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("WriteFhirPatient:682  --  Error marshaling JSON:", err)
		return err
	}
	return nil
}

type PatientFilter struct {
	Id         string    `json:"id" schema:"_id"`
	MRN        string    `json:"mrn" schema:"mrn"`
	SSN        string    `json:"ssn" schema:"ssn"`
	Identifier string    `json:"identifier" schema:"identifier"`
	Gender     string    `json:"gender" schema:"gender"`
	BirthDate  string    `json:"birthdate" schema:"birthdate"`
	Name       string    `json:"name" schema:"name"`
	Family     string    `json:"family" schema:"family"`
	Given      string    `json:"given" schema:"given"`
	Phone      string    `json:"phone" schema:"phone"`
	Email      string    `json:"email" schema:"email"`
	PostalCode string    `json:"address-postalcode" schema:"address-postalcode"`
	Active     string    `json:"active" schema:"active"`
	DOB        time.Time `json:"dob" schema:"dob"`
	BaseUrl    string    `json:"base_url"`
	RequestURI string    `json:"request_uri"`
	Limit      uint32    `json:"limit"`
	Skip       uint32    `json:"skip" schema:"skip"`
	Count      uint32    `json:"count" schema:"_count"`
	OffSet     uint32    `json:"offset" schema:"_offset"`
	// Order      SearchParam `json:"order" schema:"_order"`
	// Sort       SearchParam `json:"sort" schema:"_sort"`
	// Page       SearchParam `json:"page" schema:"_page"`
}

type SearchParam struct {
	Schema   string
	Modifier string
	Value    string
}

func findPatient(w http.ResponseWriter, r *http.Request) {
	JWToken := r.Header.Get("Authorization")
	Payload, status, err := jw_token.ValidateToken(JWToken, "")
	if err != nil {
		errMsg := err.Error()
		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	userId := Payload.UserId
	log.Info(" UserId: " + userId)
	r.ParseForm()
	decoder.RegisterConverter(time.Now(), func(value string) reflect.Value {
		result := reflect.Value{}
		if t, err := time.Parse("2006-01-02", value); err == nil {
			result = reflect.ValueOf(t)
		}
		return result
	})
	cp, err := GetConnectorPayload(r)
	if err != nil {
		errMsg := log.ErrMsg("GetConnectorPayload error: " + err.Error())
		WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeForbidden, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	//fmt.Printf("findPatient:645  --  ConnectorPayload: %s\n", spew.Sdump(cp))
	//fmt.Printf("findPatient:646  --  r.Form = %v\n", r.Form)
	//fmt.Printf("\nfindPatient:647  --  qry = %v   len of URL.Query = %d\n\n", r.URL.Query(), len(r.URL.Query()))
	patFilter := new(PatientFilter)
	u := r.URL.Query()
	fmt.Printf("findPatient:779  --  u: %v\n", u)
	// qry, err := CreateFhirQuery(r)
	// if err != nil {
	// 	if err != nil {
	// 		status := 400
	// 		errMsg := fmt.Sprintf("findPatient:655  --  CreateFhirQuery err: %s\n", err.Error())
	// 		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
	// 		return
	// 	}
	// }
	err = decoder.Decode(patFilter, r.Form)
	if err != nil {
		status := 400
		errMsg := log.ErrMsg(" decode err: " + err.Error())
		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	fmt.Printf("findPatient:796  --  newSearchParams: %s\n", spew.Sdump(patFilter))
	fmt.Printf("findPatient:797  --  raw: %v\n", r.URL.RawQuery)
	fmt.Printf("findPatient:798  --  u: %v\n", u)
	//fmt.Printf("findPatient:673  --  SearchParams: %s\n", spew.Sdump(newSearchParams))
	fmt.Printf("findPatient:800  --  r.Form: %v\n\n\n\n", r.Form)
	if patFilter.Id != "" {
		pat, err := patFilter.FindById()
		if err != nil {
			status := 400
			errMsg := fmt.Sprintf("findPatient:805  --  decode err: %s\n", err.Error())
			WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
			return
		}
		resp := common.ResourceResponse{}
		resp.Patient = pat
		WriteFhirResource(w, 200, &resp)
		return
	}
	fmt.Printf("findPatient:814  --  Calling patFilter.Find\n")
	qry := r.URL.RawQuery
	bundle, err := PatientSearch(cp, qry, "Patient", JWToken)
	if err != nil {
		fmt.Printf("findPatient:818  -- PatientSearch returned err: %s\n", err.Error())
		status := 400
		errMsg := fmt.Sprintf("findPatient:820  --  decode err: %s\n", err.Error())
		WriteFhirOperationOutcome(w, status, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	resp := common.ResourceResponse{}
	resp.Bundle = bundle
	FillResourceResponse(&resp, "patient") // Fills the response resource fields
	//fmt.Printf("\n\nfindPatient:830  --  patients: %s\n", spew.Sdump(resp.Patients))
	fmt.Printf("\n\nfindPatient:828  --  resp: %s\n", spew.Sdump(resp))
	WriteFhirResponse(w, 200, &resp)
}

func GetConnectorPayload(r *http.Request) (*common.ConnectorPayload, error) {
	body, err := io.ReadAll(r.Body) // Should be ConnectorPayload
	if err != nil {
		return nil, log.Errorf("ReadAll FhirSystem error " + err.Error())
	}
	//mt.Printf("GetConnectorPayload:717  -- Got Body Now Unmarshal ConnectorPayload\n")
	b := string(body)
	fmt.Printf("GetConnectorPayload:845  Body: %s\n", b)
	conPayload := &common.ConnectorPayload{}
	err = json.Unmarshal(body, &conPayload)
	if err != nil {
		fmt.Printf("\nGetConnectorPayload:849  --  unmarshal err = %s\n", err.Error())
		// errMsg := err.Error()
		// WriteFhirOperationOutcome(w, 400, CreateOperationOutcome(fhir.IssueTypeProcessing, fhir.IssueSeverityFatal, &errMsg))
		return nil, err
	}
	log.Info("Check ConPayload")
	if conPayload == nil {
		return nil, log.Errorf("conPayload is nil ")
	}
	//fmt.Printf("GetConnectorPayload:860  --  ConnectorPayLoad = %s\n", spew.Sdump(conPayload))
	return conPayload, err
}

func CreateFhirQuery(r *http.Request) (string, error) {
	query := ""
	values := r.URL.Query()
	log.Info(fmt.Sprintf("CreateFhirQuery  values : %v", values))
	if len(values) < 1 {
		err := log.Errorf("Url.Querys are missing")
		return "", err
	}
	//fmt.Printf("\nCreateFhirQuery:713  --  Keys : %v\n\n", keys)
	for k, v := range values {
		log.Info("Key:  " + k + " => " + v[0])
		s := strings.TrimLeft(v[0], "[]")
		if query == "" {
			//for _, kv := range v {
			query = fmt.Sprintf("%s=%s", k, s)
			//}
		} else {
			query = fmt.Sprintf("%s&%s=%s", query, k, s)
		}
		log.Info("CreateFhirQuery = " + query)
	}
	return query, nil
}

func HandleOperationOutcome(w http.ResponseWriter, body []byte) {
	fmt.Printf("HandleOperationOutcome:892  --  body = %s\n", string(body))
	opOutcome := &fhir.OperationOutcome{}
	err := json.Unmarshal(body, &opOutcome)
	if err != nil {
		errMsg := log.ErrMsg(" Error: " + err.Error())
		WriteFhirOperationOutcome(w, 401, CreateOperationOutcome(fhir.IssueTypeForbidden, fhir.IssueSeverityFatal, &errMsg))
		return
	}
	log.Info("opOutcome = " + spew.Sdump(opOutcome))
	issue := opOutcome.Issue[0]
	code := opOutcome.Issue[0].Code
	fmt.Printf("HandleOperationOutcome:904  --  code = %s,  Issue %s\n", code.Display(), *issue.Details.Text)
	//if code.Display() == "Duplicate" {
	//fmt.Printf("HandleOperationOutcome:906  --  code.Display = %s,  Issue %s\n", code.Display(), *issue.Details.Text)
	//WriteFhirOperationOutcome(w, 409, CreateOpOutcome(fhir.IssueTypeDuplicate, fhir.IssueSeverity(fhir.IssueTypeConflict), issue.Details.Text))

	WriteFhirOperationOutcome(w, 409, CreateOpOutcome(opOutcome.Issue))
}

func DetermineOutComeErr(body []byte) error {
	log.Info("HandleOperationOutcome  body n" + string(body))
	opOutcome := &fhir.OperationOutcome{}
	err := json.Unmarshal(body, &opOutcome)
	if err != nil {
		return log.Errorf("Unmarshal err = " + err.Error())
	}
	if opOutcome.Id == nil {
		return log.Errorf("opOutcome.Id is nil")
	}
	//fmt.Printf("HandleOperationOutcome:924  --  opOutcome = %s\n", spew.Sdump(opOutcome))
	issue := opOutcome.Issue[0]
	code := opOutcome.Issue[0].Code
	log.Info(fmt.Sprintf("HandleOperationOutcome:   code = %s,  Issue %s\n", code.Display(), *issue.Details.Text))
	return nil
}
