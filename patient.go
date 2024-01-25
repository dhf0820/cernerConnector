package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	//"io/ioutil"

	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	//"github.com/dhf0820/fhir4"
	fhir "github.com/dhf0820/fhir4"
	//"github.com/dhf0820/uc_core/common"
	//"github.com/samply/golang-fhir-models/fhir-models/fhir"
	common "github.com/dhf0820/uc_common"
	log "github.com/dhf0820/vslog"

	//"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	jw_token "github.com/dhf0820/golangJWT"
	//"io/ioutil"
	"net/http"
	//"os"
	//"strings"
)

type Interface interface{}
type PostPatientPayload struct {
	MRN     string       `json:"mrn"`
	Patient fhir.Patient `json:"patient"`
}

// patId, patMrn, text, err
func SavePatient(mrn string, cp *common.ConnectorPayload, JWToken string) (*http.Response, error) {
	var patient fhir.Patient
	patBytes := []byte{}
	//var err error
	if cp.SavePayload.SrcResource != nil { // Actual patient is provided use it
		patBytes = cp.SavePayload.SrcResource
	} else {
		return nil, fmt.Errorf("No patient information provided.")
	}
	patient, err := fhir.UnmarshalPatient(patBytes)
	//fmt.Printf("SavePatient:43  --  patient: %s\n", spew.Sdump(patient))
	id := primitive.NewObjectID().Hex()
	ident := CreateIdentifier(id)
	//fmt.Printf("SavePatient:46 --  Current Patient.Ident: %s\n", spew.Sdump(patient.Identifier))
	//fmt.Printf("SavePatient:47 --  New Identifier: %s\n", spew.Sdump(ident))
	patient.Identifier = append(patient.Identifier, ident)
	//fmt.Printf("SavePatient:49 --  New Identifiers: %s\n", spew.Sdump(patient.Identifier))
	url := fmt.Sprintf("/%s", "Patient")
	log.Debug3("final Query: " + url)
	//log.Infof("SavePatient:52  --  cp: %s\n", spew.Sdump(cp.ConnectorConfig)) // cp.ConnectorConfig.HostUrl)
	c := New(cp.ConnectorConfig.HostUrl, "application/json+fhir")
	log.Debug3(fmt.Sprintf("Calling postFhir at %s  with %s\n", c.BaseURL, url))
	resp, err := c.postFhir(url, "Patient", JWToken, &patient)
	if err != nil {
		return resp, log.Errorf("!!!fhir query returned err: " + err.Error())
	}
	//fmt.Printf("SavePatient:61  --  postFhir returned: %s\n", spew.Sdump(resp))
	fmt.Printf("SavePatient:62  --  resp.Status: %s\n", resp.Status)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		fmt.Printf("SavePatient:64  --  !!!fhir query returned err: %s\n", err)
		return resp, errors.New(resp.Status)
	} else {
		fmt.Printf("SavePatient:67  --  !!!fhir query returned successful PostPatient\n")
		return resp, nil
	}

	// defer resp.Body.Close()
	// bodyBytes, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("SavePatient:66  --  Error readying body: %s\n", err.Error())
	// 	return nil, fmt.Errorf("SavePatient:67  --  Error readying body: %s", err.Error())
	// }
	// fmt.Printf("SavePatient:69  --  Raw body: %s\n", string(bodyBytes))
	// resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// fmt.Printf("SavePatient:71  --  ReSet resp.Body to initial value\n")

	// // byte, err := ioutil.ReadAll(resp.Body)
	// // if err != nil {
	// // 	fmt.Printf("SavePatient:69  --  Error reading body: %s\n", err.Error())
	// // 	return nil, fmt.Errorf("SavePatient:69  --  Error readying body: %s", err.Error())
	// // }
	// fmt.Printf("SavePatient:74  --  Raw body: %s\n", string(bodyBytes))
	// return resp, nil
	// if resp.StatusCode < 200 || resp.StatusCode > 299 {
	// 	log.Errorf("postFhir:274  --  returned error of %d - %s\n", resp.StatusCode, resp.Status)
	// 	err = fmt.Errorf("%d|postFhir:275 %s", resp.StatusCode, resp.Status)
	// 	//log.Errorf("%s", err.Error())
	// 	return nil, err
	// }
	// defer resp.Body.Close()
	// byte, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("postFhir:282  --  Error readying body: %s", err.Error())
	// }
	// //fmt.Printf("postFhir:284  --  Raw body: %s\n", string(byte))

	// pat, err := fhir.UnmarshalPatient(byte)
	// if err != nil {
	// 	fmt.Printf("postFhir:288  --  Error Decoding Patient: %s\n", err.Error())
	// 	return nil, err
	// }
	// fmt.Printf("postFhir:291  --  Patient =  %s\n", spew.Sdump(pat))
	// return byte, nil
	// fmt.Printf("SavePatient:57  --  postFhir returned: %s\n", spew.Sdump(resp))

	// rawMessage, err := c.PostFhir(url, "Patient", JWToken, patient)
	// if err != nil {
	// 	fmt.Printf("SavePatient:57  --  PostFhir failed: %v\n", err.Error())
	// 	return nil, err
	// }
	// err = nil
	// fmt.Printf("SavePatient:61  --  Patient %s\n", spew.Sdump(rawMessage))
	// pat, err := fhir.UnmarshalPatient(rawMessage)
	// if err != nil {
	// 	err = fmt.Errorf("SavePatient:64  --  UnmarshalPatient failed: %v", err.Error())
	// 	return nil, err
	// }
	// //mrn = GetMrn(&pat, "urn:oid:1.3.6.1.4.1.54392.5.1593.1", "mrn")
	// //text := pat.Text.Div
	// return &pat, nil

	//TODO: CALL CHARTARCHIVE
	// fmt.Printf("SavePatient:65  --  Calling InseretOne\n")
	// result, err := collection.InsertOne(context.TODO(), patient)
	// if err != nil {
	// 	err = fmt.Errorf("savePatient:68  --  insert Patient InsertOne failed: %v", err.Error())
	// 	return "", "", "", err
	// }
	// if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
	// 	//GetMrn(patient, "http://terminology.hl7.org/CodeSystem/v2-0203", "OurMrn")
	// 	fmt.Printf("SavePatient:73  --  Insert Successful: %s\n", oid.Hex())
	// 	GetMrn(patient, "https://fhir.vertisoft.com/6329112852f3616990e2f763/codeSet/4", "OurMrn")
	// 	return oid.Hex(), *ident.Value, patient.Text.Div, err
	// } else {
	// 	err := fmt.Errorf("Invalid objectId")
	// 	return "", "", "", err
	// }
}

func CreateMRN(id string) string {
	return string(id[len(id)-6:])
}

func CreateIdentifier(id string) fhir.Identifier {
	layout := "2006-01-02T15:04:05.000Z"
	ident := fhir.Identifier{}
	ident.Id = StrPtr(primitive.NewObjectID().Hex())
	ident.Use = nil
	cc := fhir.CodeableConcept{}
	code := fhir.Coding{}
	//code.System = StrPtr("urn:oid:1.3.6.1.4.1.54392.5.1593.1")
	//code.System = StrPtr("https://fhir.vertisoft.com/640ba66cbd4105586a6dda75/codeSet/4")
	code.System = StrPtr("http://terminology.hl7.org/CodeSystem/v2-0203")
	code.Code = StrPtr("MR")
	code.Display = StrPtr("Medical Record Number")
	code.UserSelected = BoolPtr(false)

	fmt.Printf("\nCreateIdentifier:107  --  ident : %s\n\n", spew.Sdump(ident))
	//coding := []fhir.
	cc.Coding = append(cc.Coding, code)
	ident.Type = &cc
	ident.Type.Text = StrPtr("mrn")
	ident.Value = StrPtr(CreateMRN(id))
	ident.System = StrPtr("urn:oid:1.3.6.1.4.1.54392.5.1593.1")
	currentTime := time.Now()
	ident.Period = &fhir.Period{}
	ident.Period.Start = StrPtr(currentTime.Format(layout))
	//fmt.Printf("\nCreateIdentifier:117  --  ident : %s\n\n", spew.Sdump(ident))
	return ident
}

//This is Generic Fhir Interface to save a patient

// func (c *Connection) SavePatient(mrn string, patient *fhir.Patient) (*fhir.Patient, error) {

// 	if mrn == "" { // For now use the provided MRN, if not there error //Generate a new MRN and insert into Identifiers.
// 		return nil, errors.New("new UNIQUE MRN for the patient must be specified")
// 	}
// 	if patient == nil {
// 		return nil, errors.New("FHIR (R4) patient must be provided")
// 	}
// 	patient.Id = StrPtr(primitive.NewObjectID().Hex())
// 	patient.Meta = &fhir.Meta{}
// 	patient.Meta.VersionId = StrPtr("1")
// 	patient.Meta.LastUpdated = StrPtr(time.Now().Format("2006-01-02T15:04:05 0000Z"))

// 	ident := fhir.Identifier{}
// 	id := primitive.NewObjectID().Hex()
// 	ident.Id = &id
// 	// idUse := fhir.IdentifierUse.Code(fhir.IdentifierUseUsual)
// 	// fhir.IdentifierUseUsual
// 	//idUse := fhir.IdentifierUseUsual
// 	code := fhir.IdentifierUseUsual
// 	ident.Use = &code
// 	ident.Value = &mrn
// 	ident.Type = &fhir.CodeableConcept{}
// 	ident.Type.Coding = []fhir.Coding{}
// 	coding := fhir.Coding{}
// 	coding.System = StrPtr("http://terminology.hl7.org/CodeSystem/v2-0203")
// 	coding.Code = StrPtr("MR")
// 	coding.Display = StrPtr("Medical record number")
// 	coding.UserSelected = BoolPtr(false)
// 	ident.Type.Coding = append(ident.Type.Coding, coding)
// 	ident.Type.Text = StrPtr("MRN")
// 	//ident.Period
// 	ident.System = StrPtr("http://terminology.hl7.org/CodeSystem/v2-0203") //TODO: Replace with our own.
// 	ident.Value = &mrn
// 	//TODO: add _value Extension  for Rendered Value
// 	patient.Identifier = []fhir.Identifier{}
// 	patient.Identifier = append(patient.Identifier, ident)
// 	fmt.Printf("\npatient: %s\n\n", spew.Sdump(patient))
// 	client := &http.Client{}
// 	fmt.Printf("Save Fhir Patient to: [%s]\n", fhirSystemURL)
// 	bstr, err := json.Marshal(patient)
// 	req, err := http.NewRequest("POST", fhirSystemURL, bytes.NewBuffer(bstr))
// 	if err != nil {
// 		fmt.Printf("NewRequest error: %s\n", err.Error())
// 	}
// 	req.Header.Set("Accept", "application/json+fhir")
// 	req.Header.Set("Content-Type", "application/json")
// 	resp, err := client.Do(req)

// 	if err != nil {
// 		log.Println("Error Posting new Patient:", err.Error())
// 		return nil, err
// 	}
// 	//fmt.Printf("length of ressponse Body = %d\n", len(resp.Body) )
// 	defer resp.Body.Close()
// 	fmt.Printf("resp.StatusCode = %d - %s\n", resp.StatusCode, resp.Status)
// 	// body, err := ioutil.ReadAll(resp.Body)
// 	// if err != nil {
// 	// 	fmt.Printf("Query Error: %v\n", err)
// 	// 	return nil, err
// 	// }

// 	//fmt.Printf("PostPatient response: %s\n", spew.Sdump(resp))
// 	return patient, nil
// }

func GetPatient(patId string) (*fhir.Patient, error) {
	log.Debug3("-- retrieving a patient by id: " + patId)

	filter := bson.M{"id": patId}
	collection, _ := GetCollection("Patients")
	pat := &fhir.Patient{}

	log.Debug3("--Calling FindOne with Filter: " + spew.Sdump(filter))
	err := collection.FindOne(context.TODO(), filter).Decode(pat) // See if the user already has a session
	if err != nil {
		return nil, log.Errorf("-- FindOne error: " + err.Error())
	}
	//fmt.Printf("GetPatient:201  -- FindOne Patient: %s\n", spew.Sdump(pat))
	return pat, err

	// qry := fmt.Sprintf("Patient/%s", patId)
	// log.Infof("Final url to query: %s\n", qry)
	// startTime := time.Now()
	// bytes, err := c.Query(qry)
	// log.Infof("Query time: %s", time.Since(startTime))

	// if err != nil {
	// 	return nil, fmt.Errorf("Query %s failed: %s", qry, err.Error())
	// }
	// patient := fhir.Patient{}
	// err = json.Unmarshal(bytes, &patient)
	// if err != nil {
	// 	return nil, err%1@CU4HZJIYk@IT

	// }
	// return &patient, err
}

func PatientSearch(cp *common.ConnectorPayload, query, token string) (*fhir.Bundle, *common.CacheHeader, error) {
	payload, _, err := jw_token.ValidateToken(JWToken, "")
	if err != nil {
		return nil,nil,  err
	}
	userId := payload.UserId
	// fhirID, err := primitive.ObjectIDFromHex(fhirId)
	// if err != nil {
	// 	return nil, err
	// }
	log.Debug3("--  queryString: %s" + query)
	qry := fmt.Sprintf("Patient?%s", query)
	log.Debug3("--  Final url to query: " + qry)

	log.Debug3("cp.ConnectorConfig = " + spew.Sdump(cp.ConnectorConfig))
	log.Debug3("URL = " + cp.ConnectorConfig.HostUrl)
	baseUrl := cp.ConnectorConfig.HostUrl
	c := New(baseUrl, "application/json+fhir")
	log.Debug3(fmt.Sprintf("CallGetFhirBundle at %s  with %s", c.BaseURL, qry))
	startTime := time.Now()
	bundle, err := c.GetFhirBundle(qry, token)
	if err != nil {
		log.Error("getFhirBundle error: " + err.Error())
	}
	log.Debug3("--  Bundle= " + spew.Sdump(bundle))

	enteries := len(bundle.Entry)

	log.Debug3("--  Number of entries in bundle: " + fmt.Sprint(enteries))
	//log.Debug5("bundle: " + spew.Sdump(bundle))
	header := &common.CacheHeader{}
	header.SystemCfg = cp.System
	connConfig := cp.ConnectorConfig
	header.ResourceType = "Patient"
	header.UserId = userId
	header.PageId = 1
	queryId := primitive.NewObjectID().Hex()
	header.QueryId = queryId
	//log.Debug3("connConfig: " + spew.Sdump(connConfig))
	header.CacheBase = fmt.Sprintf("%s/%s", connConfig.CacheUrl, header.SystemCfg.ID.Hex())
	log.Debug3("Header:" + spew.Sdump(header))
	//header.ResourceCacheBase = fmt.Sprintf("%s/%s/%s/BundleTransaction", connConfig.CacheUrl, header.FhirSystem.ID.Hex())
	header.GetBundleCacheBase = fmt.Sprintf("%s/%s/BundleTransaction", header.CacheBase, header.SystemCfg.ID.Hex())
	header.GetResourceCacheBase = fmt.Sprintf("%s/%s/CachePage", header.CacheBase, header.SystemCfg.ID.Hex())

	cacheBundle := common.CacheBundle{}
	cacheBundle.PageId = header.PageId
	cacheBundle.Header = header
	cacheBundle.ID = primitive.NewObjectID()
	//fmt.Printf("\n\n\n\n$$$ FindResource:110 calling CacheResourceBundleAndEntries (without bundle) - %s \n", spew.Sdump(cacheBundle))
	//fmt.Printf("FindResource:126  --  bundle = %s\n", spew.Sdump(bundle))
	//Cache the first bundle(page)
	log.Debug3(fmt.Sprintf("--  Query %s for %ss took %s\n\n\n", cp.ConnectorConfig.Label, "Patient", time.Since(startTime)))
	log.Debug3("--  UnmarshalBundle")
	// bundle := fhir4.Bundle{}
	// bundle, err = fhir4.UnmarshalBundle(byte)
	// if err != nil {
	// 	return 0, nil, nil, err
	// }
	//log.Debug5("bundle: " + spew.Sdump(bundle))
	cacheBundle.Bundle = bundle
	//startTime = time.Now()
	bundle.ResourceType = StrPtr("Bundle")
	if UseCache() {
		log.Debug3("calling CacheResourceBundleAndEntries")
		pg, err := CacheResourceBundleAndEntries(&cacheBundle, JWToken, 0)
		log.Debug3(fmt.Sprintf("CacheResourceBundleAndEntries returned %d %ss in page: %d for %s  took %s", len(cacheBundle.Bundle.Entry), resource, page, sysCfg.DisplayName, time.Since(startTime)))
		if err != nil {
			//return err and done
			return int64(bundle, cacheBundle.Header, err // cacheBundle.Header, err
		}
		log.Debug3("links: " + spew.Sdump(bundle.Link))
		//Follow the bundle links to retrieve all bundles(pages) in the query response
		nextURL := GetNextObservationUrl(bundle.Link)
		total := int64(0)
		if nextURL == "" {
			log.Debug3(fmt.Sprintf("GetNext%sUrl initialy No Next - One page only ", "Patient"))
			total, err = TotalCacheForQuery(cacheBundle.QueryId)
			cacheBundle.Header.PageId = pg
			log.Debug3("total: " + fmt.Sprint(total))
			//page++
			return  bundle, cacheBundle.Header, err
		}
		page++
		go c.GetNextObservation(header, nextURL, resource, JWToken, page)
	} else {
		log.Info("Not Using Caching")
	}

	// cb := uc_core/common.CacheBundle{}
	// cb.
	// 	CacheResourceBundleAndEntries(bundle, JWToken, 1)
	//fmt.Printf("PatientSearch:237  --  Bundle= %s\n\n\n", spew.Sdump(bundle))
	return bundle, err
	/*
		if err != nil {

			return nil, fmt.Errorf("Query %s failed: %s", query, err.Error())
		}

		//fmt.Printf("\n\n\n@@@ RAW Patient: %s\n\n\n", pretty.Pretty(b))
		// prettyJSON, err := json.MarshalIndent(b, "", "    ")
		// if err != nil {
		// 	fmt.Printf("MarshalIndent failed: %s\n", err.Error())
		// 	return nil, err
		// }

		startTime = time.Now()
		bundle := &fhir.Bundle{}
		//data := PatientResult{}
		if err := json.Unmarshal(bytes, &bundle); err != nil {
			return nil, fmt.Errorf("PatientSearch ummarshal : %s", err.Error())
		}
		log.Infof("Unmarshal time: %s", time.Since(startTime))
		//fmt.Printf("Response: %s\n", spew.Sdump(bundle))
		//resourceCache := common.ResourceCache

		for _, entry := range bundle.Entry {
			resourceCache := common.ResourceCache{}
			resourceJson := entry.Resource
			patient := fhir.Patient{}
			json.Unmarshal(resourceJson, &patient)
			resourceCache.Resource = entry.Resource
			resourceCache.ResourceType = "Patient"
			fmt.Printf("PatientSearch:160  --  PatientId = %s\n", *patient.Id)

		}
		header := &common.CacheHeader{}
		header.FhirSystem = fhirSystem
		cacheBundle := common.CacheBundle{}
		cacheBundle.ID = primitive.NewObjectID() //Each cach bundle gets a new header. The queryId ties all pages together.

		header.FhirId = fhirSystem.ID.Hex()            // Uniquely identifies the real url fo the fhir server
		header.QueryId = primitive.NewObjectID().Hex() //Does not change on each page
		header.PatientId = ""                          // Not used for patient cache sine each entry is a different patient
		header.ResourceType = "Patient"
		tn := time.Now()
		header.CreatedAt = &tn

		cacheBundle.Header = header
		cacheBundle.Bundle = bundle
		cacheBundle.Header.PageId = 1

		//TODO: Call Core CacheResources to cachhe the resources(patients)
		fmt.Printf("PatientSearch:179 calling Insert %d Patients for now\n", len(cacheBundle.Bundle.Entry))

		err = Insert(context.Background(), &cacheBundle, token)
		if err != nil {
			msg := fmt.Sprintf("CacheInsert initial error %s", err.Error())
			fmt.Println(msg)
			log.Error(msg)
			return nil, errors.New(msg)
		}
		nextURL := GetNextResourceUrl(bundle.Link)
		if nextURL == "" {
			msg := fmt.Sprintf("GetNextResourceUrl initial No Next ")
			// fmt.Println(msg)
			log.Warn(msg)
			//return nil, errors.New(msg)
			return bundle, nil
		}
		go c.GetNextResource(header, nextURL, resource, token)
	*/
	return bundle, nil

}

// func GetNextResourceUrl(link []fhir.BundleLink) string {
// 	for _, lnk := range link {
// 		if lnk.Relation == "next" {
// 			return lnk.Url
// 		}
// 	}
// 	return ""
// }
// func (c *Connection) GetNextResource(header *common.CacheHeader, url, token string) {
// 	startTime := time.Now()
// 	bytes, err := c.GetFhir(url)
// 	fmt.Printf("Query Next Set time: %s\n", time.Since(startTime))
// 	if err != nil {
// 		msg := fmt.Sprintf("c.GetFhir error: %s", err.Error())
// 		fmt.Println(msg)
// 		log.Error(msg)
// 		return
// 	}
// 	bundle := &fhir.Bundle{}

// 	if err := json.Unmarshal(bytes, bundle); err != nil {
// 		msg := fmt.Sprintf("PatientSearch next unmarshal : %s", err.Error())
// 		log.Error(msg)
// 		fmt.Println(msg)
// 		return
// 	}
// 	header.PageId += 1
// 	tn := time.Now()
// 	header.CreatedAt = &tn
// 	cacheBundle := common.CacheBundle{}
// 	cacheBundle.ID = primitive.NewObjectID()
// 	cacheBundle.Header = header
// 	cacheBundle.Bundle = bundle

// 	err = Insert(context.Background(), &cacheBundle)
// 	if err != nil {
// 		msg := fmt.Sprintf("CacheInsert error %s", err.Error())
// 		fmt.Println(msg)
// 		log.Error(msg)
// 	}
// 	fmt.Printf("Link: %s\n", spew.Sdump(bundle.Link))
// 	nextURL := GetNextResourceUrl(bundle.Link)
// 	if nextURL == "" {
// 		msg := fmt.Sprintf("GetNextResourceUrl Last page had %d Resources processed ", len(bundle.Entry))
// 		// fmt.Println(msg)
// 		log.Warn(msg)
// 		fmt.Printf("GetNext Resources should return\n")
// 		return
// 	} else {
// 		fmt.Printf("GetNextResources is being called in the background\n")
// 		go c.GetNextResources(header, nextURL, token)
// 		fmt.Printf("GetNextResources was called in the background\n")
// 	}
// 	fmt.Printf("GetNext Resource is returning\n")
// 	return
// }

func GetMrn(pat *fhir.Patient, system string, code string) string {
	idents := pat.Identifier
	for _, ident := range idents {
		if *ident.Type.Text == code {
			value := *ident.Value
			fmt.Printf("GetMrn:381  --  MRN Code : %s = %s\n", code, value)
			return value
		}
	}
	fmt.Printf("GetMrn:385  -- Identifier for System: %s   Code: %s was not found\n", system, code)
	return ""
}

func (pf *PatientFilter) Find() ([]fhir.Patient, error) {
	fmt.Printf("Find:390  -- PatientFilter: %s\n", spew.Sdump(pf))
	newErr := fmt.Errorf("pf.Find:391  ==  returning Find not implemented")
	fmt.Printf("%s\n", newErr.Error())
	return nil, newErr
}

func (pf *PatientFilter) FindById() (*fhir.Patient, error) {
	fmt.Printf("FindById:397  -- PatientFilter: %s\n", spew.Sdump(pf))
	if pf.Id == "" {
		return nil, errors.New("FindById:400  --  Id is required")
	}
	filter := bson.D{{"id", pf.Id}}
	collection, _ := GetCollection("Patients")
	pat := &fhir.Patient{}
	fmt.Printf("FindById:404  --  Calling FindOne with Filter: %v\n", filter)
	err := collection.FindOne(context.TODO(), filter).Decode(pat) // See if the user already has a session
	if err != nil {
		return nil, log.Errorf("FindOne error: " + err.Error())
	}
	//fmt.Printf("GetPatient:158  -- FindOne Patient: %s\n", spew.Sdump(pat))
	return pat, err
	//return nil, fmt.Errorf("Find not implemented")
}

func (c *Connection) PostPatient(cp *common.ConnectorPayload, mrn string, patient *fhir.Patient) (*fhir.Patient, error) {
	if cp == nil {
		return nil, log.Errorf("ConnectorPayload must be provided")
	}
	systemURL := cp.ConnectorConfig.HostUrl
	if systemURL == "" {
		//if systemURL == "" {
		return nil, log.Errorf("cp.SystemUrl to add patient to must be specified")
	}
	if mrn == "" { // For now use the provided MRN, if not there error //Generate a new MRN and insert into Identifiers.
		return nil, log.Errorf("new UNIQUE MRN for the patient must be specified")
	}
	if patient == nil {
		return nil, log.Errorf("FHIR (R4) patient must be provided")
	}
	patient.Id = StrPtr(primitive.NewObjectID().Hex())
	patient.Meta = &fhir.Meta{}
	patient.Meta.VersionId = StrPtr("1")
	patient.Meta.LastUpdated = StrPtr(time.Now().Format("2006-01-02T15:04:05 0000Z"))

	ident := fhir.Identifier{}
	id := primitive.NewObjectID().Hex()
	ident.Id = &id
	// idUse := fhir.IdentifierUse.Code(fhir.IdentifierUseUsual)
	// fhir.IdentifierUseUsual
	//idUse := fhir.IdentifierUseUsual
	code := fhir.IdentifierUseUsual
	ident.Use = &code
	ident.Value = &mrn
	ident.Type = &fhir.CodeableConcept{}
	ident.Type.Coding = []fhir.Coding{}
	coding := fhir.Coding{}
	coding.System = StrPtr("http://terminology.hl7.org/CodeSystem/v2-0203")
	coding.Code = StrPtr("MR")
	coding.Display = StrPtr("Medical record number")
	coding.UserSelected = BoolPtr(false)
	ident.Type.Coding = append(ident.Type.Coding, coding)
	ident.Type.Text = StrPtr("MRN")
	//ident.Period
	ident.System = StrPtr("urn:oid:1.3.6.1.4.1.54392.5.1593.1")
	ident.Value = &mrn
	//TODO: add _value Extension  for Rendered Value
	patient.Identifier = []fhir.Identifier{}
	patient.Identifier = append(patient.Identifier, ident)
	fmt.Printf("\nPostPatient:454  --  %s\n\n", spew.Sdump(patient))
	client := &http.Client{}
	fmt.Printf("PostPatient:456  --  Save Fhir Patient to: [%s]\n", systemURL)
	bstr, err := json.Marshal(patient)
	if err != nil {
		fmt.Printf("PostPatient:459  --  Marshal error: %s\n", err.Error())
		return nil, err
	}
	req, err := http.NewRequest("POST", systemURL, bytes.NewBuffer(bstr))
	if err != nil {
		fmt.Printf("NewRequest error: %s\n", err.Error())
	}
	req.Header.Set("Accept", "application/json+fhir")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		return nil, log.Errorf("Error Posting new Patient: " + err.Error())
	}
	//fmt.Printf("length of ressponse Body = %d\n", len(resp.Body) )
	defer resp.Body.Close()
	log.Debug3(fmt.Sprintf("resp.StatusCode = %d - %s", resp.StatusCode, resp.Status))
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Printf("Query Error: %v\n", err)
	// 	return nil, err
	// }

	//fmt.Printf("PostPatient response: %s\n", spew.Sdump(resp))
	return patient, nil
}
