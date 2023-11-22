#include "rightclientview.h"
#include "ui_RightClientView.h"
#include "mainwindow.h"

RightClientView::RightClientView(MainWindow *mainWindow) :
    QWidget(mainWindow), ui(new Ui::RightClientView) {
    ui->setupUi(this);
    m_MainWindow = mainWindow;

    m_TabWidget = new QTabWidget(this);

    m_MainLayout = new QVBoxLayout(this);
    m_MainLayout->addWidget(m_TabWidget);
    this->setLayout(m_MainLayout);


}



RightClientView::~RightClientView() {
    delete ui;
}

void RightClientView::openProjectMetaInfo(const std::string &projectName) {


}

void RightClientView::openSwcMetaInfo(const std::string &swcName) {

}

void RightClientView::openDailyStatisticsMetaInfo(const std::string &dailyStatisticsName) {

}
