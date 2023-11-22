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

};
