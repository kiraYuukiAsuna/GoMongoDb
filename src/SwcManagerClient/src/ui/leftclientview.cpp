#include "leftclientview.h"

#include <QMessageBox>
#include "mainwindow.h"
#include "ui_LeftClientView.h"
#include "src/framework/service/CachedProtoData.h"
#include "src/framework/service/RpcCall.h"
#include "src/framework/defination/TypeDef.h"
#include "src/framework/defination/ImageDefination.h"
#include "src/framework/service/WrappedCall.h"

LeftClientView::LeftClientView(MainWindow* mainWindow) :
    QWidget(mainWindow), ui(new Ui::LeftClientView) {
    ui->setupUi(this);
    m_MainWindow = mainWindow;

    m_ControlBtnLayout = new QHBoxLayout;

    m_RefreshBtn = new QPushButton(this);
    m_RefreshBtn->setText("Refresh");
    connect(m_RefreshBtn,&QPushButton::clicked,this,&LeftClientView::onRefreshBtnClicked);

    m_ControlBtnLayout->addWidget(m_RefreshBtn);

    m_TreeWidget = new QTreeWidget(this);
    m_TreeWidget->setHeaderLabel("MetaInfo");
    connect(m_TreeWidget,&QTreeWidget::itemDoubleClicked, this, [&](QTreeWidgetItem* item, int column){
        if(column == 0){
            if(item){
                auto metaInfo = item->data(0,Qt::UserRole).value<LeftClientViewTreeWidgetMetaInfo>();
                switch(metaInfo.type) {
                    case MetaInfoType::eProjectContainer:
                        break;
                    case MetaInfoType::eProject: {
                        m_MainWindow->getRightClientView().openProjectMetaInfo(metaInfo.name);
                        break;
                    }
                    case MetaInfoType::eSwcContainer:
                        break;
                    case MetaInfoType::eSwc: {
                        m_MainWindow->getRightClientView().openSwcMetaInfo(metaInfo.name);
                        break;
                    }
                    case MetaInfoType::eDailyStatisticsContainer:
                        break;
                    case MetaInfoType::eDailyStatistics: {
                        m_MainWindow->getRightClientView().openDailyStatisticsMetaInfo(metaInfo.name);
                        break;
                    }
                    case MetaInfoType::eUserMetaInfo:
                        break;
                    case MetaInfoType::ePermissionGroupMetaInfo:
                        break;
                    case MetaInfoType::eUserManagerMetaInfo:
                        break;
                    case MetaInfoType::eSwcData:
                        break;
                    case MetaInfoType::eUnknown:
                        break;
                }
            }
        }
    });

    m_MainLayout = new QVBoxLayout(this);
    m_MainLayout->addLayout(m_ControlBtnLayout);
    m_MainLayout->addWidget(m_TreeWidget);
    this->setLayout(m_MainLayout);

    clearAll();
    getProjectMetaInfo();
    getSwcMetaInfo();
    getAllDailyStatisticsMetaInfo();
}

LeftClientView::~LeftClientView() {
    delete ui;
}

void LeftClientView::getProjectMetaInfo() {
    proto::GetAllProjectResponse response;
    WrappedCall::getAllProjectMetaInfo(response, this);
    auto projectInfoList = response.mutable_projectinfo();
    for(int i=0; i<projectInfoList->size();i++) {
        auto& projectInfo = projectInfoList->Get(i);
        auto* item = new QTreeWidgetItem;
        item->setText(0,QString::fromStdString(projectInfo.name()));
        item->setIcon(0,QIcon(Image::ImageProject));
        LeftClientViewTreeWidgetMetaInfo metaInfo{};
        metaInfo.type = MetaInfoType::eProject;
        metaInfo.name = projectInfo.name();
        item->setData(0,Qt::UserRole,QVariant::fromValue(metaInfo));
        m_TopProjectItem->addChild(item);
    }
}

void LeftClientView::getSwcMetaInfo() {
    proto::GetAllSwcMetaInfoResponse response;
    WrappedCall::getAllSwcMetaInfo(response, this);
    auto swcMetaInfo = response.mutable_swcinfo();
    for(int i=0; i<swcMetaInfo->size();i++) {
        auto& swcInfo = swcMetaInfo->Get(i);
        auto* item = new QTreeWidgetItem;
        item->setText(0,QString::fromStdString(swcInfo.name()));
        item->setIcon(0,QIcon(Image::ImageNode));
        LeftClientViewTreeWidgetMetaInfo metaInfo{};
        metaInfo.type = MetaInfoType::eSwc;
        metaInfo.name = swcInfo.name();
        item->setData(0,Qt::UserRole,QVariant::fromValue(metaInfo));
        m_TopSwcItem->addChild(item);
    }
}

void LeftClientView::getAllDailyStatisticsMetaInfo() {
    proto::GetAllDailyStatisticsResponse response;
    WrappedCall::getAllDailyStatisticsMetaInfo(response, this);
    auto dailyStatisticsMetaInfoList = response.mutable_dailystatisticsinfo();
    for(int i=0; i<dailyStatisticsMetaInfoList->size();i++) {
        auto& dailyStatisticsMetaInfo = dailyStatisticsMetaInfoList->Get(i);
        auto* item = new QTreeWidgetItem;
        item->setText(0,QString::fromStdString(dailyStatisticsMetaInfo.name()));
        item->setIcon(0,QIcon(Image::ImageDaily));
        LeftClientViewTreeWidgetMetaInfo metaInfo{};
        metaInfo.type = MetaInfoType::eDailyStatistics;
        metaInfo.name = dailyStatisticsMetaInfo.name();
        item->setData(0,Qt::UserRole,QVariant::fromValue(metaInfo));
        m_TopDailyStatisticsItem->addChild(item);
    }
}

void LeftClientView::onRefreshBtnClicked(bool checked) {
    clearAll();
    getProjectMetaInfo();
    getSwcMetaInfo();
    getAllDailyStatisticsMetaInfo();
}

void LeftClientView::clearAll() {
    m_TreeWidget->clear();

    m_TopProjectItem = new QTreeWidgetItem(m_TreeWidget);
    m_TopProjectItem->setText(0,"Project");
    m_TopProjectItem->setIcon(0,QIcon(Image::ImageProject));
    LeftClientViewTreeWidgetMetaInfo metaInfoProject{};
    metaInfoProject.type = MetaInfoType::eProjectContainer;
    metaInfoProject.name = "Project";
    m_TopProjectItem->setData(0,Qt::UserRole,QVariant::fromValue(metaInfoProject));

    m_TopSwcItem= new QTreeWidgetItem(m_TreeWidget);
    m_TopSwcItem->setText(0,"Swc");
    m_TopSwcItem->setIcon(0,QIcon(Image::ImageNode));
    LeftClientViewTreeWidgetMetaInfo metaInfoSwc{};
    metaInfoSwc.type = MetaInfoType::eSwcContainer;
    metaInfoProject.name = "Swc";
    m_TopSwcItem->setData(0,Qt::UserRole,QVariant::fromValue(metaInfoSwc));

    m_TopDailyStatisticsItem= new QTreeWidgetItem(m_TreeWidget);
    m_TopDailyStatisticsItem->setText(0,"DailyStatistics");
    m_TopDailyStatisticsItem->setIcon(0,QIcon(Image::ImageDaily));
    LeftClientViewTreeWidgetMetaInfo metaInfoDailyStatistic{};
    metaInfoDailyStatistic.type = MetaInfoType::eDailyStatisticsContainer;
    metaInfoProject.name = "DailyStatistics";
    m_TopDailyStatisticsItem->setData(0,Qt::UserRole,QVariant::fromValue(metaInfoDailyStatistic));

    m_TreeWidget->addTopLevelItem(m_TopProjectItem);
    m_TreeWidget->addTopLevelItem(m_TopSwcItem);
    m_TreeWidget->addTopLevelItem(m_TopDailyStatisticsItem);
}
