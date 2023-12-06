package UnitTest

import (
	"DBMS/config"
	"DBMS/dal"
	"DBMS/dbmodel"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func InitializeDb() {
	// create db
	dal.InitializeNewDataBaseIfNotExist(dal.DefaultMetaInfoDataBaseName, dal.DefaultSwcDataBaseName)

	// init db
	var createInfo dal.MongoDbConnectionCreateInfo
	createInfo.Host = config.AppConfig.MongodbIP
	createInfo.Port = config.AppConfig.MongodbPort
	createInfo.User = ""
	createInfo.Password = ""
	connectionInfo := dal.ConnectToMongoDb(createInfo)

	if connectionInfo.Err != nil {
		log.Fatal(connectionInfo.Err)
	}

	databaseInstance := dal.ConnectToDataBase(connectionInfo, dal.DefaultMetaInfoDataBaseName, dal.DefaultSwcDataBaseName)

	dal.SetDbInstance(databaseInstance)

	var defaultAdminSystemUser dbmodel.UserMetaInfoV1
	defaultAdminSystemUser.Name = dal.DefaultAdminSystemUserName
	defaultAdminSystemUser.Password = dal.DefaultAdminSystemUserPassword

	dal.CreateUser(defaultAdminSystemUser, dal.GetDbInstance())
}

func TestUserInfo() {
	fmt.Println("TestUser:")
	var info1 dbmodel.UserMetaInfoV1
	info1.Base.Id = primitive.NewObjectID()
	info1.Name = "Hanasaka"
	info1.Description = "Test user"

	var info2 dbmodel.UserMetaInfoV1
	info2.Base.Id = primitive.NewObjectID()
	info2.Name = "Hanasaka2"
	info2.Description = "Test user"

	failedNumber := 0

	if dal.CreateUser(info1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 1 Failed")
	}

	if dal.CreateUser(info1, dal.GetDbInstance()).Status == true {
		failedNumber++
		fmt.Println("Test 2 Failed")
	}

	if dal.CreateUser(info2, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 3 Failed")
	}

	if dal.DeleteUser(info2, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 4 Failed")
	}

	if dal.DeleteUser(info2, dal.GetDbInstance()).Status == true {
		failedNumber++
		fmt.Println("Test 5 Failed")
	}

	info1.Description = "Test Modify UserInfo1"
	if dal.ModifyUser(info1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 6 Failed")
	}

	info1.Description = "None"
	if dal.QueryUser(&info1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 7 Failed")
	}

	if info1.Description != "Test Modify UserInfo1" {
		failedNumber++
		fmt.Println("Test 8 Failed")
	}

	var userMetaInfoList []dbmodel.UserMetaInfoV1
	if dal.QueryAllUser(&userMetaInfoList, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 9 Failed")
	}

	fmt.Printf("failed number: %d \n", failedNumber)
}

func TestProjectInfo() {
	fmt.Println("TestProject:")
	var info1 dbmodel.ProjectMetaInfoV1
	info1.Base.Id = primitive.NewObjectID()
	info1.Name = "Hanasaka"
	info1.Description = "Test user"

	var info2 dbmodel.ProjectMetaInfoV1
	info2.Base.Id = primitive.NewObjectID()
	info2.Name = "Hanasaka2"
	info2.Description = "Test user"

	failedNumber := 0

	if dal.CreateProject(info1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 1 Failed")
	}

	if dal.CreateProject(info1, dal.GetDbInstance()).Status == true {
		failedNumber++
		fmt.Println("Test 2 Failed")
	}

	if dal.CreateProject(info2, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 3 Failed")
	}

	if dal.DeleteProject(info2, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 4 Failed")
	}

	if dal.DeleteProject(info2, dal.GetDbInstance()).Status == true {
		failedNumber++
		fmt.Println("Test 5 Failed")
	}

	info1.Description = "Test Modify UserInfo1"
	if dal.ModifyProject(info1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 6 Failed")
	}

	info1.Description = "None"
	if dal.QueryProject(&info1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 7 Failed")
	}

	if info1.Description != "Test Modify UserInfo1" {
		failedNumber++
		fmt.Println("Test 8 Failed")
	}

	var projectMetaInfoList []dbmodel.ProjectMetaInfoV1
	if dal.QueryAllProject(&projectMetaInfoList, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 9 Failed")
	}

	fmt.Printf("Failed number: %d \n", failedNumber)
}

func TestSwcInfo() {
	fmt.Println("TestSwc:")
	var info1 dbmodel.SwcMetaInfoV1
	info1.Base.Id = primitive.NewObjectID()
	info1.Name = "Hanasaka"
	info1.Description = "Test user"

	var info2 dbmodel.SwcMetaInfoV1
	info2.Base.Id = primitive.NewObjectID()
	info2.Name = "Hanasaka2"
	info2.Description = "Test user"

	failedNumber := 0

	if dal.CreateSwc(info1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 1 Failed")
	}

	if dal.CreateSwc(info1, dal.GetDbInstance()).Status == true {
		failedNumber++
		fmt.Println("Test 2 Failed")
	}

	if dal.CreateSwc(info2, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 3 Failed")
	}

	if dal.DeleteSwc(info2, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 4 Failed")
	}

	if dal.DeleteSwc(info2, dal.GetDbInstance()).Status == true {
		failedNumber++
		fmt.Println("Test 5 Failed")
	}

	info1.Description = "Test Modify UserInfo1"
	if dal.ModifySwc(info1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 6 Failed")
	}

	info1.Description = "None"
	if dal.QuerySwc(&info1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 7 Failed")
	}

	if info1.Description != "Test Modify UserInfo1" {
		failedNumber++
		fmt.Println("Test 8 Failed")
	}

	var swcMetaInfoList []dbmodel.SwcMetaInfoV1
	if dal.QueryAllSwc(&swcMetaInfoList, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 9 Failed")
	}

	fmt.Printf("Failed number: %d \n", failedNumber)
}

func TestPermissionGroupInfo() {
	fmt.Println("TestPermissionGroup:")
	var info1 dbmodel.PermissionGroupMetaInfoV1
	info1.Base.Id = primitive.NewObjectID()
	info1.Name = "Hanasaka"
	info1.Description = "Test user"

	var info2 dbmodel.PermissionGroupMetaInfoV1
	info2.Base.Id = primitive.NewObjectID()
	info2.Name = "Hanasaka2"
	info2.Description = "Test user"

	failedNumber := 0

	if dal.CreatePermissionGroup(info1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 1 Failed")
	}

	if dal.CreatePermissionGroup(info1, dal.GetDbInstance()).Status == true {
		failedNumber++
		fmt.Println("Test 2 Failed")
	}

	if dal.CreatePermissionGroup(info2, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 3 Failed")
	}

	if dal.DeletePermissionGroup(info2, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 4 Failed")
	}

	if dal.DeletePermissionGroup(info2, dal.GetDbInstance()).Status == true {
		failedNumber++
		fmt.Println("Test 5 Failed")
	}

	info1.Description = "Test Modify UserInfo1"
	if dal.ModifyPermissionGroup(info1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 6 Failed")
	}

	info1.Description = "None"
	if dal.QueryPermissionGroup(&info1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 7 Failed")
	}

	if info1.Description != "Test Modify UserInfo1" {
		failedNumber++
		fmt.Println("Test 8 Failed")
	}

	var permissionGroupList []dbmodel.PermissionGroupMetaInfoV1
	if dal.QueryAllPermissionGroup(&permissionGroupList, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 9 Failed")
	}

	fmt.Printf("Failed number: %d \n", failedNumber)
}

func TestDailyStatisticsInfo() {
	fmt.Println("TestDailyStatistics:")
	var info1 dbmodel.DailyStatisticsMetaInfoV1
	info1.Base.Id = primitive.NewObjectID()
	info1.Name = "Hanasaka"
	info1.Description = "Test user"

	var info2 dbmodel.DailyStatisticsMetaInfoV1
	info2.Base.Id = primitive.NewObjectID()
	info2.Name = "Hanasaka2"
	info2.Description = "Test user"

	failedNumber := 0

	if dal.CreateDailyStatistics(info1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 1 Failed")
	}

	if dal.CreateDailyStatistics(info1, dal.GetDbInstance()).Status == true {
		failedNumber++
		fmt.Println("Test 2 Failed")
	}

	if dal.CreateDailyStatistics(info2, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 3 Failed")
	}

	if dal.DeleteDailyStatistics(info2, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 4 Failed")
	}

	if dal.DeleteDailyStatistics(info2, dal.GetDbInstance()).Status == true {
		failedNumber++
		fmt.Println("Test 5 Failed")
	}

	info1.Description = "Test Modify UserInfo1"
	if dal.ModifyDailyStatistics(info1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 6 Failed")
	}

	info1.Description = "None"
	if dal.QueryDailyStatistics(&info1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 7 Failed")
	}

	if info1.Description != "Test Modify UserInfo1" {
		failedNumber++
		fmt.Println("Test 8 Failed")
	}

	var dailyStatisticsGroupList []dbmodel.DailyStatisticsMetaInfoV1
	if dal.QueryAllDailyStatistics(&dailyStatisticsGroupList, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 9 Failed")
	}

	fmt.Printf("Failed number: %d \n", failedNumber)
}

func TestSwcData() {
	fmt.Println("TestSwcData:")
	var info1 dbmodel.SwcNodeDataV1
	info1.Base.Id = primitive.NewObjectID()
	info1.Base.Uuid = uuid.NewString()

	var info2 dbmodel.SwcNodeDataV1
	info2.Base.Id = primitive.NewObjectID()
	info2.Base.Uuid = uuid.NewString()

	failedNumber := 0

	var swcMetaInfo dbmodel.SwcMetaInfoV1
	swcMetaInfo.Base.Id = primitive.NewObjectID()
	swcMetaInfo.Name = "Hanasaka"

	var swcData1 dbmodel.SwcDataV1
	swcData1 = append(swcData1, info1)

	var swcData2 dbmodel.SwcDataV1
	swcData2 = append(swcData2, info2)

	if dal.CreateSwcData(swcMetaInfo, swcData1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 1 Failed")
	}

	if dal.CreateSwcData(swcMetaInfo, swcData1, dal.GetDbInstance()).Status == true {
		failedNumber++
		fmt.Println("Test 2 Failed")
	}

	if dal.CreateSwcData(swcMetaInfo, swcData2, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 3 Failed")
	}

	if dal.DeleteSwcData(swcMetaInfo, swcData2, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 4 Failed")
	}

	if dal.DeleteSwcData(swcMetaInfo, swcData2, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 5 Failed")
	}
	info1.Creator = "Test Modify"
	if dal.ModifySwcData(swcMetaInfo, info1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 6 Failed")
	}
	info1.Creator = ""
	if dal.QuerySwcData(swcMetaInfo, &swcData1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 7 Failed")
	}

	if swcData1[0].Creator != "Test Modify" {
		failedNumber++
		fmt.Println("Test 8 Failed")
	}

	if dal.QuerySwcDataByUserAndTime(swcMetaInfo, "", time.Time{}, time.Time{}, &swcData1, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 9 Failed")
	}

	var swcDataList dbmodel.SwcDataV1
	if dal.QueryAllSwcData(swcMetaInfo, &swcDataList, dal.GetDbInstance()).Status == false {
		failedNumber++
		fmt.Println("Test 9 Failed")
	}

	fmt.Printf("Failed number: %d \n", failedNumber)
}
