package main

import (
	"fmt"

	//"github.com/davecgh/go-spew/spew"
	//common "github.com/dhf0820/uc_common"
	//"github.com/joho/godotenv"
	//"github.com/ory/dockertest/docker/types/versions"
	"github.com/joho/godotenv"
	//"net/http"
	"os"

	//service "github.com/dhf0820/baseConnector/services"
	log "github.com/sirupsen/logrus"
	// "strings"
)

//var err error
var err error

func main() {
	version := "230627_1"
	switch os.Getenv("MODE") {
	case "local":
		err = godotenv.Load("./.env.ca3_conn")
	case "test":
		err = godotenv.Load("./.env.ca3_conn_test")
	default:
		err = godotenv.Load("./.env")
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
	log.Printf("main:41  --  Starting ca3_conn %s\n\n", version)
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "12113"
	// 	fmt.Printf("No port set up using: %s\n ", port)
	// }
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Printf("main:49  --  Environment is not set using %s\n", os.Getenv("MODE"))
		err := godotenv.Load("./.env.ca3_conn_test")
		if err != nil {
			fmt.Printf("Main:50  --  Get environment: %s  err: %s\n", ".env.ca3_conn_test", err.Error())
			os.Exit(1)
		}
		serviceName = os.Getenv("SERVICE_NAME")
		//serviceName = "baseConnector" // baseConnector
	}
	secretKey := os.Getenv("ACCESS_SECRET")
	if secretKey == "" {
		secretKey = "I am so blessed Debbie loves me!"
		os.Setenv("ACCESS_SECRET", secretKey)
	}

	log.Printf("main:64  --  Calling service Start for %s  version: %s\n", serviceName, version)
	Start(serviceName, version) //Should Not Return

}
