#include <QMessageBox>
#include "editorusersettings.h"
#include "ui_EditorUserSettings.h"
#include "Service//Service.pb.h"
#include "Service/Service.grpc.pb.h"
#include "src/framework/service/RpcCall.h"
#include "src/framework/service/CachedProtoData.h"
#include "src/framework/defination/ImageDefination.h"

EditorUserSettings::EditorUserSettings(QWidget *parent) :
        QDialog(parent), ui(new Ui::EditorUserSettings) {
    ui->setupUi(this);
    setWindowIcon(QIcon(Image::ImageUser));

    ui->ChangeHeadPhoto->setIcon(QIcon(Image::ImageEdit));
    connect(ui->ChangeHeadPhoto,&QPushButton::clicked,this,[&](){

    });

    getUserMetaInfo();
}

EditorUserSettings::~EditorUserSettings() {
    delete ui;
}

void EditorUserSettings::getUserMetaInfo() {
    proto::GetUserRequest request;
    request.mutable_userinfo()->set_name(CachedProtoData::getInstance().CachedUserMetaInfo.name());

    proto::GetUserResponse response;

    grpc::ClientContext context;

    auto result = RpcCall::getInstance().Stub()->GetUser(&context,request,&response);
    if(result.ok()) {
        CachedProtoData::getInstance().CachedUserMetaInfo.CopyFrom(response.userinfo());

        ui->Id->setText(QString::fromStdString(response.userinfo().base()._id()));
        ui->Uuid->setText(QString::fromStdString(response.userinfo().base().uuid()));
        ui->ApiVersion->setText(QString::fromStdString(response.userinfo().base().apiversion()));
        ui->Name->setText(QString::fromStdString(response.userinfo().name()));
        ui->Description->setText(QString::fromStdString(response.userinfo().description()));
        ui->Password->setText(QString::fromStdString(response.userinfo().password()));
        ui->CreateTime->setDateTime(QDateTime::fromSecsSinceEpoch(response.userinfo().createtime().seconds()));
        QPixmap pixmap;
        pixmap.loadFromData(QByteArray::fromStdString(response.userinfo().headphotobindata()));
        ui->HeadPhoto->setPixmap(pixmap);
        ui->PermissionGroup->setText(QString::fromStdString(response.userinfo().userpermissiongroup()));

    }else{
        QMessageBox::critical(this,"Error",QString::fromStdString(result.error_message()));
    }
}
