package main

import (
	"DBMS/UnitTest"
	"DBMS/apihandler"
	"DBMS/bll"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	bll.Initialize()
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

	r.GET("/InitializeNewDataBase", apihandler.InitializeNewDataBaseHandler)
	r.GET("/CreateUser", apihandler.CreateUserHandler)
	r.GET("/OneTimeInitialize", apihandler.OneTimeInitializeHandler)

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
