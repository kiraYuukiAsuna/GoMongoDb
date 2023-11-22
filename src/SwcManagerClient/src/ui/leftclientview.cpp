#include "leftclientview.h"

#include <QMessageBox>

#include "ui_LeftClientView.h"
#include "src/framework/service/CachedProtoData.h"
#include "src/framework/service/RpcCall.h"


LeftClientView::LeftClientView(QWidget *parent) :
    QWidget(parent), ui(new Ui::LeftClientView) {
    ui->setupUi(this);

    m_ControlBtnLayout = new QHBoxLayout;

    m_RefreshBtn = new QPushButton(this);
    m_RefreshBtn->setText("Refresh");
    connect(m_RefreshBtn,&QPushButton::clicked,this,&LeftClientView::onRefreshBtnClicked);

    m_ControlBtnLayout->addWidget(m_RefreshBtn);

    m_TreeWidget = new QTreeWidget(this);
    m_TreeWidget->setHeaderLabel("MetaInfo");

    m_TopProjectItem = new QTreeWidgetItem(m_TreeWidget);
    m_TopProjectItem->setText(0,"Project");

    m_TopSwcItem= new QTreeWidgetItem(m_TreeWidget);
    m_TopSwcItem->setText(0,"Swc");

    m_TopDailyStatisticsItem= new QTreeWidgetItem(m_TreeWidget);
    m_TopDailyStatisticsItem->setText(0,"DailyStatistics");


    m_TreeWidget->addTopLevelItem(m_TopProjectItem);
    m_TreeWidget->addTopLevelItem(m_TopSwcItem);
    m_TreeWidget->addTopLevelItem(m_TopDailyStatisticsItem);

    m_MainLayout = new QVBoxLayout(this);
    m_MainLayout->addLayout(m_ControlBtnLayout);
    m_MainLayout->addWidget(m_TreeWidget);
    this->setLayout(m_MainLayout);

    getProjectMetaInfo();
    getSwcMetaInfo();
    getDailyStatisticsMetaInfo();
}

LeftClientView::~LeftClientView() {
    delete ui;
}

void LeftClientView::getProjectMetaInfo() {
    grpc::ClientContext context;
    proto::GetAllProjectRequest request;
    auto* userInfo = request.mutable_userinfo();
    userInfo->CopyFrom(CachedProtoData::getInstance().CachedUserMetaInfo);

    auto& rpcCall = RpcCall::getInstance();
    proto::GetAllProjectResponse response;
    auto status = rpcCall.Stub()->GetAllProject(&context, request, &response);
    if(status.ok()){
        if(response.status()) {
            auto projectInfoList = response.mutable_projectinfo();
            for(int i=0; i<projectInfoList->size();i++) {
                auto& projectInfo = projectInfoList->Get(i);
                auto* item = new QTreeWidgetItem;
                item->setText(0,QString::fromStdString(projectInfo.name()));
                m_TopProjectItem->addChild(item);
            }
        }else {
            QMessageBox::warning(this,"Info","GetAllProjectMetaInfo Failed!" + QString::fromStdString(response.message()));
        }

    }else{
        QMessageBox::critical(this,"Error",QString::fromStdString(status.error_message()));
    }
}

void LeftClientView::getSwcMetaInfo() {
    grpc::ClientContext context;
    proto::GetAllSwcMetaInfoRequest request;
    auto* userInfo = request.mutable_userinfo();
    userInfo->CopyFrom(CachedProtoData::getInstance().CachedUserMetaInfo);

    auto& rpcCall = RpcCall::getInstance();
    proto::GetAllSwcMetaInfoResponse response;
    auto status = rpcCall.Stub()->GetAllSwcMetaInfo(&context, request, &response);
    if(status.ok()){
        if(response.status()) {
            auto swcMetaInfo = response.mutable_swcinfo();
            for(int i=0; i<swcMetaInfo->size();i++) {
                auto& swcInfo = swcMetaInfo->Get(i);
                auto* item = new QTreeWidgetItem;
                item->setText(0,QString::fromStdString(swcInfo.name()));
                m_TopSwcItem->addChild(item);
            }
        }else {
            QMessageBox::warning(this,"Info","GetAllSwcMetaInfo Failed!" + QString::fromStdString(response.message()));
        }

    }else{
        QMessageBox::critical(this,"Error",QString::fromStdString(status.error_message()));
    }
}

void LeftClientView::getDailyStatisticsMetaInfo() {
    grpc::ClientContext context;
    proto::GetAllDailyStatisticsRequest request;
    auto* userInfo = request.mutable_userinfo();
    userInfo->CopyFrom(CachedProtoData::getInstance().CachedUserMetaInfo);

    auto& rpcCall = RpcCall::getInstance();
    proto::GetAllDailyStatisticsResponse response;
    auto status = rpcCall.Stub()->GetAllDailyStatistics(&context, request, &response);
    if(status.ok()){
        if(response.status()) {
            auto dailyStatisticsMetaInfoList = response.mutable_dailystatisticsinfo();
            for(int i=0; i<dailyStatisticsMetaInfoList->size();i++) {
                auto& dailyStatisticsMetaInfo = dailyStatisticsMetaInfoList->Get(i);
                auto* item = new QTreeWidgetItem;
                item->setText(0,QString::fromStdString(dailyStatisticsMetaInfo.name()));
                m_TopDailyStatisticsItem->addChild(item);
            }
        }else {
            QMessageBox::warning(this,"Info","GetAllDailyStatistics Failed!" + QString::fromStdString(response.message()));
        }

    }else{
        QMessageBox::critical(this,"Error",QString::fromStdString(status.error_message()));
    }
}

void LeftClientView::onRefreshBtnClicked(bool checked) {
    getProjectMetaInfo();
    getSwcMetaInfo();
    getDailyStatisticsMetaInfo();
}
