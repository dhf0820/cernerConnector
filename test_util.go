package main

import (
	//"bytes"
	"context"
	//"encoding/json"
	//"io"

	//"github.com/gorilla/mux"
	//"github.com/davecgh/go-spew/spew"
	//"github.com/joho/godotenv"

	//"gitlab.com/dhf0820/ids_model/common"

	"github.com/davecgh/go-spew/spew"
	log "github.com/dhf0820/vslog"
	//"github.com/sirupsen/logrus"
	//. "github.com/smartystreets/goconvey/convey"
	"fmt"
	//"net/http"

	//"net/http/httptest"
	//"os"
	//"strings"
	//"testing"

	//"time"

	//jw_token "github.com/dhf0820/jwToken"
	common "github.com/dhf0820/uc_common"

	//"github.com/dhf0820/uc_core/service"

	//"github.com/davecgh/go-spew/spew"
	//fhir "github.com/dhf0820/fhir4"
	"go.mongodb.org/mongo-driver/bson"

	//fhirR4go "github.com/dhf0820/fhirR4go"
	//. "github.com/smartystreets/goconvey/convey"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSystemConfigById(id primitive.ObjectID) (*common.SystemConfig, error) {
	strId := id.Hex()
	log.Debug3("--   Id: " + strId)
	collection, err := GetCollection("systemConfig")
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": id}
	sysConfig := &common.SystemConfig{}
	fmt.Printf("GetSystemConfig:24  --  Calling FindOne SystemConfig: %v\n", filter)
	err = collection.FindOne(context.Background(), filter).Decode(&sysConfig)
	return sysConfig, err
}

func CreateCP(includeSave bool) *common.ConnectorPayload {
	log.Debug3("Entering CreateCP")
	id, _ := primitive.ObjectIDFromHex("640ba5e3bd4105586a6dda74")
	sc, err := GetSystemConfigById(id)
	if err != nil {
		log.Error("--  GetSystemConfigById error: " + err.Error())
		return nil
	}
	// sc := uc_core/common.SystemConfig{}
	// sc.Name = "uc_ca3"
	// ids := []*uc_core/common.KVData{}

	// mrnIdent := uc_core/common.KVData{}
	// mrnIdent.Name = "mrn"
	// mrnIdent.Value = "mrn|"
	// ids = append(ids, &mrnIdent)
	// ssnIdent := uc_core/common.KVData{}
	// ssnIdent.Name = "ssn"
	// ssnIdent.Value = "ssn|"
	// ids = append(ids, &ssnIdent)
	// idIdent := uc_core/common.KVData{}
	// idIdent.Name = "id"
	// idIdent.Value = "id|"
	// ids = append(ids, &idIdent)
	// sc.Identifiers = ids

	cp := common.ConnectorPayload{}
	cc := common.ConnectorConfig{}
	cc.ID, _ = primitive.ObjectIDFromHex("62f1c5dab3070d0b40e7aac1")
	cc.Name = "uc_ca3"
	cc.Version = "local"
	cc.Label = "CA3FhirConnector"
	cc.Credentials = ""
	cc.HostUrl = "http://ca_3backend:4000/api/rest/v1/"
	cc.URL = "http://uc_cernerConnector:20103"
	data := []*common.KVData{}
	cacheServer := common.KVData{}
	cacheServer.Name = "cacheServer"
	cacheServer.Value = "http://universalcharts.com:30201"
	data = append(data, &cacheServer)
	hostServer := common.KVData{}
	hostServer.Name = "cacheHost"
	hostServer.Value = "http://ucCache:9200"
	data = append(data, &hostServer)
	cc.Data = data
	cc.CacheUrl = "http://universalcharts.com:30201"

	//TODO: AddFhirAuthToken
	cp.ConnectorConfig = &cc
	cp.System = sc
	// if includeSave {
	// 	cp.SavePayload = &common.SavePayload{}
	// 	cp.SavePayload.SrcResource = SamplePatient()
	// 	cp.SavePayload.ResourceType = "Patient"
	// 	cp.SavePayload.SrcPatient = SampleFhirPatient()
	// }

	//cp.BaseAddress = "http://192.168.1.148:4100"
	// fmt.Printf("CreateCP:617  --  returning cp: %s\n\n", spew.Sdump(cp))
	// fmt.Println()
	// finalJson, err := json.Marshal(cp)
	// if err != nil {
	// 	fmt.Printf("CreateCP:621  --  json.Marshal error: %s\n", err.Error())
	// 	return nil
	// }
	//fmt.Printf("\n\n\nCreateCP:623  --  finalJson: %s\n", finalJson)
	log.Debug3("CreateCP Returning: " + spew.Sdump(cp))
	return &cp
}

func CreateCernerCP(includeSave bool) *common.ConnectorPayload {
	fmt.Printf("CreateCernerCP:744  --  includeSave: %v\n", includeSave)
	id, _ := primitive.ObjectIDFromHex("640ba5e3bd4105586a6dda74")
	sc, err := GetSystemConfigById(id)
	if err != nil {
		fmt.Printf("CreateCernerCP:729  --  GetSystemConfigById error: %s\n", err.Error())
		return nil
	}
	//ids := []*uc_core/common.KVData{}

	// mrnIdent := uc_core/common.KVData{}
	// mrnIdent.Name = "mrn"
	// mrnIdent.Value = "mrn|"
	// ids = append(ids, &mrnIdent)
	// ssnIdent := uc_core/common.KVData{}
	// ssnIdent.Name = "ssn"
	// ssnIdent.Value = "ssn|"
	// ids = append(ids, &ssnIdent)
	// idIdent := uc_core/common.KVData{}
	// idIdent.Name = "id"
	// idIdent.Value = "id|"
	// ids = append(ids, &idIdent)
	// sc.Identifiers = ids

	cp := common.ConnectorPayload{}
	cc := common.ConnectorConfig{}
	cc.ID, _ = primitive.ObjectIDFromHex("6488a9580403ff647fca2f7e")
	cc.Name = "uc_cerner"
	cc.Version = "local"
	cc.Label = "CernerConnector"
	cc.Credentials = ""
	cc.HostUrl = "https://fhir-open.cerner.com/r4/ec2458f2-1e24-41c8-b71b-0e701af7583d"
	//cc.HostUrl = "http://192.168.1.148:4100/system/640ba66cbd4105586a6dda75"
	cc.URL = "192.168.1.152:20103"
	data := []*common.KVData{}
	cacheServer := common.KVData{}
	cacheServer.Name = "cacheServer"
	cacheServer.Value = "http://universalcharts.com:30201"
	data = append(data, &cacheServer)
	hostServer := common.KVData{}
	hostServer.Name = "cacheHost"
	hostServer.Value = "http://ucCache:9200"
	data = append(data, &hostServer)
	cc.Data = data
	cc.CacheUrl = "http://universalcharts.com:30201"

	//TODO: AddFhirAuthToken
	cp.ConnectorConfig = &cc
	cp.System = sc
	// if includeSave {
	// 	cp.SavePayload = &common.SavePayload{}
	// 	cp.SavePayload.SrcResource = SamplePatient()
	// 	cp.SavePayload.ResourceType = "Patient"
	// 	cp.SavePayload.SrcPatient = SampleFhirPatient()
	// }

	//cp.BaseAddress = "http://192.168.1.148:4100"
	// fmt.Printf("CreateCP:617  --  returning cp: %s\n\n", spew.Sdump(cp))
	// fmt.Println()
	// finalJson, err := json.Marshal(cp)
	// if err != nil {
	// 	fmt.Printf("CreateCP:383  --  json.Marshal error: %s\n", err.Error())
	// 	return nil
	// }
	//fmt.Printf("\n\n\nCreateCP:386  --  finalJson: %s\n", finalJson)
	fmt.Printf("\n\n\n")
	return &cp
}
