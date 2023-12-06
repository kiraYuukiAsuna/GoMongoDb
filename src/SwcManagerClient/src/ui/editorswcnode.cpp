#include <QListWidgetItem>
#include "editorswcnode.h"
#include "ui_EditorSwcNode.h"
#include "src/framework/service/WrappedCall.h"
#include "src/framework/defination/ImageDefination.h"
#include "viewswcnodedata.h"

EditorSwcNode::EditorSwcNode(const std::string &swcName, QWidget *parent) :
        QWidget(parent), ui(new Ui::EditorSwcNode) {
    ui->setupUi(this);
    std::string stylesheet = std::string("QListWidget::indicator:checked{image:url(")
                             + Image::ImageCheckBoxChecked + ");}" +
                             "QListWidget::indicator:unchecked{image:url(" +
                             Image::ImageCheckBoxUnchecked + ");}";
    ui->UserList->setStyleSheet(QString::fromStdString(stylesheet));

    m_SwcName = swcName;

    connect(ui->AddData,&QPushButton::clicked,this,[this](){
        ViewSwcNodeData editor(this);
        if(editor.exec() == QDialog::Accepted){
            auto swcNodeInternalData = editor.getSwcNodeInternalData();
            proto::SwcDataV1 swcData;
            auto* newData = swcData.add_swcdata();
            newData->mutable_swcnodeinternaldata()->CopyFrom(swcNodeInternalData);

            proto::CreateSwcNodeDataResponse response;
            if(WrappedCall::addSwcNodeData(m_SwcName, swcData, response, this)){
                QMessageBox::information(this,"Info","Create Swc node successfully!");
                refreshAll();
            }
        }
    });

    connect(ui->ModifyData,&QPushButton::clicked,this,[this](){
        ViewSwcNodeData editor(true,this);

        int currectRow = ui->SwcNodeDataTable->currentRow();
        if(currectRow<0){
            QMessageBox::information(this,"Info","You need to select one row first!");
            return;
        }

        if(m_SwcData.swcdata_size() <= currectRow){
            QMessageBox::critical(this,"Error","Swc Data outdated! Please refresh query result!");
        }
        auto InitSwcNodeData = m_SwcData.swcdata().Get(currectRow);
        auto* InitSwcNodeInternalData= InitSwcNodeData.mutable_swcnodeinternaldata();
        editor.setSwcNodeInternalData(*InitSwcNodeInternalData);
        if(editor.exec() == QDialog::Accepted){
            auto swcNodeInternalData = editor.getSwcNodeInternalData();
            proto::SwcNodeDataV1 swcNodeData;
            swcNodeData.CopyFrom(InitSwcNodeData);
            swcNodeData.mutable_swcnodeinternaldata()->CopyFrom(swcNodeInternalData);

            proto::UpdateSwcNodeDataResponse response;
            if(WrappedCall::modifySwcNodeData(m_SwcName, swcNodeData, response, this)){
                QMessageBox::information(this,"Info","Modify Swc node successfully!");
                refreshAll();
            }
        }
    });

    connect(ui->DeleteData,&QPushButton::clicked,this,[this](){
        int currectRow = ui->SwcNodeDataTable->currentRow();
        if(currectRow<0){
            QMessageBox::information(this,"Info","You need to select one row first!");
            return;
        }

        if(m_SwcData.swcdata_size() <= currectRow){
            QMessageBox::critical(this,"Error","Swc Data outdated! Please refresh query result!");
        }
        auto InitSwcNodeData = m_SwcData.swcdata().Get(currectRow);

        auto result = QMessageBox::information(this,"Info","Are your sure to delete this swc node?",
                                               QMessageBox::StandardButton::Ok,QMessageBox::StandardButton::Cancel);
        if(result == QMessageBox::Ok){
            proto::SwcDataV1 swcData;
            auto* newData = swcData.add_swcdata();
            newData->CopyFrom(InitSwcNodeData);

            proto::DeleteSwcNodeDataResponse response;
            if(WrappedCall::deleteSwcNodeData(m_SwcName, swcData, response, this)){
                QMessageBox::information(this,"Info","Delete Swc node successfully!");
                refreshAll();
            }
        }
    });

    connect(ui->QueryAll,&QPushButton::clicked,this,[this](){
        refreshAll();
    });

    connect(ui->QueryByUserAndTime,&QPushButton::clicked,this,[this](){

    });


    refreshUserArea();
}

EditorSwcNode::~EditorSwcNode() {
    delete ui;
}

void EditorSwcNode::refreshAll() {
    proto::GetSwcFullNodeDataResponse response;
    WrappedCall::getSwcFullNodeData(m_SwcName, response, this);

    ui->SwcNodeDataTable->clear();
    ui->SwcNodeDataTable->setRowCount(response.swcnodedata().swcdata_size());
    ui->SwcNodeDataTable->setColumnCount(12);
    QStringList headerLabels;
    headerLabels
            << "n"
            << "type"
            << "x"
            << "y"
            << "z"
            << "radius"
            << "parent"
            << "seg_id"
            << "level"
            << "mode"
            << "timestamp"
            << "feature_value";
    ui->SwcNodeDataTable->setHorizontalHeaderLabels(headerLabels);

    for (int i = 0; i < response.swcnodedata().swcdata_size(); i++) {
        auto info = response.swcnodedata().swcdata().Get(i);
        ui->SwcNodeDataTable->setItem(i, 0,
                                      new QTableWidgetItem(QString::fromStdString(
                                              std::to_string(info.mutable_swcnodeinternaldata()->n()))));
        ui->SwcNodeDataTable->setItem(i, 1,
                                      new QTableWidgetItem(
                                              QString::fromStdString(
                                                      std::to_string(info.mutable_swcnodeinternaldata()->type()))));
        ui->SwcNodeDataTable->setItem(i, 2,
                                      new QTableWidgetItem(
                                              QString::fromStdString(
                                                      std::to_string(info.mutable_swcnodeinternaldata()->x()))));
        ui->SwcNodeDataTable->setItem(i, 3,
                                      new QTableWidgetItem(
                                              QString::fromStdString(
                                                      std::to_string(info.mutable_swcnodeinternaldata()->y()))));
        ui->SwcNodeDataTable->setItem(i, 4,
                                      new QTableWidgetItem(
                                              QString::fromStdString(
                                                      std::to_string(info.mutable_swcnodeinternaldata()->z()))));
        ui->SwcNodeDataTable->setItem(i, 5,
                                      new QTableWidgetItem(
                                              QString::fromStdString(
                                                      std::to_string(info.mutable_swcnodeinternaldata()->radius()))));
        ui->SwcNodeDataTable->setItem(i, 6,
                                      new QTableWidgetItem(
                                              QString::fromStdString(
                                                      std::to_string(info.mutable_swcnodeinternaldata()->parent()))));
        ui->SwcNodeDataTable->setItem(i, 7,
                                      new QTableWidgetItem(
                                              QString::fromStdString(
                                                      std::to_string(info.mutable_swcnodeinternaldata()->seg_id()))));
        ui->SwcNodeDataTable->setItem(i, 8,
                                      new QTableWidgetItem(
                                              QString::fromStdString(
                                                      std::to_string(info.mutable_swcnodeinternaldata()->level()))));
        ui->SwcNodeDataTable->setItem(i, 9,
                                      new QTableWidgetItem(
                                              QString::fromStdString(
                                                      std::to_string(info.mutable_swcnodeinternaldata()->mode()))));
        ui->SwcNodeDataTable->setItem(i, 10,
                                      new QTableWidgetItem(
                                              QString::fromStdString(std::to_string(
                                                      info.mutable_swcnodeinternaldata()->timestamp()))));
        ui->SwcNodeDataTable->setItem(i, 11,
                                      new QTableWidgetItem(
                                              QString::fromStdString(std::to_string(
                                                      info.mutable_swcnodeinternaldata()->feature_value()))));
    }

    m_SwcData.CopyFrom(response.swcnodedata());

    ui->SwcNodeDataTable->resizeColumnsToContents();
}

void EditorSwcNode::refreshUserArea() {
    proto::GetAllUserResponse response;
    WrappedCall::getAllUserMetaInfo(response, this);
    for (int i = 0; i < response.userinfo_size(); i++) {
        auto userInfo = response.userinfo().Get(i);
        auto *item = new QListWidgetItem;
        item->setText(QString::fromStdString(userInfo.name()));
        item->setCheckState(Qt::Checked);
        ui->UserList->addItem(item);
    }
}
