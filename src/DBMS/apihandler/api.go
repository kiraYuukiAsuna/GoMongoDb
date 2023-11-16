package apihandler

import (
	"DBMS/dal"
	"DBMS/dbmodel"
	context2 "context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

//func SayHelloAgain(context *gin.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
//	return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
//}

func InitializeNewDataBaseHandler(context *gin.Context) {
	dal.InitializeNewDataBase(dal.DefaultMetaInfoDataBaseName, dal.DefaultSwcDataBaseName)
}

func CreateUserHandler(context *gin.Context) {

	var userInfo dbmodel.UserMetaInfoV1
	userInfo.Base.Id = primitive.NewObjectID()
	userInfo.Name = "Hanasaka"
	userInfo.Description = "Test user"

	_, err := dal.GetDbInstance().MetaInfoDb.Collection(dbmodel.UserMetaInfoCollectionString).InsertOne(context2.TODO(), userInfo)
	if err != nil {
		return
	}
}

func OneTimeInitializeHandler(context *gin.Context) {
	var createInfo dal.MongoDbConnectionCreateInfo
	createInfo.Host = "127.0.0.1"
	createInfo.Port = 27017
	createInfo.User = ""
	createInfo.Password = ""
	connectionInfo := dal.ConnectToMongoDb(createInfo)

	if connectionInfo.Err != nil {
		log.Fatal(connectionInfo.Err)
	}

	databaseInfo := dal.ConnectToDataBase(connectionInfo, dal.DefaultMetaInfoDataBaseName, dal.DefaultSwcDataBaseName)

	dal.SetDbInstance(databaseInfo)
}
