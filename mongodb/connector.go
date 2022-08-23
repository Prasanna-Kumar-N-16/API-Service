package mongodb

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConfig struct {
	DBURL    string `json:"uri" yaml:"uri"`
	Username string `json:"userName" yaml:"userName"`
	Password string `json:"password" yaml:"password"`
	Timeout  int64  `json:"timeout" yaml:"timeout"`
}

type DBConnector struct {
	dbConfig DBConfig
	Client   *mongo.Client
}

var dbConnector *DBConnector

type MongoDBManager struct {
	Instance *mongo.Collection
	ParamID  string
}

func (dbConfig DBConfig) NewConnection() (*DBConnector, error) {
	if dbConnector == nil || dbConnector.Client == nil {
		dbConnector = &DBConnector{}
		return dbConfig.openConnection()
	}
	return dbConnector, nil
}

func GetConnection() (*DBConnector, error) {
	if dbConnector == nil || dbConnector.Client == nil {
		return nil, errors.New("connection is not opened")
	}
	return dbConnector, nil
}

func (dbConnector *DBConnector) GetHistoryManager(paramID, collectionName string) (*MongoDBManager, error) {
	if collectionName == "" {
		return nil, errors.New("unable to get the collection name")
	}

	collection, collectionError := getCollection(paramID, collectionName)
	if collectionError != nil {
		return nil, collectionError
	}
	mongoDBManager := &MongoDBManager{}
	mongoDBManager.Instance = collection
	return mongoDBManager, nil
}

func (dbConnector *DBConnector) GetManager(db, collectionName string) (*MongoDBManager, error) {
	if collectionName == "" {
		return nil, errors.New("unable to get the collection name")
	}

	collection, collectionError := getCollection(db, collectionName)
	if collectionError != nil {
		return nil, collectionError
	}
	mongoDBManager := &MongoDBManager{}
	mongoDBManager.Instance = collection
	return mongoDBManager, nil
}

func isConnected() bool {
	if dbConnector == nil || dbConnector.Client == nil {
		return false
	}
	return true
}

func (dbConfig DBConfig) openConnection() (*DBConnector, error) {
	var client options.ClientOptions
	clientOptions := client.ApplyURI(dbConfig.DBURL)
	if dbConfig.Username != "" && dbConfig.Password != "" {
		credential := options.Credential{
			Username: dbConfig.Username,
			Password: dbConfig.Password,
		}
		clientOptions = clientOptions.SetAuth(credential)
	}
	mongoClient, connectionError := mongo.NewClient(clientOptions)
	if connectionError != nil {
		return nil, connectionError
	}
	if dbConfig.Timeout == 0 {
		dbConfig.Timeout = 10
	}
	mongoConTimeout := time.Duration(dbConfig.Timeout * int64(time.Second))
	ctx, _ := context.WithTimeout(context.Background(), mongoConTimeout)
	//defer cancel()
	connectionError = mongoClient.Connect(ctx)
	if connectionError != nil {
		return nil, connectionError
	}
	dbConnector.Client = mongoClient
	dbConnector.dbConfig = dbConfig
	return dbConnector, nil
}

func getCollection(paramID, collectionName string) (*mongo.Collection, error) {
	if !isConnected() {
		return nil, errors.New("unable to connect db")
	}
	db := dbConnector.Client.Database(paramID)
	collection := db.Collection(collectionName)
	return collection, nil
}
