package mongodb

import (
	"bytes"
	"time"
)

type MetaInfoBase struct {
	_id        string
	apiversion string
	uuid       string
}

type UserMetaInfoV1 struct {
	base                MetaInfoBase
	name                string
	description         string
	createTime          time.Time
	HeadPhotoBinData    bytes.Buffer
	UserPermissionGroup string
}

type GlobalPermissionMetaInfoV1 struct {
	writePermission_CreateProject bool
	writePermission_ModifyProject bool
	writePermission_DeleteProject bool

	ReadPerimission_Query bool
}

type ProjectPermissionMetaInfoV1 struct {
	writePermission_AddData    bool
	writePermission_ModifyData bool
	writePermission_DeleteData bool

	ReadPerimission_Query bool
}

type PermissionGroupMetaInfoV1 struct {
	base MetaInfoBase

	name        string
	description string

	global  GlobalPermissionMetaInfoV1
	project ProjectPermissionMetaInfoV1
}

type UserPremissionOverrideMetaInfoV1 struct {
	global  GlobalPermissionMetaInfoV1
	project ProjectPermissionMetaInfoV1
}

type ProjectMetaInfoV1 struct {
	base MetaInfoBase

	name                   string
	description            string
	creator                string
	createTime             time.Time
	lastModifiedTime       time.Time
	swcList                []string
	userPermissionOverride []ProjectPermissionMetaInfoV1
	workMode               string
}

type SwcNodeDataV1 struct {
	base              MetaInfoBase
	swcData           string // replace with actually content def
	creator           string
	createTime        time.Time
	lastModifiedTime  time.Time
	annotatorUserUuid string // unknow
	checkerUuid       string // unknow
}

type SwcDataV1 = []SwcNodeDataV1

type DailyStatisticsMetaInfoV1 struct {
	base        MetaInfoBase
	name        string
	description string
	day         string

	CreatedProjectNumber int32
	CreatedSwcNumber     int32
	CreateSwcNodeNumber  int32

	DeletedProjectNumber int32
	DeletedSwcNumber     int32
	DeletedSwcNodeNumber int32

	ModifiedProjectNumber int32
	ModifiedSwcNumber     int32
	ModifiedSwcNodeNumber int32

	projectQueryNumber int32
	swcQueryNumber     int32
	nodeQueryNumber    int32

	activeUserNumber int32
}

SWCmetainfov1
