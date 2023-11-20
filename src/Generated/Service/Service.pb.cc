// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: Service/Service.proto

#include "Service/Service.pb.h"

#include <algorithm>

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/wire_format_lite.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>

PROTOBUF_PRAGMA_INIT_SEG

namespace _pb = ::PROTOBUF_NAMESPACE_ID;
namespace _pbi = _pb::internal;

namespace proto {
}  // namespace proto
static constexpr ::_pb::EnumDescriptor const** file_level_enum_descriptors_Service_2fService_2eproto = nullptr;
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_Service_2fService_2eproto = nullptr;
const uint32_t TableStruct_Service_2fService_2eproto::offsets[1] = {};
static constexpr ::_pbi::MigrationSchema* schemas = nullptr;
static constexpr ::_pb::Message* const* file_default_instances = nullptr;

const char descriptor_table_protodef_Service_2fService_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n\025Service/Service.proto\022\005proto\032\037google/p"
  "rotobuf/timestamp.proto\032\025Message/Message"
  ".proto\032\025Message/Request.proto\032\026Message/R"
  "esponse.proto2\356\026\n\004DBMS\022C\n\nCreateUser\022\030.p"
  "roto.CreateUserRequest\032\031.proto.CreateUse"
  "rResponse\"\000\022C\n\nDeleteUser\022\030.proto.Delete"
  "UserRequest\032\031.proto.DeleteUserResponse\"\000"
  "\022C\n\nUpdateUser\022\030.proto.UpdateUserRequest"
  "\032\031.proto.UpdateUserResponse\"\000\022:\n\007GetUser"
  "\022\025.proto.GetUserRequest\032\026.proto.GetUserR"
  "esponse\"\000\022C\n\nGetAllUser\022\030.proto.GetAllUs"
  "erRequest\032\031.proto.GetAllUserResponse\"\000\022@"
  "\n\tUserLogin\022\027.proto.UserLoginRequest\032\030.p"
  "roto.UserLoginResponse\"\000\022C\n\nUserLogout\022\030"
  ".proto.UserLogoutRequest\032\031.proto.UserLog"
  "outResponse\"\000\022p\n UserOnlineHeartBeatNoti"
  "fications\022&.proto.UserOnlineHeartBeatNot"
  "ification\032\".proto.UserOnlineHeartBeatRes"
  "ponse\"\000\022g\n\026GetUserPermissionGroup\022$.prot"
  "o.GetUserPermissionGroupRequest\032%.proto."
  "GetUserPermissionGroupResponse\"\000\022[\n\022GetP"
  "ermissionGroup\022 .proto.GetPermissionGrou"
  "pRequest\032!.proto.GetPermissionGroupRespo"
  "nse\"\000\022d\n\025GetAllPermissionGroup\022#.proto.G"
  "etAllPermissionGroupRequest\032$.proto.GetA"
  "llPermissionGroupResponse\"\000\022p\n\031ChangeUse"
  "rPermissionGroup\022\'.proto.ChangeUserPermi"
  "ssionGroupRequest\032(.proto.ChangeUserPerm"
  "issionGroupResponse\"\000\022L\n\rCreateProject\022\033"
  ".proto.CreateProjectRequest\032\034.proto.Crea"
  "teProjectResponse\"\000\022L\n\rDeleteProject\022\033.p"
  "roto.DeleteProjectRequest\032\034.proto.Delete"
  "ProjectResponse\"\000\022L\n\rUpdateProject\022\033.pro"
  "to.UpdateProjectRequest\032\034.proto.UpdatePr"
  "ojectResponse\"\000\022C\n\nGetProject\022\030.proto.Ge"
  "tProjectRequest\032\031.proto.GetProjectRespon"
  "se\"\000\022L\n\rGetAllProject\022\033.proto.GetAllProj"
  "ectRequest\032\034.proto.GetAllProjectResponse"
  "\"\000\022@\n\tCreateSwc\022\027.proto.CreateSwcRequest"
  "\032\030.proto.CreateSwcResponse\"\000\022@\n\tDeleteSw"
  "c\022\027.proto.DeleteSwcRequest\032\030.proto.Delet"
  "eSwcResponse\"\000\022@\n\tUpdateSwc\022\027.proto.Upda"
  "teSwcRequest\032\030.proto.UpdateSwcResponse\"\000"
  "\022O\n\016GetSwcMetaInfo\022\034.proto.GetSwcMetaInf"
  "oRequest\032\035.proto.GetSwcMetaInfoResponse\""
  "\000\022X\n\021GetAllSwcMetaInfo\022\037.proto.GetAllSwc"
  "MetaInfoRequest\032 .proto.GetAllSwcMetaInf"
  "oResponse\"\000\022X\n\021CreateSwcNodeData\022\037.proto"
  ".CreateSwcNodeDataRequest\032 .proto.Create"
  "SwcNodeDataResponse\"\000\022X\n\021DeleteSwcNodeDa"
  "ta\022\037.proto.DeleteSwcNodeDataRequest\032 .pr"
  "oto.DeleteSwcNodeDataResponse\"\000\022X\n\021Updat"
  "eSwcNodeData\022\037.proto.UpdateSwcNodeDataRe"
  "quest\032 .proto.UpdateSwcNodeDataResponse\""
  "\000\022O\n\016GetSwcNodeData\022\034.proto.GetSwcNodeDa"
  "taRequest\032\035.proto.GetSwcNodeDataResponse"
  "\"\000\022[\n\022GetSwcFullNodeData\022 .proto.GetSwcF"
  "ullNodeDataRequest\032!.proto.GetSwcFullNod"
  "eDataResponse\"\000\022\202\001\n\037GetSwcNodeDataListBy"
  "TimeAndUser\022-.proto.GetSwcNodeDataListBy"
  "TimeAndUserRequest\032..proto.GetSwcNodeDat"
  "aListByTimeAndUserResponse\"\000\022[\n\022BackupFu"
  "llDatabase\022 .proto.BackupFullDatabaseReq"
  "uest\032!.proto.BackupFullDatabaseResponse\""
  "\000\022d\n\025CreateDailyStatistics\022#.proto.Creat"
  "eDailyStatisticsRequest\032$.proto.CreateDa"
  "ilyStatisticsResponse\"\000\022d\n\025DeleteDailySt"
  "atistics\022#.proto.DeleteDailyStatisticsRe"
  "quest\032$.proto.DeleteDailyStatisticsRespo"
  "nse\"\000\022d\n\025UpdateDailyStatistics\022#.proto.U"
  "pdateDailyStatisticsRequest\032$.proto.Upda"
  "teDailyStatisticsResponse\"\000\022[\n\022GetDailyS"
  "tatistics\022 .proto.GetDailyStatisticsRequ"
  "est\032!.proto.GetDailyStatisticsResponse\"\000"
  "\022d\n\025GetAllDailyStatistics\022#.proto.GetAll"
  "DailyStatisticsRequest\032$.proto.GetAllDai"
  "lyStatisticsResponse\"\000B\036Z\034DBMS/Generated"
  "/proto/serviceP\000b\006proto3"
  ;
static const ::_pbi::DescriptorTable* const descriptor_table_Service_2fService_2eproto_deps[4] = {
  &::descriptor_table_Message_2fMessage_2eproto,
  &::descriptor_table_Message_2fRequest_2eproto,
  &::descriptor_table_Message_2fResponse_2eproto,
  &::descriptor_table_google_2fprotobuf_2ftimestamp_2eproto,
};
static ::_pbi::once_flag descriptor_table_Service_2fService_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_Service_2fService_2eproto = {
    false, false, 3104, descriptor_table_protodef_Service_2fService_2eproto,
    "Service/Service.proto",
    &descriptor_table_Service_2fService_2eproto_once, descriptor_table_Service_2fService_2eproto_deps, 4, 0,
    schemas, file_default_instances, TableStruct_Service_2fService_2eproto::offsets,
    nullptr, file_level_enum_descriptors_Service_2fService_2eproto,
    file_level_service_descriptors_Service_2fService_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_Service_2fService_2eproto_getter() {
  return &descriptor_table_Service_2fService_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_Service_2fService_2eproto(&descriptor_table_Service_2fService_2eproto);
namespace proto {

// @@protoc_insertion_point(namespace_scope)
}  // namespace proto
PROTOBUF_NAMESPACE_OPEN
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
