package bll

import (
	"DBMS/Generated/proto/message"
	"DBMS/Generated/proto/request"
	"DBMS/Generated/proto/response"
	"DBMS/Generated/proto/service"
	"DBMS/dal"
	"DBMS/dbmodel"
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DBMSServerController struct {
	service.UnimplementedDBMSServer
}

func (D DBMSServerController) CreateUser(ctx context.Context, request *request.CreateUserRequest) (*response.CreateUserResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)
	userMetaInfo.Base.Id = primitive.NewObjectID()
	userMetaInfo.Base.Uuid = uuid.NewString()
	userMetaInfo.Base.ApiVersion = "V1"

	userMetaInfo.Name = request.UserInfo.Name
	userMetaInfo.Password = request.UserInfo.Password
	userMetaInfo.Description = request.UserInfo.Description
	userMetaInfo.UserPermissionGroup = "default"
	userMetaInfo.CreateTime = time.Now()
	userMetaInfo.HeadPhotoBinData = bytes.NewBuffer(request.UserInfo.HeadPhotoBinData)

	result := dal.CreateUser(*userMetaInfo, dal.GetDbInstance())
	if result.Status == true {
		fmt.Println("User " + request.UserInfo.Name + " Created")
		return &response.CreateUserResponse{
			Status:   true,
			Message:  result.Message,
			UserInfo: UserMetaInfoV1DbmodelToProtobuf(userMetaInfo),
		}, nil
	} else {
		return &response.CreateUserResponse{
			Status:   false,
			Message:  result.Message,
			UserInfo: UserMetaInfoV1DbmodelToProtobuf(userMetaInfo),
		}, nil
	}
}

func (D DBMSServerController) DeleteUser(ctx context.Context, request *request.DeleteUserRequest) (*response.DeleteUserResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)

	result := dal.DeleteUser(*userMetaInfo, dal.GetDbInstance())
	if result.Status == true {
		fmt.Println("User " + request.UserInfo.Name + " Deleted")
		return &response.DeleteUserResponse{
			Status:   true,
			Message:  result.Message,
			UserInfo: UserMetaInfoV1DbmodelToProtobuf(userMetaInfo),
		}, nil
	} else {
		return &response.DeleteUserResponse{
			Status:   false,
			Message:  result.Message,
			UserInfo: UserMetaInfoV1DbmodelToProtobuf(userMetaInfo),
		}, nil
	}
}

func (D DBMSServerController) UpdateUser(ctx context.Context, request *request.UpdateUserRequest) (*response.UpdateUserResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)

	result := dal.ModifyUser(*userMetaInfo, dal.GetDbInstance())
	if result.Status == true {
		fmt.Println("User " + request.UserInfo.Name + " Updated")
		return &response.UpdateUserResponse{
			Status:   true,
			Message:  result.Message,
			UserInfo: UserMetaInfoV1DbmodelToProtobuf(userMetaInfo),
		}, nil
	} else {
		return &response.UpdateUserResponse{
			Status:   false,
			Message:  result.Message,
			UserInfo: UserMetaInfoV1DbmodelToProtobuf(userMetaInfo),
		}, nil
	}
}

func (D DBMSServerController) GetUser(ctx context.Context, request *request.GetUserRequest) (*response.GetUserResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)

	result := dal.QueryUser(userMetaInfo, dal.GetDbInstance())
	if result.Status == true {
		fmt.Println("User " + request.UserInfo.Name + " Get")
		return &response.GetUserResponse{
			Status:   true,
			Message:  result.Message,
			UserInfo: UserMetaInfoV1DbmodelToProtobuf(userMetaInfo),
		}, nil
	} else {
		return &response.GetUserResponse{
			Status:   false,
			Message:  result.Message,
			UserInfo: UserMetaInfoV1DbmodelToProtobuf(userMetaInfo),
		}, nil
	}
}

func (D DBMSServerController) GetAllUser(ctx context.Context, request *request.GetAllUserRequest) (*response.GetAllUserResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)

	fmt.Println("TODO: Add permission check" + userMetaInfo.Name)

	var userMetaInfoList []dbmodel.UserMetaInfoV1
	var protoMessage []*message.UserMetaInfoV1

	for _, userMetaInfo := range userMetaInfoList {
		protoMessage = append(protoMessage, UserMetaInfoV1DbmodelToProtobuf(&userMetaInfo))
	}

	result := dal.QueryAllUser(&userMetaInfoList, dal.GetDbInstance())
	if result.Status == true {
		fmt.Println("User " + request.UserInfo.Name + " Try Get AllUser")
		return &response.GetAllUserResponse{
			Status:   true,
			Message:  result.Message,
			UserInfo: protoMessage,
		}, nil
	} else {
		return &response.GetAllUserResponse{
			Status:   false,
			Message:  result.Message,
			UserInfo: protoMessage,
		}, nil
	}
}

func (D DBMSServerController) UserLogin(ctx context.Context, request *request.UserLoginRequest) (*response.UserLoginResponse, error) {
	var userMetaInfo dbmodel.UserMetaInfoV1
	userMetaInfo.Name = request.UserName

	result := dal.QueryUser(&userMetaInfo, dal.GetDbInstance())
	if result.Status == true {
		if userMetaInfo.Password == request.Password {
			fmt.Println("User " + request.UserName + " Login")
			return &response.UserLoginResponse{
				Status:   true,
				Message:  result.Message,
				UserInfo: UserMetaInfoV1DbmodelToProtobuf(&userMetaInfo),
			}, nil
		} else {
			userMetaInfo.Password = ""
			return &response.UserLoginResponse{
				Status:   false,
				Message:  result.Message,
				UserInfo: UserMetaInfoV1DbmodelToProtobuf(&userMetaInfo),
			}, nil
		}
	} else {
		userMetaInfo.Password = ""
		return &response.UserLoginResponse{
			Status:   false,
			Message:  result.Message,
			UserInfo: UserMetaInfoV1DbmodelToProtobuf(&userMetaInfo),
		}, nil
	}
}

func (D DBMSServerController) UserLogout(ctx context.Context, request *request.UserLogoutRequest) (*response.UserLogoutResponse, error) {
	var userMetaInfo dbmodel.UserMetaInfoV1
	userMetaInfo.Name = request.UserInfo.Name

	result := dal.QueryUser(&userMetaInfo, dal.GetDbInstance())
	if result.Status == true {
		if userMetaInfo.Password == request.UserInfo.Password {
			fmt.Println("User " + request.UserInfo.Name + " Logout")
			return &response.UserLogoutResponse{
				Status:  true,
				Message: result.Message,
			}, nil
		} else {
			userMetaInfo.Password = ""
			return &response.UserLogoutResponse{
				Status:  false,
				Message: result.Message,
			}, nil
		}
	} else {
		userMetaInfo.Password = ""
		return &response.UserLogoutResponse{
			Status:  false,
			Message: result.Message,
		}, nil
	}
}

func (D DBMSServerController) UserOnlineHeartBeatNotifications(ctx context.Context, notification *request.UserOnlineHeartBeatNotification) (*response.UserOnlineHeartBeatResponse, error) {
	var userMetaInfo dbmodel.UserMetaInfoV1
	userMetaInfo.Name = notification.UserInfo.Name

	result := dal.QueryUser(&userMetaInfo, dal.GetDbInstance())
	if result.Status == true {
		if userMetaInfo.Password == notification.UserInfo.Password {
			fmt.Println("User " + notification.UserInfo.Name + " OnlineHeartBeatNotifications")
			return &response.UserOnlineHeartBeatResponse{
				Status:  true,
				Message: result.Message,
			}, nil
		} else {
			userMetaInfo.Password = ""
			return &response.UserOnlineHeartBeatResponse{
				Status:  false,
				Message: result.Message,
			}, nil
		}
	} else {
		userMetaInfo.Password = ""
		return &response.UserOnlineHeartBeatResponse{
			Status:  false,
			Message: result.Message,
		}, nil
	}
}

func (D DBMSServerController) GetUserPermissionGroup(ctx context.Context, request *request.GetUserPermissionGroupRequest) (*response.GetUserPermissionGroupResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) GetPermissionGroup(ctx context.Context, request *request.GetPermissionGroupRequest) (*response.GetPermissionGroupResponse, error) {
	//TODO implement me
	panic("implement me")
}

func GetAllPermissionGroup(context.Context, *request.GetAllPermissionGroupRequest) (*response.GetAllPermissionGroupResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) ChangeUserPermissionGroup(ctx context.Context, request *request.ChangeUserPermissionGroupRequest) (*response.ChangeUserPermissionGroupResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) CreateProject(ctx context.Context, request *request.CreateProjectRequest) (*response.CreateProjectResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) DeleteProject(ctx context.Context, request *request.DeleteProjectRequest) (*response.DeleteProjectResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) UpdateProject(ctx context.Context, request *request.UpdateProjectRequest) (*response.UpdateProjectResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) GetProject(ctx context.Context, request *request.GetProjectRequest) (*response.GetProjectResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) GetAllProject(ctx context.Context, request *request.GetAllProjectRequest) (*response.GetAllProjectResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) CreateSwc(ctx context.Context, request *request.CreateSwcRequest) (*response.CreateSwcResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) DeleteSwc(ctx context.Context, request *request.DeleteSwcRequest) (*response.DeleteSwcResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) UpdateSwc(ctx context.Context, request *request.UpdateSwcRequest) (*response.UpdateSwcResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) GetSwcMetaInfo(ctx context.Context, request *request.GetSwcMetaInfoRequest) (*response.GetSwcMetaInfoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) GetAllSwcMetaInfo(ctx context.Context, request *request.GetAllSwcMetaInfoRequest) (*response.GetAllSwcMetaInfoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) CreateSwcNodeData(ctx context.Context, request *request.CreateSwcNodeDataRequest) (*response.CreateSwcNodeDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) DeleteSwcNodeData(ctx context.Context, request *request.DeleteSwcNodeDataRequest) (*response.DeleteSwcNodeDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) UpdateSwcNodeData(ctx context.Context, request *request.UpdateSwcNodeDataRequest) (*response.UpdateSwcNodeDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) GetSwcNodeData(ctx context.Context, request *request.GetSwcNodeDataRequest) (*response.GetSwcNodeDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) GetSwcFullNodeData(ctx context.Context, request *request.GetSwcFullNodeDataRequest) (*response.GetSwcFullNodeDataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) GetSwcNodeDataListByTimeAndUser(ctx context.Context, request *request.GetSwcNodeDataListByTimeAndUserRequest) (*response.GetSwcNodeDataListByTimeAndUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) BackupFullDatabase(ctx context.Context, request *request.BackupFullDatabaseRequest) (*response.BackupFullDatabaseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) CreateDailyStatistics(ctx context.Context, request *request.CreateDailyStatisticsRequest) (*response.CreateDailyStatisticsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) DeleteDailyStatistics(ctx context.Context, request *request.DeleteDailyStatisticsRequest) (*response.DeleteDailyStatisticsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) UpdateDailyStatistics(ctx context.Context, request *request.UpdateDailyStatisticsRequest) (*response.UpdateDailyStatisticsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) GetDailyStatistics(ctx context.Context, request *request.GetDailyStatisticsRequest) (*response.GetDailyStatisticsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) GetDailyStatisticsList(ctx context.Context, request *request.GetAllDailyStatisticsRequest) (*response.GetAllDailyStatisticsResponse, error) {
	//TODO implement me
	panic("implement me")
}
