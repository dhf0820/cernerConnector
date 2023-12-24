package main

import (
	"context"
	"fmt"

	//"github.com/davecgh/go-spew/spew"
	"os"
	"strconv"
	"time"

	//"github.com/davecgh/go-spew/spew"
	common "github.com/dhf0820/uc_common"
	log "github.com/dhf0820/vslog"

	//"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client       *mongo.Client
	DatabaseName string
	URL          string
	Database     *mongo.Database
	Session      mongo.Session
	Collection   *mongo.Collection
}

var DB MongoDB
var mongoClient *mongo.Client
var dbConnector *common.DataConnector

// var DbConnector *DataConnector
var insertResult *mongo.InsertOneResult

func OpenDBUrl(dbURL string) *MongoDB {
	var err error
	//svcConfig := GetConfig()
	//if svcConfig == nil {
	//	fmt.Printf("\n---$$$Config is not initialized\n\n")
	//}
	//startTime := time.Now()
	uri := dbURL
	log.Debug3("Opening database uri: " + uri)
	//uri := dbURL + databaseName
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	// client, error := vsmongo.NewClient(options.Client().ApplyURI("ur_Database_uri"))
	// error = client.Connect(ctx)

	// //Checking the connection
	// error = client.Ping(context.TODO(), nil)
	// fmt.Println("Database connected")

	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	log.Debug3("-- Connecting to mongo")
	if mongoClient, err = mongo.Connect(ctx, opts); err != nil {
		msg := log.ErrMsg("mongo.Connect error: " + err.Error())
		log.Error(msg)
		return nil
	}
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		//fmt.Printf("Database did not connect:62 %v\n", err)
		log.Error("--Database did not connect:   " + err.Error())
		return nil
	}
	Company := os.Getenv("COMPANY")
	log.Info("Company(DatabaseName: " + Company)

	DB.Client = mongoClient
	DB.DatabaseName = Company //.Getenv("COMPANY") //DbConnector.Database  //databaseName
	DB.Database = mongoClient.Database(DB.DatabaseName)
	DB.URL = dbURL
	log.Info("Database: " + DB.DatabaseName + " Connected")

	//fmt.Printf("Client: %s\n", spew.Sdump(client))

	DB.Collection = DB.Client.Database(DB.DatabaseName).Collection(GetDbField("collection"))
	//fmt.Printf("DBOpen-77 took %d ms\n", time.Since(startTime).Milliseconds())
	return &DB
}

func OpenMongoDB() (*MongoDB, error) {
	var err error
	//svcConfig := GetConfig()
	//if svcConfig == nil {
	//	fmt.Printf("\n---$$$Config is not initialized\n\n")
	//}
	//startTime := time.Now()
	//DbConnector = svcConfig.DataConnector
	//dbURL := DbConnector.Server
	// DbConnector, err := common.GetDatabaseByName(Conf.DataConnectors, "mongo")
	// if err != nil {
	// 	return nil, err
	// }
	dbURL := DbConnector.Server
	//dbURL := DBUrl() //os.Getenv("CORE_DB")
	uri := dbURL
	log.Debug3("Opening Mongo URI: " + uri)
	//uri := dbURL + databaseName
	//ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	//defer cancel()

	// client, error := vsmongo.NewClient(options.Client().ApplyURI("ur_Database_uri"))
	// error = client.Connect(ctx)

	// //Checking the connection
	// error = client.Ping(context.TODO(), nil)
	// fmt.Println("Database connected")

	clientOptions := options.Client()
	clientOptions.ApplyURI(uri)
	clientOptions.SetMaxPoolSize(5)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Info("Using new connect routine from atlas")
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		msg := log.ErrMsg(" Mongo.Connect error: " + err.Error())
		log.Fatal(msg)
	}

	// if mongoClient, err = mongo.Connect(ctx, clientOptions); err != nil {
	// 	msg := fmt.Sprintf("Mongo.Connect error: %s\n", err.Error())
	// 	log.Error(msg)
	// 	return nil
	// }
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		//fmt.Printf("Database did not connect: %v\n", err)
		err = log.Errorf("did not connect by ping: " + err.Error())
		return nil, err
	}

	DB.Client = mongoClient
	DB.DatabaseName = DbConnector.Database
	//DB.DatabaseName = os.Getenv("COMPANY") //DbConnector.Database  //databaseName
	//DB.DatabaseName = "test"
	//fmt.Printf("DatabaseName:140 -  %s\n", DB.DatabaseName)
	DB.Database = mongoClient.Database(DB.DatabaseName)
	DB.URL = dbURL
	log.Info("Database: " + DB.DatabaseName + " connected")
	DB.Collection = DB.Client.Database(DB.DatabaseName).Collection(GetDbField("collection"))
	//fmt.Printf("vsmongo:153 - DBOpen -- took %d ms\n", time.Since(startTime).Milliseconds())
	return &DB, err
}

func DBUrl() string {

	//cacheDB := os.Getenv("CACHE_DB")
	var err error
	//log.Debug3(" GetDatabaseByName Conf: " + spew.Sdump(Conf))
	dbConnector, err = common.GetDatabaseByName(Conf.DataConnectors, "mongo")
	if err != nil {
		log.Error("GetDatabaseByName error: " + err.Error())
		return ""
	}
	//fmt.Printf("vsmongo.DBUrl:162 - GotDataConnector: %s\n", spew.Sdump(dbConnector))
	return dbConnector.Server
}

// ConnectToDB starts a new database connection and returns a reference to it
func ConnectToMongoDB() (*MongoDB, error) {
	// DbConnector = GetConfig().DataConnector
	// databaseName := DbConnector.Database
	// url := DbConnector.Server
	// coreDB := os.Getenv("CORE_DB")
	// if coreDB == "" {
	// 	coreDB = "LOCAL_DB"
	// }
	// fmt.Printf("Using database: %s\n", coreDB)
	url := DBUrl()
	if url == "" {
		log.Error("coreDB is not defined. Should contain the name of the actual Database to use")
		panic("coreDB is not defined")
	}

	fmt.Printf("Use DB URL: %s\n", url)
	//databaseName := os.Getenv("COMPANY")
	//databaseName := Company
	databaseName := dbConnector.Database
	log.Info("Using DB: " + databaseName)
	//if url == "" {
	//	url = settings.DbURL()
	//}
	//fmt.Printf("Mongo URL: %s\n", url)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	options := options.Client().ApplyURI(url)

	options.SetMaxPoolSize(DbPoolSize())
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}
	DB.Client = client
	DB.DatabaseName = databaseName
	DB.Database = client.Database(DB.DatabaseName)
	DB.URL = url
	return &DB, nil
}

func DbPoolSize() uint64 {
	var poolSize uint64
	poolSizeString := GetDbField("poolsize")
	if poolSizeString == "" {
		poolSizeString = "100"
	}
	poolSizeInt64, err := strconv.ParseInt(poolSizeString, 10, 64)
	if err == nil {
		poolSize = uint64(poolSizeInt64)
	} else {
		poolSize = 100
	}
	return poolSize
}
func Current() (*MongoDB, error) {
	if DB.Client != nil {
		return &DB, nil
	}
	_, err := ConnectToMongoDB()
	//client, err := Open("")
	return &DB, err
}

func (db *MongoDB) Close() error {
	err := db.Client.Disconnect(context.TODO())
	return err
}

func GetCollection(collection string) (*mongo.Collection, error) {
	if collection == "" {
		collection = CollectionName()
		log.Info("Using default Collection: " + collection)
	}
	db, err := Current() //"mongodb://admin:Sacj0nhat1@cat.vertisoft.com:27017")
	if err != nil {

		log.Fatal(err.Error())
		//return nil, err
	}
	client := db.Client
	coll := client.Database(DB.DatabaseName).Collection(collection)
	log.Info("Changed to Collection: " + collection)
	return coll, nil
}

func CollectionName() string {
	return "srv_config"
}

func GetDbField(key string) string {
	return ""
	// //LogMessage(&payload, "Detailed", "Info", "Checking config value for field: "+field, payload.Config.Core_log_url)
	// flds := mod.DataConnector.Fields
	// for _, fld := range flds {
	// 	switch {
	// 	case fld.Name == key:
	// 		return fld.Value
	// 	}
	// }
	// return ""

}

// IsDup returns whether err informs of a duplicate key error because
// a primary key index or a secondary unique index already has an entry
// with the given value.
func IsDup(err error) bool {
	if wes, ok := err.(mongo.WriteException); ok {
		for i := range wes.WriteErrors {
			if wes.WriteErrors[i].Code == 11000 || wes.WriteErrors[i].Code == 11001 || wes.WriteErrors[i].Code == 12582 || wes.WriteErrors[i].Code == 16460 {
				return true
			}
		}
	}
	return false
}
