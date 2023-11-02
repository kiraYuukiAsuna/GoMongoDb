package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func connectToMongoDb(createInfo MongoDbConnectionCreateInfo) MongoDbConnectionInfo {
	url := "mongo://" + createInfo.host + ":" + string(createInfo.port)
	var connectionInfo MongoDbConnectionInfo
	connectionInfo.client, connectionInfo.err = mongo.Connect(context.TODO(), options.Client().ApplyURI(url).SetConnectTimeout(10*time.Second))
	if connectionInfo.err != nil {
		log.Fatal(connectionInfo.err)
		return connectionInfo
	}

	if connectionInfo.client.Ping(context.TODO(), nil) != nil {
		log.Fatal(connectionInfo.err)
		return connectionInfo
	}

	return connectionInfo
}

func connectToDataBase(connectionInfo MongoDbConnectionInfo, metainfoDataBaseName string, swcDataBaseName string) MongoDbDataBaseInfo {
	if connectionInfo.err != nil {
		log.Fatal(connectionInfo.err)
		return MongoDbDataBaseInfo{}
	}

	var dbInfo MongoDbDataBaseInfo

	dbInfo.metaInfoDb = connectionInfo.client.Database(metainfoDataBaseName)
	dbInfo.swcDb = connectionInfo.client.Database(swcDataBaseName)

	return dbInfo
}
