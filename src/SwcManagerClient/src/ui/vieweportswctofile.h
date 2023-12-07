#pragma once

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
