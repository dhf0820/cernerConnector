package main

import (
	common "github.com/dhf0820/uc_core/common"
	"github.com/gorilla/schema"
)

var (
	//Conf        *common.ServiceConfig
	Mongo       *MongoDB
	Company     string
	DbConnector *common.DataConnector
)

func init() {
	decoder = schema.NewDecoder()
}
