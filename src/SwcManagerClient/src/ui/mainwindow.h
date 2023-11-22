#pragma once

#include <QMainWindow>
#include <QSplitter>
#include "leftclientview.h"
#include "rightclientview.h"

QT_BEGIN_NAMESPACE
namespace Ui { class MainWindow; }
QT_END_NAMESPACE

class MainWindow : public QMainWindow {
Q_OBJECT

public:
    explicit MainWindow(QWidget *parent = nullptr);
    ~MainWindow() override;

    LeftClientView& getLeftClientView();
    RightClientView& getRightClientView();

private:
    Ui::MainWindow *ui;

    QSplitter* m_Splitter;
    LeftClientView* m_LeftClientView;
    RightClientView* m_RightClientView;

};
