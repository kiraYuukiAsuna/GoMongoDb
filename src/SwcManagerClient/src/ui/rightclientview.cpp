#include "rightclientview.h"

#include "editordailystatisticsmetainfo.h"
#include "ui_RightClientView.h"
#include "mainwindow.h"
#include "src/framework/service/WrappedCall.h"
#include "editorprojectmetainfo.h"
#include "editorswcmetainfo.h"
#include "src/framework/defination/ImageDefination.h"

RightClientView::RightClientView(MainWindow *mainWindow) :
    QWidget(mainWindow), ui(new Ui::RightClientView) {
    ui->setupUi(this);
    m_MainWindow = mainWindow;

    m_TabWidget = new QTabWidget(this);
    m_TabWidget->setTabsClosable(true);
    connect(m_TabWidget,&QTabWidget::tabCloseRequested,this,[&](int index) {
        m_TabWidget->removeTab(index);
    });

    m_MainLayout = new QVBoxLayout(this);
    m_MainLayout->addWidget(m_TabWidget);
    this->setLayout(m_MainLayout);

}

RightClientView::~RightClientView() {
    delete ui;
}

void RightClientView::openProjectMetaInfo(const std::string &projectName) {
    auto index = findIfTabAlreadOpenned(projectName, MetaInfoType::eProject);
    if(index != -1) {
        m_TabWidget->setCurrentIndex(index);
        return;
    }

    proto::GetProjectResponse response;
    if(WrappedCall::getProjectMetaInfoByName(projectName, response, this)){
        auto* editor = new EditorProjectMetaInfo(response, m_TabWidget);
        auto newIndex = m_TabWidget->addTab(editor, QIcon(Image::ImageProject), QString::fromStdString(response.projectinfo().name()));
        m_TabWidget->setCurrentIndex(newIndex);
    }
}

void RightClientView::openSwcMetaInfo(const std::string &swcName) {
    auto index = findIfTabAlreadOpenned(swcName, MetaInfoType::eSwc);
    if(index != -1) {
        m_TabWidget->setCurrentIndex(index);
        return;
    }

    proto::GetSwcMetaInfoResponse response;
    if(WrappedCall::getSwcMetaInfoByName(swcName, response,this)){
        auto* editor = new EditorSwcMetaInfo(response, m_TabWidget);
        auto newIndex = m_TabWidget->addTab(editor, QIcon(Image::ImageNode), QString::fromStdString(response.swcinfo().name()));
        m_TabWidget->setCurrentIndex(newIndex);
    }
}

void RightClientView::openDailyStatisticsMetaInfo(const std::string &dailyStatisticsName) {
    auto index = findIfTabAlreadOpenned(dailyStatisticsName, MetaInfoType::eDailyStatistics);
    if(index != -1) {
        m_TabWidget->setCurrentIndex(index);
        return;
    }

    proto::GetDailyStatisticsResponse response;
    if(WrappedCall::getDailyStatisticsmMetaInfoByName(dailyStatisticsName, response, this)) {
        auto *editor = new EditorDailyStatisticsMetaInfo(response, m_TabWidget);
        auto newIndex = m_TabWidget->addTab(editor, QIcon(Image::ImageDaily),
                                            QString::fromStdString(response.dailystatisticsinfo().name()));
        m_TabWidget->setCurrentIndex(newIndex);
    }
}

int RightClientView::findIfTabAlreadOpenned(const std::string& name, MetaInfoType metaInfoType) {
    for (int i=0;i<m_TabWidget->count();i++) {
        auto editorBase = dynamic_cast<EditorBase*>(m_TabWidget->widget(i));
        if(!editorBase) {
            continue;
        }

        if(editorBase->getName() == name && editorBase->getMetaInfoType() == metaInfoType) {
            return i;
        }
    }
    return -1;
}

