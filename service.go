package main

import (
	// "context"
	"encoding/json"

	//"errors"
	"fmt"
	//jw_token "github.com/dhf0820/jwToken"
	//"github.com/davecgh/go-spew/spew"

	"io"
	"strings"

	common "github.com/dhf0820/uc_common"
	log "github.com/sirupsen/logrus"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"

	//"github.com/dhf0820/ids_model/common"
	//"net"
	"net/http"
	"os"
	"time"
	//"github.com/sirupsen/logrus"
)

// var Log = logrus.New()
var (
	// DateMods      map[string]string
	// SqlMods       map[string]string
	// StringMods    map[string]string
	// NumericMods   map[string]string
	Conf   *common.ServiceConfig //*config
	CoreEp *common.EndPoint
	CaEp   *common.EndPoint
	//DbConnector 	*common.DataConnector
	// serverAddress string
	// baseAddress   string
	// port          string
	// tlsMode       string
	// deployMode    string
	// lis           net.Listener
	PatientData string
	//*messaging.MessagingClient
	//GWConfig *m.DeliveryConfig
)

type ConfigResp struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Config  common.ServiceConfig `json:"config"`
}

func Start(serviceName, codeVersion string) {
	baseAddress := os.Getenv("BASE_ADDRESS")
	if baseAddress == "" {
		baseAddress = "localhost"
	}
	listenPort := os.Getenv("LISTEN_PORT")
	if listenPort == "" {
		log.Warnf("Start:61  --  LISTEN_PORT not set in environment. Meaning  running standalone.  Using default port 40103")
		listenPort = "40103" // Default for Ca3Connector local host standalone
	}
	restAddress := fmt.Sprintf("%s:%s", "0.0.0.0", listenPort)
	log.Printf("Start:64  --  restAddress: %s  -  port: %s", restAddress, listenPort)
	router := NewRouter()
	configVersion := os.Getenv("CONFIG_VERSION")

	log.Printf("Start:68  --  %s CodeVersion: [%s]  ConfigVersion: [%s]  is listening for restful requests at %s", serviceName, codeVersion, configVersion, restAddress)
	err := http.ListenAndServe(restAddress, router)
	log.Printf("Start:70  --  This should not happen err = %s", err.Error())

}

func Initialize(serviceName, version string) (*common.ServiceConfig, error) {
	var err error
	log.Printf("Initiallizing %s version %s\n", serviceName, version)
	if serviceName == "" {
		if os.Getenv("SERVICE_NAME") == "" {
			serviceName = "cerner_conn"
			os.Setenv("SERVICE_NAME", "cerner_conn")
		}
	}
	if os.Getenv("SERVICE_VERSION") == "" {
		os.Setenv("SERVICE_VERSION", "local_test")
	}

	if os.Getenv("SERVICE_COMPANY") == "" {
		os.Setenv("SERVICE_COMPANY", "test")
	}

	fmt.Printf("Initialize:91  --  Service: %s  SERVICE_VERSION: %s  SERVICE_COMPANY = [%s]\n\n", serviceName, os.Getenv("SERVICE_VERSION"), os.Getenv("SERVICE_COMPANY"))
	// PatientData := strings.ToLower(os.Getenv("PATIENT_DATA"))
	// if strings.Trim(PatientData, " ") == "" {
	// 	PatientData = "postgres"
	// }
	Conf, err = GetServiceConfig(strings.ToLower(serviceName), strings.ToLower(os.Getenv("SERVICE_VERSION")),
		strings.ToLower(os.Getenv("SERVICE_COMPANY")))
	if err != nil {
		return nil, fmt.Errorf("could not retieve Configuration : %s", err.Error())
	}
	//DbConnector, err = common.GetDatabaseByName(Conf.DataConnectors, "mongo")
	//fmt.Printf("\n----config: %s]\n", spew.Sdump(Conf))

	// OpenCaDB(Conf.DataConnectors)

	setEndPoints()
	fmt.Printf("Initilized %s\n\n", serviceName)
	return Conf, err
}

func GetServiceConfig(name, version, company string) (*common.ServiceConfig, error) {
	log.Printf("GetServiceConfig:111   Starting\n")
	//_, _ = GetCollection("sys_config")
	var cfg ConfigResp
	var err error
	//var unmarshalErr *json.UnmarshalTypeError
	var bdy []byte
	cfg = ConfigResp{}
	log.Printf("GetServiceConfig:118  --  name: %s, version: %s, company: %s", name, version, company)
	//coreName := strings.ReplaceAll(os.Getenv("CORE_NAME_PORT"), " ","")
	//api := os.Getenv("API")
	//Log.Infof("API: [%s]\n", api)
	configAddr := os.Getenv("CONFIG_ADDRESS")
	// configAddr = "http://docker1.ihids.com:19101/api/rest/v1"
	log.Printf("GetServiceConfig:124  --  ConfigAddress: %s", configAddr)
	url := fmt.Sprintf("%s/config?name=%s&version=%s&company=%s", configAddr, name, version, company)
	fmt.Printf("GetServiceConfig:126  --  core url: %s\n", url)
	startTime := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("GetServiceConfig:130  --  http.Get(%s)  returned error: %s", url, err.Error())
		return nil, err
	}
	fmt.Printf("GetServiceConfig:133  --  Elapsed Time: %s\n", time.Since(startTime))
	fmt.Printf("GetServiceConfig:134  --  status : %d\n", resp.StatusCode)
	defer resp.Body.Close()
	//cfg = mod.ServiceConfig{}
	bdy, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("raw string: %s\n", string(bdy))
	err = json.Unmarshal(bdy, &cfg)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("GetServiceConfig:224  --  Config = %s\n", spew.Sdump(cfg))
	//fmt.Printf("GetServiceConfig:225  --  Elapsed Time: %s\n", time.Since(startTime))
	//fmt.Printf("GetServiceConfig:226 - status : %d\n", resp.StatusCode)
	//_, err = GetCollection("sys_config")
	//}
	//Conf := &cfg.Config
	//DatabaseConnector, err := common.GetDatabaseByName(Conf.DataConnectors, "mongo")
	//fmt.Printf("CFG: = %s\n", spew.Sdump(cfg.Config))
	//.Printf("\n\n###Conf: = %s\n", spew.Sdump(Conf))
	// dbName := os.Getenv("DBNAME")
	// if dbName == "" {
	// 	dbName = "test"
	// 	fmt.Printf("GetServiceConfig:238  --  DBNAME not provideded using %s\n", dbName)
	// }
	// DbConnector, err = common.GetDatabaseByName(Conf.DataConnectors, dbName)
	// fmt.Printf("GetServiceConfig:241  -  Opening MongoDb Name: %s\n", dbName)
	// OpenMongoDB()
	return &cfg.Config, err
}

//func GetServiceEndpoint(value string) *mod.EndPoint {
//	fmt.Printf("GetServiceEndpoint\n")
//	endPoints := GetConfig().ServiceEndpoints
//	for _, ep := range endPoints {
//		fmt.Printf("Looking at %s for %s\n", ep.Name, value)
//		if ep.Name == value {
//			fmt.Printf("Found Endpoint: %s\n", ep.Name)
//			return ep
//		}
//	}
//	fmt.Printf("--Endpoint not found\n")
//	return nil
//}

func GetConfig() *common.ServiceConfig {

	return Conf
}

func setEndPoints() {
	//fmt.Printf("\n--In setEndPoints\n")
	// CoreEp = GetServiceEndpoint(Conf.ServiceEndPoints, "core")
	// if CoreEp == nil {
	// 	Log.Errorf("---Core EndPoint was not found in configuration: %s\n", spew.Sdump(Conf))
	// } else {
	// 	//fmt.Printf("---CoreEP: %s\n", spew.Sdump(CoreEp))
	// }

	// CaEp = GetServiceEndpoint(Conf.ServiceEndPoints, "ca_api")
	// if CaEp == nil {
	// 	Log.Errorf("---ca_api EndPoint was not found in configuration: %s\n", spew.Sdump(Conf))
	// } else {
	// 	//fmt.Printf("---CaEP: %s\n", spew.Sdump(CaEp))
	// }
}

func GetConfigDataElement(name string) string {
	data := Conf.Data
	for _, elem := range data {
		if elem.Name == name {
			return elem.Value
		}
	}
	return ""

}

// func GetFhirSystem(id string) (*common.FhirSystem, error) {
// 	//fmt.Printf("\n\n\nGetFhirSystem:277 - for %s\n\n", id)
// 	//fmt.Printf("GetFhirSystem:278 - GetCollection\n")
// 	startTime := time.Now()
// 	collection, err := GetCollection("fhirSystem")
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Printf("GetFhirSystem:221 - Elapsed time: %s\n", time.Since(startTime))
// 	oid, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, fmt.Errorf("GetFhirSystem:224  -  invalid FhirId: [%s]", id)
// 	}
// 	query := bson.D{{"_id", oid}}
// 	fmt.Printf("Query: %v\n", query)
// 	//filter := bson.D{{"name", "demo.Cerner"}}
// 	filter := bson.D{{"_id", oid}}
// 	//filterM := bson.M{"_id": oid}
// 	fhirSystem := &common.FhirSystem{}
// 	startTime = time.Now()
// 	fmt.Printf("\n\n\nGetFhirSystem:233 -  FindOne fhirConnector: bson.D %v\n", filter)
// 	err = collection.FindOne(context.Background(), filter).Decode(fhirSystem)
// 	//fmt.Printf("GetFhirSystem:290 - Elapsed Time: %s\n", time.Since(startTime))
// 	// if err != nil {
// 	// 	fmt.Printf("   Now Calling GetFhirConnector FindOne SvcConfig: bson.M %v\n", filterM)
// 	// 	err = collection.FindOne(context.Background(), filterM).Decode(fhirConfig)
// 	// }
// 	if err != nil {
// 		log.Errorf("GetFhirSystem:241 FindOne %v NotFound\n", filter)
// 		return nil, fmt.Errorf("GetFhirSystem:304  FindOne %v NotFound\n", filter)
// 	}
// 	return fhirSystem, nil
// }

// func CreateJWToken() (string, error) {
// 	err := os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
// 	if err != nil {
// 		return "", err
// 	}
// 	os.Setenv("ACCESS_SECRET", "I am so blessed Debbie loves me!")
// 	dur := time.Duration(300) * time.Second
// 	jwt, err := token.CreateToken("192.168.1.2", "DHarman", dur, "dharman0127", "Debbie Harman", "Physician")
// 	//maker, err := token.NewJWTMaker(os.Getenv("ACCESS_SECRET"))
// 	if err != nil {
// 		return "", err
// 	}
// 	return jwt, nil
// }
