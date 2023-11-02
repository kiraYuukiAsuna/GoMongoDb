package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func initializeNewDataBase(metaInfoDataBaseName string, swcDataBaseName string) {
	createInfo := MongoDbConnectionCreateInfo{
		host:     "127.0.0.1",
		port:     1425,
		user:     "userdbmanager",
		password: "userdbmanagerpassword",
	}

	connectionInfo := connectToMongoDb(createInfo)

	if connectionInfo.err != nil {
		log.Fatal(connectionInfo.err)
	}

	databaseNames, err := connectionInfo.client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	databaseExists1 := false
	for _, dbName := range databaseNames {
		if dbName == metaInfoDataBaseName {
			databaseExists1 = true
			break
		}
	}
	if databaseExists1 {
		fmt.Printf("Database %s exists!\n", metaInfoDataBaseName)
	} else {
		fmt.Printf("Database %s does not exist.\n", metaInfoDataBaseName)
	}

	databaseExists2 := false
	for _, dbName := range databaseNames {
		if dbName == swcDataBaseName {
			databaseExists2 = true
			break
		}
	}
	if databaseExists2 {
		fmt.Printf("Database %s exists!\n", swcDataBaseName)
	} else {
		fmt.Printf("Database %s does not exist.\n", swcDataBaseName)
	}

}
