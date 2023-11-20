package bll

import (
	"DBMS/Generated/proto/message"
	"DBMS/dbmodel"
	"bytes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserMetaInfoV1ProtobufToDbmodel(protoMessage *message.UserMetaInfoV1) *dbmodel.UserMetaInfoV1 {
	var userMetaInfo dbmodel.UserMetaInfoV1

	if protoMessage.Base != nil {
		userMetaInfo.Base.Id, _ = primitive.ObjectIDFromHex(protoMessage.Base.XId)
		userMetaInfo.Base.Uuid = protoMessage.Base.Uuid
		userMetaInfo.Base.ApiVersion = protoMessage.Base.ApiVersion
	}

	userMetaInfo.Name = protoMessage.Name
	userMetaInfo.Password = protoMessage.Password
	userMetaInfo.Description = protoMessage.Description
	userMetaInfo.UserPermissionGroup = protoMessage.UserPermissionGroup
	userMetaInfo.OnlineStatus = protoMessage.OnlineStatus

	if protoMessage.CreateTime != nil {
		userMetaInfo.CreateTime = protoMessage.CreateTime.AsTime()
	}

	if protoMessage.HeadPhotoBinData != nil {
		userMetaInfo.HeadPhotoBinData = bytes.NewBuffer(protoMessage.HeadPhotoBinData)
	}

	return &userMetaInfo
}

func UserMetaInfoV1DbmodelToProtobuf(userMetaInfoV1 *dbmodel.UserMetaInfoV1) *message.UserMetaInfoV1 {
	var protoMessage message.UserMetaInfoV1
	protoMessage.Base = &message.MetaInfoBase{}
	protoMessage.Base.XId = userMetaInfoV1.Base.Id.String()
	protoMessage.Base.Uuid = userMetaInfoV1.Base.Uuid
	protoMessage.Base.ApiVersion = userMetaInfoV1.Base.ApiVersion

	protoMessage.Name = userMetaInfoV1.Name
	protoMessage.Password = userMetaInfoV1.Password
	protoMessage.Description = userMetaInfoV1.Description
	protoMessage.UserPermissionGroup = userMetaInfoV1.UserPermissionGroup
	protoMessage.CreateTime = timestamppb.New(userMetaInfoV1.CreateTime)
	if userMetaInfoV1.HeadPhotoBinData != nil {
		protoMessage.HeadPhotoBinData = userMetaInfoV1.HeadPhotoBinData.Bytes()
	}

	return &protoMessage
}
