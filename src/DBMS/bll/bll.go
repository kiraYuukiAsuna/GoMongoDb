package bll

import (
	"DBMS/Generated/proto/request"
	"DBMS/Generated/proto/response"
	"DBMS/Generated/proto/service"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type CreateUserController struct {
	service.UnimplementedDBMSServer
}

func (c CreateUserController) mustEmbedUnimplementedDBMSServer() {
	return
}

func (c CreateUserController) CreateUser(ctx context.Context, request *request.CreateUserRequest) (*response.CreateUserResponse, error) {
	fmt.Println(request.UserInfo.Name)
	return &response.CreateUserResponse{
		Status:  true,
		Message: "Hello " + request.UserInfo.Name,
	}, nil
}

func NewGrpcServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()

	var createUserController CreateUserController
	service.RegisterDBMSServer(s, createUserController)

	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}
