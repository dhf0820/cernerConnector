package main
import (
	"fmt"
	fhir "github.com/dhf0820/fhir4"

)

func SamplePatient() []byte {
	return []byte(`{
		"resourceType": "Patient",
		"id": "12742611",

		"text": {
			"status": "generated",
			"div": "<div><p><b>Patient</b></p><p><b>Name</b>: SMART, Debra</p><p><b>DOB</b>: Jan 27, 1950</p><p><b>Administrative Gender</b>: Female</p><p><b>Status</b>: Active</p></div>"
		},
		"identifier": [
			{
				"use": "usual",
				"type": {
					"coding": [
						{
							"system": "http://hl7.org/fhir/v2/0203",
							"code": "MR",
							"display": "Medical record number",
							"userSelected": false
						}
					],
					"text": "MRN"
				},
				"system": "urn:oid:2.16.840.1.113883.6.1000",
				"value": "106979",
				"_value": {
					"extension": [
						{
							"url": "http://hl7.org/fhir/StructureDefinition/rendered-value",
							"valueString": "00000106979"
						}
					]
				},
				"period": {
					"start": "2023-01-02T19:34:51.000Z"
				}
			}
		],
		"active": true,
		"name": [
			{
				"use": "official",
				"text": "SMART, Debra",
				"family": [
					"SMART"
				],
				"given": [
					"Debra"
				],
				"period": {
					"start": "2023-01-01T19:15:48.000Z"
				}
			}
		],
		"telecom": [
			{
				"system": "email",
				"value": "dsmart@yopmail.com",
				"use": "home",
				"period": {
					"start": "2023-01-01T19:15:47.000Z"
				}
			}
		],
		"gender": "female",
		"birthDate": "1958-01-27"
	}`)
}

func SampleFhirPatient() *fhir.Patient {
	//fp := fhir.Patient{}
	//fmt.Printf("\nfhirPatient: %s\n\n", spew.Sdump(fp))
	fhirPat, err := fhir.UnmarshalPatient(SamplePatient())
	if err != nil {
		fmt.Printf("UnmarshalPatient error : %s\n", err.Error())
		return nil
	}
	return &fhirPat
}
