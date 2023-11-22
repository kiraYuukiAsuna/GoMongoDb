#include "editordailystatisticsmetainfo.h"
#include "ui_EditorDailyStatisticsMetaInfo.h"


EditorDailyStatisticsMetaInfo::EditorDailyStatisticsMetaInfo(proto::GetDailyStatisticsResponse& response, QWidget *parent) :
    QWidget(parent), ui(new Ui::EditorDailyStatisticsMetaInfo) {
    ui->setupUi(this);
    m_DailyStatisticsMetaInfo.CopyFrom(response.dailystatisticsinfo());

    ui->Id->setText(QString::fromStdString(m_DailyStatisticsMetaInfo.base()._id()));
    ui->Uuid->setText(QString::fromStdString(m_DailyStatisticsMetaInfo.base().uuid()));
    ui->ApiVersion->setText(QString::fromStdString(m_DailyStatisticsMetaInfo.base().apiversion()));
    ui->Name->setText(QString::fromStdString(m_DailyStatisticsMetaInfo.name()));
    ui->Description->setText(QString::fromStdString(m_DailyStatisticsMetaInfo.description()));
    ui->Day->setText(QString::fromStdString(m_DailyStatisticsMetaInfo.day()));
    ui->CreatedProjectNumber->setText(QString::fromStdString(std::to_string(m_DailyStatisticsMetaInfo.createdprojectnumber())));
    ui->CreatedSwcNumber->setText(QString::fromStdString(std::to_string(m_DailyStatisticsMetaInfo.createdswcnumber())));
    ui->CreateSwcNodeNumber->setText(QString::fromStdString(std::to_string(m_DailyStatisticsMetaInfo.createswcnodenumber())));
    ui->DeletedProjectNumber->setText(QString::fromStdString(std::to_string(m_DailyStatisticsMetaInfo.deletedprojectnumber())));
    ui->DeletedSwcNumber->setText(QString::fromStdString(std::to_string(m_DailyStatisticsMetaInfo.deletedswcnumber())));
    ui->DeletedSwcNodeNumber->setText(QString::fromStdString(std::to_string(m_DailyStatisticsMetaInfo.deletedswcnodenumber())));
    ui->ModifiedProjectNumber->setText(QString::fromStdString(std::to_string(m_DailyStatisticsMetaInfo.modifiedprojectnumber())));
    ui->ModifiedSwcNumber->setText(QString::fromStdString(std::to_string(m_DailyStatisticsMetaInfo.modifiedswcnodenumber())));
    ui->ModifiedSwcNodeNumber->setText(QString::fromStdString(std::to_string(m_DailyStatisticsMetaInfo.modifiedswcnodenumber())));
    ui->ProjectQueryNumber->setText(QString::fromStdString(std::to_string(m_DailyStatisticsMetaInfo.projectquerynumber())));
    ui->SwcQueryNumber->setText(QString::fromStdString(std::to_string(m_DailyStatisticsMetaInfo.swcquerynumber())));
    ui->NodeQueryNumber->setText(QString::fromStdString(std::to_string(m_DailyStatisticsMetaInfo.nodequerynumber())));
    ui->ActiveUserNumber->setText(QString::fromStdString(std::to_string(m_DailyStatisticsMetaInfo.activeusernumber())));

}

EditorDailyStatisticsMetaInfo::~EditorDailyStatisticsMetaInfo() {
    delete ui;
}
