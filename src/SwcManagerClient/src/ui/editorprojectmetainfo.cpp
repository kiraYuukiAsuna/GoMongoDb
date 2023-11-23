//
// Created by KiraY on 2023/11/22.
//

// You may need to build the project (run Qt uic code generator) to get "ui_EditorProjectMetaInfo.h" resolved

#include "editorprojectmetainfo.h"
#include "ui_EditorProjectMetaInfo.h"
#include <QDateTime>

#include "MainWindow.h"
#include "src/framework/defination/ImageDefination.h"
#include "src/framework/service/WrappedCall.h"

EditorProjectMetaInfo::EditorProjectMetaInfo(proto::GetProjectResponse&response, QWidget* parent) : QWidget(parent),
                                                                                                    ui(new Ui::EditorProjectMetaInfo) {
    ui->setupUi(this);
    setWindowIcon(QIcon(Image::ImageProject));

    refresh(response);
}

EditorProjectMetaInfo::~EditorProjectMetaInfo() {
    delete ui;
}

void EditorProjectMetaInfo::refresh(proto::GetProjectResponse& response) {
        m_ProjectMetaInfo.CopyFrom(response.projectinfo());

    ui->Id->setText(QString::fromStdString(m_ProjectMetaInfo.base()._id()));
    ui->Uuid->setText(QString::fromStdString(m_ProjectMetaInfo.base().uuid()));
    ui->ApiVersion->setText(QString::fromStdString(m_ProjectMetaInfo.base().apiversion()));
    ui->Name->setText(QString::fromStdString(m_ProjectMetaInfo.name()));
    ui->Creator->setText(QString::fromStdString(m_ProjectMetaInfo.creator()));
    ui->WorkMode->setText(QString::fromStdString(m_ProjectMetaInfo.workmode()));
    ui->CreateTime->setDateTime(QDateTime::fromSecsSinceEpoch(m_ProjectMetaInfo.createtime().seconds()));
    ui->LastModifiedTime->setDateTime(QDateTime::fromSecsSinceEpoch(m_ProjectMetaInfo.lastmodifiedtime().seconds()));
    ui->Dsecription->setText(QString::fromStdString(m_ProjectMetaInfo.description()));
    
    ui->SwcList->clear();

    std::string stylesheet = std::string("QListWidget::indicator:checked{image:url(")
                            + Image::ImageCheckBoxChecked + ");}" +
                                "QListWidget::indicator:unchecked{image:url(" +
                                    Image::ImageCheckBoxUnchecked+");}";
    ui->SwcList->setStyleSheet(QString::fromStdString(stylesheet));

    proto::GetAllSwcMetaInfoResponse responseAllSwc;
    WrappedCall::getAllSwcMetaInfo(responseAllSwc, this);
    for (int i = 0; i < responseAllSwc.swcinfo_size(); i++) {
        auto swcInfo = responseAllSwc.swcinfo().Get(i);
        bool bFind=false;
        auto* item = new QListWidgetItem;
        item->setText(QString::fromStdString(swcInfo.name()));
        for (int j = 0; j < m_ProjectMetaInfo.swclist().size(); j++) {
            auto name = m_ProjectMetaInfo.swclist().Get(j);
            if(name == swcInfo.name()) {
                bFind = true;
            }
        }
        if(bFind) {
            item->setCheckState(Qt::Checked);
        }else {
            item->setCheckState(Qt::Unchecked);
        }
        ui->SwcList->addItem(item);
    }

    proto::GetAllUserResponse responseAllUser;
    WrappedCall::getAllUserMetaInfo(responseAllUser, this);
    ui->UserPermissionOverride->clear();
    ui->UserPermissionOverride->setRowCount(responseAllUser.userinfo_size());
    ui->UserPermissionOverride->setColumnCount(5);
    QStringList headerLabels;
    headerLabels
            << "UserName"
            << "ProjectWritePermissionAddData"
            << "ProjectWritePermissionModifyData"
            << "ProjectWritePermissionDeleteData"
            << "ProjectReadPerimissionQuery";
    ui->UserPermissionOverride->setHorizontalHeaderLabels(headerLabels);

    for (int i = 0; i < responseAllUser.userinfo_size(); i++) {
        auto userInfo = responseAllUser.userinfo().Get(i);
        auto userNameItem = new QTableWidgetItem(
            QString::fromStdString(userInfo.name()));
        ui->UserPermissionOverride->setItem(i, 0, userNameItem);
        ui->UserPermissionOverride->setItem(i, 1,
                                 new QTableWidgetItem(
                                     QString::fromStdString(
                                         std::to_string(1))));
        ui->UserPermissionOverride->setItem(i, 2,
                                 new QTableWidgetItem(
                                     QString::fromStdString(
                                         std::to_string(1))));
        ui->UserPermissionOverride->setItem(i, 3,
                                 new QTableWidgetItem(
                                     QString::fromStdString(
                                         std::to_string(1))));
        ui->UserPermissionOverride->setItem(i, 4,
                                 new QTableWidgetItem(
                                     QString::fromStdString(
                                         std::to_string(1))));
        bool bFind = false;
        for (int j = 0; j < m_ProjectMetaInfo.userpermissionoverride().size(); j++) {
            auto permissionOverride = m_ProjectMetaInfo.userpermissionoverride().Get(i);
            if(permissionOverride.username() == userInfo.name()) {
                bFind = true;
            }
        }
        if(bFind) {
            userNameItem->setCheckState(Qt::Checked);
        }else {
            userNameItem->setCheckState(Qt::Unchecked);
        }
    }

    ui->UserPermissionOverride->resizeColumnsToContents();
}
