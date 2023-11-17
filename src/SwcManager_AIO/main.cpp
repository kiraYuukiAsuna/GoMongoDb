#include <QApplication>
#include <QPushButton>
#include <iostream>
#include "grpcpp/grpcpp.h"
#include "Service/Service.grpc.pb.h"

class DBMSServiceImpl final : public proto::DBMS::Service{
public:
    virtual ::grpc::Status CreateUser(::grpc::ServerContext* context,
                                      const ::proto::CreateUserRequest* request,
                                      ::proto::CreateUserResponse* response){
        ::proto::UserMetaInfoV1* metaInfoV1 = new ::proto::UserMetaInfoV1;
        metaInfoV1->set_name("Hello World!");
        auto req = request->New();
        req->set_allocated_userinfo(metaInfoV1);
        delete metaInfoV1;
    }
};

int main(int argc, char** argv) {
    setbuf(stdout, nullptr);

    QApplication a(argc, argv);
    QPushButton button("Hello world!", nullptr);
    button.resize(200, 100);
    button.show();
    return QApplication::exec();

    auto ch = grpc::CreateChannel("127.0.0.1:8080", grpc::InsecureChannelCredentials());

    auto cli = proto::DBMS::NewStub(ch);

    grpc::ClientContext context;

    proto::CreateUserRequest req;
    proto::UserMetaInfoV1 metaInfo;
    req.mutable_userinfo()->set_name("WRL");

    proto::CreateUserResponse rsp;
    auto status = cli->CreateUser(&context, req, &rsp);
    if(status.ok()){
        std::cout<<rsp.message()<<"\n";

    }else{
        std::cout<<status.error_message()<<"\n";
    }



}
