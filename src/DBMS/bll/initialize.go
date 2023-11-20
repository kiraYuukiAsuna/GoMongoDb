package bll

import (
	"DBMS/Generated/proto/service"
	"DBMS/dal"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Initialize() {
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

func NewGrpcServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()

	var instanceDBMSServerController DBMSServerController
	service.RegisterDBMSServer(s, instanceDBMSServerController)

	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
