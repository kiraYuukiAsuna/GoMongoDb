package UnitTest

import (
	"DBMS/dal"
	"DBMS/dbmodel"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func InitializeDb() {
	// create db
	dal.InitializeNewDataBase(dal.DefaultMetaInfoDataBaseName, dal.DefaultSwcDataBaseName)

	// init db
	var createInfo dal.MongoDbConnectionCreateInfo
	createInfo.Host = "127.0.0.1"
	createInfo.Port = 27017
	createInfo.User = ""
	createInfo.Password = ""
	connectionInfo := dal.ConnectToMongoDb(createInfo)

	if connectionInfo.Err != nil {
		log.Fatal(connectionInfo.Err)
	}

	databaseInstance := dal.ConnectToDataBase(connectionInfo, dal.DefaultMetaInfoDataBaseName, dal.DefaultSwcDataBaseName)

	dal.SetDbInstance(databaseInstance)
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

	fmt.Printf("Failed number: %d \n", failedNumber)
}
