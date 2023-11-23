#pragma once

#include <QWidget>
#include <QTreeWidget>
#include <QPushButton>
#include <QVBoxLayout>
#include <QHBoxLayout>

class MainWindow;

QT_BEGIN_NAMESPACE
namespace Ui { class LeftClientView; }
QT_END_NAMESPACE

class LeftClientView : public QWidget {
Q_OBJECT

public:
    explicit LeftClientView(MainWindow *mainWindow);
    ~LeftClientView() override;

    void clearAll();
    void getProjectMetaInfo();
    void getSwcMetaInfo();
    void getAllDailyStatisticsMetaInfo();

public slots:
    void onRefreshBtnClicked(bool checked);
private:
    Ui::LeftClientView *ui;
    QVBoxLayout* m_MainLayout;
    QHBoxLayout* m_ControlBtnLayout;
    MainWindow* m_MainWindow;

    QPushButton* m_UserSettingBtn;
    QPushButton* m_RefreshBtn;

    QTreeWidget* m_TreeWidget;
    QTreeWidgetItem* m_TopProjectItem;
    QTreeWidgetItem* m_TopSwcItem;
    QTreeWidgetItem* m_TopDailyStatisticsItem;
};
