#pragma once

#include <QWidget>


QT_BEGIN_NAMESPACE
namespace Ui { class LeftClientView; }
QT_END_NAMESPACE

class LeftClientView : public QWidget {
Q_OBJECT

public:
    explicit LeftClientView(QWidget *parent = nullptr);
    ~LeftClientView() override;

private:
    Ui::LeftClientView *ui;
};
