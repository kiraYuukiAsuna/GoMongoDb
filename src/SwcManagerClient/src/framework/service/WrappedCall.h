#pragma once

#include <QMessageBox>
#include <QString>
#include "grpcpp/client_context.h"
#include "Message/Request.pb.h"
#include "RpcCall.h"
#include "CachedProtoData.h"

class WrappedCall{
public:
    static bool getAllProjectMetaInfo(proto::GetAllProjectResponse& response, QWidget* parent){
        grpc::ClientContext context;
        proto::GetAllProjectRequest request;
        auto* userInfo = request.mutable_userinfo();
        userInfo->CopyFrom(CachedProtoData::getInstance().CachedUserMetaInfo);

        auto& rpcCall = RpcCall::getInstance();
        auto status = rpcCall.Stub()->GetAllProject(&context, request, &response);
        if(status.ok()){
            if(response.status()) {
                return true;
            }else {
                QMessageBox::warning(parent,"Info","GetAllProjectMetaInfo Failed!" + QString::fromStdString(response.message()));
            }

        }else{
            QMessageBox::critical(parent,"Error",QString::fromStdString(status.error_message()));
        }
        return false;
    }

    static bool getAllSwcMetaInfo(proto::GetAllSwcMetaInfoResponse& response, QWidget* parent){
        grpc::ClientContext context;
        proto::GetAllSwcMetaInfoRequest request;
        auto* userInfo = request.mutable_userinfo();
        userInfo->CopyFrom(CachedProtoData::getInstance().CachedUserMetaInfo);

        auto& rpcCall = RpcCall::getInstance();
        auto status = rpcCall.Stub()->GetAllSwcMetaInfo(&context, request, &response);
        if(status.ok()){
            if(response.status()) {
                return true;
            }else {
                QMessageBox::warning(parent,"Info","GetAllSwcMetaInfo Failed!" + QString::fromStdString(response.message()));
            }

        }else{
            QMessageBox::critical(parent,"Error",QString::fromStdString(status.error_message()));
        }
        return false;
    }

    static bool getAllDailyStatisticsMetaInfo(proto::GetAllDailyStatisticsResponse& response, QWidget* parent){
        grpc::ClientContext context;
        proto::GetAllDailyStatisticsRequest request;
        auto* userInfo = request.mutable_userinfo();
        userInfo->CopyFrom(CachedProtoData::getInstance().CachedUserMetaInfo);

        auto& rpcCall = RpcCall::getInstance();
        auto status = rpcCall.Stub()->GetAllDailyStatistics(&context, request, &response);
        if(status.ok()){
            if(response.status()) {
                return true;
            }else {
                QMessageBox::warning(parent,"Info","GetAllDailyStatistics Failed!" + QString::fromStdString(response.message()));
            }

        }else{
            QMessageBox::critical(parent,"Error",QString::fromStdString(status.error_message()));
        }
        return false;
    }

    static bool getAllUserMetaInfo(proto::GetAllUserResponse& response, QWidget* parent){
        grpc::ClientContext context;
        proto::GetAllUserRequest request;
        auto* userInfo = request.mutable_userinfo();
        userInfo->CopyFrom(CachedProtoData::getInstance().CachedUserMetaInfo);

        auto& rpcCall = RpcCall::getInstance();
        auto status = rpcCall.Stub()->GetAllUser(&context, request, &response);
        if(status.ok()){
            if(response.status()) {
                return true;
            }else {
                QMessageBox::warning(parent,"Info","GetAllUserMetaInfo Failed!" + QString::fromStdString(response.message()));
            }

        }else{
            QMessageBox::critical(parent,"Error",QString::fromStdString(status.error_message()));
        }
        return false;
    }

    static bool getProjectMetaInfoByName(const std::string& projectName, proto::GetProjectResponse& response, QWidget* parent) {
        proto::GetProjectRequest request;
        request.mutable_userinfo()->CopyFrom(CachedProtoData::getInstance().CachedUserMetaInfo);
        request.mutable_projectinfo()->set_name(projectName);

        grpc::ClientContext context;
        auto result = RpcCall::getInstance().Stub()->GetProject(&context, request,&response);
        if(result.ok()){
            if(response.status()) {
                return true;
            }else {
                QMessageBox::warning(parent,"Info","GetProjectMetaInfo Failed!" + QString::fromStdString(response.message()));
            }
        }else{
            QMessageBox::critical(parent,"Error",QString::fromStdString(result.error_message()));
        }
        return false;
    }

    static bool getSwcMetaInfoByName(const std::string& swcName, proto::GetSwcMetaInfoResponse& response, QWidget* parent){
        proto::GetSwcMetaInfoRequest request;
        request.mutable_userinfo()->CopyFrom(CachedProtoData::getInstance().CachedUserMetaInfo);
        request.mutable_swcinfo()->set_name(swcName);


        grpc::ClientContext context;
        auto result = RpcCall::getInstance().Stub()->GetSwcMetaInfo(&context, request,&response);
        if(result.ok()){
            if(response.status()) {
                return true;
            }else {
                QMessageBox::warning(parent,"Info","GetSwcMetaInfo Failed!" + QString::fromStdString(response.message()));
            }

        }else{
            QMessageBox::critical(parent,"Error",QString::fromStdString(result.error_message()));
        }
        return false;
    }

    static bool getDailyStatisticsmMetaInfoByName(const std::string& dailyStatisticsName, proto::GetDailyStatisticsResponse& response, QWidget* parent) {
        proto::GetDailyStatisticsRequest request;
        request.mutable_userinfo()->CopyFrom(CachedProtoData::getInstance().CachedUserMetaInfo);
        request.mutable_dailystatisticsinfo()->set_name(dailyStatisticsName);

        grpc::ClientContext context;
        auto result = RpcCall::getInstance().Stub()->GetDailyStatistics(&context, request,&response);
        if(result.ok()){
            if(response.status()) {
                return true;
            }else {
                QMessageBox::warning(parent,"Info","GetSwcMetaInfo Failed!" + QString::fromStdString(response.message()));
            }

        }else{
            QMessageBox::critical(parent,"Error",QString::fromStdString(result.error_message()));
        }
        return false;
    }

};
