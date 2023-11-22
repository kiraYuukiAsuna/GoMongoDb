#pragma once

#include <QWidget>
#include <QTabWidget>
#include <QVBoxLayout>

class MainWindow;

QT_BEGIN_NAMESPACE
namespace Ui { class RightClientView; }
QT_END_NAMESPACE

class RightClientView : public QWidget {
Q_OBJECT

public:
    explicit RightClientView(MainWindow *mainWindow);
    ~RightClientView() override;

    void openProjectMetaInfo(const std::string& projectName);
    void openSwcMetaInfo(const std::string& swcName);
    void openDailyStatisticsMetaInfo(const std::string& dailyStatisticsName);

private:
    Ui::RightClientView *ui;
    QVBoxLayout* m_MainLayout;
    MainWindow *m_MainWindow;

    QTabWidget* m_TabWidget;
};
