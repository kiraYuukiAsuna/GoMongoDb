package dbmodel

import (
	"bytes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	ProjectMetaInfoCollectionString         = "ProjectMetaInfoCollectionString"
	UserMetaInfoCollectionString            = "UserMetaInfoCollection"
	PermissionGroupMetaInfoCollectioString  = "PermissionGroupMetaInfoCollection"
	SwcMetaInfoCollectionString             = "SwcMetaInfoCollection"
	DailyStatisticsMetaInfoCollectionString = "DailyStatisticsMetaInfoCollection"
)

type MetaInfoBase struct {
	Id         primitive.ObjectID `bson:"_id"`
	ApiVersion string             `bson:"ApiVersion"`
	Uuid       string             `bson:"uuid"`
}

type UserMetaInfoV1 struct {
	Base                MetaInfoBase  `bson:"Base,inline"`
	Name                string        `bson:"Name"`
	Password            string        `bson:"Password"`
	Description         string        `bson:"Description"`
	CreateTime          time.Time     `bson:"CreateTime"`
	HeadPhotoBinData    *bytes.Buffer `bson:"HeadPhotoBinData"`
	UserPermissionGroup string        `bson:"UserPermissionGroup"`
}

type GlobalPermissionMetaInfoV1 struct {
	WritePermissionCreateProject bool `bson:"WritePermissionCreateProject"`
	WritePermissionModifyProject bool `bson:"WritePermissionModifyProject"`
	WritePermissionDeleteProject bool `bson:"WritePermissionDeleteProject"`

	ReadPerimissionQuery bool `bson:"ReadPerimissionQuery"`
}

type ProjectPermissionMetaInfoV1 struct {
	WritePermissionAddData    bool `bson:"WritePermissionAddData"`
	WritePermissionModifyData bool `bson:"WritePermissionModifyData"`
	WritePermissionDeleteData bool `bson:"WritePermissionDeleteData"`

	ReadPerimissionQuery bool `bson:"ReadPerimissionQuery"`
}

type PermissionGroupMetaInfoV1 struct {
	Base MetaInfoBase `bson:"Base,inline"`

	Name        string `bson:"Name"`
	Description string `bson:"Description"`

	Global  GlobalPermissionMetaInfoV1  `bson:"Global"`
	Project ProjectPermissionMetaInfoV1 `bson:"Project"`
}

type UserPermissionOverrideMetaInfoV1 struct {
	Project  ProjectPermissionMetaInfoV1 `bson:"Project"`
	UserName string                      `bson:"UserName"`
}

type ProjectMetaInfoV1 struct {
	Base MetaInfoBase `bson:"Base,inline"`

	Name                   string                             `bson:"Name"`
	Description            string                             `bson:"Description"`
	Creator                string                             `bson:"Creator"`
	CreateTime             time.Time                          `bson:"CreateTime"`
	LastModifiedTime       time.Time                          `bson:"LastModifiedTime"`
	SwcList                []string                           `bson:"SwcList"`
	UserPermissionOverride []UserPermissionOverrideMetaInfoV1 `bson:"UserPermissionOverride"`
	WorkMode               string                             `bson:"WorkMode"`
}

type SwcMetaInfoV1 struct {
	Base             MetaInfoBase `bson:"Base,inline"`
	Name             string       `bson:"Name"`
	Description      string       `bson:"Description"`
	Creator          string       `bson:"Creator"`
	CreateTime       time.Time    `bson:"CreateTime"`
	LastModifiedTime time.Time    `bson:"LastModifiedTime"`
}

type SwcNodeDataV1 struct {
	Base              MetaInfoBase `bson:"Base,inline"`
	SwcData           string       `bson:"SwcData"` // replace with actually content def
	Creator           string       `bson:"Creator"`
	CreateTime        time.Time    `bson:"CreateTime"`
	LastModifiedTime  time.Time    `bson:"LastModifiedTime"`
	AnnotatorUserUuid string       `bson:"AnnotatorUserUuid"`
	CheckerUserUuid   string       `bson:"CheckerUserUuid"`
}

type SwcDataV1 = []SwcNodeDataV1

type DailyStatisticsMetaInfoV1 struct {
	Base        MetaInfoBase `bson:"Base,inline"`
	Name        string       `bson:"Name"`
	Description string       `bson:"Description"`
	Day         string       `bson:"Day"`

	CreatedProjectNumber int32 `bson:"CreatedProjectNumber"`
	CreatedSwcNumber     int32 `bson:"CreatedSwcNumber"`
	CreateSwcNodeNumber  int32 `bson:"CreateSwcNodeNumber"`

	DeletedProjectNumber int32 `bson:"DeletedProjectNumber"`
	DeletedSwcNumber     int32 `bson:"DeletedSwcNumber"`
	DeletedSwcNodeNumber int32 `bson:"DeletedSwcNodeNumber"`

	ModifiedProjectNumber int32 `bson:"ModifiedProjectNumber"`
	ModifiedSwcNumber     int32 `bson:"ModifiedSwcNumber"`
	ModifiedSwcNodeNumber int32 `bson:"ModifiedSwcNodeNumber"`

	ProjectQueryNumber int32 `bson:"ProjectQueryNumber"`
	SwcQueryNumber     int32 `bson:"SwcQueryNumber"`
	NodeQueryNumber    int32 `bson:"NodeQueryNumber"`

	ActiveUserNumber int32 `bson:"ActiveUserNumber"`
}
