//
// Created by KiraY on 2023/11/23.
//

#ifndef VIEWCREATESWC_H
#define VIEWCREATESWC_H

#include <QWidget>


QT_BEGIN_NAMESPACE
namespace Ui { class ViewCreateSwc; }
QT_END_NAMESPACE

class ViewCreateSwc : public QWidget {
Q_OBJECT

public:
    explicit ViewCreateSwc(QWidget *parent = nullptr);
    ~ViewCreateSwc() override;

private:
    Ui::ViewCreateSwc *ui;
};


#endif //VIEWCREATESWC_H
