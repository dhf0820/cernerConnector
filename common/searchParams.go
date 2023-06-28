package common

import (
	//"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/schema"
	"github.com/oleiade/reflections"
	"net/http"
	//"reflect"
	//"strconv"
	"log"

	"strings"
)

type PSearchParams struct {
	OffSet     SearchParam   `json:"offset" schema:"_offset"`
	Count      SearchParam   `json:"count" schema:"_count"`
	Order      SearchParam   `json:"order" schema:"_order"`
	Sort       SearchParam   `json:"sort" schema:"_sort"`
	Page       SearchParam   `json:"page" schema:"_page"`
	Id         SearchParam   `json:"id" schema:"id"`
	MRN        SearchParam   `json:"mrn" schema:"mrn"`
	SSN        SearchParam   `json:"ssn" schema:"ssn"`
	Identifier SearchParam   `json:"identifier" schema:"identifier"`
	Gender     SearchParam   `json:"gender" schema:"gender"`
	BirthDate  SearchParam   `json:"birthdate" schema:"birthdate"`
	Family     SearchParam   `json:"family" schema:"family"`
	Given      SearchParam   `json:"given" schema:"given"`
	Facility   SearchParam   `json:"facility" schema:"facility"`
	Active     SearchParam   `json:"active" schema:"active"`
	DOB        []SearchParam `json:"dob" schema:"dob"`
	BaseUrl    string        `json:"base_url"`
	RequestURI string        `json:"request_uri"`
	Limit      uint32        `json:"limit"`
	Skip       uint32        `json:"skip"`
}

type PatientFhirSearchParams struct {
	Count      SearchParam `json:"count" schema:"_count"`
	OffSet     SearchParam `json:"offset" schema:"_offset"`
	Order      SearchParam `json:"order" schema:"_order"`
	Sort       SearchParam `json:"sort" schema:"_sort"`
	Page       SearchParam `json:"page" schema:"_page"`
	Id         SearchParam `json:"id" schema:"_id"`
	MRN        SearchParam `json:"mrn" schema:"mrn"`
	SSN        SearchParam `json:"ssn" schema:"ssn"`
	Identifier SearchParam `json:"identifier" schema:"identifier"`
	Gender     SearchParam `json:"gender" schema:"gender"`
	BirthDate  SearchParam `json:"birthdate" schema:"birthdate"`
	Name       SearchParam `json:"name" schema:"name"`
	Family     SearchParam `json:"family" schema:"family"`
	Given      SearchParam `json:"given" schema:"given"`
	Phone      SearchParam `json:"phone" schema:"phone"`
	Email      SearchParam `json:"email" schema:"email"`
	PostalCode SearchParam `json:"address-postalcode" schema:"address-postalcode"`

	Facility   SearchParam   `json:"facility" schema:"facility"`
	Active     SearchParam   `json:"active" schema:"active"`
	DOB        []SearchParam `json:"dob" schema:"dob"`
	BaseUrl    string        `json:"base_url"`
	RequestURI string        `json:"request_uri"`
	Limit      uint32        `json:"limit"`
	Skip       uint32        `json:"skip"`
}

type DocumentSearchParams struct {
	Count      SearchParam   `json:"count" schema:"_count"`
	OffSet     SearchParam   `json:"offset" schema:"_offset"`
	Order      SearchParam   `json:"order" schema:"_order"`
	Sort       SearchParam   `json:"sort" schema:"_sort"`
	Page       SearchParam   `json:"page" schema:"_page"`
	Id         SearchParam   `json:"id" schema:"id"`
	Patient    SearchParam   `json:"patient" schema:"patient"`
	Subject    SearchParam   `json:"subject" schema:"subject"`
	Encounter  SearchParam   `json:"encounter" schema:"encounter"`
	Created    []SearchParam `json:"created" schema:"created"` // need to handle two of these
	Facility   SearchParam   `json:"facility" schema:"facility"`
	BaseUrl    string        `json:"base_url"`
	RequestURI string        `json:"request_uri"`
	Limit      uint32        `json:"limit"`
	Skip       uint32        `json:"skip"`
}

type SearchParam struct {
	Schema   string
	Modifier string
	Value    string
}

var decoder = schema.NewDecoder()

func FhirPatientSearch(r *http.Request) (*PatientFhirSearchParams, error) {
	fmt.Printf("FhirPatientSearch:93  --  rawQuery: %s\n", r.URL.RawQuery)
	err := r.ParseForm()
	if err != nil {
		return nil, fmt.Errorf("FhirPatientSearch:96  --  error: %s", err.Error())
		// Handle error
	}

	var pat PatientFhirSearchParams

	// r.PostForm is a map of our POST form values
	err = decoder.Decode(&pat, r.PostForm)
	if err != nil {
		return nil, err
	}
	fmt.Printf("\nFhirPatientSearch:107  --  PatFilter: %s\n", spew.Sdump(pat))
	return &pat, err
}

func PatientSearchparams(r *http.Request) (*PSearchParams, error) {
	var pspTags map[string]string
	tagFields := make(map[string]string)

	// //buildFieldsByTagMap("schema", *psp)
	// //facility := "demo"
	// fmt.Printf("searchPatient called\n")
	// if err := r.ParseForm(); err != nil {
	// 	err = fmt.Errorf("Error parsing query: %s", err.Error())
	// 	return nil, err
	// }
	psp := new(PSearchParams)
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("ParseForm error: %s\n", err.Error())
		return nil, err
	}
	fmt.Printf("query: %s\n", r.URL.RawQuery)
	pspTags, _ = reflections.Tags(psp, "schema")
	for k, v := range pspTags {
		//	fmt.Printf("key: = %s;  value: %s\n", k, v)
		tagFields[v] = k
	}
	fmt.Printf("\ntagFields: %s\n\n", spew.Sdump(tagFields))
	qryParams := strings.Split(r.URL.RawQuery, "&")
	for _, param := range qryParams {
		fmt.Printf("patientSearchParam:98  -- param: %s\n", param)
		keyValue := strings.Split(param, "=") //split key and value
		value := keyValue[1]
		fmt.Printf("parts : %v\n", keyValue[0])   // key is elem 0 value is elem 1
		keyMod := strings.Split(keyValue[0], ":") // separate the modifier from the key if any
		key := keyMod[0]
		mod := ""
		if len(keyMod) == 2 { //There is a modifier
			mod = keyMod[1]
		}
		if len(keyMod) == 2 { //There is a modifier
			mod = keyMod[1]
		}
		param := SearchParam{}
		param.Modifier = mod
		param.Value = strings.Trim(value, " ")
		param.Schema = key
		fmt.Printf("Key: %s,  mod: %s  value: %s  Field: %s\n", key, mod, value, spew.Sdump(param))
		fieldName := tagFields[key]
		fmt.Printf("Setting Field Data\n")
		err := reflections.SetField(psp, fieldName, param)
		if err != nil {
			fmt.Printf("Set Field %s error: %s\n", spew.Sdump(param), err.Error())
		}
		fmt.Printf("Interim psp := %s\n", spew.Sdump(psp))
	}
	fmt.Printf("psp := %s\n", spew.Sdump(psp))
	if psp.Id.Value != "" {
		fmt.Printf("Find Patient by ID: %s\n", psp.Id.Value)
	}

	// decoder := schema.NewDecoder()
	// // r.PostForm is a map of our POST form values
	// err = decoder.Decode(psp, r.PostForm)

	// if err != nil {
	// 	fmt.Printf("decoder.Decode error: %s\n", err.Error())
	// 	return nil, err
	// }
	//fmt.Printf("psp := %s\n", spew.Sdump(psp))
	// //TODO: Include the facility as part of the base url not as a parameter
	// fmt.Printf("tls: %v\n", r.TLS)
	// protocol := "http://"
	// psp.BaseUrl = fmt.Sprintf("%s%s/api/rest/v1", protocol, r.Host)
	// psp.RequestURI = r.RequestURI
	// // psp.Facility.Value = strings.Trim(psp.Facility.Value, " ")
	// // if psp.Facility.Value == "" {
	// // 	err := fmt.Errorf("Faciity is required")
	// // 	return nil, err
	// // }

	// pspTags, _ = reflections.Tags(psp, "schema")
	// for k, v := range pspTags {
	// 	//	fmt.Printf("key: = %s;  value: %s\n", k, v)
	// 	tagFields[v] = k
	// }
	// fmt.Printf("\ntagFields: %s\n\n", spew.Sdump(tagFields))

	// var decoder = schema.NewDecoder()
	// decoder.IgnoreUnknownKeys(true)
	// fmt.Printf("query: %s\n", r.URL.RawQuery)
	// qryParams := strings.Split(r.URL.RawQuery, "&")
	// for _, param := range qryParams {
	// 	fmt.Printf("patientSearchParam:98  -- param: %s\n", param)
	// 	keyValue := strings.Split(param, "=") //split key and value
	// 	value := keyValue[1]
	// 	fmt.Printf("parts : %v\n", keyValue[0])   // key is elem 0 value is elem 1
	// 	keyMod := strings.Split(keyValue[0], ":") // separate the modifier from the key if any
	// 	key := keyMod[0]
	// 	mod := ""
	// 	if len(keyMod) == 2 { //There is a modifier
	// 		mod = keyMod[1]
	// 	}
	// 	param := SearchParam{}
	// 	param.Modifier = mod
	// 	param.Value = strings.Trim(value, " ")
	// 	param.Schema = key
	// 	fmt.Printf("Key: %s,  mod: %s  value: %s  Field: %s\n", key, mod, value, spew.Sdump(param))
	// 	fieldName := tagFields[key]
	// 	fmt.Printf("Setting Field Data\n")
	// 	err := reflections.SetField(psp, fieldName, param)
	// 	if err != nil {
	// 		fmt.Printf("Set Field %s error: %s\n", spew.Sdump(param), err.Error())
	// 	}
	// }

	// if psp.Count.Value == "" { // set ount to default 20
	// 	psp.Count.Value = "20"
	// 	psp.Count.Schema = "_count"
	// }
	// count, err := strconv.ParseUint(psp.Count.Value, 10, 32)
	// if err != nil {
	// 	err = fmt.Errorf("invalid _count Err: %s", err.Error())
	// 	return nil, err
	// }
	return psp, nil
}

type User struct {
	email    string
	password string
	created  string
}

func registrationHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(2048)
	if err != nil {
		http.Error(w, "Invalid or missing fields", http.StatusBadRequest)
		return
	}

	log.Printf("registrationHandler: %+v\n", r.PostForm)
	user := new(User)

	// r.PostForm is a map of our POST form values
	err = decoder.Decode(user, r.PostForm)

	if err != nil {
		log.Printf("Error: %+v\n", err)
		http.Error(w, "Error Invalid or Missing Params", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%+v\n", user)
	log.Printf("%+v\n", user)

	log.Printf("Email is %s", user.email)
	//Setting this as a session cookie so we don't have to parameterize it.
	// session, _ := store.Get(r, "user")
	// session.Values["user"] = user
	// session.Save(r, w)

	//We now have a user object with an email and a password
	//Query the users collection to determine if there is a user with that email already

}
