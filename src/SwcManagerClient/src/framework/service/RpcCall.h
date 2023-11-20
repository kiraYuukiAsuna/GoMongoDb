#pragma once
#include <string>
#include <grpcpp/grpcpp.h>
#include <Service/Service.grpc.pb.h>

class RpcCall {
public:
    RpcCall(RpcCall&) = delete;
    RpcCall& operator=(RpcCall&) = delete;

    void initialize(const std::string& endpoint) {
        m_Endpoint = endpoint;
        m_Channel = grpc::CreateChannel(m_Endpoint, grpc::InsecureChannelCredentials());
        m_Stub = proto::DBMS::NewStub(m_Channel);

    }

    static RpcCall& getInstance() {
        static RpcCall instance;
        return instance;
    }

    auto& Endpoint() {
        return m_Endpoint;
    }

    auto& Channel() {
        return m_Channel;
    }

    auto& Stub() {
        return m_Stub;
    }

private:
    RpcCall() {
    }

    std::string m_Endpoint;
    std::shared_ptr<grpc::Channel> m_Channel;
    std::unique_ptr<proto::DBMS::Stub> m_Stub;
};
