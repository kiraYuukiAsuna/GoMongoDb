//
// Created by KiraY on 2023/11/23.
//

#ifndef VIEWEPORTSWCTOFILE_H
#define VIEWEPORTSWCTOFILE_H

#include <QDialog>


QT_BEGIN_NAMESPACE
namespace Ui { class ViewEportSwcToFile; }
QT_END_NAMESPACE

class ViewEportSwcToFile : public QDialog {
Q_OBJECT

public:
    explicit ViewEportSwcToFile(QWidget *parent = nullptr);
    ~ViewEportSwcToFile() override;

private:
    Ui::ViewEportSwcToFile *ui;
};


#endif //VIEWEPORTSWCTOFILE_H
