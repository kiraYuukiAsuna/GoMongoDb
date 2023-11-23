//
// Created by KiraY on 2023/11/22.
//

// You may need to build the project (run Qt uic code generator) to get "ui_EditorSwcMetaInfo.h" resolved

#include "editorswcmetainfo.h"

#include <Message/Response.pb.h>

#include "ui_EditorSwcMetaInfo.h"
#include "src/framework/defination/ImageDefination.h"


EditorSwcMetaInfo::EditorSwcMetaInfo(proto::GetSwcMetaInfoResponse& response, QWidget *parent) :
    QWidget(parent), ui(new Ui::EditorSwcMetaInfo) {
    ui->setupUi(this);
    setWindowIcon(QIcon(Image::ImageNode));

    m_SwcMetaInfo.CopyFrom(response.swcinfo());

    ui->Id->setText(QString::fromStdString(m_SwcMetaInfo.base()._id()));
    ui->Uuid->setText(QString::fromStdString(m_SwcMetaInfo.base().uuid()));
    ui->ApiVersion->setText(QString::fromStdString(m_SwcMetaInfo.base().apiversion()));
    ui->Name->setText(QString::fromStdString(m_SwcMetaInfo.name()));
    ui->Description->setText(QString::fromStdString(m_SwcMetaInfo.description()));
    ui->Creator->setText(QString::fromStdString(m_SwcMetaInfo.creator()));
    ui->CreateTime->setDateTime(QDateTime::fromSecsSinceEpoch(m_SwcMetaInfo.createtime().seconds()));
    ui->LastModifiedTime->setDateTime(QDateTime::fromSecsSinceEpoch(m_SwcMetaInfo.lastmodifiedtime().seconds()));
}

EditorSwcMetaInfo::~EditorSwcMetaInfo() {
    delete ui;
}
