package dal

import (
	"DBMS/dbmodel"
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var g_MongoDbDataBaseInfo MongoDbDataBaseInfo

func SetDbInstance(instance MongoDbDataBaseInfo) {
	g_MongoDbDataBaseInfo = instance
}

func GetDbInstance() MongoDbDataBaseInfo {
	return g_MongoDbDataBaseInfo
}

func InitializeNewDataBase(metaInfoDataBaseName string, swcDataBaseName string) {
	createInfo := MongoDbConnectionCreateInfo{
		Host:     "127.0.0.1",
		Port:     27017,
		User:     "defaultuser",
		Password: "defaultpassword",
	}

	connectionInfo := ConnectToMongoDb(createInfo)

	if connectionInfo.Err != nil {
		log.Fatal(connectionInfo.Err)
	}

	var dbInfo MongoDbDataBaseInfo
	dbInfo.MetaInfoDb = connectionInfo.Client.Database(metaInfoDataBaseName)
	dbInfo.SwcDb = connectionInfo.Client.Database(swcDataBaseName)

	databaseNames, err := connectionInfo.Client.ListDatabaseNames(context.TODO(), bson.M{})
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
		log.Fatalf("Database %s exists! Check your database!\n", metaInfoDataBaseName)

	} else {
		log.Printf("Database %s does not exist. Start to create a new one!\n", metaInfoDataBaseName)

		var err error
		err = dbInfo.MetaInfoDb.CreateCollection(context.TODO(), dbmodel.ProjectMetaInfoCollectionString)
		if err != nil {
			log.Fatal(err)
		}

		err = dbInfo.MetaInfoDb.CreateCollection(context.TODO(), dbmodel.UserMetaInfoCollectionString)
		if err != nil {
			log.Fatal(err)
		}

		err = dbInfo.MetaInfoDb.CreateCollection(context.TODO(), dbmodel.PermissionGroupMetaInfoCollectioString)
		if err != nil {
			log.Fatal(err)
		}

		var permissionGroup = dbmodel.PermissionGroupMetaInfoV1{
			Base: dbmodel.MetaInfoBase{
				Id:         primitive.NewObjectID(),
				ApiVersion: "V1",
				Uuid:       uuid.NewString(),
			},
			Name:        "Admin",
			Description: "Admin Permission Group",
			Global: dbmodel.GlobalPermissionMetaInfoV1{
				WritePermissionCreateProject: true,
				WritePermissionModifyProject: true,
				WritePermissionDeleteProject: true,
				ReadPerimissionQuery:         true,
			},
			Project: dbmodel.ProjectPermissionMetaInfoV1{
				WritePermission_AddData:    true,
				WritePermission_ModifyData: true,
				WritePermission_DeleteData: true,
				ReadPerimission_Query:      true,
			},
		}

		CreatePermissionGroup(permissionGroup, dbInfo)

		err = dbInfo.MetaInfoDb.CreateCollection(context.TODO(), dbmodel.SwcMetaInfoCollectionString)
		if err != nil {
			log.Fatal(err)
		}

		opts := options.CreateCollection().SetCapped(true).SetMaxDocuments(1000).SetSizeInBytes(100 * 1024 * 1025)
		err = dbInfo.MetaInfoDb.CreateCollection(context.TODO(), dbmodel.DailyStatisticsMetaInfoCollectionString, opts)
		if err != nil {
			log.Fatal(err)
		}
	}

	databaseExists2 := false
	for _, dbName := range databaseNames {
		if dbName == swcDataBaseName {
			databaseExists2 = true
			break
		}
	}
	if databaseExists2 {
		log.Fatalf("Database %s exists! Check your database!\n", swcDataBaseName)
	} else {
		log.Printf("Database %s does not exist. Will create new one when needed!\n", swcDataBaseName)
	}

}
