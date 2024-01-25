package main

import (
	"fmt"

	//"github.com/davecgh/go-spew/spew"
	//common "github.com/dhf0820/uc_core/common"
	//"github.com/joho/godotenv"
	//"github.com/ory/dockertest/docker/types/versions"
	"github.com/joho/godotenv"
	//"net/http"
	"os"

	//service "github.com/dhf0820/baseConnector/services"
	log "github.com/dhf0820/vslog"
	//"github.com/sirupsen/logrus"
	// "strings"
)

// var err error
var err error
var version string
var Mode string
var Env string
var ResponseType string
var QueryString string

func main() {
	version = "240124.0"
	log.SetDebuglevel("DEBUG3")
	log.Info("run mode: " + os.Getenv("MODE"))
	switch os.Getenv("MODE") {
	case "local":
		err = godotenv.Load("./.env.cerner_conn")
		if err == nil {
			Mode = "local"
			Env = "./.env.cerner_conn"
		}
	case "test":
		err = godotenv.Load("./.env.cerner_conn_test")
		if err == nil {
			Mode = "test"
			Env = "./.env.cerner_conn_test"
		}
	case "go_test":
		err = godotenv.Load("./.env.cerner_conn_go_test")
		if err == nil {
			Mode = "go_test"
			Env = "./.env.cerner_conn_go_test"
		}

	default:
		err = godotenv.Load("./.env")
		if err != nil {
			log.Error("Error loading environment: " + err.Error())
			os.Exit(1)
		}
		if err == nil {
			Mode = os.Getenv("MODE")
			Env = "./.env"
		}
		log.Info("run mode: default")
	}
	// if err != nil {
	// 	fmt.Printf("Main:41  --  Error loading environment: %s\n", err.Error())
	// 	os.Exit(1)
	// }
	// if os.Getenv("MODE") == "go_test" {
	// 	godotenv.Load("./.env.ca3_conn_test")
	// } else {
	// 	godotenv.Load("./.env_ca3_conn")
	// }

	os.Setenv("CodeVersion", version)
	log.Debug3("Starting cernerConnector version" + version)
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "12113"
	// 	fmt.Printf("No port set up using: %s\n ", port)
	// }
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Printf("main:51  --  Environment is not set using %s\n", os.Getenv("MODE"))
		err := godotenv.Load("./.env.cerner_conn_go_test")
		if err != nil {
			fmt.Printf("Main:54  --  Get environment: %s  err: %s\n", ".env.ca3_conn_test", err.Error())
			os.Exit(1)
		}
		serviceName = os.Getenv("SERVICE_NAME")
	}
	fmt.Printf("main:61  --  Service Name: %s\n", serviceName)
	secretKey := os.Getenv("ACCESS_SECRET")
	if secretKey == "" {
		secretKey = "I am so blessed Debbie loves me!"
		os.Setenv("ACCESS_SECRET", secretKey)
	}

	log.Debug3(fmt.Sprintf("Calling service Start for %s  version: %s", serviceName, version))
	log.Info("Service Starting with mode: " + os.Getenv("MODE") + "  env: " + Env)
	Start(serviceName, version) //Should Not Return

}
