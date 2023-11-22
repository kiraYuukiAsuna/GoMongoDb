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

var OnlineUserList dbmodel.UserMetaInfoV1

var DailyStatisticsData dbmodel.DailyStatisticsMetaInfoV1

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

	result := dal.QueryAllUser(&userMetaInfoList, dal.GetDbInstance())
	if result.Status == true {
		fmt.Println("User " + request.UserInfo.Name + " Try Get AllUser")
		for _, userMetaInfo := range userMetaInfoList {
			protoMessage = append(protoMessage, UserMetaInfoV1DbmodelToProtobuf(&userMetaInfo))
		}
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
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)

	var permissionGroupMetaInfo dbmodel.PermissionGroupMetaInfoV1

	result := dal.QueryUser(userMetaInfo, dal.GetDbInstance())
	if result.Status == true {
		permissionGroupMetaInfo.Name = userMetaInfo.UserPermissionGroup
		result = dal.QueryPermissionGroup(&permissionGroupMetaInfo, dal.GetDbInstance())
		if result.Status == true {
			fmt.Println("User " + request.UserInfo.Name + " GetUserPermissionGroup")
			return &response.GetUserPermissionGroupResponse{
				Status:          true,
				Message:         result.Message,
				PermissionGroup: PermissionGroupMetaInfoV1DbmodelToProtobuf(&permissionGroupMetaInfo),
			}, nil
		} else {
			return &response.GetUserPermissionGroupResponse{
				Status:          false,
				Message:         result.Message,
				PermissionGroup: PermissionGroupMetaInfoV1DbmodelToProtobuf(&permissionGroupMetaInfo),
			}, nil
		}

	} else {
		return &response.GetUserPermissionGroupResponse{
			Status:          false,
			Message:         result.Message,
			PermissionGroup: PermissionGroupMetaInfoV1DbmodelToProtobuf(&permissionGroupMetaInfo),
		}, nil
	}
}

func (D DBMSServerController) GetPermissionGroup(ctx context.Context, request *request.GetPermissionGroupRequest) (*response.GetPermissionGroupResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)

	permissionGroupMetaInfo := PermissionGroupMetaInfoV1ProtobufToDbmodel(request.PermissionGroup)

	result := dal.QueryUser(userMetaInfo, dal.GetDbInstance())
	if result.Status == true {
		result = dal.QueryPermissionGroup(permissionGroupMetaInfo, dal.GetDbInstance())
		if result.Status == true {
			fmt.Println("User " + request.UserInfo.Name + " GetPermissionGroup")
			return &response.GetPermissionGroupResponse{
				Status:          true,
				Message:         result.Message,
				PermissionGroup: PermissionGroupMetaInfoV1DbmodelToProtobuf(permissionGroupMetaInfo),
			}, nil
		} else {
			return &response.GetPermissionGroupResponse{
				Status:          false,
				Message:         result.Message,
				PermissionGroup: PermissionGroupMetaInfoV1DbmodelToProtobuf(permissionGroupMetaInfo),
			}, nil
		}

	} else {
		return &response.GetPermissionGroupResponse{
			Status:          false,
			Message:         result.Message,
			PermissionGroup: PermissionGroupMetaInfoV1DbmodelToProtobuf(permissionGroupMetaInfo),
		}, nil
	}
}

func GetAllPermissionGroup(ctx context.Context, request *request.GetAllPermissionGroupRequest) (*response.GetAllPermissionGroupResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)

	var permissionGroupList []dbmodel.PermissionGroupMetaInfoV1
	var protoMessage []*message.PermissionGroupMetaInfoV1

	result := dal.QueryUser(userMetaInfo, dal.GetDbInstance())
	if result.Status == true {
		result = dal.QueryAllPermissionGroup(&permissionGroupList, dal.GetDbInstance())
		if result.Status == true {
			fmt.Println("User " + request.UserInfo.Name + " GetAllPermissionGroup")
			for _, permissionGroupMetaInfo := range permissionGroupList {
				protoMessage = append(protoMessage, PermissionGroupMetaInfoV1DbmodelToProtobuf(&permissionGroupMetaInfo))
			}
			return &response.GetAllPermissionGroupResponse{
				Status:              true,
				Message:             result.Message,
				PermissionGroupList: protoMessage,
			}, nil
		} else {
			return &response.GetAllPermissionGroupResponse{
				Status:              false,
				Message:             result.Message,
				PermissionGroupList: protoMessage,
			}, nil
		}

	} else {
		return &response.GetAllPermissionGroupResponse{
			Status:              false,
			Message:             result.Message,
			PermissionGroupList: protoMessage,
		}, nil
	}
}

func (D DBMSServerController) ChangeUserPermissionGroup(ctx context.Context, request *request.ChangeUserPermissionGroupRequest) (*response.ChangeUserPermissionGroupResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)

	var permissionGroupMetaInfo dbmodel.PermissionGroupMetaInfoV1

	result := dal.QueryUser(userMetaInfo, dal.GetDbInstance())
	if result.Status == true {
		permissionGroupMetaInfo.Name = userMetaInfo.UserPermissionGroup
		result = dal.QueryPermissionGroup(&permissionGroupMetaInfo, dal.GetDbInstance())
		if result.Status == true {
			fmt.Println("User " + request.UserInfo.Name + " GetUserPermissionGroup")
			return &response.ChangeUserPermissionGroupResponse{
				Status:          true,
				Message:         result.Message,
				PermissionGroup: PermissionGroupMetaInfoV1DbmodelToProtobuf(&permissionGroupMetaInfo),
			}, nil
		} else {
			return &response.ChangeUserPermissionGroupResponse{
				Status:          false,
				Message:         result.Message,
				PermissionGroup: PermissionGroupMetaInfoV1DbmodelToProtobuf(&permissionGroupMetaInfo),
			}, nil
		}

	} else {
		return &response.ChangeUserPermissionGroupResponse{
			Status:          false,
			Message:         result.Message,
			PermissionGroup: PermissionGroupMetaInfoV1DbmodelToProtobuf(&permissionGroupMetaInfo),
		}, nil
	}
}

func (D DBMSServerController) CreateProject(ctx context.Context, request *request.CreateProjectRequest) (*response.CreateProjectResponse, error) {
	_ = UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)

	projectMetaInfo := ProjectMetaInfoV1ProtobufToDbmodel(request.ProjectInfo)
	projectMetaInfo.Base.Id = primitive.NewObjectID()
	projectMetaInfo.Base.Uuid = uuid.NewString()
	projectMetaInfo.Base.ApiVersion = "V1"

	projectMetaInfo.Name = request.ProjectInfo.Name
	projectMetaInfo.Description = request.ProjectInfo.Description
	projectMetaInfo.Creator = request.ProjectInfo.Creator

	projectMetaInfo.CreateTime = time.Now()
	projectMetaInfo.LastModifiedTime = time.Now()
	projectMetaInfo.SwcList = []string{}

	projectMetaInfo.WorkMode = request.ProjectInfo.WorkMode

	if request.ProjectInfo.UserPermissionOverride != nil {
		for _, protoPermissionOverride := range request.ProjectInfo.UserPermissionOverride {
			var dbmodelPermissionOverride dbmodel.ProjectPermissionMetaInfoV1
			dbmodelPermissionOverride.ReadPerimissionQuery = protoPermissionOverride.ReadPerimissionQuery
			dbmodelPermissionOverride.WritePermissionAddData = protoPermissionOverride.WritePermissionAddData
			dbmodelPermissionOverride.WritePermissionModifyData = protoPermissionOverride.WritePermissionModifyData
			dbmodelPermissionOverride.WritePermissionDeleteData = protoPermissionOverride.WritePermissionDeleteData

			projectMetaInfo.UserPermissionOverride = append(projectMetaInfo.UserPermissionOverride, dbmodelPermissionOverride)
		}
	}

	result := dal.CreateProject(*projectMetaInfo, dal.GetDbInstance())
	if result.Status == true {
		fmt.Println("Project " + request.ProjectInfo.Name + " Created")
		return &response.CreateProjectResponse{
			Status:      true,
			Message:     result.Message,
			ProjectInfo: ProjectMetaInfoV1DbmodelToProtobuf(projectMetaInfo),
		}, nil
	} else {
		return &response.CreateProjectResponse{
			Status:      false,
			Message:     result.Message,
			ProjectInfo: ProjectMetaInfoV1DbmodelToProtobuf(projectMetaInfo),
		}, nil
	}
}

func (D DBMSServerController) DeleteProject(ctx context.Context, request *request.DeleteProjectRequest) (*response.DeleteProjectResponse, error) {
	_ = UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)
	projectMetaInfo := ProjectMetaInfoV1ProtobufToDbmodel(request.ProjectInfo)

	result := dal.DeleteProject(*projectMetaInfo, dal.GetDbInstance())
	if result.Status == true {
		fmt.Println("Project " + request.ProjectInfo.Name + " Deleted")
		return &response.DeleteProjectResponse{
			Status:      true,
			Message:     result.Message,
			ProjectInfo: ProjectMetaInfoV1DbmodelToProtobuf(projectMetaInfo),
		}, nil
	} else {
		return &response.DeleteProjectResponse{
			Status:      false,
			Message:     result.Message,
			ProjectInfo: ProjectMetaInfoV1DbmodelToProtobuf(projectMetaInfo),
		}, nil
	}
}

func (D DBMSServerController) UpdateProject(ctx context.Context, request *request.UpdateProjectRequest) (*response.UpdateProjectResponse, error) {
	_ = UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)
	projectMetaInfo := ProjectMetaInfoV1ProtobufToDbmodel(request.ProjectInfo)

	result := dal.ModifyProject(*projectMetaInfo, dal.GetDbInstance())
	if result.Status == true {
		fmt.Println("Project " + request.UserInfo.Name + " Updated")
		return &response.UpdateProjectResponse{
			Status:      true,
			Message:     result.Message,
			ProjectInfo: ProjectMetaInfoV1DbmodelToProtobuf(projectMetaInfo),
		}, nil
	} else {
		return &response.UpdateProjectResponse{
			Status:      false,
			Message:     result.Message,
			ProjectInfo: ProjectMetaInfoV1DbmodelToProtobuf(projectMetaInfo),
		}, nil
	}
}

func (D DBMSServerController) GetProject(ctx context.Context, request *request.GetProjectRequest) (*response.GetProjectResponse, error) {
	_ = UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)
	projectMetaInfo := ProjectMetaInfoV1ProtobufToDbmodel(request.ProjectInfo)

	result := dal.QueryProject(projectMetaInfo, dal.GetDbInstance())
	if result.Status == true {
		fmt.Println("Project " + request.UserInfo.Name + " Get")
		return &response.GetProjectResponse{
			Status:      true,
			Message:     result.Message,
			ProjectInfo: ProjectMetaInfoV1DbmodelToProtobuf(projectMetaInfo),
		}, nil
	} else {
		return &response.GetProjectResponse{
			Status:      false,
			Message:     result.Message,
			ProjectInfo: ProjectMetaInfoV1DbmodelToProtobuf(projectMetaInfo),
		}, nil
	}
}

func (D DBMSServerController) GetAllProject(ctx context.Context, request *request.GetAllProjectRequest) (*response.GetAllProjectResponse, error) {
	_ = UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)

	var projectMetaInfoList []dbmodel.ProjectMetaInfoV1
	var protoMessage []*message.ProjectMetaInfoV1

	result := dal.QueryAllProject(&projectMetaInfoList, dal.GetDbInstance())
	if result.Status == true {
		fmt.Println("User " + request.UserInfo.Name + " Try Get AllProject")
		for _, projectMetaInfo := range projectMetaInfoList {
			protoMessage = append(protoMessage, ProjectMetaInfoV1DbmodelToProtobuf(&projectMetaInfo))
		}
		return &response.GetAllProjectResponse{
			Status:      true,
			Message:     result.Message,
			ProjectInfo: protoMessage,
		}, nil
	} else {
		return &response.GetAllProjectResponse{
			Status:      false,
			Message:     result.Message,
			ProjectInfo: protoMessage,
		}, nil
	}
}

func (D DBMSServerController) CreateSwc(ctx context.Context, request *request.CreateSwcRequest) (*response.CreateSwcResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)
	swcMetaInfo := SwcMetaInfoV1ProtobufToDbmodel(request.SwcInfo)
	swcMetaInfo.Base.Id = primitive.NewObjectID()
	swcMetaInfo.Base.Uuid = uuid.NewString()
	swcMetaInfo.Base.ApiVersion = "V1"
	swcMetaInfo.Creator = userMetaInfo.Name
	swcMetaInfo.LastModifiedTime = time.Now()
	swcMetaInfo.CreateTime = time.Now()
	swcMetaInfo.Name = request.SwcInfo.Name
	swcMetaInfo.Description = request.SwcInfo.Description

	result := dal.CreateSwc(*swcMetaInfo, dal.GetDbInstance())
	if result.Status {
		fmt.Println("User " + request.UserInfo.Name + "Create Swc " + swcMetaInfo.Name)
		return &response.CreateSwcResponse{
			Status:  true,
			Message: result.Message,
			SwcInfo: SwcMetaInfoV1DbmodelToProtobuf(swcMetaInfo),
		}, nil
	} else {
		return &response.CreateSwcResponse{
			Status:  false,
			Message: result.Message,
			SwcInfo: SwcMetaInfoV1DbmodelToProtobuf(swcMetaInfo),
		}, nil
	}
}

func (D DBMSServerController) DeleteSwc(ctx context.Context, request *request.DeleteSwcRequest) (*response.DeleteSwcResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)
	swcMetaInfo := SwcMetaInfoV1ProtobufToDbmodel(request.SwcInfo)

	result := dal.DeleteSwc(*swcMetaInfo, dal.GetDbInstance())
	if result.Status {
		fmt.Println("User " + request.UserInfo.Name + "Delete Swc " + swcMetaInfo.Name)
		return &response.DeleteSwcResponse{
			Status:  true,
			Message: result.Message,
			SwcInfo: SwcMetaInfoV1DbmodelToProtobuf(swcMetaInfo),
		}, nil
	} else {
		return &response.DeleteSwcResponse{
			Status:  false,
			Message: result.Message,
			SwcInfo: SwcMetaInfoV1DbmodelToProtobuf(swcMetaInfo),
		}, nil
	}
}

func (D DBMSServerController) UpdateSwc(ctx context.Context, request *request.UpdateSwcRequest) (*response.UpdateSwcResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)
	swcMetaInfo := SwcMetaInfoV1ProtobufToDbmodel(request.SwcInfo)

	result := dal.ModifySwc(*swcMetaInfo, dal.GetDbInstance())
	if result.Status {
		fmt.Println("User " + request.UserInfo.Name + "Update SwcMetaInfo " + swcMetaInfo.Name)
		return &response.UpdateSwcResponse{
			Status:  true,
			Message: result.Message,
			SwcInfo: SwcMetaInfoV1DbmodelToProtobuf(swcMetaInfo),
		}, nil
	} else {
		return &response.UpdateSwcResponse{
			Status:  false,
			Message: result.Message,
			SwcInfo: SwcMetaInfoV1DbmodelToProtobuf(swcMetaInfo),
		}, nil
	}
}

func (D DBMSServerController) GetSwcMetaInfo(ctx context.Context, request *request.GetSwcMetaInfoRequest) (*response.GetSwcMetaInfoResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)
	swcMetaInfo := SwcMetaInfoV1ProtobufToDbmodel(request.SwcInfo)

	result := dal.QuerySwc(swcMetaInfo, dal.GetDbInstance())
	if result.Status {
		fmt.Println("User " + request.UserInfo.Name + "Query SwcMetaInfo " + swcMetaInfo.Name)
		return &response.GetSwcMetaInfoResponse{
			Status:  true,
			Message: result.Message,
			SwcInfo: SwcMetaInfoV1DbmodelToProtobuf(swcMetaInfo),
		}, nil
	} else {
		return &response.GetSwcMetaInfoResponse{
			Status:  false,
			Message: result.Message,
			SwcInfo: SwcMetaInfoV1DbmodelToProtobuf(swcMetaInfo),
		}, nil
	}
}

func (D DBMSServerController) GetAllSwcMetaInfo(ctx context.Context, request *request.GetAllSwcMetaInfoRequest) (*response.GetAllSwcMetaInfoResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)

	var dbmodelMessage []dbmodel.SwcMetaInfoV1

	var protoMessage []*message.SwcMetaInfoV1
	result := dal.QueryAllSwc(&dbmodelMessage, dal.GetDbInstance())
	if result.Status {
		fmt.Println("User " + request.UserInfo.Name + "Query All SwcMetaInfo ")
		for _, dbMessage := range dbmodelMessage {
			protoMessage = append(protoMessage, SwcMetaInfoV1DbmodelToProtobuf(&dbMessage))
		}
		return &response.GetAllSwcMetaInfoResponse{
			Status:  true,
			Message: result.Message,
			SwcInfo: protoMessage,
		}, nil
	} else {
		return &response.GetAllSwcMetaInfoResponse{
			Status:  false,
			Message: result.Message,
			SwcInfo: protoMessage,
		}, nil
	}
}

func (D DBMSServerController) CreateSwcNodeData(ctx context.Context, request *request.CreateSwcNodeDataRequest) (*response.CreateSwcNodeDataResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)
	swcMetaInfo := SwcMetaInfoV1ProtobufToDbmodel(request.SwcInfo)

	var swcData dbmodel.SwcDataV1
	for _, swcNodeData := range request.SwcNodeData.SwcData {
		swcData = append(swcData, *SwcNodeDataV1ProtobufToDbmodel(swcNodeData))
	}

	result := dal.CreateSwcData(*swcMetaInfo, swcData, dal.GetDbInstance())
	if result.Status {
		fmt.Println("User " + request.UserInfo.Name + "Create Swc " + swcMetaInfo.Name)
		return &response.CreateSwcNodeDataResponse{
			Status:      true,
			Message:     result.Message,
			SwcNodeData: request.SwcNodeData,
		}, nil
	} else {
		return &response.CreateSwcNodeDataResponse{
			Status:      false,
			Message:     result.Message,
			SwcNodeData: request.SwcNodeData,
		}, nil
	}
}

func (D DBMSServerController) DeleteSwcNodeData(ctx context.Context, request *request.DeleteSwcNodeDataRequest) (*response.DeleteSwcNodeDataResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)
	swcMetaInfo := SwcMetaInfoV1ProtobufToDbmodel(request.SwcInfo)

	var swcData dbmodel.SwcDataV1
	for _, swcNodeData := range request.SwcNodeData.SwcData {
		swcData = append(swcData, *SwcNodeDataV1ProtobufToDbmodel(swcNodeData))
	}

	result := dal.DeleteSwcData(*swcMetaInfo, swcData, dal.GetDbInstance())
	if result.Status {
		fmt.Println("User " + request.UserInfo.Name + "Delete Swc " + swcMetaInfo.Name)
		return &response.DeleteSwcNodeDataResponse{
			Status:      true,
			Message:     result.Message,
			SwcNodeData: request.SwcNodeData,
		}, nil
	} else {
		return &response.DeleteSwcNodeDataResponse{
			Status:      false,
			Message:     result.Message,
			SwcNodeData: request.SwcNodeData,
		}, nil
	}
}

func (D DBMSServerController) UpdateSwcNodeData(ctx context.Context, request *request.UpdateSwcNodeDataRequest) (*response.UpdateSwcNodeDataResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)
	swcMetaInfo := SwcMetaInfoV1ProtobufToDbmodel(request.SwcInfo)

	swcNodeData := SwcNodeDataV1ProtobufToDbmodel(request.SwcNodeData)
	result := dal.ModifySwcData(*swcMetaInfo, *swcNodeData, dal.GetDbInstance())
	if result.Status {
		fmt.Println("User " + request.UserInfo.Name + "Update Swc " + swcMetaInfo.Name)
		return &response.UpdateSwcNodeDataResponse{
			Status:      true,
			Message:     result.Message,
			SwcNodeData: request.SwcNodeData,
		}, nil
	} else {
		return &response.UpdateSwcNodeDataResponse{
			Status:      false,
			Message:     result.Message,
			SwcNodeData: request.SwcNodeData,
		}, nil
	}
}

func (D DBMSServerController) GetSwcNodeData(ctx context.Context, request *request.GetSwcNodeDataRequest) (*response.GetSwcNodeDataResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)
	swcMetaInfo := SwcMetaInfoV1ProtobufToDbmodel(request.SwcInfo)

	var dbmodelMessage dbmodel.SwcDataV1

	var protoMessage message.SwcDataV1

	for _, swcNodeData := range request.SwcNodeData.SwcData {
		dbmodelMessage = append(dbmodelMessage, *SwcNodeDataV1ProtobufToDbmodel(swcNodeData))
	}

	result := dal.DeleteSwcData(*swcMetaInfo, dbmodelMessage, dal.GetDbInstance())
	if result.Status {
		fmt.Println("User " + request.UserInfo.Name + "Get SwcData " + swcMetaInfo.Name)

		for _, swcNodeData := range dbmodelMessage {
			protoMessage.SwcData = append(protoMessage.SwcData, SwcNodeDataV1DbmodelToProtobuf(&swcNodeData))
		}

		return &response.GetSwcNodeDataResponse{
			Status:      true,
			Message:     result.Message,
			SwcNodeData: &protoMessage,
		}, nil
	} else {
		return &response.GetSwcNodeDataResponse{
			Status:      false,
			Message:     result.Message,
			SwcNodeData: &protoMessage,
		}, nil
	}
}

func (D DBMSServerController) GetSwcFullNodeData(ctx context.Context, request *request.GetSwcFullNodeDataRequest) (*response.GetSwcFullNodeDataResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)
	swcMetaInfo := SwcMetaInfoV1ProtobufToDbmodel(request.SwcInfo)

	var dbmodelMessage dbmodel.SwcDataV1
	var protoMessage message.SwcDataV1

	result := dal.DeleteSwcData(*swcMetaInfo, dbmodelMessage, dal.GetDbInstance())
	if result.Status {
		fmt.Println("User " + request.UserInfo.Name + "Get SwcFullNodeData " + swcMetaInfo.Name)

		for _, swcNodeData := range dbmodelMessage {
			protoMessage.SwcData = append(protoMessage.SwcData, SwcNodeDataV1DbmodelToProtobuf(&swcNodeData))
		}

		return &response.GetSwcFullNodeDataResponse{
			Status:      true,
			Message:     result.Message,
			SwcNodeData: &protoMessage,
		}, nil
	} else {
		return &response.GetSwcFullNodeDataResponse{
			Status:      false,
			Message:     result.Message,
			SwcNodeData: &protoMessage,
		}, nil
	}
}

func (D DBMSServerController) GetSwcNodeDataListByTimeAndUser(ctx context.Context, request *request.GetSwcNodeDataListByTimeAndUserRequest) (*response.GetSwcNodeDataListByTimeAndUserResponse, error) {
	userMetaInfo := UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)
	swcMetaInfo := SwcMetaInfoV1ProtobufToDbmodel(request.SwcInfo)

	var dbmodelMessage dbmodel.SwcDataV1

	var protoMessage message.SwcDataV1

	for _, swcNodeData := range request.SwcNodeData.SwcData {
		dbmodelMessage = append(dbmodelMessage, *SwcNodeDataV1ProtobufToDbmodel(swcNodeData))
	}

	result := dal.DeleteSwcData(*swcMetaInfo, dbmodelMessage, dal.GetDbInstance())
	if result.Status {
		fmt.Println("User " + request.UserInfo.Name + "Get SwcData " + swcMetaInfo.Name)

		for _, swcNodeData := range dbmodelMessage {
			protoMessage.SwcData = append(protoMessage.SwcData, SwcNodeDataV1DbmodelToProtobuf(&swcNodeData))
		}

		return &response.GetSwcNodeDataListByTimeAndUserResponse{
			Status:      true,
			Message:     result.Message,
			SwcNodeData: &protoMessage,
		}, nil
	} else {
		return &response.GetSwcNodeDataListByTimeAndUserResponse{
			Status:      false,
			Message:     result.Message,
			SwcNodeData: &protoMessage,
		}, nil
	}
}

func (D DBMSServerController) BackupFullDatabase(ctx context.Context, request *request.BackupFullDatabaseRequest) (*response.BackupFullDatabaseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBMSServerController) CreateDailyStatistics(ctx context.Context, request *request.CreateDailyStatisticsRequest) (*response.CreateDailyStatisticsResponse, error) {
	_ = UserMetaInfoV1ProtobufToDbmodel(request.UserInfo)

	dailyStatisticsInfo := DailyStatisticsMetaInfoV1ProtobufToDbmodel(request.DailyStatisticsInfo)
	dailyStatisticsInfo.Base.Id = primitive.NewObjectID()
	dailyStatisticsInfo.Base.Uuid = uuid.NewString()
	dailyStatisticsInfo.Base.ApiVersion = "V1"

	result := dal.CreateDailyStatistics(*dailyStatisticsInfo, dal.GetDbInstance())
	if result.Status == true {
		fmt.Println("DailyStatistics " + request.DailyStatisticsInfo.Name + " Created")
		return &response.CreateDailyStatisticsResponse{
			Status:              true,
			Message:             result.Message,
			DailyStatisticsInfo: DailyStatisticsMetaInfoV1DbmodelToProtobuf(dailyStatisticsInfo),
		}, nil
	} else {
		return &response.CreateDailyStatisticsResponse{
			Status:              false,
			Message:             result.Message,
			DailyStatisticsInfo: DailyStatisticsMetaInfoV1DbmodelToProtobuf(dailyStatisticsInfo),
		}, nil
	}
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
