package dal

import "go.mongodb.org/mongo-driver/mongo"

type ReturnWrapper struct {
	Status  bool
	Message string
}

type MongoDbConnectionCreateInfo struct {
	Host     string
	Port     int32
	User     string
	Password string
}

type MongoDbConnectionInfo struct {
	Client *mongo.Client
	Err    error
}

type MongoDbDataBaseInfo struct {
	SwcDb      *mongo.Database
	MetaInfoDb *mongo.Database
}

const (
	DefaultMetaInfoDataBaseName string = "MetaInfoDataBase"
	DefaultSwcDataBaseName      string = "SwcDataBase"
)
