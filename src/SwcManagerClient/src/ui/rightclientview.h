#pragma once

#include <QWidget>


QT_BEGIN_NAMESPACE
namespace Ui { class RightClientView; }
QT_END_NAMESPACE

class RightClientView : public QWidget {
Q_OBJECT

public:
    explicit RightClientView(QWidget *parent = nullptr);
    ~RightClientView() override;

private:
    Ui::RightClientView *ui;
};
