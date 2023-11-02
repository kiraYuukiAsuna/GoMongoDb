package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type MongoDbConnectionCreateInfo struct {
	host     string
	port     int32
	user     string
	password string
}

type MongoDbConnectionInfo struct {
	client *mongo.Client
	err    error
}

type MongoDbDataBaseInfo struct {
	swcDb      *mongo.Database
	metaInfoDb *mongo.Database
}
