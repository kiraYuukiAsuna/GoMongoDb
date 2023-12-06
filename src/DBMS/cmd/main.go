package main

import (
	"DBMS/UnitTest"
	"DBMS/apihandler"
	"DBMS/bll"
	"DBMS/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	config.SetDafaultAppConfig()
	config.ReadConfig()
	bll.Initialize()
	bll.CronAutoSaveDailyStatistics()
	bll.CronHeartBeatValidationAndRefresh()
	bll.NewGrpcServer()
	return

	UnitTest.InitializeDb()
	UnitTest.TestUserInfo()
	UnitTest.TestProjectInfo()
	UnitTest.TestPermissionGroupInfo()
	UnitTest.TestSwcInfo()
	UnitTest.TestDailyStatisticsInfo()
	UnitTest.TestSwcData()

	return

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/InitializeNewDataBaseIfNotExist", apihandler.InitializeNewDataBaseIfNotExistHandler)
	r.GET("/CreateUser", apihandler.CreateUserHandler)

	err := r.Run("0.0.0.0:8088")
	if err != nil {
		return
	}
}
