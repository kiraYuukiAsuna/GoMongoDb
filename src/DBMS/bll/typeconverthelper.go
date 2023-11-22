package bll

import (
	"DBMS/Generated/proto/message"
	"DBMS/dbmodel"
	"bytes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserMetaInfoV1ProtobufToDbmodel(protoMessage *message.UserMetaInfoV1) *dbmodel.UserMetaInfoV1 {
	var dbmodelMessage dbmodel.UserMetaInfoV1

	if protoMessage.Base != nil {
		dbmodelMessage.Base.Id, _ = primitive.ObjectIDFromHex(protoMessage.Base.XId)
		dbmodelMessage.Base.Uuid = protoMessage.Base.Uuid
		dbmodelMessage.Base.ApiVersion = protoMessage.Base.ApiVersion
	}

	dbmodelMessage.Name = protoMessage.Name
	dbmodelMessage.Password = protoMessage.Password
	dbmodelMessage.Description = protoMessage.Description
	dbmodelMessage.UserPermissionGroup = protoMessage.UserPermissionGroup
	dbmodelMessage.OnlineStatus = protoMessage.OnlineStatus

	if protoMessage.CreateTime != nil {
		dbmodelMessage.CreateTime = protoMessage.CreateTime.AsTime()
	}

	if protoMessage.HeadPhotoBinData != nil {
		dbmodelMessage.HeadPhotoBinData = bytes.NewBuffer(protoMessage.HeadPhotoBinData)
	}

	return &dbmodelMessage
}

func UserMetaInfoV1DbmodelToProtobuf(dbmodelMessage *dbmodel.UserMetaInfoV1) *message.UserMetaInfoV1 {
	var protoMessage message.UserMetaInfoV1
	protoMessage.Base = &message.MetaInfoBase{}
	protoMessage.Base.XId = dbmodelMessage.Base.Id.String()
	protoMessage.Base.Uuid = dbmodelMessage.Base.Uuid
	protoMessage.Base.ApiVersion = dbmodelMessage.Base.ApiVersion

	protoMessage.Name = dbmodelMessage.Name
	protoMessage.Password = dbmodelMessage.Password
	protoMessage.Description = dbmodelMessage.Description
	protoMessage.UserPermissionGroup = dbmodelMessage.UserPermissionGroup
	protoMessage.CreateTime = timestamppb.New(dbmodelMessage.CreateTime)
	if dbmodelMessage.HeadPhotoBinData != nil {
		protoMessage.HeadPhotoBinData = dbmodelMessage.HeadPhotoBinData.Bytes()
	}

	return &protoMessage
}

func PermissionGroupMetaInfoV1ProtobufToDbmodel(protoMessage *message.PermissionGroupMetaInfoV1) *dbmodel.PermissionGroupMetaInfoV1 {
	var dbmodelMessage dbmodel.PermissionGroupMetaInfoV1
	if protoMessage.Base != nil {
		dbmodelMessage.Base.Id, _ = primitive.ObjectIDFromHex(protoMessage.Base.XId)
		dbmodelMessage.Base.Uuid = protoMessage.Base.Uuid
		dbmodelMessage.Base.ApiVersion = protoMessage.Base.ApiVersion
	}

	dbmodelMessage.Name = protoMessage.Name
	dbmodelMessage.Description = protoMessage.Description
	if protoMessage.GlobalPermission != nil {
		dbmodelMessage.Global.ReadPerimissionQuery = protoMessage.GlobalPermission.ReadPerimissionQuery
		dbmodelMessage.Global.WritePermissionCreateProject = protoMessage.GlobalPermission.WritePermissionCreateProject
		dbmodelMessage.Global.WritePermissionModifyProject = protoMessage.GlobalPermission.WritePermissionModifyProject
		dbmodelMessage.Global.WritePermissionCreateProject = protoMessage.GlobalPermission.WritePermissionCreateProject
	}
	if protoMessage.ProjectPermission != nil {
		dbmodelMessage.Project.ReadPerimissionQuery = protoMessage.ProjectPermission.ReadPerimissionQuery
		dbmodelMessage.Project.WritePermissionAddData = protoMessage.ProjectPermission.WritePermissionAddData
		dbmodelMessage.Project.WritePermissionModifyData = protoMessage.ProjectPermission.WritePermissionModifyData
		dbmodelMessage.Project.WritePermissionDeleteData = protoMessage.ProjectPermission.WritePermissionDeleteData
	}

	return &dbmodelMessage
}

func PermissionGroupMetaInfoV1DbmodelToProtobuf(dbmodelMessage *dbmodel.PermissionGroupMetaInfoV1) *message.PermissionGroupMetaInfoV1 {
	var protoMessage message.PermissionGroupMetaInfoV1

	protoMessage.Base = &message.MetaInfoBase{}
	protoMessage.Base.XId = dbmodelMessage.Base.Id.String()
	protoMessage.Base.Uuid = dbmodelMessage.Base.Uuid
	protoMessage.Base.ApiVersion = dbmodelMessage.Base.ApiVersion

	protoMessage.Name = dbmodelMessage.Name
	protoMessage.Description = dbmodelMessage.Description

	protoMessage.GlobalPermission = &message.GlobalPermissionMetaInfoV1{}
	protoMessage.GlobalPermission.ReadPerimissionQuery = dbmodelMessage.Global.ReadPerimissionQuery
	protoMessage.GlobalPermission.WritePermissionCreateProject = dbmodelMessage.Global.WritePermissionCreateProject
	protoMessage.GlobalPermission.WritePermissionModifyProject = dbmodelMessage.Global.WritePermissionModifyProject
	protoMessage.GlobalPermission.WritePermissionDeleteProject = dbmodelMessage.Global.WritePermissionDeleteProject

	protoMessage.ProjectPermission = &message.ProjectPermissionMetaInfoV1{}
	protoMessage.ProjectPermission.ReadPerimissionQuery = dbmodelMessage.Project.ReadPerimissionQuery
	protoMessage.ProjectPermission.WritePermissionAddData = dbmodelMessage.Project.WritePermissionAddData
	protoMessage.ProjectPermission.WritePermissionModifyData = dbmodelMessage.Project.WritePermissionModifyData
	protoMessage.ProjectPermission.WritePermissionDeleteData = dbmodelMessage.Project.WritePermissionDeleteData

	return &protoMessage
}

func ProjectMetaInfoV1ProtobufToDbmodel(protoMessage *message.ProjectMetaInfoV1) *dbmodel.ProjectMetaInfoV1 {
	var dbmodelMessage dbmodel.ProjectMetaInfoV1

	if protoMessage.Base != nil {
		dbmodelMessage.Base.Id, _ = primitive.ObjectIDFromHex(protoMessage.Base.XId)
		dbmodelMessage.Base.Uuid = protoMessage.Base.Uuid
		dbmodelMessage.Base.ApiVersion = protoMessage.Base.ApiVersion
	}

	dbmodelMessage.Name = protoMessage.Name
	dbmodelMessage.Description = protoMessage.Description
	dbmodelMessage.Creator = protoMessage.Creator

	if protoMessage.CreateTime != nil {
		dbmodelMessage.CreateTime = protoMessage.CreateTime.AsTime()
	}

	if protoMessage.LastModifiedTime != nil {
		dbmodelMessage.LastModifiedTime = protoMessage.LastModifiedTime.AsTime()
	}

	if protoMessage.SwcList != nil {
		dbmodelMessage.SwcList = protoMessage.SwcList
	}

	dbmodelMessage.WorkMode = protoMessage.WorkMode

	if protoMessage.UserPermissionOverride != nil {
		for _, protoPermissionOverride := range protoMessage.UserPermissionOverride {
			var dbmodelPermissionOverride dbmodel.ProjectPermissionMetaInfoV1
			dbmodelPermissionOverride.ReadPerimissionQuery = protoPermissionOverride.ReadPerimissionQuery
			dbmodelPermissionOverride.WritePermissionAddData = protoPermissionOverride.WritePermissionAddData
			dbmodelPermissionOverride.WritePermissionModifyData = protoPermissionOverride.WritePermissionModifyData
			dbmodelPermissionOverride.WritePermissionDeleteData = protoPermissionOverride.WritePermissionDeleteData

			dbmodelMessage.UserPermissionOverride = append(dbmodelMessage.UserPermissionOverride, dbmodelPermissionOverride)
		}
	}

	return &dbmodelMessage
}

func ProjectMetaInfoV1DbmodelToProtobuf(dbmodelMessage *dbmodel.ProjectMetaInfoV1) *message.ProjectMetaInfoV1 {
	var protoMessage message.ProjectMetaInfoV1
	protoMessage.Base = &message.MetaInfoBase{}
	protoMessage.Base.XId = dbmodelMessage.Base.Id.String()
	protoMessage.Base.Uuid = dbmodelMessage.Base.Uuid
	protoMessage.Base.ApiVersion = dbmodelMessage.Base.ApiVersion

	protoMessage.Name = dbmodelMessage.Name
	protoMessage.Description = dbmodelMessage.Description
	protoMessage.Creator = dbmodelMessage.Creator

	protoMessage.CreateTime = timestamppb.New(dbmodelMessage.CreateTime)
	protoMessage.LastModifiedTime = timestamppb.New(dbmodelMessage.LastModifiedTime)
	protoMessage.SwcList = dbmodelMessage.SwcList
	protoMessage.WorkMode = dbmodelMessage.WorkMode

	for _, dbmodelPermissionOverride := range dbmodelMessage.UserPermissionOverride {
		var protoPermissionOverride message.ProjectPermissionMetaInfoV1
		protoPermissionOverride.ReadPerimissionQuery = dbmodelPermissionOverride.ReadPerimissionQuery
		protoPermissionOverride.WritePermissionAddData = dbmodelPermissionOverride.WritePermissionAddData
		protoPermissionOverride.WritePermissionModifyData = dbmodelPermissionOverride.WritePermissionModifyData
		protoPermissionOverride.WritePermissionDeleteData = dbmodelPermissionOverride.WritePermissionDeleteData

		protoMessage.UserPermissionOverride = append(protoMessage.UserPermissionOverride, &protoPermissionOverride)
	}

	return &protoMessage
}

func SwcMetaInfoV1ProtobufToDbmodel(protoMessage *message.SwcMetaInfoV1) *dbmodel.SwcMetaInfoV1 {

}

func SwcMetaInfoV1DbmodelToProtobuf(dbmodelMessage *dbmodel.SwcMetaInfoV1) *message.SwcMetaInfoV1 {

}
