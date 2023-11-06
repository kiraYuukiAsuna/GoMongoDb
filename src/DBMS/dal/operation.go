package dal

import (
	"DBMS/dbmodel"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"time"
)

func ConnectToMongoDb(createInfo MongoDbConnectionCreateInfo) MongoDbConnectionInfo {
	//mongodb://defaultuser:defaultpassword@localhost:27017/?authMechanism=DEFAULT
	//url := "mongodb://" + createInfo.user + ":" + createInfo.password + "@" + createInfo.host + ":" + string(createInfo.port) + "/?authMechanism=DEFAULT"
	url := "mongodb://" + createInfo.Host + ":" + strconv.Itoa(int(createInfo.Port))
	var connectionInfo MongoDbConnectionInfo
	connectionInfo.Client, connectionInfo.Err = mongo.Connect(context.TODO(), options.Client().ApplyURI(url).SetConnectTimeout(10*time.Second))
	if connectionInfo.Err != nil {
		log.Fatal(connectionInfo.Err)
		return connectionInfo
	}

	var err = connectionInfo.Client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return connectionInfo
	}

	return connectionInfo
}

func ConnectToDataBase(connectionInfo MongoDbConnectionInfo, metainfoDataBaseName string, swcDataBaseName string) MongoDbDataBaseInfo {
	if connectionInfo.Err != nil {
		log.Fatal(connectionInfo.Err)
		return MongoDbDataBaseInfo{}
	}

	var dbInfo MongoDbDataBaseInfo

	dbInfo.MetaInfoDb = connectionInfo.Client.Database(metainfoDataBaseName)
	dbInfo.SwcDb = connectionInfo.Client.Database(swcDataBaseName)

	return dbInfo
}

func CreateProject(projectMetaInfo dbmodel.ProjectMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var projectCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.ProjectMetaInfoCollectionString)

	result := projectCollection.FindOne(context.TODO(), bson.D{
		{"Name", projectMetaInfo.Name}})

	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			_, err := projectCollection.InsertOne(context.TODO(), projectMetaInfo)
			if err != nil {
				return ReturnWrapper{false, "Create user failed! Error:" + err.Error()}
			}
			return ReturnWrapper{true, "Create project successfully!"}
		}
		return ReturnWrapper{false, "Unknown error!"}
	} else {
		// find one means already exist
		return ReturnWrapper{false, "Project already exit!"}
	}

}

func DeleteProject(projectMetaInfo dbmodel.ProjectMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var projectCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.ProjectMetaInfoCollectionString)

	result := projectCollection.FindOneAndDelete(context.TODO(), bson.D{
		{"Name", projectMetaInfo.Name}})

	if result.Err() != nil {
		return ReturnWrapper{false, result.Err().Error()}
	} else {
		return ReturnWrapper{true, "Delete successfully!"}
	}
}

func ModifyProject(projectMetaInfo dbmodel.ProjectMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var projectCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.ProjectMetaInfoCollectionString)

	result := projectCollection.FindOneAndReplace(
		context.TODO(),
		bson.D{{"Name", projectMetaInfo.Name}},
		projectMetaInfo)

	if result.Err() != nil {
		return ReturnWrapper{false, "Update project info failed! Error:" + result.Err().Error()}
	} else {
		return ReturnWrapper{true, "Update project info success!"}
	}

}

func QueryProject(projectMetaInfo *dbmodel.ProjectMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var projectCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.ProjectMetaInfoCollectionString)

	result := projectCollection.FindOne(
		context.TODO(),
		bson.D{{"Name", projectMetaInfo.Name}})

	if result.Err() != nil {
		return ReturnWrapper{false, "Cannot find target project!"}
	} else {
		err := result.Decode(projectMetaInfo)
		if err != nil {
			return ReturnWrapper{false, err.Error()}
		} else {
			return ReturnWrapper{true, ""}
		}
	}
}

func CreateUser(userMetaInfo dbmodel.UserMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var userCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.UserMetaInfoCollectionString)

	result := userCollection.FindOne(context.TODO(), bson.D{
		{"Name", userMetaInfo.Name},
	})

	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			_, err := userCollection.InsertOne(context.TODO(), userMetaInfo)
			if err != nil {
				return ReturnWrapper{false, "Create user failed! Error:" + err.Error()}
			}
			return ReturnWrapper{true, "Create user successfully!"}
		}
		return ReturnWrapper{false, "Unknown error!"}
	} else {
		// find one means already exist
		return ReturnWrapper{false, "User already exit!"}
	}

}

func DeleteUser(userMetaInfo dbmodel.UserMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var userCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.UserMetaInfoCollectionString)

	result := userCollection.FindOneAndDelete(context.TODO(), bson.D{
		{"Name", userMetaInfo.Name},
	})

	if result.Err() != nil {
		return ReturnWrapper{false, result.Err().Error()}
	} else {
		return ReturnWrapper{true, "Delete successfully!"}
	}
}

func ModifyUser(userMetaInfo dbmodel.UserMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var userCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.UserMetaInfoCollectionString)

	result := userCollection.FindOneAndReplace(
		context.TODO(),
		bson.D{{"Name", userMetaInfo.Name}},
		userMetaInfo)

	if result.Err() != nil {
		return ReturnWrapper{false, "Update user info failed! Error:" + result.Err().Error()}
	} else {
		return ReturnWrapper{true, "Update user info success!"}
	}

}

func QueryUser(userMetaInfo *dbmodel.UserMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var userCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.UserMetaInfoCollectionString)

	result := userCollection.FindOne(
		context.TODO(),
		bson.D{{"Name", userMetaInfo.Name}})

	if result.Err() != nil {
		return ReturnWrapper{false, "Cannot find target user!"}
	} else {
		err := result.Decode(userMetaInfo)
		if err != nil {
			return ReturnWrapper{false, err.Error()}
		} else {
			return ReturnWrapper{true, ""}
		}
	}
}

func CreatePermissionGroup(permissionGroupMetaInfo dbmodel.PermissionGroupMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var permissionGroupCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.PermissionGroupMetaInfoCollectioString)

	result := permissionGroupCollection.FindOne(context.TODO(), bson.D{
		{"Name", permissionGroupMetaInfo.Name},
	})

	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			_, err := permissionGroupCollection.InsertOne(context.TODO(), permissionGroupMetaInfo)
			if err != nil {
				return ReturnWrapper{false, "Create user failed! Error:" + err.Error()}
			}
			return ReturnWrapper{true, "Create permission group successfully!"}
		}
		return ReturnWrapper{false, "Unknown error!"}
	} else {
		// find one means already exist
		return ReturnWrapper{false, "Permission group already exit!"}
	}

}

func DeletePermissionGroup(permissionGroupMetaInfo dbmodel.PermissionGroupMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var permissionGroupCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.PermissionGroupMetaInfoCollectioString)

	result := permissionGroupCollection.FindOneAndDelete(context.TODO(), bson.D{
		{"Name", permissionGroupMetaInfo.Name},
	})

	if result.Err() != nil {
		return ReturnWrapper{false, result.Err().Error()}
	} else {
		return ReturnWrapper{true, "Delete successfully!"}
	}
}

func ModifyPermissionGroup(permissionGroupMetaInfo dbmodel.PermissionGroupMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var permissionGroupCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.PermissionGroupMetaInfoCollectioString)

	result := permissionGroupCollection.FindOneAndReplace(
		context.TODO(),
		bson.D{{"Name", permissionGroupMetaInfo.Name}},
		permissionGroupMetaInfo)

	if result.Err() != nil {
		return ReturnWrapper{false, "Update permission group failed! Error:" + result.Err().Error()}
	} else {
		return ReturnWrapper{true, "Update permission group success!"}
	}

}

func QueryPermissionGroup(permissionGroupMetaInfo *dbmodel.PermissionGroupMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var permissionGroupCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.PermissionGroupMetaInfoCollectioString)

	result := permissionGroupCollection.FindOne(
		context.TODO(),
		bson.D{{"Name", permissionGroupMetaInfo.Name}})

	if result.Err() != nil {
		return ReturnWrapper{false, "Cannot find target permission group!"}
	} else {
		err := result.Decode(permissionGroupMetaInfo)
		if err != nil {
			return ReturnWrapper{false, err.Error()}
		} else {
			return ReturnWrapper{true, ""}
		}
	}
}

func CreateSwc(swcMetaInfo dbmodel.SwcMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var swcCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.SwcMetaInfoCollectionString)

	result := swcCollection.FindOne(context.TODO(), bson.D{
		{"Name", swcMetaInfo.Name},
	})

	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			_, err := swcCollection.InsertOne(context.TODO(), swcMetaInfo)
			if err != nil {
				return ReturnWrapper{false, "Create swc failed! Error:" + err.Error()}
			}
			return ReturnWrapper{true, "Create swc successfully!"}
		}
		return ReturnWrapper{false, "Unknown error!"}
	} else {
		// find one means already exist
		return ReturnWrapper{false, "Swc already exit!"}
	}

}

func DeleteSwc(swcMetaInfo dbmodel.SwcMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var swcCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.SwcMetaInfoCollectionString)

	result := swcCollection.FindOneAndDelete(context.TODO(), bson.D{
		{"Name", swcMetaInfo.Name},
	})

	if result.Err() != nil {
		return ReturnWrapper{false, result.Err().Error()}
	} else {
		return ReturnWrapper{true, "Delete successfully!"}
	}
}

func ModifySwc(swcMetaInfo dbmodel.SwcMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var swcCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.SwcMetaInfoCollectionString)

	result := swcCollection.FindOneAndReplace(
		context.TODO(),
		bson.D{{"Name", swcMetaInfo.Name}},
		swcMetaInfo)

	if result.Err() != nil {
		return ReturnWrapper{false, "Update swc failed! Error:" + result.Err().Error()}
	} else {
		return ReturnWrapper{true, "Update swc success!"}
	}

}

func QuerySwc(swcMetaInfo *dbmodel.SwcMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var swcCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.SwcMetaInfoCollectionString)

	result := swcCollection.FindOne(
		context.TODO(),
		bson.D{{"Name", swcMetaInfo.Name}})

	if result.Err() != nil {
		return ReturnWrapper{false, "Cannot find target swc!"}
	} else {
		err := result.Decode(swcMetaInfo)
		if err != nil {
			return ReturnWrapper{false, err.Error()}
		} else {
			return ReturnWrapper{true, ""}
		}
	}
}

func CreateDailyStatistics(dailyStatisticsMetaInfo dbmodel.DailyStatisticsMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var dailyStatisticsCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.DailyStatisticsMetaInfoCollectionString)

	result := dailyStatisticsCollection.FindOne(context.TODO(), bson.D{
		{"Name", dailyStatisticsMetaInfo.Name},
	})

	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			_, err := dailyStatisticsCollection.InsertOne(context.TODO(), dailyStatisticsMetaInfo)
			if err != nil {
				return ReturnWrapper{false, "Create daily statistics failed! Error:" + err.Error()}
			}
			return ReturnWrapper{true, "Create daily statistics successfully!"}
		}
		return ReturnWrapper{false, "Unknown error!"}
	} else {
		// find one means already exist
		return ReturnWrapper{false, "Daily statistics already exit!"}
	}

}

func DeleteDailyStatistics(dailyStatisticsMetaInfo dbmodel.DailyStatisticsMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var dailyStatisticsCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.DailyStatisticsMetaInfoCollectionString)

	result := dailyStatisticsCollection.FindOneAndDelete(context.TODO(), bson.D{
		{"Name", dailyStatisticsMetaInfo.Name},
	})

	if result.Err() != nil {
		return ReturnWrapper{false, result.Err().Error()}
	} else {
		return ReturnWrapper{true, "Delete successfully!"}
	}
}

func ModifyDailyStatistics(dailyStatisticsMetaInfo dbmodel.DailyStatisticsMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var dailyStatisticsCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.DailyStatisticsMetaInfoCollectionString)

	result := dailyStatisticsCollection.FindOneAndReplace(
		context.TODO(),
		bson.D{{"Name", dailyStatisticsMetaInfo.Name}},
		dailyStatisticsMetaInfo)

	if result.Err() != nil {
		return ReturnWrapper{false, "Update daily statistics failed! Error:" + result.Err().Error()}
	} else {
		return ReturnWrapper{true, "Update daily statistics success!"}
	}

}

func QueryDailyStatistics(permissionGroupMetaInfo *dbmodel.DailyStatisticsMetaInfoV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	var permissionGroupCollection = databaseInfo.MetaInfoDb.Collection(dbmodel.PermissionGroupMetaInfoCollectioString)

	result := permissionGroupCollection.FindOne(
		context.TODO(),
		bson.D{{"Name", permissionGroupMetaInfo.Name}})

	if result.Err() != nil {
		return ReturnWrapper{false, "Cannot find target daily statistics!"}
	} else {
		err := result.Decode(permissionGroupMetaInfo)
		if err != nil {
			return ReturnWrapper{false, err.Error()}
		} else {
			return ReturnWrapper{true, ""}
		}
	}
}

func CreateSwcData(swcMetaInfo dbmodel.SwcMetaInfoV1, swcData dbmodel.SwcDataV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	collection := databaseInfo.SwcDb.Collection(swcMetaInfo.Name)
	var interfaceSlice dbmodel.SwcDataInterfaceV1 = make(dbmodel.SwcDataInterfaceV1, len(swcData))
	for i, v := range swcData {
		interfaceSlice[i] = v
	}
	result, err := collection.InsertMany(context.TODO(), interfaceSlice)
	if err != nil {
		if result != nil {
			return ReturnWrapper{false,
				"Insert many node failed! Inserted:" + strconv.Itoa(len(result.InsertedIDs)) +
					" , Error:" + strconv.Itoa(len(interfaceSlice)-len(result.InsertedIDs)) +
					" Total:" + strconv.Itoa(len(interfaceSlice))}
		} else {
			return ReturnWrapper{false, "Insert many node failed!"}
		}
	}
	return ReturnWrapper{true, "Create many node Success"}
}

func DeleteSwcData(swcMetaInfo dbmodel.SwcMetaInfoV1, swcData dbmodel.SwcDataV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	collection := databaseInfo.SwcDb.Collection(swcMetaInfo.Name)

	// 创建一个存储所有uuids的切片
	var uuids []interface{}
	for _, v := range swcData {
		uuids = append(uuids, v.Base.Uuid)
	}

	// 创建一个使用$in操作符的过滤器
	filterInterface := bson.D{{Key: "uuid", Value: bson.M{"$in": uuids}}}

	// 使用这个过滤器来删除所有匹配的文档
	result, err := collection.DeleteMany(context.TODO(), filterInterface)

	fmt.Print(err.Error())
	if err != nil {
		if result != nil {
			return ReturnWrapper{false,
				"Delete many node failed! Deleted:" + strconv.Itoa(int(result.DeletedCount)) +
					" , Error:" + strconv.Itoa(len(filterInterface)-int(result.DeletedCount)) +
					" Total:" + strconv.Itoa(len(filterInterface))}
		} else {
			return ReturnWrapper{false, "Delete many node failed!"}
		}

	}
	return ReturnWrapper{true, "Delete many node Success"}
}

func ModifySwcData(swcMetaInfo dbmodel.SwcMetaInfoV1, swcData dbmodel.SwcDataV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	collection := databaseInfo.SwcDb.Collection(swcMetaInfo.Name)
	var interfaceSlice dbmodel.SwcDataInterfaceV1 = make(dbmodel.SwcDataInterfaceV1, len(swcData))
	for i, v := range swcData {
		interfaceSlice[i] = v
	}

	filter := bson.A{}
	for _, v := range swcData {
		filter = append(filter, bson.D{{"uuid", v.Base.Uuid}})
	}
	filterInterface := bson.D{{
		"$or", filter,
	}}

	result, err := collection.UpdateMany(context.TODO(), filterInterface, interfaceSlice)
	if err != nil {
		if result != nil {
			return ReturnWrapper{false,
				"Modify many node failed! Matched:" + strconv.Itoa(int(result.MatchedCount)) +
					" , Error:" + strconv.Itoa(len(interfaceSlice)-int(result.MatchedCount)) +
					" Total:" + strconv.Itoa(len(interfaceSlice))}
		} else {
			return ReturnWrapper{false, "Modify many node failed!"}
		}

	}
	return ReturnWrapper{true, "Modify many node Success"}
}

func QuerySwcData(swcMetaInfo dbmodel.SwcMetaInfoV1, swcData *dbmodel.SwcDataV1, databaseInfo MongoDbDataBaseInfo) ReturnWrapper {
	collection := databaseInfo.SwcDb.Collection(swcMetaInfo.Name)
	var interfaceSlice dbmodel.SwcDataInterfaceV1 = make(dbmodel.SwcDataInterfaceV1, len(*swcData))
	for i, v := range *swcData {
		interfaceSlice[i] = v
	}

	filter := bson.A{}
	for _, v := range *swcData {
		filter = append(filter, bson.D{{"uuid", v.Base.Uuid}})
	}
	filterInterface := bson.D{{
		"$or", filter,
	}}

	result, err := collection.Find(context.TODO(), filterInterface)
	if err != nil {
		return ReturnWrapper{false, "Query many node failed!"}
	}

	err2 := result.Decode(interfaceSlice)
	if err2 != nil {
		return ReturnWrapper{false, "Decode many node query result failed!"}
	}
	return ReturnWrapper{true, "Query many node Success"}
}
