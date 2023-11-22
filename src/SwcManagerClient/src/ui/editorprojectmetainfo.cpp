//
// Created by KiraY on 2023/11/22.
//

// You may need to build the project (run Qt uic code generator) to get "ui_EditorProjectMetaInfo.h" resolved

#include "editorprojectmetainfo.h"
#include "ui_EditorProjectMetaInfo.h"
#include <QDateTime>

#include "MainWindow.h"

EditorProjectMetaInfo::EditorProjectMetaInfo(proto::GetProjectResponse&response, QWidget* parent) : QWidget(parent),
    ui(new Ui::EditorProjectMetaInfo) {
    ui->setupUi(this);
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
    ui->Uuid->setText(QString::fromStdString(m_ProjectMetaInfo.name()));

    ui->SwcList->clear();
    for (int i = 0; i < m_ProjectMetaInfo.swclist().size(); i++) {
        auto name = m_ProjectMetaInfo.swclist().Get(i);
        ui->SwcList->addItem(QString::fromStdString(name));
    }

    ui->UserPermissionOverride->clear();
    ui->UserPermissionOverride->setRowCount(m_ProjectMetaInfo.userpermissionoverride().size());
    ui->UserPermissionOverride->setColumnCount(8);
    QStringList headerLabels;
    headerLabels
            << "UserName"
            << "Value"
            << "GlobalWritePermissionCreateProject"
            << "GlobalWritePermissionModifyProject"
            << "GlobalPermissionDeleteProject"
            << "GlobalReadPerimissionQuery"
            << "ProjectWritePermissionAddData"
            << "ProjectWritePermissionModifyData"
            << "ProjectWritePermissionDeleteData"
            << "ProjectReadPerimissionQuery";

    ui->UserPermissionOverride->setHorizontalHeaderLabels(headerLabels);
    for (int i = 0; i < m_ProjectMetaInfo.userpermissionoverride().size(); i++) {
        auto permissionOverride = m_ProjectMetaInfo.userpermissionoverride().Get(i);
        ui->UserPermissionOverride->setItem(i, 0,
                                            new QTableWidgetItem(
                                                QString::fromStdString(permissionOverride.username())));
        ui->UserPermissionOverride->setItem(i, 1,
                                            new QTableWidgetItem(
                                                QString::fromStdString(std::to_string(permissionOverride.projectpermission().writepermissionadddata()))));
        ui->UserPermissionOverride->setItem(i, 2,
                                            new QTableWidgetItem(
                                            QString::fromStdString(std::to_string(permissionOverride.projectpermission().writepermissionmodifydata()))));
        ui->UserPermissionOverride->setItem(i, 3,
                                            new QTableWidgetItem(
                                            QString::fromStdString(std::to_string(permissionOverride.projectpermission().writepermissiondeletedata()))));
        ui->UserPermissionOverride->setItem(i, 4,
                                            new QTableWidgetItem(
                                            QString::fromStdString(std::to_string(permissionOverride.projectpermission().readperimissionquery()))));
    }
    ui->UserPermissionOverride->resizeColumnsToContents();
}
